def generate(numRows: int) -> list[list[int]]:
    result = []

    i = 1
    while i <= numRows:
        match i:
            case 1:
                result.append([1])
            case 2:
                result.append([1, 1])
            case _:
                prevRow = result[i - 2]
                newRow = [0] * i
                newRow[0], newRow[i - 1] = 1, 1

                idx = 1
                j = 0
                while j < len(prevRow) - 1:
                    newRow[idx] = prevRow[j] + prevRow[j + 1]
                    j += 1
                    idx += 1

                result.append(newRow)

        i += 1
    return result


print(generate(5))
