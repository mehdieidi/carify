import pandas as pd

df = pd.read_csv("../pycsv/pcars_data.csv")

# Remove the "cost" column because it is not a feature but a label.
# Also remove the "id" column and "token" column that are useless in model training.
x = df.drop(labels="cash_cost", axis=1).iloc[:,2:-1]
y = df['cash_cost']

from sklearn.model_selection import train_test_split

# random_state as the name suggests, is used for initializing the internal random number generator,
# which will decide the splitting of data into train and test indices in your case
X_train, X_test, y_train, y_test = train_test_split(x,y , random_state=4,test_size=0.2, shuffle=True)

from sklearn.ensemble import RandomForestRegressor
from sklearn.metrics import mean_absolute_error

# n_estimators is the number of the trees in the forrest. max_depth is set to prevent over growing of trees.
regressor = RandomForestRegressor(criterion='absolute_error',n_estimators = 200, max_depth=10)
regressor.fit(X_train.values,y_train)

y_pred = regressor.predict(X_test)
print(mean_absolute_error(y_test, y_pred))

from joblib import dump

dump(regressor, '../model/model.joblib') 
