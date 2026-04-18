class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def sortedArrayToBST(nums: list[int]) -> TreeNode | None:
    return addBinaryNode(nums, 0, len(nums) - 1)


# Time -> 0(N) Space -> 0(log(N))
def addBinaryNode(nums: list[int], start: int, stop: int) -> TreeNode | None:
    if start > stop:
        return None

    mid = (start + stop) // 2

    node = TreeNode(val=nums[mid])

    node.left = addBinaryNode(nums, start, mid - 1)
    node.right = addBinaryNode(nums, mid + 1, stop)

    return node


def viewTreeNodeInOrder(root: TreeNode | None):
    if root is None:
        return

    viewTreeNodeInOrder(root.left)
    print(root.val)
    viewTreeNodeInOrder(root.right)


root = sortedArrayToBST([-4, -1, 0, 1, 4])
viewTreeNodeInOrder(root)
