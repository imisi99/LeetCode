class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N)
def levelOrder(root: TreeNode | None) -> list[list[int]]:
    result = []
    activeQueue = []
    activeQueue.append([root])
    while len(activeQueue) > 0:
        currQueue = activeQueue.pop(0)
        nextQueue = []
        currResult = []
        while len(currQueue) > 0:
            node = currQueue.pop(0)

            if node is None:
                continue

            currResult.append(node.val)

            nextQueue.append(node.left)
            nextQueue.append(node.right)

        if len(nextQueue) > 0:
            activeQueue.append(nextQueue)

        if len(currResult) > 0:
            result.append(currResult)

    i, j = 0, len(result) - 1
    while i < j:
        result[i], result[j] = result[j], result[i]
        i += 1
        j -= 1

    return result


left = TreeNode(val=1)
right = TreeNode(val=3)
root = TreeNode(2, left, right)
print(levelOrder(root))
