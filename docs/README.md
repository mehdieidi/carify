# Carify documentation

This documentation describes the repository as it exists today. Start with the architecture overview, then use the setup and component guides for hands-on work.

## Guides

| Guide | What it covers |
| --- | --- |
| [Architecture overview](architecture/overview.md) | Components, data flow, storage tables, and boundaries. |
| [Development setup](development/setup.md) | Prerequisites, environment configuration, start commands, and current setup constraints. |
| [Backend service](services/backend.md) | HTTP server configuration, routes, and Python inference handoff. |
| [Data pipeline](services/data-pipeline.md) | Divar collection, PostgreSQL storage, preprocessing, and CSV export. |
| [Frontend service](services/frontend.md) | React app setup, API connection, and current UI behavior. |
| [Machine learning](machine-learning.md) | Dataset shape, training scripts, evaluation, and model artifact expectations. |
| [API reference](reference/api.md) | Routes, request fields, response examples, and value codes. |
| [Configuration reference](reference/configuration.md) | Environment variables and runtime configuration by component. |

## Repository map

```text
backend/       Go HTTP API and Swagger files
data/          Go listing collector, PostgreSQL access, migrations, preprocessing
frontend/      React web application
images/        Architecture and historical model plots
orchestrator/  Legacy pipeline runner
plot/          Python evaluation and plotting script
predictor/     Python inference adapter
pycsv/         PostgreSQL-to-CSV export script
regressor/     Python model training scripts and notebook
docs/          Project documentation
```

The Go projects each have their own `go.mod`; there is no root Go workspace. Python dependencies are not declared in a requirements file. Environment files are loaded relative to the process working directory, so run each component from its own directory unless a guide says otherwise.
