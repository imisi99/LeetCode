# Time -> 0(N) Space -> 0(N)
def insertIntervals(
    intervals: list[list[int]], newInterval: list[int]
) -> list[list[int]]:
    insertedIntervals: list[list[int]] = []

    for idx in range(len(intervals)):
        if intervals[idx][0] > newInterval[0]:
            insertedIntervals.append(newInterval)
            insertedIntervals.extend(intervals[idx : len(intervals)])
            break
        insertedIntervals.append(intervals[idx])

    if len(intervals) == len(insertedIntervals):
        insertedIntervals.append(newInterval)

    mergedIntervals: list[list[int]] = []

    i = 0
    while i < len(insertedIntervals):
        start = i
        i += 1
        while (
            i < len(insertedIntervals)
            and insertedIntervals[i][0] <= insertedIntervals[start][1]
        ):
            insertedIntervals[start][1] = max(
                insertedIntervals[start][1], insertedIntervals[i][1]
            )
            i += 1

        mergedIntervals.append(insertedIntervals[start])
    return mergedIntervals


print(insertIntervals([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]))
