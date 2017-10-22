import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

# "index", "chosen_arm", "reward", "cumulative_rewards"}

# df = pd.read_csv("ucb1.csv")
df = pd.read_csv("annealing_softmax.csv")

df["ratio"] = [ v / (k + 1) for k,v in enumerate(df["cumulative_rewards"])]
df1 = df.drop("cumulative_rewards", axis=1) # axis 1 means column, not row
df2 = df1.drop("index", axis=1)
df3 = df2.drop("reward", axis=1)
df4 = df3.drop("chosen_arm", axis=1)

df4.plot()
plt.show()

