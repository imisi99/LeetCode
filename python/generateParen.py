def generateParenthesis(n: int) -> list[str]:
    array = []
    recurse(0, n, [], array)
    return array


def recurse(open: int, n: int, trace: list[str], array: list[str]):
    if open < 0 or open > n:
        return

    if len(trace) == (2 * n):
        if open == 0:
            array.append("".join(trace))
        return

    trace.append("(")
    recurse(open + 1, n, trace, array)
    trace.pop()

    trace.append(")")
    recurse(open - 1, n, trace, array)
    trace.pop()


print(generateParenthesis(1))
print(generateParenthesis(2))
print(generateParenthesis(3))
