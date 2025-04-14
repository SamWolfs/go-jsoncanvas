package jsoncanvas

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type DecoderOption func(*json.Decoder)

type jsonCanvas struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

func (c jsonCanvas) toCanvas() (*Canvas, error) {
	var typedNodes []*TypedNode
	var nodeErrors []error

	for _, n := range c.Nodes {
		typedNode, err := n.ToTypedNode()
		if err != nil {
			nodeErrors = append(nodeErrors, err)
		}
		typedNodes = append(typedNodes, &typedNode)
	}

	if len(nodeErrors) > 0 {
		return nil, fmt.Errorf("validation errors:\n%s", joinErrors(nodeErrors))
	}

	return &Canvas{
		Nodes: typedNodes,
		Edges: c.Edges,
	}, nil
}

func DisallowUnknownFields() DecoderOption {
	return func(decoder *json.Decoder) {
		decoder.DisallowUnknownFields()
	}
}

func Decode(r io.Reader, opts ...DecoderOption) (*Canvas, error) {
	decoder := json.NewDecoder(r)
	for _, opt := range opts {
		opt(decoder)
	}

	c := new(jsonCanvas)
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("can't decode canvas file: %w", err)
	}

	return c.toCanvas()
}

func DecodeFile(path string) (*Canvas, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error when closing file:", err)
		}
	}()

	return Decode(f)
}

func Encode(c *Canvas, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(c.toJsonCanvas()); err != nil {
		return fmt.Errorf("can't encode canvas object: %w", err)
	}
	return nil
}

func EncodeFile(c *Canvas, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		// TODO: create file if it does not exist?
		return fmt.Errorf("can't stat file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("got dir %s, please specify a file instead", path)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 0644)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", path, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error when closing file:", err)
		}
	}()

	return Encode(c, f)
}
