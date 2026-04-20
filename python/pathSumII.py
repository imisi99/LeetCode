class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def pathSum(root: TreeNode, targetSum: int) -> list[list[int]]:
    result = []
    findPathSums(root, targetSum, [], result)
    return result


def findPathSums(
    root: TreeNode | None, targetSum: int, path: list[int], result: list[list[int]]
):
    if root is None:
        return

    newTargetSum = targetSum - root.val
    path.append(root.val)

    if newTargetSum == 0 and root.left is None and root.right is None:
        result.append(path[:])
        path.pop()
        return

    findPathSums(root.left, newTargetSum, path, result)
    findPathSums(root.right, newTargetSum, path, result)

    path.pop()


left = TreeNode(val=2)
a = TreeNode(val=-1)
right = TreeNode(val=3, left=a)
root = TreeNode(1, left, right)

print(pathSum(root, 3))
