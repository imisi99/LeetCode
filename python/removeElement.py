class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def removeElement(head: ListNode | None, val: int) -> ListNode | None:
    curr = head
    prev = None
    start = None
    while curr is not None:
        if curr.val == val:
            if prev is not None:
                prev.next = curr.next
            curr = curr.next
        else:
            if start is None:
                start = curr
            prev = curr
            curr = curr.next

    return start
