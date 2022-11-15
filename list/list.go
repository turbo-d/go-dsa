package list

type SinglyNode struct {
	key  int
	next *SinglyNode
}

type SinglyList struct {
	head *SinglyNode
}

func (sl SinglyList) Head() *SinglyNode {
	return sl.head
}

func (sl SinglyList) Insert(node *SinglyNode) {
	if node == nil {
		// No can do
	}

	node.next = sl.head.next
	sl.head.next = node
}

func (sl SinglyList) Delete(node *SinglyNode) {
	if node == nil {
		// No can do
	}

	prev := sl.head
	for prev != nil && prev != node {
		prev = prev.next
	}

	if prev == nil {
		// No can do
	}

	prev.next = node.next
}

func (sl SinglyList) Search(key int) *SinglyNode {
	cur := sl.head
	for cur != nil {
		if cur.key == key {
			break
		}
		cur = cur.next
	}
	return cur
}

type DoublyNode struct {
	key  int
	next *DoublyNode
	prev *DoublyNode
}

type DoublyList struct {
	head *DoublyNode
	tail *DoublyNode
}

func (dl DoublyList) Head() *DoublyNode {
	return dl.head
}

func (dl DoublyList) Tail() *DoublyNode {
	return dl.tail
}

func (dl DoublyList) Insert(node *DoublyNode) {
	if node == nil {
		// No can do
	}

	prev := dl.head
	next := dl.head.next

	prev.next = node
	node.prev = prev

	node.next = next
	next.prev = node
}

func (dl DoublyList) Delete(node *DoublyNode) {
	if node == nil {
		// No can do
	}

	node.prev.next = node.next
}

func (dl DoublyList) Search(key int) *DoublyNode {
	cur := dl.head
	for cur != nil {
		if cur.key == key {
			break
		}
		cur = cur.next
	}
	return cur
}
