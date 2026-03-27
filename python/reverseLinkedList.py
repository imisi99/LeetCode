class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def reverseList(head) -> ListNode | None:
    if head is None:
        return None
    prev = None
    while head.next is not None:
        tmp = head.next
        head.next = prev
        prev = head
        head = tmp

    head.next = prev
    return head
