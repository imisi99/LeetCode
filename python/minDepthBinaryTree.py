class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(log(N))
def minDepth(root) -> int:
    if root is None:
        return 0
    return findDepth(root, 1)


def findDepth(root, depth) -> int:
    if root.left is None and root.right is None:
        return depth

    newDepth = depth + 1
    leftDepth = findDepth(root.left, newDepth) if root.left is not None else -1
    rightDepth = findDepth(root.right, newDepth) if root.right is not None else -1

    if leftDepth == -1 or rightDepth == -1:
        if leftDepth == -1:
            return rightDepth
        return leftDepth
    return min(leftDepth, rightDepth)


right = TreeNode(val=20)
left = TreeNode(val=9)
root = TreeNode(val=3, left=left, right=right)
print(minDepth(root))
