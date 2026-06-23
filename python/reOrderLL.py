class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def reOrderLL(head: ListNode | None) -> ListNode | None:
    count = 0
    tmp = head
    while tmp is not None:
        count += 1
        tmp = tmp.next

    mid = count // 2
    reverse = head
    prev = None
    while mid >= 0 and reverse is not None:
        prev = reverse
        reverse = reverse.next
        mid -= 1

    if prev:
        prev.next = None

    tail = None
    while reverse is not None:
        tmp = reverse.next
        reverse.next = tail
        tail = reverse
        reverse = tmp

    curr = head
    while tail is not None and head is not None:
        headTmp = head.next
        tailTmp = tail.next

        head.next = tail
        tail.next = headTmp
        head = headTmp
        tail = tailTmp

    return curr


n0 = ListNode(val=4)
n1 = ListNode(val=3, next=n0)
n2 = ListNode(val=2, next=n1)
n3 = ListNode(val=1, next=n2)

val = reOrderLL(n3)

while val is not None:
    print(val.val)
