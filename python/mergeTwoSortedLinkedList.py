class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


def mergeTwoLists(list1, list2) -> ListNode:
    if list1 is None:
        return list2
    if list2 is None:
        return list1

    head, tail = list1, list2
    if head.val > tail.val:
        head, tail = tail, head

    merged = head
    mergedHead = merged
    head = head.next

    while head is not None and tail is not None:
        if head.val <= tail.val:
            merged.next = head
            head = head.next
        else:
            merged.next = tail
            tail = tail.next
        merged = merged.next
        merged.next = None

    if head is None and tail is not None:
        merged.next = tail
    elif head is not None and tail is None:
        merged.next = head

    return mergedHead

