# Python program to read data from a PostgreSQL table and load into a pandas DataFrame.
import pandas as pds
from sqlalchemy import create_engine

# Create an engine instance.
alchemyEngine = create_engine('postgresql+psycopg2://postgres:1234@127.0.0.1:5440/carify_db', pool_recycle=3600);

# Connect to PostgreSQL server.
dbConnection = alchemyEngine.connect();

# Read data from PostgreSQL database table and load into a DataFrame instance.
dataFrame = pds.read_sql("select * from \"pcars\"", dbConnection);

pds.set_option('display.expand_frame_repr', False);

dataFrame.to_csv("pcars_data.csv")

dbConnection.close();
