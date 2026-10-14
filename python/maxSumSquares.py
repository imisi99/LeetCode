# Time => 0(N) Space => 0(N)
def maxSumOfSquares(num: int, sum: int) -> str:
    nines = sum // 9
    noNine = sum % 9

    count = nines

    if noNine > 0:
        count += 1
    if count > num:
        return ""

    val = [""] * num

    i = 0
    for _ in range(nines):
        val[i] = "9"
        i += 1

    if i < len(val):
        val[i] = str(noNine)
        i += 1

        while i < len(val):
            val[i] = "0"
            i += 1

    return "".join(val)


print(maxSumOfSquares(2, 18))
