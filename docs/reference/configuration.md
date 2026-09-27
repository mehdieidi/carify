# Configuration reference

The Go services load a `.env` file from their current working directory with `godotenv`. The backend and data service use separate files. Values below describe keys read by the code; do not copy private credentials into documentation or commit local secrets.

Both component `.env` files are currently tracked by Git. Treat their contents as repository-visible, and replace any real credentials before sharing the repository.

## Backend (`backend/.env`)

| Variable | Purpose |
| --- | --- |
| `SERVER_ADDR` | Address passed to the HTTP server, such as `127.0.0.1:8082`. |
| `LOG_NAME` | File path/name opened for append logging. Relative paths resolve from `backend/`. |

The backend also reads categorical options from `backend/settings/settings.json`. Prediction expects `python3`, `../predictor/predictor.py`, and the repository-root `model/model.joblib` path to resolve from the backend working directory.

## Data service (`data/.env`)

| Variable | Purpose |
| --- | --- |
| `POSTGRES_HOST` | PostgreSQL host. |
| `POSTGRES_PORT` | PostgreSQL port. |
| `POSTGRES_USER` | Database user. |
| `POSTGRES_PASSWORD` | Database password. |
| `POSTGRES_DB` | Database name. |
| `LOG_NAME` | Log file path/name. |
| `MIGRATIONS_PATH` | Migration source understood by `golang-migrate`; the checked-in value is relative to the `data/` directory. |
| `CITIES` | Space-separated Divar city codes used by `-fetch`. |
| `BRAND_MODEL` | Search value passed to Divar by `-fetch`. |

The data executable must be built or run with the `postgres` build tag to use the PostgreSQL implementation.

## Other tools

- `pycsv/to_csv.py` contains its PostgreSQL SQLAlchemy connection string in source rather than reading environment variables.
- `frontend/src/components/Home.js` contains the API origin in source rather than reading a frontend environment variable.
- The Python scripts have no checked-in requirements or environment specification.
