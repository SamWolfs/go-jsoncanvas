// Copyright (c) 2024 supersonicpineapple

package jsoncanvas

import (
	"fmt"
	"regexp"
	"strings"
)

type Canvas struct {
	Nodes []TypedNode
	Edges []Edge
}

type CanvasOpt func(*Canvas)

func WithNodes(nodes ...TypedNode) CanvasOpt {
	return func(c *Canvas) {
		c.AddNodes(nodes...)
	}
}

func WithEdges(edges ...Edge) CanvasOpt {
	return func(c *Canvas) {
		c.AddEdges(edges...)
	}
}

func NewCanvas(opts ...CanvasOpt) *Canvas {
	c := &Canvas{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Canvas) toJsonCanvas() *jsonCanvas {
	var nodes []Node

	for _, n := range c.Nodes {
		node := n.ToNode()
		nodes = append(nodes, node)
	}

	return &jsonCanvas{
		Nodes: nodes,
		Edges: c.Edges,
	}
}

func (c *Canvas) FileNodes() []FileNode {
	return filterNodes[FileNode](c)
}

func (c *Canvas) GroupNodes() []GroupNode {
	return filterNodes[GroupNode](c)
}

func (c *Canvas) LinkNodes() []LinkNode {
	return filterNodes[LinkNode](c)
}

func (c *Canvas) TextNodes() []TextNode {
	return filterNodes[TextNode](c)
}

func filterNodes[T TypedNode](c *Canvas) []T {
	var nodes []T
	for _, node := range c.Nodes {
		switch node := node.(type) {
		case T:
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (c *Canvas) NodeGroup(n GroupNode) []TypedNode {
	var nodes []TypedNode
	for _, node := range c.Nodes {
		if node.ToNode().ID == n.ID {
			continue
		}
		if n.Contains(node.ToNode().BaseNode) {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (c *Canvas) ChildNodes(n BaseNode) []TypedNode {
	var nodes []TypedNode

	fromEdges, _ := c.nodeEdges(n)

	for _, edge := range fromEdges {
		nodes = append(nodes, c.GetNodeById(edge.ToNode))
	}

	return nodes
}

func (c *Canvas) ParentNodes(n BaseNode) []TypedNode {
	var nodes []TypedNode

	_, toEdges := c.nodeEdges(n)

	for _, edge := range toEdges {
		nodes = append(nodes, c.GetNodeById(edge.FromNode))
	}

	return nodes
}

func (c *Canvas) nodeEdges(n BaseNode) ([]Edge, []Edge) {
	var fromEdges []Edge
	var toEdges []Edge

	for _, edge := range c.Edges {
		if edge.FromNode == n.ID {
			fromEdges = append(fromEdges, edge)
		}

		if edge.ToNode == n.ID {
			toEdges = append(toEdges, edge)
		}
	}

	return fromEdges, toEdges
}

func (c *Canvas) GetNodeById(id string) TypedNode {
	var node TypedNode
	for _, n := range c.Nodes {
		if n.ToNode().ID == id {
			node = n
		}
	}
	return node
}

func (c *Canvas) GetNodesByTag(tag string) []TextNode {
	var nodes []TextNode
	r := regexp.MustCompile("(^|\\s)#" + regexp.QuoteMeta(tag) + "(\\s|$)")
	for _, n := range c.TextNodes() {
		if r.MatchString(*n.Text) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func (c *Canvas) Validate() error {
	if c == nil {
		return nil
	}

	var nodeErrors, edgeErrors []error

	for _, node := range c.Nodes {
		if err := node.Validate(); err != nil {
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

func (c *Canvas) AddNodes(nodes ...TypedNode) *Canvas {
	c.Nodes = append(c.Nodes, nodes...)
	return c
}

func (c *Canvas) AddEdges(edges ...Edge) *Canvas {
	c.Edges = append(c.Edges, edges...)
	return c
}
