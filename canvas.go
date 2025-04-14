package jsoncanvas

import (
	"fmt"
	"strings"
)

type Canvas struct {
	Nodes []*TypedNode
	Edges []*Edge
}

func NewCanvas() *Canvas {
	return &Canvas{}
}

func (c *Canvas) toJsonCanvas() *jsonCanvas {
	var nodes []*Node

	for _, n := range c.Nodes {
		node := (*n).ToNode()
		nodes = append(nodes, &node)
	}

	return &jsonCanvas{
		Nodes: nodes,
		Edges: c.Edges,
	}
}

func (c *Canvas) Validate() error {
	if c == nil {
		return nil
	}

	var nodeErrors, edgeErrors []error

	for _, node := range c.Nodes {
		if err := (*node).Validate(); err != nil {
			nodeErrors = append(nodeErrors, err)
		}
	}

	for _, edge := range c.Edges {
		if err := edge.Validate(); err != nil {
			edgeErrors = append(edgeErrors, err)
		}
	}

	if len(nodeErrors) > 0 || len(edgeErrors) > 0 {
		return fmt.Errorf("validation errors:\n%s\n%s", joinErrors(nodeErrors), joinErrors(edgeErrors))
	} else {
		return nil
	}
}

func joinErrors(errors []error) string {
	var sb strings.Builder

	for i, err := range errors {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, err))
	}

	return sb.String()
}

func (c *Canvas) AddNodes(nodes ...*TypedNode) *Canvas {
	c.Nodes = append(c.Nodes, nodes...)
	return c
}

func (c *Canvas) AddEdges(edges ...*Edge) *Canvas {
	c.Edges = append(c.Edges, edges...)
	return c
}
