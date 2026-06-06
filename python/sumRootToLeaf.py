class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def sumNumbers(root: TreeNode | None) -> int:
    sumPaths = [0]
    traverse(root, "", sumPaths)
    return sumPaths[0]


def traverse(root: TreeNode | None, path: str, pathsSum: list[int]):
    if root is None:
        return

    path += str(root.val)
    if root.left is None and root.right is None:
        pathsSum[0] += int(path)
        return

    traverse(root.left, path, pathsSum)
    traverse(root.right, path, pathsSum)


left = TreeNode(val=2)
right = TreeNode(val=3)
root = TreeNode(val=1, left=left, right=right)
print(sumNumbers(root))
