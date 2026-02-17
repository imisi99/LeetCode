# knapsack -> time 0(nlogn) space -> 0(n)
def knapsack(knapsack: int, items: list[list[int]]):
    totalValue = 0
    ratios = []

    for i, item in enumerate(items):
        weight, value = item[0], item[1]
        ratio = value / weight
        ratios.append([i, ratio])

    ratios.sort(key=lambda x: x[1], reverse=True)

    print(ratios)

    weight = knapsack

    for val in ratios:
        currWeight, currVal = items[val[0]][0], items[val[0]][1]
        if currWeight <= weight:
            weight -= currWeight
            totalValue += currVal
        else:
            frac = weight / currWeight
            totalValue += frac * currVal
            break

    return totalValue


print(knapsack(20, [[10, 70], [5, 50], [7, 20], [4, 50]]))
