import pandas as pd
import matplotlib.pyplot as plt  
import seaborn as seabornInstance 

df = pd.read_csv("../pycsv/pcars_data.csv")

plt.figure(figsize=(15,10))
plt.tight_layout()
seabornInstance.distplot(df["cash_cost"])

# Remove the "cost" column because it is not a feature but a label.
# Also remove the "id" column and "token" column that are useless in model training.
x = df.drop(labels="cash_cost", axis=1).iloc[:,2:-1]
y = df['cash_cost']

from sklearn.model_selection import train_test_split

# random_state as the name suggests, is used for initializing the internal random number generator,
# which will decide the splitting of data into train and test indices in your case
X_train, X_test, y_train, y_test = train_test_split(x,y , test_size=0.2, shuffle=True)

from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_absolute_error

regressor = LinearRegression() 
regressor.fit(X_train.values,y_train)

y_pred = regressor.predict(X_test)
print(mean_absolute_error(y_test, y_pred))

from joblib import dump

dump(regressor, '../model/model.joblib') 

import numpy as np

y_test = np.array(list(y_test))
y_pred = np.array(y_pred)

df = pd.DataFrame({'Actual': y_test.flatten(), 'Predicted': y_pred.flatten()})
df1 = df.head(25)
df1.plot(kind='bar',figsize=(16,10))
plt.grid(which='major', linestyle='-', linewidth='0.5', color='green')
plt.grid(which='minor', linestyle=':', linewidth='0.5', color='black')
plt.show()

plt.scatter(X_test, y_test,  color='gray')
plt.plot(X_test, y_pred, color='red', linewidth=2)
plt.show()