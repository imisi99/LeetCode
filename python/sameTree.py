class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def isSameTree(p, q) -> bool:
    return recurse(p, q)


def recurse(p, q) -> bool:
    if p is None and q is None:
        return True
    elif p is None or q is None:
        return False

    if p.val != q.val:
        return False

    return recurse(p.left, q.left) and recurse(p.right, q.right)
