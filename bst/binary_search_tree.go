package bst

import "fmt"

type Node struct {
	Left, Right, Parent *Node
	Key int
}

type Tree struct {
	Root *Node
}

func (T *Tree) InsertNode(z *Node) {
	var y *Node
	var x = T.Root
	
	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y == nil {
		T.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}

func (T *Tree) Transplant(u, v *Node) {
	if u.Parent == nil {
		T.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

func (T *Tree) Minimum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func (T *Tree) TreeDelete(z *Node) {
	if z.Left == nil {
		T.Transplant(z, z.Right)
	} else if z.Right == nil {
		T.Transplant(z, z.Right)
	} else {
		y := T.Minimum(z.Right)
		if y.Parent != z {
			T.Transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		T.Transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

func TreeSearch(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return TreeSearch(x.Left, k)
	} else {
		return TreeSearch(x.Right, k)
	}
}


func InOrderTreeWalk(x *Node) {
	if x != nil {
		InOrderTreeWalk(x.Left)
		fmt.Println(x.Key)
		InOrderTreeWalk(x.Right)
	}
}

func PreOrderTreeWalk(x *Node) {
	if x != nil {
		fmt.Println(x.Key)
		PreOrderTreeWalk(x.Left)
		PreOrderTreeWalk(x.Right)
	}
}

func PostOrderTreeWalk(x *Node) {
	if x != nil {
		PostOrderTreeWalk(x.Left)
		PostOrderTreeWalk(x.Right)
		fmt.Println(x.Key)
	}
}