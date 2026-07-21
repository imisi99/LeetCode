package main

type DLL struct {
	Val  int
	Next *DLL
	Prev *DLL
}

type LRUCache struct {
	Cap  int
	Size int
	Head *DLL
	Tail *DLL
	Map  map[int]*DLL
}

func NewLRUCache(cap int) LRUCache {
	return LRUCache{
		cap,
		0,
		nil,
		nil,
		make(map[int]*DLL),
	}
}

// Get Time -> 0(1) Space -> 0(1)
func (l *LRUCache) Get(key int) int {
	if val, exist := l.Map[key]; exist {
		if l.Size == 1 || l.Tail == val {
			return val.Val
		}

		if l.Head == val {
			val.Next.Prev = nil
			l.Head = val.Next
		} else {
			val.Prev.Next = val.Next
			val.Next.Prev = val.Prev
		}

		l.Tail.Next = val
		val.Prev = l.Tail
		l.Tail = val
		l.Tail.Next = nil
		return val.Val
	}
	return -1
}

// Put Time -> 0(1) Space -> 0(1)
func (l *LRUCache) Put(key int, value int) {
	if val, exist := l.Map[key]; exist {
		val.Val = value
		return
	}

	if l.Size == l.Cap {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		l.Size--
	}

	node := &DLL{Val: value}
	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node
	l.Size++
}
