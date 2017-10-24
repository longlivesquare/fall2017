package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var time int
var tree []Edge
var back []Edge
var forw []Edge
var cros []Edge

type Vertex struct {
	path  []int
	adj   []*Vertex
	color int // -1 = white, 0 = grey, 1 =black
	dist  float64
	id    int
	disc  int // Vertex first discovered in DFS
	fin   int // Vertex finished in DFS
}

type Edge struct {
	x int
	y int
}

type Graph struct {
	verts  []*Vertex
	topOrd []int
}

type sortGraph []*Vertex

func (v sortGraph) Len() int {
	return len(v)
}

func (v sortGraph) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v sortGraph) Less(i, j int) bool {
	return v[i].id < v[j].id
}

type sortEdge []Edge

func (e sortEdge) Len() int {
	return len(e)
}

func (e sortEdge) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e sortEdge) Less(i, j int) bool {
	if e[i].x < e[j].x {
		return true
	} else if e[i].x == e[j].x {
		return e[i].y < e[j].y
	} else {
		return false
	}
}

func BFS(g *Graph, s *Vertex) {
	for i := range g.verts {
		g.verts[i].color = -1 // White
		g.verts[i].dist = math.Inf(1.0)
		g.verts[i].path = make([]int, 0)
	}
	s.color = 0
	s.dist = 0

	queue := make([]*Vertex, 0)

	queue = append(queue, s)

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		for i, v := range u.adj {
			if v.color == -1 {
				u.adj[i].color = 0
				u.adj[i].dist = u.dist + 1
				u.adj[i].path = append(u.adj[i].path, u.path...)
				u.adj[i].path = append(u.adj[i].path, u.id)
				queue = append(queue, u.adj[i])
			}
		}
		u.color = 1
		u.path = append(u.path, u.id)
	}
}

func DFS(g *Graph) {
	for i := range g.verts {
		g.verts[i].color = -1
		g.verts[i].path = make([]int, 0)
	}
	time = 0
	for i, v := range g.verts {
		if v.color == -1 {
			DFSVisit(g, g.verts[i])
		}
	}
}

func DFSVisit(g *Graph, u *Vertex) {
	time++
	u.disc = time
	u.color = 0

	for i, v := range u.adj {
		if v.color == -1 {
			u.adj[i].path = append(u.adj[i].path, u.id)
			addTree(Edge{u.id, v.id})
			DFSVisit(g, u.adj[i])
		} else if v.color == 0 {
			addBack(Edge{u.id, v.id})
		} else { //v.color == 1
			if v.disc > u.disc {
				addForw(Edge{u.id, v.id})
			} else {
				addCros(Edge{u.id, v.id})
			}
		}
	}

	u.color = 1
	time++
	u.fin = time
	g.topOrd = append([]int{u.id}, g.topOrd...)
}

// Makes a graph from a passed in filename
func makeGraph(s string) Graph {
	g := Graph{}
	g.topOrd = make([]int, 0)
	f, err := os.Open(s)
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		v := &Vertex{}
		a := &Vertex{}

		str := scanner.Text()
		split := strings.Split(str, " ")
		id, _ := strconv.Atoi(split[0])
		adj, _ := strconv.Atoi(split[1])

		vexists := false // Vertex exists
		aexists := false // Adjacent vertex exists
		for i, u := range g.verts {
			if u.id == id {
				vexists = true
				v = g.verts[i]
			}
			if u.id == adj {
				aexists = true
				a = g.verts[i]
			}
		}

		if !aexists {
			a.id = adj
			g.verts = append(g.verts, a)
		}

		v.adj = append(v.adj, a)
		if !vexists {
			v.id = id
			g.verts = append(g.verts, v)
		}

	}
	f.Close()
	sort.Sort(sortGraph(g.verts))
	return g
}

func printBFS(g *Graph) {
	BFS(g, g.verts[0])

	fmt.Println("Breadth First Search")
	fmt.Println("Vertex: Distance [Path]")

	for _, v := range g.verts {
		fmt.Printf("%d : %f %v\n", v.id, math.Trunc(v.dist), v.path) // id, dist, path
	}
}

func printDFS(g *Graph) {
	tree = make([]Edge, 0)
	cros = make([]Edge, 0)
	back = make([]Edge, 0)
	forw = make([]Edge, 0)

	DFS(g)

	fmt.Println("\nDepth First Search")
	fmt.Println("Vertex: Discover/Finish")
	for _, v := range g.verts {
		fmt.Printf("%d : %d/%d\n", v.id, v.disc, v.fin)
	}

	sort.Sort(sortEdge(tree))
	fmt.Print("Tree: [")
	for _, e := range tree {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	sort.Sort(sortEdge(back))
	fmt.Print("Back: [")
	for _, e := range back {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	sort.Sort(sortEdge(forw))
	fmt.Print("Forward: [")
	for _, e := range forw {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	sort.Sort(sortEdge(cros))
	fmt.Print("Cross: [")
	for _, e := range cros {
		fmt.Printf("(%d, %d) ", e.x, e.y)
	}
	fmt.Println("]")

	fmt.Printf("Vertices in topological order: %v\n\n", g.topOrd)
}

func addTree(e Edge) {
	tree = append(tree, e)
}

func addForw(e Edge) {
	forw = append(forw, e)
}

func addCros(e Edge) {
	cros = append(cros, e)
}

func addBack(e Edge) {
	back = append(back, e)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Graph 1")
	g1 := makeGraph("graph1.dat")
	printBFS(&g1)
	printDFS(&g1)

	fmt.Println("\nGraph 2")
	g2 := makeGraph("graph2.dat")
	printBFS(&g2)
	printDFS(&g2)

	fmt.Println("\nGraph 3")
	g3 := makeGraph("graph3.dat")
	printBFS(&g3)
	printDFS(&g3)
}
