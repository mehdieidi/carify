import pandas as pd

df = pd.read_csv("../pycsv/pcars_data.csv")

x = df.drop(labels="cash_cost", axis=1).iloc[:,2:-1]
y = df['cash_cost']

from sklearn.model_selection import train_test_split

X_train, X_test, y_train, y_test = train_test_split(x,y , random_state=4,test_size=0.2, shuffle=True)

from sklearn.ensemble import RandomForestRegressor
from sklearn.metrics import mean_absolute_error

regressor = RandomForestRegressor(criterion='absolute_error',n_estimators = 200, max_depth=10)
regressor.fit(X_train.values,y_train)
y_pred = regressor.predict(X_test)
print(mean_absolute_error(y_test, y_pred))

from joblib import dump

dump(regressor, '../model/model.joblib') 
