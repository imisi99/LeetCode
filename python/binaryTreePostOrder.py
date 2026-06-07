class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def postorderTraversal(root: TreeNode | None) -> list:
    array = []
    postorder(root, array)
    return array


def postorder(root: TreeNode | None, array: list):
    if root is None:
        return

    postorder(root.left, array)
    postorder(root.right, array)
    array.append(root.val)


def postorderTraversalIterative(root: TreeNode | None) -> list:
    if root is None:
        return []

    array = []
    stack = []
    stack.append(root)

    while len(stack) != 0:
        node = stack.pop()
        array.append(node.val)
        if node.left is not None:
            stack.append(node.left)
        if node.right is not None:
            stack.append(node.right)

    i, j = 0, len(array) - 1

    while i < j:
        array[i], array[j] = array[j], array[i]
        i, j = i + 1, j - 1

    return array


node2 = TreeNode(val=3)
node1 = TreeNode(val=2)
root = TreeNode(val=1, left=node1, right=node2)
print(postorderTraversal(root))
print(postorderTraversalIterative(root))
