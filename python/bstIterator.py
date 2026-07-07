class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class BSTIterator:
    def __init__(self, root: TreeNode | None):
        self.array = self.get_inorder(root)
        self.curr_idx = 0

    # Time -> 0(N) Space -> 0(N)
    def get_inorder(self, root: TreeNode | None, array: list = []) -> list:
        if root is None:
            return []
        
        self.get_inorder(root.left, array)
        array.append(root.val)
        self.get_inorder(root.right, array)

        return array
    
    # Time -> 0(1) Space -> 0(1)
    def next(self) -> int:
        self.curr_idx += 1
        return self.array[self.curr_idx-1]
    
    # Time -> 0(1) Space -> 0(1)
    def has_next(self) -> bool:
        return not self.curr_idx == len(self.array)


class BSTIteratorI:
    def __init__(self, root: TreeNode | None):
        self.array = []
        self.leftmost(root)
        
    # Time -> 0(h) Space -> 0(h)
    def leftmost(self, root: TreeNode | None):
        while root is not None:
            self.array.append(root)
            root = root.left

    # Time -> 0(h) Space -> 0(h) on avg 0(1)
    def next(self) -> int:
        curr = self.array.pop()
        self.leftmost(curr.right)
        return curr.val
    
    # Time -> 0(h) Space -> 0(1)
    def has_next(self) -> bool:
        return len(self.array) > 0