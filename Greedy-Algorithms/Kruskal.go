package main

import "fmt"

type Weight int
type Edges []EdgeStructure

type EdgeStructure struct {
	v, w   int
	weight Weight
}

type Graph struct { // adj array graph
	vertexNum int
	edgeArray [][]Weight
}

func (e Edges) QuickSort(left, right int) {
	pivot := right
	l := left
	r := right - 1
	swap := func(a, b int) {
		temp := e[a]
		e[a] = e[b]
		e[b] = temp
	}

	fmt.Println("function called", l, r)
	for l <= r {
		fmt.Println(l, r)
		for l <= r && e[l].weight < e[pivot].weight {
			l++
			fmt.Println("l :", l)
		}
		for r >= l && e[r].weight > e[pivot].weight {
			r--
			fmt.Println("r :", r)
		}

		if l < r {
			swap(l, r)
			l++
			r--
			fmt.Println("swap", l, r)
		}
	}
	swap(l, pivot)

	if left < l-1 {
		e.QuickSort(left, l-1)
	}
	if right > l+1 {
		e.QuickSort(l+1, right)
	}
}

func (e Edges) testQuicksort() {
	e.QuickSort(0, len(e)-1)
	fmt.Println(e)
}

func (g *Graph) Graph(n int) {
	g.vertexNum = n
	g.edgeArray = make([][]Weight, n)

	for i := 0; i < n; i++ {
		g.edgeArray[i] = make([]Weight, n)
	}
}

func MinMax(x, y int) (min, max int) {
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}
	return
}

func (g *Graph) InsertEdge(v, w int, weight Weight) {

	if v == w {
		fmt.Printf(" %d == %d", v, w)
		return
	}

	min, max := MinMax(v, w)
	g.edgeArray[min][max] = weight
}

func (g Graph) kruskal() (mst Graph) {
	mst.Graph(g.vertexNum)

	//================================================================================
	edges := Edges{}
	for i := 0; i < g.vertexNum; i++ {
		for j := 0; j < g.vertexNum; j++ {
			if g.edgeArray[i][j] != 0 {
				edges = append(edges, EdgeStructure{i, j, g.edgeArray[i][j]})
			}
		}
	}

	//edges sort ascending
	fmt.Println(edges)
	edges.QuickSort(0, len(edges)-1)
	fmt.Println("sorting done")
	//================================================================================
	// 힙으로 바꿀 것

	//union-find
	var edgeCount int
	var parent = []int{}
	var setSize = []int{}
	const (
		ROOT = -1
	)

	func() {
		for i := 0; i < mst.vertexNum; i++ {
			parent = append(parent, ROOT)
			setSize = append(setSize, 1)
		}
	}()

	find := func(a int) int {
		root := a
		for parent[root] != ROOT {
			root = parent[root]
		}

		var p int
		i := a
		for parent[i] != ROOT {
			p = parent[i]
			parent[i] = root
			i = p
		}

		return root
	}

	union := func(a, b int) {
		if setSize[a] > setSize[b] {
			parent[b] = a
			setSize[a] += setSize[b]
		} else {
			parent[a] = b
			setSize[b] += setSize[a]
		}
	}

	for i := 0; edgeCount < mst.vertexNum-1; i++ {
		if find(edges[i].v) != find(edges[i].w) {
			mst.InsertEdge(edges[i].v, edges[i].w, edges[i].weight)
			union(edges[i].v, edges[i].w)
			edgeCount++
		}
	}
	return
}

func main() {
	var g Graph
	g.Graph(5)
	fmt.Println("graph init")

	g.InsertEdge(0, 1, 10)
	g.InsertEdge(0, 3, 8)
	g.InsertEdge(0, 2, 5)
	g.InsertEdge(1, 4, 9)
	g.InsertEdge(1, 3, 3)

	g.InsertEdge(2, 4, 6)
	g.InsertEdge(3, 4, 2)
	fmt.Println(g)
	fmt.Println("edge insertion done")

	mst := g.kruskal()
	fmt.Println(mst)

}
