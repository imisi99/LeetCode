class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(N)
def kthSmallest(root: TreeNode | None, k: int) -> int:
    array = []
    recurse(root, array)
    return array[k - 1]


def recurse(root: TreeNode | None, array: list[int]):
    if root is None:
        return

    recurse(root.left, array)
    array.append(root.val)
    recurse(root.right, array)


# Time -> 0 max(logN, k) Space -> 0(logN)
def kthSmallestI(root: TreeNode | None, k: int) -> int:
    stack: list[TreeNode] = []

    while root is not None:
        stack.append(root)
        root = root.left

    while k > 1:
        node = stack.pop()
        k -= 1

        if node.right is not None:
            stack.append(node.right)
            curr = node.right.left
            while curr is not None:
                stack.append(curr)
                curr = curr.left

    return stack[-1].val


lright = TreeNode(2)
left = TreeNode(1, None, lright)
right = TreeNode(4)
root = TreeNode(3, left, right)

print(kthSmallest(root, 3))
print(kthSmallestI(root, 3))
