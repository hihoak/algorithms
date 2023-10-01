package dijkstras_algorithm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	cases := []struct {
		name              string
		graph             Graph
		startNode         NodeID
		expectedDistances map[NodeID]EdgeValue
	}{
		{
			name: "first case",
			graph: Graph{
				Nodes: map[NodeID]interface{}{
					"1": nil,
					"2": nil,
					"3": nil,
				},
				Edges: map[NodeID]map[NodeID]EdgeValue{
					"1": {
						"2": 1,
						"3": 4,
					},
					"2": {
						"1": 2,
						"3": 1,
					},
				},
			},
			startNode: "1",
			expectedDistances: map[NodeID]EdgeValue{
				"1": 0,
				"2": 1,
				"3": 2,
			},
		},
		{
			name: "second case",
			graph: Graph{
				Nodes: map[NodeID]interface{}{
					"A": nil,
					"B": nil,
					"C": nil,
					"D": nil,
					"E": nil,
					"F": nil,
					"G": nil,
					"H": nil,
					"I": nil,
				},
				Edges: map[NodeID]map[NodeID]EdgeValue{
					"A": {
						"B": 2,
						"H": 15,
					},
					"B": {
						"A": 2,
						"C": 1,
						"D": 5,
					},
					"C": {
						"B": 1,
						"D": 3,
						"F": 2,
						"G": 1,
					},
					"D": {
						"B": 5,
						"C": 3,
						"F": 4,
						"E": 6,
					},
					"E": {
						"D": 6,
						"F": 7,
						"I": 2,
					},
					"F": {
						"C": 2,
						"D": 4,
						"E": 7,
						"H": 3,
						"G": 1,
					},
					"G": {
						"C": 1,
						"F": 1,
					},
					"H": {
						"A": 15,
						"F": 3,
						"I": 12,
					},
					"I": {
						"E": 2,
						"H": 12,
					},
				},
			},
			startNode: "A",
			expectedDistances: map[NodeID]EdgeValue{
				"A": 0,
				"B": 2,
				"C": 3,
				"D": 6,
				"E": 12,
				"F": 5,
				"G": 4,
				"H": 8,
				"I": 14,
			},
		},
		{
			name: "second case",
			graph: Graph{
				Nodes: map[NodeID]interface{}{
					"A": nil,
					"B": nil,
					"C": nil,
					"D": nil,
					"E": nil,
					"F": nil,
					"G": nil,
					"H": nil,
					"I": nil,
				},
				Edges: map[NodeID]map[NodeID]EdgeValue{
					"A": {
						"B": 2,
						"H": 15,
					},
					"B": {
						"A": 2,
						"C": 1,
						"D": 5,
					},
					"C": {
						"B": 1,
						"D": 3,
						"F": 1,
						"G": 1,
					},
					"D": {
						"B": 5,
						"C": 3,
						"F": 4,
						"E": 6,
					},
					"E": {
						"D": 6,
						"F": 7,
						"I": 2,
					},
					"F": {
						"C": 2,
						"D": 1,
						"E": 7,
						"H": 3,
						"G": 1,
					},
					"G": {
						"C": 1,
						"F": 1,
					},
					"H": {
						"A": 15,
						"F": 3,
						"I": 12,
					},
					"I": {
						"E": 2,
						"H": 12,
					},
				},
			},
			startNode: "A",
			expectedDistances: map[NodeID]EdgeValue{
				"A": 0,
				"B": 2,
				"C": 3,
				"D": 5,
				"E": 11,
				"F": 4,
				"G": 4,
				"H": 7,
				"I": 13,
			},
		},
	}

	for _, tc := range cases {
		require.Equalf(t, tc.expectedDistances, DijkstrasAlgorithm(tc.graph, tc.startNode), tc.name)
	}
}
