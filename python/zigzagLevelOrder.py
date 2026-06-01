class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def zigzagLevelOrder(root: TreeNode | None) -> list[list[int]]:
    if not root:
        return []

    queue = []
    result = []
    queue.append([root])
    dir = "left"
    while len(queue) != 0:
        activeQueue = queue.pop(0)
        currResult = []
        nextQueue = []
        for node in activeQueue:
            if node is None:
                continue

            currResult.append(node.val)

            nextQueue.append(node.left)
            nextQueue.append(node.right)

        if len(currResult) > 0:
            if dir == "right":
                i, j = 0, len(currResult) - 1
                while i < j:
                    currResult[i], currResult[j] = currResult[j], currResult[i]
                    i += 1
                    j -= 1

                dir = "left"
                result.append(currResult)
            else:
                dir = "right"
                result.append(currResult)

        if len(nextQueue) > 0:
            queue.append(nextQueue)

    return result


left = TreeNode(val=1)
right = TreeNode(val=3)
root = TreeNode(2, left, right)
print(zigzagLevelOrder(root))
