class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def deleteMiddle(head: ListNode | None) -> ListNode | None:
    count = 0
    headI, headII = head, head

    while head is not None:
        count += 1
        head = head.next

    mid = count // 2
    prev = None
    while mid != 0 and headI is not None:
        prev = headI
        headI = headI.next
        mid -= 1

    if prev and headI:
        prev.next = headI.next

    return headII
