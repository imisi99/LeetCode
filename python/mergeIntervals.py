# Time -> 0(NlogN) Space -> 0(N)
def merge(intervals: list[list[int]]) -> list[list[int]]:
    merged = []

    intervals.sort()

    idx = 0
    while idx < len(intervals):
        start = idx
        idx += 1

        while idx < len(intervals) and intervals[start][1] >= intervals[idx][0]:
            intervals[start][1] = max(intervals[start][1], intervals[idx][1])
            idx += 1

        merged.append(intervals[start])

    return merged


print(merge([[4, 7], [1, 4]]))
