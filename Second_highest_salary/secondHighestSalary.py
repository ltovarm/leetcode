#! /usr/bin/env python3
import pandas as pd
import numpy as np


def second_highest_salary(employee: pd.DataFrame) -> pd.DataFrame:
    firstValue = employee['salary'].sort_values(ascending=False).iloc[0]
    for value in employee['salary'].sort_values(ascending=False).iloc[1:]:
        if firstValue == value:
            continue
        return pd.DataFrame({'SecondHighestSalary': [value]})
    return pd.DataFrame({'SecondHighestSalary': [np.nan]})


if __name__ == "__main__":
    employee = pd.read_json(example2.json')
    print(second_highest_salary(employee=employee))
