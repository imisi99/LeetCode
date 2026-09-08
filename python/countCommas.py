# Time -> 0(N) Space -> 0(1)
def countCommas(n: int) -> int:
    num = n // 1000

    if num < 1:
        return 0

    count = ((num - 1) * 1000) + 1
    count += n % 1000

    return count


print(countCommas(998))
print(countCommas(1003))
print(countCommas(4003))
