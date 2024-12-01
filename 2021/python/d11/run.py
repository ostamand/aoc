from argparse import ArgumentParser
from typing import Text

import numpy as np


def get_data(file_path: Text):
    data = []
    with open(file_path, "r") as f:
        for text in f:
            line = []
            for n in text.strip():
                line.append(int(n))
            data.append(line)
    return np.array(data)


def simulate(file_path: Text, n_steps: int):
    data = get_data(file_path)
    n, m = data.shape

    def flash(i,j):
        data[i,j] = 0
        cond = (lines >= i - 1) & (lines <= i+1) & (columns >= j-1) & (columns <= j+1 ) & (data != 0)
        data[cond] += 1

    flashes_sync = []
    n_flash = 0
    lines = np.array([ [i for _ in range(m)] for i in range(n)])
    columns = np.array([ [i for i in range(m)] for _ in range(n)])
    for step_n in range(n_steps):
        data = data + 1
        cond = data > 9
        flag = np.sum(cond.astype(np.int32)) > 0
        while flag:
            i = lines[cond][0]
            j = columns[cond][0]

            flash(i,j)
            n_flash+=1

            # check if we should continue
            cond = (data != 0) & (data > 9)
            if np.sum(cond.astype(np.int32)) == 0:
                flag = False

        # check for sync flash
        if (data == 0).all():
            flashes_sync.append(step_n+1)

    return n_flash, flashes_sync


if __name__ == "__main__":
    parser = ArgumentParser()
    parser.add_argument("-p", type=int, default=1)
    parser.add_argument("-data", type=str, default="2021/inputs/day11.txt")
    args = parser.parse_args()

    n_steps = 100 if args.p == 1 else 300

    n_flash, flashes_sync = simulate(args.data, n_steps)

    if args.p == 1:
        print(f"{n_flash} flashes after step {n_steps}")

    if args.p == 2:
        print(f"First step during which all octopuses flash: {flashes_sync[0]}")
