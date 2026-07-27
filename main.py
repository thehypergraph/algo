from hash import HashTable, fnv1, fnv1a


def main():
    input = "Fire up the engines and let's go baby go!"
    print(fnv1(input))
    print(fnv1a(input))

    table = HashTable()
    table.insert("engine", 1)
    table.insert("thruster", 2)
    table["booster"] = 3

    print(table.get("engine"), table["thruster"], table["booster"])
    print("booster" in table, len(table))

    table.remove("booster")
    print("booster" in table, len(table))


if __name__ == "__main__":
    main()
