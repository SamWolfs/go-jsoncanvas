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

type TextNode struct {
	BaseNode
	Text *string
}

type FileNode struct {
	BaseNode
	File    *string
	Subpath *string
}

type LinkNode struct {
	BaseNode
	URL *string
}

type GroupNode struct {
	BaseNode
	Label           *string
	Background      *string
	BackgroundStyle *string // one of "cover", "ratio", "repeat"
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
func NewNode() *Node {
	n := Node{
		BaseNode: BaseNode{
			ID:     util.NewID(),
			X:      0,
			Y:      0,
			Width:  DefaultWidth,
			Height: DefaultHeight,
		},
	}
	return &n
}

func (n *Node) SetPosition(x, y int) *Node {
	n.X = x
	n.Y = y
	return n
}

func (n *Node) TranslateX(x int) *Node {
	n.X += x
	return n
}

func (n *Node) TranslateY(y int) *Node {
	n.Y += y
	return n
}

func (n *Node) SetWidth(width int) *Node {
	n.Width = width
	return n
}

func (n *Node) SetHeight(height int) *Node {
	n.Height = height
	return n
}
