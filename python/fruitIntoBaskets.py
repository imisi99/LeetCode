# Time -> 0(N^2) Space -> 0(1)
def totalFruit(fruits: list[int]) -> int:
    totalSum = 0

    for i in range(len(fruits)):
        basket1 = -1
        basket2 = -1
        currSum = 0
        for j in range(i, len(fruits)):
            if basket1 == -1:
                basket1 = fruits[j]
                currSum += 1
            elif basket2 == -1:
                basket2 = fruits[j]
                currSum += 1
            elif fruits[j] == basket2 or fruits[j] == basket1:
                currSum += 1
            else:
                break

        if currSum > totalSum:
            totalSum = currSum

    return totalSum


print(totalFruit([1, 2, 3, 2, 2]))
print(totalFruit([0, 1, 2, 2]))
