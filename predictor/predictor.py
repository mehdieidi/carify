from joblib import load

clf = load('../model/model.joblib') 

inputFields = input()
inputFields = inputFields.split(',')

output = clf.predict([inputFields])

print(int(output[0]))
