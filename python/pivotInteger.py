# Time -> 0(N) Space -> 0(1)
def pivotInteger(n: int) -> int:
    totSum = 0
    for idx in range(1, n + 1):
        totSum += idx

    runSum = 0
    for idx in range(1, n + 1):
        runSum += idx
        if runSum == totSum:
            return idx
        totSum -= idx

    return -1


print(pivotInteger(4))
