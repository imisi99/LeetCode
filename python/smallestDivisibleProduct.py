# Time -> 0(1) Space -> 0(1)
def smallest(n: int, t: int) -> int:
    if n % 10 == 0:
        return n

    k = (n + 10) % 10
    end = k * 10

    while n < end:
        val = str(n)
        prod = 1
        for v in val:
            prod *= int(v)

        if prod % t == 0:
            return n
        n += 1

    return end


print(smallest(10, 3))
print(smallest(14, 3))
