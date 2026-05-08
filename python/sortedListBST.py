class ListNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(N)
def sortedListToBST(head: ListNode | None) -> TreeNode | None:
    array = []
    while head is not None:
        array.append(head.val)
        head = head.next

    return createBST(array, 0, len(array) - 1)


def createBST(array: list[int], start: int, stop: int) -> TreeNode | None:
    if start > stop:
        return None

    mid = (start + stop) // 2

    node = TreeNode(val=array[mid])
    node.left = createBST(array, start, mid - 1)
    node.right = createBST(array, mid + 1, stop)

    return node
