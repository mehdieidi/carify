# API reference

The Go backend exposes JSON routes under `/v1`. With the checked-in backend configuration, the base URL is `http://127.0.0.1:8082`. The generated Swagger UI is available at `/docs/swagger/index.html` while the backend is running.

## `GET /v1/site/settings/get`

Returns option arrays read from `backend/settings/settings.json`.

```json
{
  "data": {
    "body": [{ "name": "بدنه سالم", "value": 1 }],
    "front_chassis_status": [{ "name": "سالم", "value": 1 }],
    "rear_chassis_status": [{ "name": "سالم", "value": 1 }],
    "gearbox": [{ "name": "دنده ای", "value": 1 }],
    "motor_status": [{ "name": "سالم", "value": 1 }],
    "color": [{ "name": "آبی", "value": 1 }]
  }
}
```

The example abbreviates each array; the live response contains all configured values.

## `POST /v1/costs/predict`

Accepts JSON with integer fields:

| Field | Meaning |
| --- | --- |
| `year` | Vehicle model year, in the same calendar convention as the training data. |
| `color_value` | `value` from the settings `color` array. |
| `usage_km` | Odometer reading in kilometers. |
| `body_status_value` | `value` from `body`. |
| `motor_status_value` | `value` from `motor_status`. |
| `rear_chassis_status_value` | `value` from `rear_chassis_status`. |
| `front_chassis_status_value` | `value` from `front_chassis_status`. |
| `insurance` | Remaining third-party insurance months in the current UI. Accepted by the request struct, but currently not copied into the model feature vector. |
| `gearbox` | `value` from `gearbox`. |

Example:

```sh
curl -X POST http://127.0.0.1:8082/v1/costs/predict \
  -H "Content-Type: application/json" \
  -d '{
    "year": 1394,
    "color_value": 20,
    "usage_km": 85000,
    "body_status_value": 1,
    "motor_status_value": 1,
    "rear_chassis_status_value": 1,
    "front_chassis_status_value": 1,
    "insurance": 6,
    "gearbox": 1
  }'
```

Response shape:

```json
{"data": 123456789}
```

The actual number depends on the local model artifact and training data. The API does not currently validate numeric ranges before converting category codes. Malformed JSON returns HTTP 400; prediction/runtime errors are generally mapped to HTTP 500.

## CORS and authentication

The backend enables all CORS origins and does not implement authentication. Do not expose it to untrusted networks without adding appropriate access controls.
