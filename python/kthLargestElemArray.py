# Time -> 0(N) Space -> (N)  N is the max num in the array
def largest(array: list[int], k: int) -> int:
    max = 0
    for val in array:
        if val > max:
            max = val

    prefill = [0] * (max + 1)

    for val in array:
        prefill[val] += 1

    idx = max
    while k > 1:
        if prefill[idx] >= k:
            break

        k -= prefill[idx]
        idx -= 1

    return idx


print(largest([2, 3, 4, 5, 6, 7, 8, 9], 3))
