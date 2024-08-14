package evmos

import (
	"context"
	"log"
	"strconv"

	"github.com/hanchon/hanchond/playground/database"
)

type NodeFromDB struct {
	Node  database.Node
	Chain database.Chain
	Ports database.Port
}

func GetNodeFromDB(queries *database.Queries, nodeID string) *NodeFromDB {
	validatorID, err := strconv.ParseInt(nodeID, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	validatorNode, err := queries.GetNode(context.Background(), validatorID)
	if err != nil {
		log.Panic(err)
	}

	validatorPorts, err := queries.GetNodePorts(context.Background(), validatorNode.ID)
	if err != nil {
		log.Panic(err)
	}
	chain, err := queries.GetChain(context.Background(), validatorNode.ChainID)
	if err != nil {
		log.Panic(err)
	}

	return &NodeFromDB{
		Node:  validatorNode,
		Chain: chain,
		Ports: validatorPorts,
	}
}

func NewEvmosFromDB(queries *database.Queries, nodeID string) *Evmos {
	data := GetNodeFromDB(queries, nodeID)
	e := NewEvmos(data.Node.Moniker, data.Node.BinaryVersion, data.Node.ConfigFolder, data.Chain.ChainID, data.Node.ValidatorKeyName, data.Chain.Denom)
	e.RestorePortsFromDB(data.Ports)
	return e
}
