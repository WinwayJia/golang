package main

import (
	"fmt"
)

type node struct {
	val   interface{}
	ptr   []*node
	steps int
	path  []*node
}

func bfs(src *node, dst *node) {
	queue := make([]*node, 0)
	queue = append(queue, src)
	src.steps = 0
	src.path = append(src.path, src)

	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		if tmp.val == dst.val {
			break
		}

		for _, n := range tmp.ptr {
			if n.steps < 0 || tmp.steps+1 < n.steps {
				n.steps = tmp.steps + 1
				n.path = append(tmp.path, n)
			}
			queue = append(queue, n)
		}
	}

	// 输出路径
	start := true
	fmt.Printf("[%d->%d]: ", src.val, dst.val)
	for _, p := range dst.path {
		if start {
			fmt.Printf("%d", p.val)
			start = false
		} else {
			fmt.Printf(" -> %d", p.val)
		}
	}
	fmt.Printf("\n")
}

func dfs(src *node, dst *node) {

}

func main() {
	n1 := &node{val: 1, steps: -1}
	n2 := &node{val: 2, steps: -1}
	n3 := &node{val: 3, steps: -1}
	n4 := &node{val: 4, steps: -1}
	n5 := &node{val: 5, steps: -1}
	n6 := &node{val: 6, steps: -1}
	n7 := &node{val: 7, steps: -1}

	n1.ptr = append(n1.ptr, n2)
	n1.ptr = append(n1.ptr, n5)

	n2.ptr = append(n2.ptr, n3)
	n2.ptr = append(n2.ptr, n4)

	n3.ptr = append(n3.ptr, n4)

	n4.ptr = append(n4.ptr, n6)

	n5.ptr = append(n5.ptr, n3)
	n5.ptr = append(n5.ptr, n6)

	n6.ptr = append(n6.ptr, n7)

	bfs(n1, n2)
	bfs(n1, n3)
	bfs(n1, n4)
	bfs(n1, n5)
	bfs(n1, n6)
	bfs(n1, n7)
}
