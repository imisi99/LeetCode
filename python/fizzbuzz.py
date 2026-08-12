# Time -> 0(N) Space -> 0(N)
def fizzBuzz(n: int) -> list[str]:
    result: list[str] = []
    idx = 1
    while idx <= n:
        fizz = idx % 3
        buzz = idx % 5

        match fizz:
            case 0 if buzz == 0:
                result.append("FizzBuzz")
            case 0:
                result.append("Fizz")
            case _ if buzz == 0:
                result.append("Buzz")
            case _:
                result.append(str(idx))

        idx += 1

    return result


print(fizzBuzz(3))
print(fizzBuzz(5))
print(fizzBuzz(15))
