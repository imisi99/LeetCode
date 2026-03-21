def isValid(s: str) -> bool:
    stack: list[str] = []
    for val in s:
        if val == "(":
            stack.append(val)
        elif val == "[":
            stack.append(val)
        elif val == "{":
            stack.append(val)
        else:
            pop = stack.pop() if len(stack) > 0 else ""
            if pop == "":
                return False
            elif pop == "(":
                if val != ")":
                    return False
            elif pop == "{":
                if val != "}":
                    return False
            else:
                if pop != "[":
                    return False

    return len(stack) == 0


print(isValid("[({})]"))
print(isValid("[({))]"))
