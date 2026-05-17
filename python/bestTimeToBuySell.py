# Time -> 0(N) Space -> 0(1)
def bestTime(stocks: list[int]) -> int:
    profit = 0
    n = len(stocks) - 1
    bestDay = stocks[n]

    while n >= 0:
        if bestDay - stocks[n] > profit:
            profit = bestDay - stocks[n]
        if stocks[n] > bestDay:
            bestDay = stocks[n]

        n -= 1

    return profit


print(bestTime([7, 1, 5, 3, 6, 4]))
