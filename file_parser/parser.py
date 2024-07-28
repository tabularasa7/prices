#!/usr/bin/env python3

import csv
import os
import sys

def parse_csv():
    print(sys.argv)
    with open(os.path.dirname(__file__) + '/../84-1321373_centennial-healthcare-plaza_standardcharges.csv') as csv_file:    
        reader = csv.reader(csv_file)

        # update = next(reader)

        # baseHeader = next(reader)
        # print(update)
        # print(baseHeader)

        prices = []
        split_prices = []

        # prevRow = reader[0]
        for row in reader:
            if row[0] in (None, ""):
                prices.append(split_prices)
                split_prices = []
                # print(row)
                # print(next(reader))
                # print(next(reader))
            else:
                split_prices.append(row)
            # else:
            #     print(len(prices))
            #     newFile = open(os.path.dirname(__file__) + '/../84-1321373_centennial-healthcare-plaza_standardcharges' + prices[0] + '.csv', 'w+')
            #     newFile.writelines(prices)
            #     prices = []

        print(len(prices))
        print(prices[0][0])
        return prices


    # results.close()

parse_csv()