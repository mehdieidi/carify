# Machine learning workflow

The Python scripts train scikit-learn regressors to estimate `cash_cost` from the one-hot encoded `pcars` dataset. The scripts are experiments rather than a packaged training application: dependency versions, model metadata, and repeatable evaluation settings are not recorded in the repository.

## Data preparation

1. Populate and preprocess PostgreSQL `cars` using the [data pipeline](services/data-pipeline.md).
2. Run one-hot encoding to populate `pcars`.
3. Run `python3 to_csv.py` from `pycsv/`. It exports the full table to `pycsv/pcars_data.csv`.

The exporter currently has a PostgreSQL connection string embedded in its source. Configure that connection before running it. The generated CSV is ignored by Git.

## Train

Install the script imports in a Python 3 environment: `pandas`, `scikit-learn`, and `joblib`. Then run one of these from `regressor/`:

```sh
python3 linear_regression.py
python3 decision_tree.py
```

Both scripts read `../pycsv/pcars_data.csv`, split rows randomly into train and test sets, print mean absolute error, fit the model, and write `../model/model.joblib`. Create the `model/` directory first. The scripts write to the same destination, so running the second overwrites the first model.

`linear_regression_notebook.ipynb` contains the notebook-based exploration. `plot/linear_regression.py` also compares linear regression and a decision tree and displays plots; it expects the exported CSV and the plotting dependencies.

## Inference contract

`predictor/predictor.py` reads a comma-separated line from standard input, loads `../model/model.joblib`, predicts one value, and prints the integer result. The Go backend builds a fixed-length vector of one-hot values plus year and usage, then calls this script.

Training selects features by CSV column position after dropping `cash_cost`, while backend inference constructs a hard-coded field order. Keep these orders aligned whenever the table schema, exporter, or model features change. The model file itself does not include the feature names or preprocessing rules.

The backend currently leaves the insurance-due feature at zero regardless of the request's `insurance` field. Reproducible deployment also requires independently supplying a model artifact compatible with the checked-in feature order; no model is tracked in Git.

## Interpreting results

The training scripts report mean absolute error for a random holdout split, but do not set a random seed or save evaluation metadata. The plots in `images/` are historical artifacts; they do not establish current model quality. Treat output as an experimental estimate, not a guaranteed market valuation.
