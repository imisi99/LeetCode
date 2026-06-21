class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(logN)
def rightSideView(root: TreeNode | None) -> list[int]:
    visible = []
    queue = []
    queue.append([root])

    while len(queue) > 0:
        curr = queue.pop(0)
        next = []

        while len(curr) > 0:
            node = curr.pop(0)

            if node.left is not None:
                next.append(node.left)

            if node.right is not None:
                next.append(node.right)

        if len(next) > 0:
            queue.append(next)
            visible.append(next[len(next) - 1].val)

    return visible
