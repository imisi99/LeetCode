# Time -> 0(N) Space -> 0(N)
def reverse(x: int) -> int:
    val = str(x)
    start = 0
    if val[0] == "-":
        start = 1
    end = len(val) - 1

    valStr = [s for s in val]

    while start < end:
        valStr[start], valStr[end] = valStr[end], valStr[start]
        start += 1
        end -= 1

    reversed = "".join(valStr)
    valInt = int(reversed)
    if valInt < pow(-2, 31) or valInt > pow(2, 31) - 1:
        return 0

    return valInt


print(reverse(-456))
