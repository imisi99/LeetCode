class Node:
    def __init__(self, val=0, next=None, random=None) -> None:
        self.val = val
        self.next = next
        self.random = random


# Time -> 0(N) Space -> 0(N)
def copyNode(head: Node | None) -> Node | None:
    nodeMap = {}

    start = None
    newHead = None
    actualHead = head

    while head is not None:
        if start is None:
            start = Node()
            newHead = start

        start.val = head.val
        nodeMap[head] = start

        if head.next is not None:
            start.next = Node()

        head = head.next
        start = start.next

    head = newHead
    while actualHead is not None and newHead is not None:
        newHead.random = nodeMap.get(actualHead.random, None)
        actualHead = actualHead.next
        newHead = newHead.next

    return head
