package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vertex struct {
	path  []int
	adj   []*Vertex
	color int // -1 = white, 0 = grey, 1 =black
	dist  int
	id    int
}

type Graph struct {
	verts []*Vertex
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

func BFS(g *Graph, s *Vertex) {
	for i := range g.verts {
		g.verts[i].color = -1 // White
		g.verts[i].dist = -1
	}
	s.color = 0
	s.dist = 0
	s.path = make([]int, 0)

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

// Makes a graph from a passed in filename
func makeGraph(s string) Graph {
	g := Graph{}
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

	fmt.Println("Vertex: Distance [Path]")

	for _, v := range g.verts {
		fmt.Printf("%d : %d %v\n", v.id, v.dist, v.path) // id, dist, path
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	g := makeGraph("graph1.dat")
	printBFS(&g)

}
