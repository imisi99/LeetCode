class LinkedList:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def removeDuplicates(head) -> LinkedList:
    stillHead = head
    while head is not None:
        currHead = head
        head = head.next
        while head is not None and currHead.val == head.val:
            head = head.next
        currHead.next = head
    return stillHead
