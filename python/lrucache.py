class DLL:
    def __init__(self, val=0, next=None, prev=None) -> None:
        self.val = val
        self.prev = prev
        self.next = next


class LRUCache:
    def __init__(self, cap=0) -> None:
        self.cap = cap
        self.size = 0
        self.head = None
        self.tail = None
        self.map: dict[int, DLL] = {}

    # Time -> 0(1) Space -> 0(1)
    def get(self, key: int) -> int:
        if self.map.get(key, None):
            val = self.map[key]
            if self.size == 1 or self.tail == val:
                return val.val

            if self.head == val:
                if val.next:
                    val.next.prev = None
                    self.head = val.next
            else:
                if val.prev and val.next:
                    val.prev.next = val.next
                    val.next.prev = val.prev

            if self.tail:
                self.tail.next = val
                val.prev = self.tail
                self.tail = val
                self.tail.next = None
                return val.val
        return -1

    # Time -> 0(1) Space -> 0(1)
    def put(self, key: int, value: int):
        if self.map.get(key, None):
            val = self.map[key]
            val.val = value
            return

        if self.size == self.cap:
            if self.head:
                self.head = self.head.next
                self.head.prev = None
                self.size -= 1

        node = DLL(val=value)
        if self.tail:
            self.tail.next = node
            node.prev = self.tail
            self.tail = node
            self.size += 1
