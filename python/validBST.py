class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def isValidBST(root: TreeNode | None) -> bool:
    return checkBST(root, float("-inf"), float("inf"))


# Time -> 0(N) Space -> 0(logN)
def checkBST(root: TreeNode | None, min: float, max: float) -> bool:
    if root is None:
        return True

    if root.val <= min or root.val >= max:
        return False

    return checkBST(root.left, min, root.val) and checkBST(root.right, root.val, max)
