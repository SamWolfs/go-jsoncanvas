package jsoncanvas

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testCanvas = `{
	"nodes":[
		{"id":"00000000000group","type":"group","x":0,"y":0,"width":100,"height":100,"label":"Group"},
		{"id":"00000000000text1","type":"text","text":"text-goes-here","x":200,"y":200,"width":10,"height":10},
		{"id":"00000000000text2","type":"text","text":"text-goes-here","x":220,"y":200,"width":10,"height":10},
		{"id":"00000000000file1","type":"file","file":"file.md","x":10,"y":10,"width":10,"height":10},
		{"id":"00000000000link1","type":"link","URL":"url","x":20,"y":20,"width":10,"height":10}
	],
	"edges":[
		{"id":"0000000000edge01","fromNode":"00000000000file1","fromSide":"right","toNode":"00000000000text1","toSide":"left"},
		{"id":"0000000000edge02","fromNode":"00000000000text1","fromSide":"right","toNode":"00000000000text2","toSide":"left"}
	]
}`

func TestDecodeEncode(t *testing.T) {
	reader := strings.NewReader(testCanvas)
	c, err := Decode(reader)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = Encode(c, &buf)
	require.NoError(t, err)
}
