class BinaryNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


# Time -> 0(N) Space -> 0(logN)
def countNodes(root: BinaryNode | None) -> int:
    count = [0]
    recurse(root, count)
    return count[0]


def recurse(root: BinaryNode | None, count: list[int]) -> tuple[int, int]:
    if root is None:
        return 0, 0

    leftSum, leftSize = recurse(root.left, count)
    rightSum, rightSize = recurse(root.right, count)

    currSum = leftSum + rightSum + root.val
    currSize = leftSize + rightSize + 1

    if (currSum // currSize) == root.val:
        count[0] += 1

    return currSum, currSize


leaf1 = BinaryNode(val=0, left=None, right=None)
leaf2 = BinaryNode(val=1, left=None, right=None)
leaf3 = BinaryNode(val=6, left=None, right=None)
mid1 = BinaryNode(val=8, left=leaf1, right=leaf2)
mid2 = BinaryNode(val=5, left=None, right=leaf3)
root = BinaryNode(val=4, left=mid1, right=mid2)

print(countNodes(root))
