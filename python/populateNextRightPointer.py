class Node:
    def __init__(self, val=0, left=None, right=None, next=None) -> None:
        self.val = val
        self.left = left
        self.right = right
        self.next = next


def populate(root: Node | None) -> Node | None:
    queue = []
    queue.append([root])
    while len(queue) > 0:
        node = queue.pop(0)
        next = []
        while len(node) > 0:
            curr = node.pop(0)

            if curr is None:
                continue

            if len(node) > 0:
                curr.next = node[0]

            next.append(curr.left)
            next.append(curr.right)

        if len(next) > 0:
            queue.append(next)

    return root
