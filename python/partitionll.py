class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(1)
def partition(head: ListNode | None, x: int) -> ListNode | None:
    curr = head
    lessHead, less, prev, pivot = None, None, None, None
    while curr is not None:
        if curr.val >= x:
            pivot = curr
            currPrev = prev
            while curr is not None:
                if curr.val < x:
                    if less is None:
                        less = curr
                        lessHead = less
                    else:
                        less.next = curr
                        less = less.next

                    if currPrev is not None:
                        currPrev.next = curr.next

                else:
                    currPrev = curr
                curr = curr.next
            break

        prev = curr
        curr = curr.next

    if less is not None:
        if prev is None:
            less.next = head
            return lessHead

        prev.next = lessHead
        less.next = pivot

    return head


head = ListNode()
currHead = head
for val in [2, 3, 4, 3, 1, 2, 5]:
    currHead.val = val
    currHead.next = ListNode()
    currHead = currHead.next

res = partition(head, 3)

array = []
while res is not None:
    array.append(res.val)
    res = res.next

print(array)
