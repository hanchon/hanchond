package database

import (
	"context"
	"fmt"
)

// GetAllNodesForChainID returns all the nodes for a given chain ID
// and extends the SQL generated query to check if there are zero nodes
// found.
func (q *Queries) GetAllNodesForChainID(ctx context.Context, chainID int64) ([]GetAllChainNodesRow, error) {
	nodes, err := q.GetAllChainNodes(ctx, chainID)
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("no nodes found for chain %d", chainID)
	}

	return nodes, nil
}
