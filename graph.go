package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	Nodes map[string]*Node
}

type Node struct {
	Name     string
	Children []*Node
}

type DuplicateInfo struct {
	ScriptName string
	Count      int
	Parents    []string
}

func NewGraph() *Graph {
	return &Graph{Nodes: make(map[string]*Node)}
}

func (g *Graph) AddNode(name string) *Node {
	if _, exists := g.Nodes[name]; !exists {
		g.Nodes[name] = &Node{Name: name}
	}
	return g.Nodes[name]
}

func (g *Graph) AddEdge(from, to string) {
	fromNode := g.AddNode(from)
	toNode := g.AddNode(to)
	fromNode.Children = append(fromNode.Children, toNode)
}

func (n *Node) Print(indent int) {
	fmt.Printf("%s- %s\n", strings.Repeat("  ", indent), n.Name)
	for _, child := range n.Children {
		child.Print(indent + 1)
	}
}
