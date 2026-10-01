# Time -> 0(N) Space -> 0(N)
def isValid(s: str) -> bool:
    stack = []

    for val in s:
        match val:
            case "(":
                stack.append("(")
            case "{":
                stack.append("{")
            case "[":
                stack.append("[")
            case ")":
                popVal = popStack(stack)
                if popVal != "(":
                    return False
                stack.pop()
            case "}":
                popVal = popStack(stack)
                if popVal != "{":
                    return False
                stack.pop()
            case "]":
                popVal = popStack(stack)
                if popVal != "[":
                    return False
                stack.pop()

    return len(stack) == 0


def popStack(stack: list[str]) -> str:
    if len(stack) == 0:
        return "nil"
    return stack[-1]


print(isValid("()"))
