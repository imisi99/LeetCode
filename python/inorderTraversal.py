class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(log(N))
def inorderTraversal(root: TreeNode) -> list[int]:
    return inorder(root, [])


def inorder(root, array: list[int]) -> list[int]:
    if root is None:
        return []

    inorder(root.left, array)
    array.append(root.val)
    inorder(root.right, array)

    return array


left = TreeNode(val=2)
right = TreeNode(val=3)
root = TreeNode(val=1, left=left, right=right)

print(inorderTraversal(root))
