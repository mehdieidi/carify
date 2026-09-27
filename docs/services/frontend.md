# Frontend service

The user interface is a React 18 application in `frontend/`, bootstrapped with Create React App. Its main screen is implemented in `src/components/Home.js`.

## Run locally

From `frontend/`:

```sh
npm ci
npm start
```

The development server opens at `http://localhost:3000`. Available npm scripts are `start`, `build`, `test`, and `eject` (ejecting is an irreversible Create React App operation).

## Current behavior

- Loads color, body, motor, chassis, and gearbox options from `GET /v1/site/settings/get`.
- Offers a fixed vehicle choice and fixed year and insurance-month choices in the React source.
- Submits the selected values to `POST /v1/costs/predict` and displays the returned number.
- Uses the API origin `http://45.67.85.169:8080` directly in source. This is not configurable through the package scripts or an environment variable. For local development, change the origin in `src/components/Home.js` to the backend address.

The insurance selection is sent in the request, but the current backend does not include it in the model input. The generated Create React App starter test also does not match the current Carify page.
