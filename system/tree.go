package system

type Tree struct {
	leaves map[string]*Tree
	data   interface{}
}

func NewTree() *Tree {
	return &Tree{
		leaves: make(map[string]*Tree),
	}
}

//添加树节点
func (t *Tree) Add(nodes []string, data interface{}) {
	lenght := len(nodes)
	if lenght > 0 {
		node := nodes[0]
		tree, ok := t.leaves[node]
		if !ok {
			tree = NewTree()
			t.leaves[node] = tree
		}

		if lenght > 1 {
			tree.Add(nodes[1:], data)
		} else {
			tree.data = data
		}
	}
}

func (t *Tree) Get(nodes []string) interface{} {
	lenght := len(nodes)
	if lenght > 0 {
		tree, ok := t.leaves[nodes[0]]
		if !ok {
			return t.data
		}
		return tree.Get(nodes[1:])
	}
	return t.data
}

func (t *Tree) Delete() {

}

func (t *Tree) Update() {

}
