package _1_array

type Node558 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node558
	TopRight    *Node558
	BottomLeft  *Node558
	BottomRight *Node558
}

func intersect(quadTree1 *Node558, quadTree2 *Node558) *Node558 {
	if quadTree1.IsLeaf {
		if quadTree1.Val { // 1和任意数字逻辑或都是1
			return &Node558{Val: true, IsLeaf: true}
		}
		return quadTree2
	} else if quadTree2.IsLeaf {
		if quadTree2.Val {
			return &Node558{Val: true, IsLeaf: true}
		}
		return quadTree1
	}

	o1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	o2 := intersect(quadTree1.TopRight, quadTree2.TopRight)
	o3 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	o4 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if o1.IsLeaf && o2.IsLeaf && o3.IsLeaf && o4.IsLeaf && o1.Val == o2.Val && o1.Val == o3.Val && o1.Val == o4.Val {
		return &Node558{Val: o1.Val, IsLeaf: true}
	}
	return &Node558{Val: false, IsLeaf: false, TopLeft: o1, TopRight: o2, BottomLeft: o3, BottomRight: o4}
}
