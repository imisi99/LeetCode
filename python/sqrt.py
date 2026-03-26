# Time -> 0(log(N)) Space 0(1)
def mySqrt(x: int) -> int:
    start, end = 0, x
    best = 0
    while start <= end:
        mid = (start + end) // 2
        midSquare = mid * mid
        if midSquare == x:
            return mid
        elif midSquare > x:
            end = mid - 1
        else:
            start = mid + 1

        if midSquare < x:
            best = mid

    return best


print(mySqrt(8))
