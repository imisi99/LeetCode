# Time -> 0(N) Space -> 0(1)
def setBit(num: int) -> int:
    count = 0
    while num != 0:
        bit = num % 2
        if bit == 1:
            count += 1
        num = num // 2
    return count


print(setBit(7))
