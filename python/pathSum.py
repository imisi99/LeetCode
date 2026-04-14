from typing import Union


class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(logN)
def hasPathSum(root: Union[TreeNode, None], targetSum: int) -> bool:
    if root is None:
        return False

    newTarget = targetSum - root.val
    if newTarget == 0 and root.left is None and root.right is None:
        return True

    left = hasPathSum(root.left, newTarget)
    right = hasPathSum(root.right, newTarget)

    return left or right


left = TreeNode(val=2)
right = TreeNode(val=3)
root = TreeNode(1, left, right)

print(hasPathSum(root, 5))
