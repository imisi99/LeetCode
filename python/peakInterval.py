# Time -> 0(N) Space -> 0(1)
def peak(intervals: list[int]) -> int:
    start = float("-inf")
    for idx in range(len(intervals)):
        curr_val = intervals[idx]
        end = float("-inf") if idx + 1 == len(intervals) else intervals[idx + 1]

        if curr_val > start and curr_val > end:
            return idx

        start = curr_val
    return -1


print(peak([]))
