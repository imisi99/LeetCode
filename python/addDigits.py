# Time -> 0(N^2) Space -> 0(N)
def addDigits(num: int) -> int:
    while len(str(num)) != 1:
        run_sum = 0
        val = str(num)
        for v in val:
            run_sum += int(v)
        num = run_sum

    return num


print(addDigits(121))
