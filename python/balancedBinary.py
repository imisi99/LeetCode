class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(log(N))
def isBalanced(root) -> bool:
    _, balanced = check(root)
    return balanced


def check(root) -> tuple[int, bool]:
    if root is None:
        return 0, True

    height1, balance1 = check(root.left)
    height2, balance2 = check(root.right)

    balanced = balance1 and balance2

    diff = height1 - height2
    if diff < -1 or diff > 1:
        balanced = False
    else:
        balanced = balanced and True

    height = max(height1, height2) + 1
    return height, balanced
