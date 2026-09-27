# Data pipeline

The `data/` Go module collects listings from Divar, stores source rows in PostgreSQL, repairs selected fields, and builds a one-hot feature table for model training.

## Storage

- `cars` contains the source record fields: year, color, usage, body condition, cash cost, motor and chassis conditions, insurance due, gearbox, and listing token.
- `pcars` contains the expanded one-hot features, numeric values, target cost, and listing token.
- SQL migrations are in `data/internal/database/migrations/`. The service applies migrations on startup when built with the `postgres` build tag.

## Configure and run

The service loads `data/.env`. Set PostgreSQL values, `MIGRATIONS_PATH`, and `LOG_NAME`. The fetch stage additionally requires `CITIES` (space-separated Divar city codes) and `BRAND_MODEL` (Divar's brand/model query value).

From `data/`:

```sh
go run -tags postgres . -fetch
go run -tags postgres . -preprocess
go run -tags postgres . -onehot
```

The stages can also be requested together because the program runs them in fetch, preprocess, then one-hot order:

```sh
go run -tags postgres . -fetch -preprocess -onehot
```

Collection searches Divar and retrieves each unseen listing with a one-second delay. Existing tokens are skipped. Network or parse failures are logged and skipped during listing retrieval.

## Preprocessing behavior

Preprocessing mutates the raw `cars` table. In the current implementation it:

- Replaces out-of-range years with the most frequent in-range year (the command uses 1389 through 1400).
- Fills missing category values with the most frequent observed value.
- Replaces non-positive or over-600,000 km usage with the average of accepted usage values.
- Deletes rows with cash cost below 100,000,000 or above 260,000,000.

The monetary bounds and imputation rules are hard-coded in `data/services/preprocess/service.go`; they are not settings. Back up the database before running preprocessing on data you need to preserve. The average-usage calculation also assumes that at least one usage value falls within its accepted range.

## One-hot encoding and export

One-hot encoding reads all rows from `cars` and inserts a transformed row for each into `pcars`. It does not clear or deduplicate `pcars`, so repeated runs create duplicates. The `pcars` schema and feature field order are part of the training and inference contract.

The separate Python exporter at `pycsv/to_csv.py` reads the full `pcars` table and writes `pycsv/pcars_data.csv`. Its PostgreSQL connection string is currently embedded in the script; update it for your environment before use. See [Machine learning](../machine-learning.md) for training details.
