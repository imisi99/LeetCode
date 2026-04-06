class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def removeNthFromEnd(head, n: int) -> ListNode | None:
    current = head
    k = 0
    while current is not None:
        current = current.next
        k += 1

    if k == n:
        return head.next

    stop = k - n - 1

    current = head
    while stop != 0:
        current = current.next
        stop -= 1

    current.next = current.next.next

    return head
