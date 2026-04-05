from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N*Max(M)) Space -> 0(1)
def meregeKLists(lists: list[Optional[ListNode]]) -> ListNode | None:
    if len(lists) <= 1:
        if len(lists) == 0:
            return None
        return lists[0]

    mergedHead = lists[0]
    for llist in lists[1:]:
        merged = mergedHead
        mergedPrev = None
        while llist is not None and merged is not None:
            if llist.val < merged.val:
                listNext = llist.next
                if mergedPrev is not None:
                    mergedPrev = llist
                else:
                    mergedHead = llist
                mergedPrev = llist
                llist.next = merged
                llist = listNext
            else:
                mergedPrev = merged
                merged = merged.next

        if llist is not None and mergedPrev is not None:
            mergedPrev.next = llist

    return mergedHead
