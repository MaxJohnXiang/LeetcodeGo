// package problem0208

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Trie 是便于 word 插入与查找的数据结构

type Node struct {
	IsWord bool
	Child  [26]*Node
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{
		Root: &Node{
			Child: [26]*Node{},
		},
	}
}

func (this *Trie) Insert(word string) {
	node := this.Root

	for _, c := range word {
		index := c - 'a'
		cNode := node.Child[index]
		if cNode == nil {
			//create a new node
			cNode = &Node{
				Child: [26]*Node{},
			}
			node.Child[index] = cNode
		}
		node = node.Child[index]
	}
	node.IsWord = true
}

func (this *Trie) Search(word string) bool {
	node := this.Root
	for _, c := range word {
		index := c - 'a'
		cNode := node.Child[index]
		if cNode == nil {
			return false
		}
		node = cNode
	}

	return node.IsWord
}

func (this *Trie) StartsWith(word string) bool {
	node := this.Root

	for _, c := range word {
		index := c - 'a'
		cNode := node.Child[index]

		if cNode == nil {
			return false
		}
		node = cNode
	}

	return true
}
