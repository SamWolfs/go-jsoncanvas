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

func shuffle[T interface{}](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
