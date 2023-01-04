import pandas as pd
import matplotlib.pyplot as plt  
import seaborn as seabornInstance 
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_absolute_error, mean_squared_error
from joblib import dump
import numpy as np
from bokeh.plotting import figure, show, output_file

df = pd.read_csv("../pycsv/pcars_data.csv")
x = df.drop(labels="cash_cost", axis=1).iloc[:,2:-1]
y = df['cash_cost']
X_train, x_test, y_train, y_test = train_test_split(x,y , test_size=0.2, shuffle=True)
regressor = LinearRegression() 
regressor.fit(X_train.values,y_train)
y_pred = regressor.predict(x_test)
mae1 = mean_absolute_error(y_test, y_pred)
print(mae1)

plt.figure(figsize=(10,10))
plt.scatter(y_test, y_pred, c='crimson')
plt.yscale('log')
plt.xscale('log')

p1 = max(max(y_pred), max(y_test))
p2 = min(min(y_pred), min(y_test))
plt.plot([p1, p2], [p1, p2])
plt.xlabel('True Values', fontsize=15)
plt.ylabel('Predictions', fontsize=15)
plt.axis('equal')
plt.show()

from sklearn.tree import DecisionTreeRegressor 
regressor = DecisionTreeRegressor(criterion='absolute_error')
regressor.fit(X_train.values,y_train)
y_pred2 = regressor.predict(x_test)
mae2 = mean_absolute_error(y_test, y_pred2)
print(mae2)

y_test = np.array(y_test)
y_pred = np.array(y_pred)
y_pred2 = np.array(y_pred2)

df = pd.DataFrame({'Actual': y_test.flatten(), 'Linear Regression Predicted': y_pred.flatten(),  'Decision Tree Predicted': y_pred2.flatten()})
df1 = df.head(25)
df1.plot(kind='bar',figsize=(16,10))
plt.grid(linestyle='-', linewidth='0.5', color='green')
plt.grid( linestyle=':', linewidth='0.5', color='black')
plt.show()

