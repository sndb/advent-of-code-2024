package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func main() {
	net := simple.NewUndirectedGraph()
	ids := map[string]int64{}
	names := map[int64]string{}
	node := func(name string) graph.Node {
		id, ok := ids[name]
		if ok {
			return net.Node(id)
		}

		n := net.NewNode()
		net.AddNode(n)
		ids[name] = n.ID()
		names[n.ID()] = name
		return n
	}

	data, _ := io.ReadAll(os.Stdin)
	links := strings.Fields(string(data))
	for _, l := range links {
		net.SetEdge(simple.Edge{F: node(l[:2]), T: node(l[3:])})
	}

	var best []graph.Node
	for _, clique := range topo.BronKerbosch(net) {
		if len(clique) > len(best) {
			best = clique
		}
	}

	var answer []string
	for _, n := range best {
		answer = append(answer, names[n.ID()])
	}
	sort.Strings(answer)
	fmt.Println(strings.Join(answer, ","))
}
