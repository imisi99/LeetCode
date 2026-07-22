def ugly(num: int) -> bool:
    while num != 1:
        if num % 2 == 0:
            num = num // 2
        elif num % 3 == 0:
            num = num // 3
        elif num % 5 == 0:
            num = num // 5
        else:
            break

    return num == 1


print(ugly(5))
print(ugly(14))
print(ugly(1))
