class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def flattenBinaryTree(root: TreeNode | None):
    if root is None:
        return
    convertLinkedList(root)


# Time -> 0(N) Space -> 0(log(N))
def convertLinkedList(root: TreeNode) -> TreeNode:
    if root.left is None and root.right is None:
        return root

    storedRight = None
    if root.left is not None:
        storedRight = root.right
        root.right, root.left = root.left, None
        root = root.right
        root = convertLinkedList(root)
        root.right = storedRight

    if root.right is not None:
        root = root.right
        root = convertLinkedList(root)

    return root


def printList(root: TreeNode | None):
    array = []
    while root is not None:
        array.append(root.val)
        root = root.right
    print(array)


a = TreeNode(val=4)
left = TreeNode(val=2)
right = TreeNode(val=3, right=a)
root = TreeNode(val=1, left=left, right=right)
flattenBinaryTree(root)
printList(root)
