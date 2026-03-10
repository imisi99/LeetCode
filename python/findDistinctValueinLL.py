class LinkedList:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(N)
def findDistinct(head):
    set = {}
    tmpHead = head
    newHead = None
    actualHead = None

    while head is not None:
        if head.val in set.keys():
            set[head.val] += 1
        else:
            set[head.val] = 1
        head = head.next

    head = tmpHead
    while head is not None:
        tmpVal = head.val
        if set[tmpVal] > 1:
            while head is not None and head.val == tmpVal:
                head = head.next
        else:
            if newHead is None:
                newHead = head
                actualHead = newHead
                head = head.next
            else:
                newHead.next = head
                head = head.next
                newHead = newHead.next
            newHead.next = None

    return actualHead
