class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def swapPairs(head) -> ListNode:
    swap = head
    prev = None
    if head is not None and head.next is not None:
        swap = head.next

    while head is not None and head.next is not None:
        next = head.next.next
        tmpHead = head
        head = head.next
        if prev is not None:
            prev.next = head
        head.next = tmpHead
        head.next.next = next
        prev = head.next
        head = next

    return swap
