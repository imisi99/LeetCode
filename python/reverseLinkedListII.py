class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def reverseBetween(head, left: int, right: int):
    if left == right:
        return head
    curr = head
    prev = None
    diff = left

    while left > 1:
        prev = curr
        curr = curr.next
        left -= 1

    reversed = None
    endNode = None
    while diff <= right:
        next = curr.next
        curr.next = reversed
        if endNode is None:
            endNode = curr
        reversed = curr
        curr = next
        diff += 1

    if endNode is not None:
        endNode.next = curr

    if prev is not None:
        prev.next = reversed
    else:
        return reversed

    return head
