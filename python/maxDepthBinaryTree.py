class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(log(N))
def maxDepth(root) -> int:
    return findDepth(root, 0)


def findDepth(root, depth) -> int:
    if root is None:
        return depth

    newDepth = depth + 1
    leftDepth = findDepth(root.left, newDepth)
    rightDepth = findDepth(root.right, newDepth)

    return max(leftDepth, rightDepth)


right = TreeNode(val=20)
left = TreeNode(val=9)
root = TreeNode(val=3, left=left, right=right)
print(maxDepth(root))
