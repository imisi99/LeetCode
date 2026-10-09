def removeOuterParentheses(s: str) -> str:
    primitives = []
    count = 0
    start = 0

    for idx in range(len(s)):
        if s[idx] == "(":
            count += 1
        else:
            count -= 1

        if count == 0:
            primitives.append(s[start + 1 : idx])
            start = idx + 1

    return "".join(primitives)


print(removeOuterParentheses("(()())(())"))
print(removeOuterParentheses("(()())(())(()(()))"))
print(removeOuterParentheses("()()"))
