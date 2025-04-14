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
			BaseNode: testBaseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": {
			BaseNode: testBaseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": {
			BaseNode: testBaseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": {
			BaseNode:        testBaseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("cover"),
		},
	}

	expected := map[string]TypedNode{
		"text": TextNode{
			BaseNode: testBaseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": FileNode{
			BaseNode: testBaseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": LinkNode{
			BaseNode: testBaseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": GroupNode{
			BaseNode:        testBaseNode("group"),
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
			BaseNode: testBaseNode("text"),
			Text:     nil,
		},
		"file": {
			BaseNode: testBaseNode("file"),
			File:     nil,
		},
		"link": {
			BaseNode: testBaseNode("link"),
			URL:      nil,
		},
		"group": {
			BaseNode: testBaseNode("group"),
			Label:    nil,
		},
		"groupBg": {
			BaseNode:        testBaseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("error"),
		},
	}

	expectedError := map[string]string{
		"text":    "text type node requires text attribute",
		"file":    "file type node requires file attribute",
		"link":    "link type node requires url attribute",
		"group":   "group type node requires label attribute",
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
			BaseNode: testBaseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": FileNode{
			BaseNode: testBaseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": LinkNode{
			BaseNode: testBaseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": GroupNode{
			BaseNode:        testBaseNode("group"),
			Label:           util.StrPtr("label"),
			Background:      util.StrPtr("1"),
			BackgroundStyle: util.StrPtr("cover"),
		},
	}

	expected := map[string]Node{
		"text": {
			BaseNode: testBaseNode("text"),
			Text:     util.StrPtr("text"),
		},
		"file": {
			BaseNode: testBaseNode("file"),
			File:     util.StrPtr("fileName"),
			Subpath:  util.StrPtr("fileSubpath"),
		},
		"link": {
			BaseNode: testBaseNode("link"),
			URL:      util.StrPtr("URL"),
		},
		"group": {
			BaseNode:        testBaseNode("group"),
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

func TestConstructors(t *testing.T) {
	inputs := map[string]TypedNode{
		"text":  NewTextNode("text"),
		"file":  NewFileNode("file", nil),
		"link":  NewLinkNode("link"),
		"group": NewGroupNode("group", nil, nil),
	}

	for _, node := range inputs {
		require.NoError(t, node.Validate())
	}
}

func TestBaseNodeOpts(t *testing.T) {
	node := newBaseNode("text", Position(100, 200), Width(10), Height(20))

	assert.Equal(t, node.X, 100)
	assert.Equal(t, node.Y, 200)
	assert.Equal(t, node.Width, 10)
	assert.Equal(t, node.Height, 20)
}

func testBaseNode(t string) BaseNode {
	return BaseNode{
		ID:     t,
		Type:   t,
		X:      0,
		Y:      0,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	}
}
