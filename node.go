package jsoncanvas

import (
	"fmt"

	"github.com/SamWolfs/go-jsoncanvas/internal/util"
)

const (
	DefaultWidth  = 250
	DefaultHeight = 60
	DefaultGap    = 30
)

type TypedNode interface {
	ToNode() Node
	Validate() error
}

type BaseNode struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"` // one of "text", "file", "link", "group"
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Color  *string `json:"color,omitempty"`
}

type Node struct {
	// generic fields
	BaseNode
	Text            *string `json:"text,omitempty"`
	File            *string `json:"file,omitempty"`
	Subpath         *string `json:"subpath,omitempty"`
	URL             *string `json:"url,omitempty"`
	Label           *string `json:"label,omitempty"`
	Background      *string `json:"background,omitempty"`
	BackgroundStyle *string `json:"backgroundStyle,omitempty"`
}

// --------------------------------- TEXT NODE ---------------------------------
type TextNode struct {
	BaseNode
	Text *string
}

func NewTextNode(text string, baseOpts ...BaseNodeOpt) TextNode {
	baseNode := newBaseNode("text", baseOpts...)
	return TextNode{
		BaseNode: baseNode,
		Text:     &text,
	}
}

func (n TextNode) ToNode() Node {
	return Node{
		BaseNode: n.BaseNode,
		Text:     n.Text,
	}
}

func (n TextNode) Validate() error {
	if n.Text == nil {
		return fmt.Errorf("text type node requires text attribute")
	}
	return nil
}

// --------------------------------- FILE NODE ---------------------------------
type FileNode struct {
	BaseNode
	File    *string
	Subpath *string
}

func NewFileNode(file string, subpath *string, baseOpts ...BaseNodeOpt) FileNode {
	baseNode := newBaseNode("text", baseOpts...)
	return FileNode{
		BaseNode: baseNode,
		File:     &file,
		Subpath:  subpath,
	}
}

func (n FileNode) ToNode() Node {
	return Node{
		BaseNode: n.BaseNode,
		File:     n.File,
		Subpath:  n.Subpath,
	}
}

func (n FileNode) Validate() error {
	if n.File == nil || *n.File == "" {
		return fmt.Errorf("file type node requires file attribute")
	}
	return nil
}

// --------------------------------- LINK NODE ---------------------------------
type LinkNode struct {
	BaseNode
	URL *string
}

func NewLinkNode(url string, baseOpts ...BaseNodeOpt) LinkNode {
	baseNode := newBaseNode("text", baseOpts...)
	return LinkNode{
		BaseNode: baseNode,
		URL:      &url,
	}
}

func (n LinkNode) ToNode() Node {
	return Node{
		BaseNode: n.BaseNode,
		URL:      n.URL,
	}
}

func (n LinkNode) Validate() error {
	if n.URL == nil || *n.URL == "" {
		return fmt.Errorf("link type node requires url attribute")
	}
	return nil
}

// --------------------------------- GROUP NODE --------------------------------
type GroupNode struct {
	BaseNode
	Label           *string
	Background      *string
	BackgroundStyle *string // one of "cover", "ratio", "repeat"
}

func NewGroupNode(label string, background, backgroundStyle *string, baseOpts ...BaseNodeOpt) GroupNode {
	baseNode := newBaseNode("text", baseOpts...)
	return GroupNode{
		BaseNode:        baseNode,
		Label:           &label,
		Background:      background,
		BackgroundStyle: backgroundStyle,
	}
}

func (n GroupNode) ToNode() Node {
	return Node{
		BaseNode:        n.BaseNode,
		Label:           n.Label,
		Background:      n.Background,
		BackgroundStyle: n.BackgroundStyle,
	}
}

func (n GroupNode) Validate() error {
	if n.Label == nil {
		return fmt.Errorf("group type node requires label attribute")
	}
	if n.BackgroundStyle != nil && *n.BackgroundStyle != "cover" && *n.BackgroundStyle != "ratio" && *n.BackgroundStyle != "repeat" {
		return fmt.Errorf("invalid background style: %s", *n.BackgroundStyle)
	}
	return nil
}

// -----------------------------------------------------------------------------

func (n *Node) ToTypedNode() (TypedNode, error) {
	switch n.Type {
	case "text":
		node := TextNode{BaseNode: n.BaseNode, Text: n.Text}
		return node, node.Validate()
	case "file":
		node := FileNode{BaseNode: n.BaseNode, File: n.File, Subpath: n.Subpath}
		return node, node.Validate()
	case "link":
		node := LinkNode{BaseNode: n.BaseNode, URL: n.URL}
		return node, node.Validate()
	case "group":
		node := GroupNode{BaseNode: n.BaseNode, Label: n.Label, Background: n.Background, BackgroundStyle: n.BackgroundStyle}
		return node, node.Validate()
	default:
		return nil, fmt.Errorf("invalid type: %s", n.Type)
	}
}

// TODO: Change to Typed Nodes
type BaseNodeOpt func(*BaseNode)

func Position(x, y int) BaseNodeOpt {
	return func(node *BaseNode) {
		node.X = x
		node.Y = y
	}
}

func TranslateX(x int) BaseNodeOpt {
	return func(node *BaseNode) {
		node.X = x
	}
}

func TranslateY(y int) BaseNodeOpt {
	return func(node *BaseNode) {
		node.Y = y
	}
}

func Width(width int) BaseNodeOpt {
	return func(node *BaseNode) {
		node.Width = width
	}
}

func Height(height int) BaseNodeOpt {
	return func(node *BaseNode) {
		node.Height = height
	}
}

func newBaseNode(t string, opts ...BaseNodeOpt) BaseNode {
	node := BaseNode{
		ID:     util.NewID(),
		Type:   t,
		X:      0,
		Y:      0,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	}

	for _, opt := range opts {
		opt(&node)
	}

	return node
}
