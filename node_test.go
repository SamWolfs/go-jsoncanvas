package jsoncanvas

import (
	"testing"

	"github.com/SamWolfs/go-jsoncanvas/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToTypedNode(t *testing.T) {
	inputs := map[string]Node{
		"text": Node{
			BaseNode: baseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": Node{
			BaseNode: baseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": Node{
			BaseNode: baseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": Node{
			BaseNode:        baseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("cover"),
		},
	}

	expected := map[string]TypedNode{
		"text": TextNode{
			BaseNode: baseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": FileNode{
			BaseNode: baseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": LinkNode{
			BaseNode: baseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": GroupNode{
			BaseNode:        baseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("cover"),
		},
	}

	for nodeType, node := range inputs {
		actual, err := node.ToTypedNode()

		require.NoError(t, err)
		assert.Equal(t, expected[nodeType], actual)
	}
}

func baseNode(t string) BaseNode {
	return BaseNode{
		ID:     t,
		Type:   t,
		X:      0,
		Y:      0,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	}
}
