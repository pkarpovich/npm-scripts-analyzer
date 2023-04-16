package main

import (
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	graph := NewGraph()
	graph.AddNode("script1")
	graph.AddNode("script2")

	if len(graph.Nodes) != 2 {
		t.Errorf("Expected 2 nodes in the graph, got %d", len(graph.Nodes))
	}

	if graph.Nodes["script1"] == nil {
		t.Errorf("Expected script1 to be in the graph")
	}

	if graph.Nodes["script2"] == nil {
		t.Errorf("Expected script2 to be in the graph")
	}
}

func TestGraph_AddEdge(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge("script1", "script2")

	if len(graph.Nodes["script1"].Children) != 1 {
		t.Errorf("Expected 1 child for script1, got %d", len(graph.Nodes["script1"].Children))
	}

	if graph.Nodes["script1"].Children[0].Name != "script2" {
		t.Errorf("Expected child to be script2, got %s", graph.Nodes["script1"].Children[0].Name)
	}

	if len(graph.Nodes["script2"].Children) != 0 {
		t.Errorf("Expected 0 children for script2, got %d", len(graph.Nodes["script2"].Children))
	}

	if len(graph.Nodes) != 2 {
		t.Errorf("Expected 2 nodes in the graph, got %d", len(graph.Nodes))
	}

	if graph.Nodes["script1"] == nil {
		t.Errorf("Expected script1 to be in the graph")
	}
}
