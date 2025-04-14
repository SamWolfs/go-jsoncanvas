package jsoncanvas

import (
	"testing"

	"github.com/SamWolfs/go-jsoncanvas/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToTypedNodeSuccess(t *testing.T) {
	inputs := map[string]Node{
		"text": {
			BaseNode: baseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": {
			BaseNode: baseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": {
			BaseNode: baseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": {
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

func TestToTypedNodeFailure(t *testing.T) {
	inputs := map[string]Node{
		"text": {
			BaseNode: baseNode("text"),
			Text:     nil,
		},
		"file": {
			BaseNode: baseNode("file"),
			File:     nil,
		},
		"link": {
			BaseNode: baseNode("link"),
			URL:      nil,
		},
		"group": {
			BaseNode:        baseNode("group"),
			Label:           nil,
		},
		"groupBg": {
			BaseNode:        baseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("error"),
		},
	}

	expectedError := map[string]string{
		"text": "text type node requires text attribute",
		"file": "file type node requires file attribute",
		"link": "link type node requires url attribute",
		"group": "group type node requires label attribute",
		"groupBg": "invalid background style",
	}

	for nodeType, node := range inputs {
		_, err := node.ToTypedNode()

		require.ErrorContains(t, err, expectedError[nodeType])
	}
}

func TestToNode(t *testing.T) {
	inputs := map[string]TypedNode{
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

	expected := map[string]Node{
		"text": {
			BaseNode: baseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": {
			BaseNode: baseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": {
			BaseNode: baseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": {
			BaseNode:        baseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("cover"),
		},
	}


	for nodeType, node := range inputs {
		actual := node.ToNode()

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
