class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def invertTree(root: TreeNode | None) -> TreeNode | None:
    return root


# Time -> 0(N) Space -> 0(log(N))
def recursiveSwap(root: TreeNode | None):
    if root is None:
        return

    root.left, root.right = root.right, root.left

    recursiveSwap(root.left)
    recursiveSwap(root.right)
