// Copyright 2022 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package framer

import (
	"context"
	"github.com/synnaxlabs/freighter/fgrpc"
	"github.com/synnaxlabs/synnax/pkg/distribution/framer"
	"github.com/synnaxlabs/synnax/pkg/distribution/framer/iterator"
	"github.com/synnaxlabs/synnax/pkg/distribution/framer/writer"
	framerv1 "github.com/synnaxlabs/synnax/pkg/distribution/transport/grpc/gen/proto/go/framer/v1"
	"google.golang.org/grpc"
)

type (
	writerClient = fgrpc.StreamClientCore[
		writer.Request,
		*framerv1.WriterRequest,
		writer.Response,
		*framerv1.WriterResponse,
	]
	writerServerCore = fgrpc.StreamServerCore[
		writer.Request,
		*framerv1.WriterRequest,
		writer.Response,
		*framerv1.WriterResponse,
	]
	iteratorClient = fgrpc.StreamClientCore[
		iterator.Request,
		*framerv1.IteratorRequest,
		iterator.Response,
		*framerv1.IteratorResponse,
	]
	iteratorServerCore = fgrpc.StreamServerCore[
		iterator.Request,
		*framerv1.IteratorRequest,
		iterator.Response,
		*framerv1.IteratorResponse,
	]
)

var (
	_ framerv1.WriterServiceServer   = (*writerServer)(nil)
	_ writer.TransportServer         = (*writerServer)(nil)
	_ writer.TransportClient         = (*writerClient)(nil)
	_ framerv1.IteratorServiceServer = (*iteratorServer)(nil)
	_ iterator.TransportServer       = (*iteratorServer)(nil)
	_ iterator.TransportClient       = (*iteratorClient)(nil)
	_ framer.Transport               = Transport{}
	_ fgrpc.BindableTransport        = Transport{}
)

// New creates a new grpc Transport that opens connections from the given pool.
func New(pool *fgrpc.Pool) Transport {
	return Transport{
		writer: writerTransport{
			client: &writerClient{
				Pool:               pool,
				RequestTranslator:  writerRequestTranslator{},
				ResponseTranslator: writerResponseTranslator{},
				ClientFunc: func(
					ctx context.Context,
					conn grpc.ClientConnInterface,
				) (fgrpc.GRPCClientStream[*framerv1.WriterRequest, *framerv1.WriterResponse], error) {
					return framerv1.NewWriterServiceClient(conn).Write(ctx)
				},
			},
			server: &writerServer{writerServerCore: writerServerCore{
				RequestTranslator:  writerRequestTranslator{},
				ResponseTranslator: writerResponseTranslator{},
				ServiceDesc:        &framerv1.WriterService_ServiceDesc,
			}},
		},
		iterator: iteratorTransport{
			server: &iteratorServer{iteratorServerCore: iteratorServerCore{
				RequestTranslator:  iteratorRequestTranslator{},
				ResponseTranslator: iteratorResponseTranslator{},
				ServiceDesc:        &framerv1.IteratorService_ServiceDesc,
			}},
			client: &iteratorClient{
				Pool:               pool,
				RequestTranslator:  iteratorRequestTranslator{},
				ResponseTranslator: iteratorResponseTranslator{},
				ClientFunc: func(
					ctx context.Context,
					conn grpc.ClientConnInterface,
				) (fgrpc.GRPCClientStream[*framerv1.IteratorRequest, *framerv1.IteratorResponse], error) {
					return framerv1.NewIteratorServiceClient(conn).Iterate(ctx)
				},
			},
		},
	}
}

type writerServer struct{ writerServerCore }

func (w *writerServer) Write(server framerv1.WriterService_WriteServer) error {
	return w.Handler(server.Context(), w.Server(server))
}

type iteratorServer struct{ iteratorServerCore }

func (t *iteratorServer) Iterate(server framerv1.IteratorService_IterateServer) error {
	return t.Handler(server.Context(), t.Server(server))
}

// Transport is a grpc backed implementation of the framer.Transport interface.
type Transport struct {
	writer   writerTransport
	iterator iteratorTransport
}

// Writer implements the framer.Transport interface.
func (t Transport) Writer() writer.Transport { return t.writer }

// Iterator implements the framer.Transport interface.
func (t Transport) Iterator() iterator.Transport { return t.iterator }

// BindTo implements the fgrpc.BindableTransport interface.
func (t Transport) BindTo(server grpc.ServiceRegistrar) {
	framerv1.RegisterWriterServiceServer(server, t.writer.server)
	framerv1.RegisterIteratorServiceServer(server, t.iterator.server)
}

type writerTransport struct {
	client *writerClient
	server *writerServer
}

// Client implements the writer.Transport interface.
func (t writerTransport) Client() writer.TransportClient { return t.client }

// Server implements the writer.Transport interface.
func (t writerTransport) Server() writer.TransportServer { return t.server }

type iteratorTransport struct {
	client *iteratorClient
	server *iteratorServer
}

// Client implements the iterator.Transport interface.
func (t iteratorTransport) Client() iterator.TransportClient { return t.client }

// Server implements the iterator.Transport interface.
func (t iteratorTransport) Server() iterator.TransportServer { return t.server }
