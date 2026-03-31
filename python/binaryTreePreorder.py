class TreeNode:
    def __init__(self, val=0, next=None) -> None:
        self.val = val
        self.next = next


# Time -> 0(N) Space -> 0(N)
def preorderTraversal(root) -> list[int]:
    array = []
    preorder(root, array)
    return array


def preorder(root, array):
    if root is None:
        return

    array.append(root.val)
    preorder(root.left, array)
    preorder(root.right, array)


def preorderIterative(root) -> list[int]:
    if root is None:
        return []
    stack = [root]
    array = []
    while len(stack) != 0:
        root = stack.pop()
        array.append(root.val)
        if root.right is not None:
            stack.append(root.right)
        if root.left is not None:
            stack.append(root.left)
    return array
