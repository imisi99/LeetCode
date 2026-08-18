# Time -> 0(N^2) Space -> 0(N)
def bulbSwitch(n: int) -> int:
    array = [0] * n

    multiplier = 1
    for idx in range(len(array)):
        while idx < len(array):
            inverse(array, idx)
            idx += multiplier
        multiplier += 1

    counter = 0
    for val in array:
        if val == 1:
            counter += 1
    return counter


def inverse(array: list[int], idx: int):
    if array[idx] == 0:
        array[idx] = 1
    else:
        array[idx] = 0


print(bulbSwitch(8))
