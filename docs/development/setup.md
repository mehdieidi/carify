# Development setup

The repository is a set of separate services and scripts rather than one root application. Run commands from the component directory because `.env` files, settings files, CSV paths, and model paths are relative to the current working directory.

## Prerequisites

- Go toolchain compatible with the module declarations (`go 1.19`) for the Go services.
- Node.js and npm for the React app. Dependencies and scripts are declared in `frontend/package.json` and `frontend/package-lock.json`.
- Python 3 available as `python3` for exporting, training, plotting, and API inference.
- PostgreSQL for the data pipeline.
- Python packages used by the scripts: `pandas`, `SQLAlchemy`, `psycopg2`, `scikit-learn`, and `joblib`. The plotting script also imports `matplotlib`, `seaborn`, `numpy`, and `bokeh`.

The repository does not include a Python dependency lock or requirements file, a Docker Compose setup, or a tracked model artifact.

## Configure the data service

The data process loads `data/.env`. The required keys are listed in the [configuration reference](../reference/configuration.md). Configure PostgreSQL connection values, the migrations source path, and a log filename. To collect listings, also provide Divar city codes and a brand/model search string.

From `data/`, run the stage or stages you need:

```sh
go run -tags postgres . -fetch
go run -tags postgres . -preprocess
go run -tags postgres . -onehot
```

The `postgres` build tag is required for the real PostgreSQL implementation. Without it, the database methods deliberately panic. The `-fetch` stage calls Divar over the network. Preprocessing modifies `cars` rows and deletes listings outside the implemented price bounds. One-hot encoding inserts new `pcars` rows; it does not clear existing rows first. Use a disposable database or back up data before repeating mutating stages.

The database schema is applied on process startup from the configured migration path. Keep that path valid from the `data/` working directory.

## Run the API

From `backend/`:

```sh
go run .
```

The process loads `backend/.env`, reads `./settings/settings.json`, and listens at `SERVER_ADDR` (the checked-in local value is `127.0.0.1:8082`). Run from `backend/` so the settings and `../predictor/predictor.py` paths resolve. See [Backend service](../services/backend.md) and the [API reference](../reference/api.md).

Predictions also require a trained artifact at `model/model.joblib` in the repository root and a Python environment with the model dependencies installed. The artifact is ignored by Git and is not included in this checkout.

## Run the frontend

From `frontend/`:

```sh
npm ci
npm start
```

Create React App opens the development app at `http://localhost:3000`. The current source uses a fixed API URL rather than a configurable environment variable. For a local end-to-end session, point that URL at the backend address and ensure the model artifact is available. See [Frontend service](../services/frontend.md).

## Train a model

After the `pcars` table has data, export it and train from the expected directories:

```sh
# From pycsv/
python3 to_csv.py

# From regressor/
python3 linear_regression.py
# or
python3 decision_tree.py
```

The training scripts expect `pycsv/pcars_data.csv` and write `model/model.joblib`; create the `model/` directory before training. Both scripts write to the same artifact path, so the last one run replaces the earlier model. See [Machine learning](../machine-learning.md).

## Current repository constraints

- The model artifact is excluded from version control, so inference needs a local training run or a separately supplied compatible model.
- The frontend source points to `http://45.67.85.169:8080`, while the checked-in backend `.env` listens on `127.0.0.1:8082`. The UI does not currently read an API base URL from environment configuration.
- The `orchestrator` executable uses fixed relative paths, ignores command failures, and tries to run `regressor/note-exp.py`, which is not present in the repository. Run the documented stages manually instead.
- Python dependencies are imported by scripts but are not captured in a repository requirements or lock file.
- The training scripts select model columns by CSV position. Changes in database/export column order can silently change the feature contract.
- The submitted `insurance` input is currently omitted from the inference feature row.
- The checked-in frontend test is still the Create React App starter test and looks for starter content rather than the Carify screen.

These notes describe the checked-in code, not a guarantee that remote listing endpoints or an external API host remain available.
