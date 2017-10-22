import pandas as pd
import numpy as np
import matplotlib.pyplot as plt


df = pd.read_csv("annealing_softmax.csv")

df["ratio"] = [ v / (k + 1) for k,v in enumerate(df["cumulative_rewards"])]

df.plot()
plt.show()

