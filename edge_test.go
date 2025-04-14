package jsoncanvas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSuccess(t *testing.T) {
	edge := Edge{
		ID:       "id",
		FromNode: "from",
		ToNode:   "to",
	}

	require.NoError(t, edge.Validate())
}

func TestValidateFailure(t *testing.T) {
	inputs := map[string]Edge{
		"noFrom": {
			ID:     "id",
			ToNode: "to",
		},
		"noTo": {
			ID:       "id",
			FromNode: "from",
		},
		"fromIsTo": {
			ID:       "id",
			FromNode: "node",
			ToNode:   "node",
		},
	}

	expected := map[string]string{
		"noFrom":   "fromNode and toNode are required",
		"noTo":     "fromNode and toNode are required",
		"fromIsTo": "fromNode and toNode cannot be the same node",
	}

	for errorType, edge := range inputs {
		require.ErrorContains(t, edge.Validate(), expected[errorType])
	}
}
