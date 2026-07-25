# Time -> 0(N^2) Space -> 0(1)
def maxProductI(n: int) -> int:
    str_n = str(n)
    max_product = 0
    for idx in range(len(str_n)):
        for j in range(len(str_n)):
            if idx == j:
                continue
            prod = int(str_n[idx]) * int(str_n[j])
            if prod > max_product:
                max_product = prod

    return max_product


# Time -> 0(NlogN) Space -> 0(N)
def maxProductII(n: int) -> int:
    str_n = str(n)
    sorted_n = sorted(str_n)
    prod = int(sorted_n[len(sorted_n) - 1]) * int(sorted_n[len(sorted_n) - 2])
    return prod


print(maxProductI(31))
print(maxProductII(31))

print(maxProductI(421))
print(maxProductII(421))
