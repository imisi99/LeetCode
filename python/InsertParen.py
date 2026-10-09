# Time => 0(N) Space => 0(1)
def minInsertions(s: str) -> int:
    stack = 0
    waiting = False
    insertions = 0

    for paren in s:
        if paren == "(":
            if waiting:
                if stack == 0:
                    insertions += 1
                else:
                    stack -= 1
                insertions += 1
            stack += 1
            waiting = False
        else:
            if waiting:
                if stack == 0:
                    insertions += 1
                else:
                    stack -= 1
                waiting = False
            else:
                waiting = True

    if stack > 0:
        if waiting:
            insertions += 1
            stack -= 1
        insertions += stack * 2
    else:
        if waiting:
            insertions += 2

    return insertions


print(minInsertions("(()))"))
print(minInsertions("())"))
print(minInsertions("))())("))
