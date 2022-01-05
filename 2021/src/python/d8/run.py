from argparse import ArgumentParser


if __name__ == "__main__":
    parser = ArgumentParser()
    parser.add_argument("-p", type=int, default=1)
    parser.add_argument("-data", type=str, default="2021/inputs/day8.txt")
    args = parser.parse_args()

    segments = [2, 3, 4, 7] # digit 1, 4, 7, 8 VS number of segments

    n_digits = 0
    with open(args.data, "r") as f:
        for line in f:
            n_digits += sum([int(len(s) in segments) for s in line.split("|")[1].lstrip().replace("\n", "").split(" ")])
    
    print(f"Times digits 1, 4, 7, or 8 appear: {n_digits}")