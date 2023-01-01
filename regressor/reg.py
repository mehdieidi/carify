import pandas as pd

df = pd.read_csv("pcars_data.csv")

df.info()

df.head()

df.isna().sum()

x = df.drop(labels="cash_cost", axis=1).iloc[:,2:-1]
y = df['cash_cost']

x

y

from sklearn.tree import DecisionTreeRegressor
from sklearn.ensemble import RandomForestRegressor
from sklearn.metrics import mean_squared_error
from sklearn.metrics import mean_absolute_error
from sklearn.metrics import r2_score

from sklearn.model_selection import train_test_split

X_train, X_test, y_train, y_test = train_test_split(x,y , random_state=4,test_size=0.2, shuffle=True)

regressor = RandomForestRegressor(criterion='absolute_error',n_estimators = 200, max_depth=10)
regressor.fit(X_train,y_train)
y_pred = regressor.predict(X_test)
print(mean_absolute_error(y_test, y_pred))
# cross_val_score(regressor, X, y, cv=10)

y_test.iloc[25]

X_test.iloc[25]

print(regressor.predict([X_test.iloc[25]]))

"""# Model import / export section"""

from joblib import dump, load

# exporting the model

dump(regressor, './filename.joblib')

# importing the model
clf = load('./filename.joblib')

# getting input - separating with commas - predicting with loaded model - printing the output
a = input()
a = a.split(',')
output = clf.predict([a])
print(output)



















pip install --use-deprecated=legacy-resolver pycaret[full]

!pip install numpy==1.20

from pycaret.regression import *

reg_experiment = setup(df.drop(labels='id',axis=1),
                       target = 'cash_cost',train_size=0.8,imputation_type="iterative")

best = compare_models(cross_validation=True)
predictions = predict_model(best,data=df)

df.head()