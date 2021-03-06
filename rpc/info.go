package rpc

import (
	"github.com/MixinNetwork/mixin/kernel"
	"github.com/MixinNetwork/mixin/storage"
)

func getInfo(store storage.Store) (map[string]interface{}, error) {
	info := make(map[string]interface{})
	graph, err := kernel.LoadRoundGraph(store)
	if err != nil {
		return info, err
	}
	cacheGraph := make(map[string]interface{})
	for n, r := range graph.CacheRound {
		cacheGraph[n.String()] = map[string]interface{}{
			"node":      r.NodeId.String(),
			"round":     r.Number,
			"start":     r.Start,
			"end":       r.End,
			"snapshots": r.Snapshots,
		}
	}
	finalGraph := make(map[string]interface{})
	for n, r := range graph.FinalRound {
		finalGraph[n.String()] = map[string]interface{}{
			"node":  r.NodeId.String(),
			"round": r.Number,
			"start": r.Start,
			"end":   r.End,
			"hash":  r.Hash.String(),
		}
	}
	info["graph"] = map[string]interface{}{
		"network":  kernel.NetworkId(),
		"node":     kernel.NodeIdForNetwork(),
		"cache":    cacheGraph,
		"final":    finalGraph,
		"topology": kernel.TopologicalOrder(),
	}
	return info, nil
}
