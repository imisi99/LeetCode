# Time -> 0(N*M) Space -> 0(N*M)
def convert(s: str, numRows: int) -> str:
    if numRows == 1 or numRows >= len(s):
        return s

    array = [["" for _ in range(len(s))] for _ in range(numRows)]

    col = 0
    sid = 0
    while sid < len(s):
        for idx in range(numRows):
            if sid == len(s):
                break
            array[idx][col] = s[sid]
            sid += 1

        col += 1
        startRow = numRows - 2

        while startRow >= 1 and sid < len(s):
            array[startRow][col] = s[sid]
            sid += 1
            startRow -= 1
            col += 1

    result = []
    row = 0
    while row < len(array):
        col = 0
        while col < len(array[0]):
            if array[row][col] != "":
                result.append(array[row][col])
            col += 1
        row += 1

    return "".join(result)


# Time -> 0(N) Space -> 0(N)
def convertI(s: str, numRows: int) -> str:
    if numRows == 1 or numRows >= len(s):
        return s
    array = [""] * numRows

    row, direction = 0, -1
    for char in s:
        array[row] += char

        if row == 0 or row == numRows - 1:
            direction *= -1

        row += direction

    return "".join(array)


print(convert("PAYPALISHIRING", 4))
print(convertI("PAYPALISHIRING", 4))
