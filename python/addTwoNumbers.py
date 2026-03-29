class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(max(M, N)) Space -> 0(1)
def addTwoNumbers(l1, l2) -> ListNode:
    head = l1
    prev = l1
    carry: int = 0
    while l1 is not None and l2 is not None:
        sum = l1.val + l2.val + carry
        val = sum % 10
        carry = sum // 10
        l1.val = val
        prev = l1
        l1 = l1.next
        l2 = l2.next

    if l1 is not None or l2 is not None:
        if l1 is None:
            prev.next = l2
            l1 = prev.next

        while carry != 0 and l1 is not None:
            sum = l1.val + carry
            val = sum % 10
            carry = sum // 10
            l1.val = val
            prev = l1
            l1 = l1.next

    if carry != 0:
        prev.next = ListNode(val=carry)
    return head
