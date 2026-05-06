class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N)
def levelOrder(root: TreeNode | None) -> list[list[int]]:
    result = []
    activeQueue = [[root]]

    while len(activeQueue) > 0:
        node = activeQueue[0]
        activeQueue = activeQueue[1:]
        currLevel = []
        nextLevel = []

        while len(node) > 0:
            currNode = node[0]
            node = node[1:]

            if currNode is None:
                continue

            currLevel.append(currNode.val)

            nextLevel.append(currNode.left)
            nextLevel.append(currNode.right)

        if len(nextLevel) > 0:
            activeQueue.append(nextLevel)

        if len(currLevel) > 0:
            result.append(currLevel)

    return result


left = TreeNode(val=1)
right = TreeNode(val=3)
root = TreeNode(2, left, right)
print(levelOrder(root))
