# Time -> 0(N) Space -> 0(N)
def reverse(num: int) -> int:
    array = []
    while num != 0:
        bit = num % 2
        array.append(bit)
        num = num // 2

    while len(array) < 32:
        array.append(0)

    power, idx = 0, len(array) - 1
    val = 0
    while idx >= 0:
        base = 2**power
        val += base * array[idx]
        power += 1
        idx -= 1

    return val


print(reverse(43261596))
