def countAndSay(n: int) -> str:
    return recurse(n)


# Time -> 0(N*M) Space -> 0(N + M)
def recurse(n: int) -> str:
    if n == 1:
        return "1"

    rle = recurse(n - 1)

    idx = 0
    builder = []
    while idx < len(rle):
        currIdx = idx
        while idx < len(rle) and rle[currIdx] == rle[idx]:
            idx += 1

        count = f"{idx - currIdx}"
        builder.append(count)
        builder.append(rle[currIdx])

    return "".join(builder)


# Time -> 0(N*M) Space -> 0(M)
def approach1(n: int) -> str:
    rle = "1"
    i = 1
    while i < n:
        builder = []
        idx = 0
        while idx < len(rle):
            currIdx = idx
            while idx < len(rle) and rle[currIdx] == rle[idx]:
                idx += 1

            count = f"{idx - currIdx}"
            builder.append(count)
            builder.append(rle[currIdx])
        rle = "".join(builder)
        i += 1
    return rle


print(countAndSay(10))
print(approach1(10))
