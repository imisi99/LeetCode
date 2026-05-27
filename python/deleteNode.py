class ListNode:
    def __init__(self, val, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def deleteNode(node: ListNode):
    prev = None
    while node.next is not None:
        node.val = node.next.val
        prev = node
        node = node.next

    if prev:
        prev.next = None


two = ListNode(val=3)
one = ListNode(val=2, next=two)
head = ListNode(val=1, next=one)

deleteNode(head)

while head is not None:
    print(head.val)
    head = head.next
