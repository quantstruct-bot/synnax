package core

import (
	"context"
	"github.com/synnaxlabs/aspen"
	aspentransport "github.com/synnaxlabs/aspen/transport/grpc"
	"github.com/synnaxlabs/synnax/pkg/storage"
	"github.com/synnaxlabs/x/config"
)

// Core is the foundational primitive for distributed compute in the delta Cluster. It exposes the following essential
// APIs:
//
//  1. LocalKey.KV - an eventually consistent distributed key-value store.
//  2. LocalKey.TS - a node local time-series engine for writing segment data.
//  3. Cluster - an API for querying information about the underlying Cluster topology.
type Core struct {
	// Config is the configuration for the distribution layer.
	Config Config
	// Cluster is the API for the delta Cluster.
	Cluster aspen.Cluster
	// Storage is the storage for the node. The distribution layer replaces the original key-value store with
	// a distributed key-value store. The caller should NOT call Close on the storage engine.
	Storage *storage.Store
}

// Open opens a new  core distribution layer. The caller is responsible for closing the distribution layer when it is
// no longer in use.
func Open(ctx context.Context, cfg Config) (c Core, err error) {
	cfg, err = config.OverrideAndValidate(DefaultConfig, cfg)
	if err != nil {
		return c, err
	}

	c.Storage, err = storage.Open(cfg.Storage)
	if err != nil {
		return c, err
	}

	clusterTransport := aspentransport.New(cfg.Pool)
	*cfg.Transports = append(*cfg.Transports, clusterTransport)

	// Since we're using our own key-value engine, the value we used for 'dirname'
	// doesn't matter.
	clusterKV, err := aspen.Open(
		ctx,
		/* dirname */ "",
		cfg.AdvertiseAddress,
		cfg.PeerAddresses,
		aspen.WithEngine(c.Storage.KV),
		aspen.WithExperiment(cfg.Experiment),
		aspen.WithLogger(cfg.Logger.Sugar()),
		aspen.WithTransport(clusterTransport),
	)
	c.Cluster = clusterKV
	// configure out storage system to use a distributed key-value store
	c.Storage.KV = clusterKV
	return c, err
}
