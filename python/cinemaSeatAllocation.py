# Time -> 0(N) Space -> 0(N) where N is the no of reserved seats
def maxNumberOfFamilies(n: int, reservedSeats: list[list[int]]) -> int:
    array = [[] for _ in range(n)]

    for pos in reservedSeats:
        array[pos[0] - 1].append(pos[1])

    count = 0
    for pos in array:
        block1 = True
        block2 = True
        block3 = True

        for idx in pos:
            if idx >= 2 and idx <= 5:
                block1 = False
            if idx >= 4 and idx <= 7:
                block2 = False
            if idx >= 6 and idx <= 9:
                block3 = False

        if block1 or block2 or block3:
            if block1 and block3:
                count += 2
            else:
                count += 1

    return count


print(maxNumberOfFamilies(4, [[4, 3], [1, 4], [4, 6], [1, 7]]))
