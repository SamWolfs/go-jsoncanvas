package jsoncanvas

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeFilters(t *testing.T) {
	files := []TypedNode{NewFileNode("File 1", nil), NewFileNode("File 2", nil)}
	groups := []TypedNode{NewGroupNode("Group 1", nil, nil), NewGroupNode("Group 2", nil, nil)}
	links := []TypedNode{NewLinkNode("Link 1"), NewLinkNode("Link 2")}
	texts := []TypedNode{NewTextNode("Text 1"), NewTextNode("Text 2")}

	nodes := slices.Concat(files, groups, links, texts)
	shuffle(nodes)

	c := NewCanvas(WithNodes(nodes...))

	assert.ElementsMatch(t, files, c.FileNodes())
	assert.ElementsMatch(t, groups, c.GroupNodes())
	assert.ElementsMatch(t, links, c.LinkNodes())
	assert.ElementsMatch(t, texts, c.TextNodes())
}

func TestGroupedNodes(t *testing.T) {
	groupNode := NewGroupNode("Group 1", nil, nil, Position(0, 0), Width(100), Height(100))
	contained := []TypedNode{
		NewTextNode("Text 1", Position(1, 1), Width(10), Height(10)),
		NewTextNode("Text 2", Position(20, 1), Width(10), Height(10)),
		NewTextNode("Text 3", Position(40, 1), Width(10), Height(10)),
	}
	rest := []TypedNode{
		NewTextNode("Text 4", Position(-1, 1), Width(10), Height(10)),
		NewTextNode("Text 5", Position(120, 1), Width(10), Height(10)),
		NewTextNode("Text 6", Position(0, 0), Width(110), Height(10)),
	}

	nodes := slices.Concat([]TypedNode{groupNode}, contained, rest)
	shuffle(nodes)

	c := NewCanvas(WithNodes(nodes...))

	assert.ElementsMatch(t, contained, c.NodeGroup(groupNode))
}

func shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
