class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def binaryTreePaths(root: TreeNode | None) -> list[str]:
    result = []
    recurse(root, "", result)
    return result


# Time -> NlogN Space -> NlogN
def recurse(root: TreeNode | None, path: str, result: list[str]):
    if root is None:
        return
    if root.left is None and root.right is None:
        result.append(path + str(root.val))
        return

    path += str(root.val) + "->"
    recurse(root.left, path, result)
    recurse(root.right, path, result)
