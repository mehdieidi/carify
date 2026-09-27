# Backend service

The backend is a Go HTTP server in `backend/`. It serves form settings from a JSON file and runs the Python predictor as a child process for cost estimates. It does not read car data or model weights directly from PostgreSQL.

## Start

Run from the `backend/` directory:

```sh
go run .
```

At startup it loads `.env`, opens the configured log file, reads `./settings/settings.json`, and listens on `SERVER_ADDR`. The checked-in address is `127.0.0.1:8082`. Relative paths mean starting the binary from the repository root will not find the expected settings and predictor paths.

## Routes

| Method | Route | Behavior |
| --- | --- | --- |
| `GET` | `/v1/site/settings/get` | Returns category option arrays from `settings/settings.json`. |
| `POST` | `/v1/costs/predict` | Converts request category codes to one-hot values, invokes Python, and returns its integer output. |
| `GET` | `/docs/swagger/index.html` | Serves the generated Swagger UI. |

All routes are unauthenticated. The server enables permissive CORS with `AllowAll`.

## Prediction handoff

The prediction service expects a local `python3` executable and `../predictor/predictor.py`. It sends a single comma-separated feature row on standard input. The adapter loads `../model/model.joblib`, predicts one value, and prints an integer. Ensure the backend process uses the expected working directory and that the model artifact and Python packages exist.

The request accepts an `insurance` integer, but the current Go feature construction does not copy it into `ThirdPartyInsuranceDue`. Inference therefore sends zero for that model feature. See the [API reference](../reference/api.md) for the full request contract.

## API definition

The generated Swagger files are in `backend/docs/`. The specification's base path is `/v1`; open the route above while the server is running. The service's error encoder returns JSON with an `error` field. Malformed JSON is reported as a bad request; internal and feature-conversion failures may be returned as server errors.
