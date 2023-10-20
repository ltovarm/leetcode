#! /usr/bin/env python3
import pandas as pd


def combine_two_tables(person: pd.DataFrame, address: pd.DataFrame) -> pd.DataFrame:

    return person.merge(address, on='personId', how='left').drop('personId', axis=1).drop('addressId', axis=1)


if __name__ == "__main__":

    person = pd.read_json('person.json')
    address = pd.read_json('address.json')

    print(combine_two_tables(person=person, address=address))
