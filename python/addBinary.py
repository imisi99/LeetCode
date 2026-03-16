# Time -> 0(max(M, N)) Space -> 0(max(M,N)+1)
def addBinary(a: str, b: str) -> str:
    i, j = len(a) - 1, len(b) - 1
    carry = 0
    binary = []
    while i >= 0 and j >= 0:
        val = int(a[i]) + int(b[j]) + carry
        if val == 3:
            binary.append("1")
            carry = 1
        elif val == 2:
            binary.append("0")
            carry = 1
        elif val == 1:
            binary.append("1")
            carry = 0
        else:
            binary.append("0")
            carry = 0
        i -= 1
        j -= 1

    while i >= 0:
        val = int(a[i]) + carry
        if val == 2:
            binary.append("0")
            carry = 1
        else:
            binary.append(f"{val}")
            carry = 0
        i -= 1

    while j >= 0:
        val = int(b[j]) + carry
        if val == 2:
            binary.append("0")
            carry = 1
        else:
            binary.append(f"{val}")
            carry = 0
        j -= 1

    if carry == 1:
        binary.append("1")
    return "".join(reversed(binary))


print(addBinary("11", "1"))
