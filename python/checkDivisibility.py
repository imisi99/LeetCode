# Time -> 0(N) Space -> 0(1)
def checkDivisibility(n: int) -> bool:
    val = str(n)

    prod, sum = 1, 0
    for s in val:
        num = int(s)

        prod *= num
        sum += num

    return n % (prod + sum) == 0


print(checkDivisibility(99))
print(checkDivisibility(23))
