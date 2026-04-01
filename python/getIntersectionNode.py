class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(M + N) Space -> 0(max(M, N))
def getIntersectionNode(headA, headB) -> ListNode | None:
    nodes = {}

    while headA is not None:
        nodes[headA] = True
        headA = headA.next

    while headB is not None:
        if headB in nodes:
            return nodes[headB]

        headB = headB.next

    return None


def getIntersectionNodeI(headA, headB) -> ListNode | None:
    a, b = headA, headB
    while a != b:
        if a is None:
            a = headB
        else:
            a = a.next

        if b is None:
            b = headA
        else:
            b = b.next

    return a
