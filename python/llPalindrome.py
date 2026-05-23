class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


def isPalindrome(head: ListNode | None) -> bool:
    return approachII(head)


# Time -> 0(N) Space -> 0(N)
def approachI(head: ListNode | None) -> bool:
    array = []

    while head is not None:
        array.append(head.val)
        head = head.next

    start, end = 0, len(array) - 1
    while start < end:
        if array[start] != array[end]:
            return False

        start += 1
        end -= 1

    return True


# Time -> 0(N) Space -> 0(1)
def approachII(head: ListNode | None) -> bool:
    count = 0
    curr = head

    while curr is not None:
        count += 1
        curr = curr.next

    mid = count // 2

    reverse = head
    while mid > 0:
        if reverse:  # This will always be true just type checking
            reverse = reverse.next
        mid -= 1

    if count % 2 == 1 and reverse:
        reverse = reverse.next

    endLL = None
    while reverse is not None:
        tmp = reverse.next
        reverse.next = endLL
        endLL = reverse
        reverse = tmp

    while endLL is not None and head:
        if endLL.val != head.val:
            return False

        head = head.next
        endLL = endLL.next

    return True
