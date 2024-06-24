// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: synnax/pkg/distribution/transport/grpc/framer/v1/ts.proto

package v1

import (
	errors "github.com/synnaxlabs/x/errors"
	telem "github.com/synnaxlabs/x/telem"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IteratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command   int32              `protobuf:"varint,1,opt,name=command,proto3" json:"command,omitempty"`
	Stamp     int64              `protobuf:"varint,2,opt,name=stamp,proto3" json:"stamp,omitempty"`
	Span      int64              `protobuf:"varint,3,opt,name=span,proto3" json:"span,omitempty"`
	Bounds    *telem.PBTimeRange `protobuf:"bytes,4,opt,name=bounds,proto3" json:"bounds,omitempty"`
	Keys      []uint32           `protobuf:"varint,6,rep,packed,name=keys,proto3" json:"keys,omitempty"`
	ChunkSize int64              `protobuf:"varint,7,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

func (x *IteratorRequest) Reset() {
	*x = IteratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IteratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IteratorRequest) ProtoMessage() {}

func (x *IteratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IteratorRequest.ProtoReflect.Descriptor instead.
func (*IteratorRequest) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{0}
}

func (x *IteratorRequest) GetCommand() int32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *IteratorRequest) GetStamp() int64 {
	if x != nil {
		return x.Stamp
	}
	return 0
}

func (x *IteratorRequest) GetSpan() int64 {
	if x != nil {
		return x.Span
	}
	return 0
}

func (x *IteratorRequest) GetBounds() *telem.PBTimeRange {
	if x != nil {
		return x.Bounds
	}
	return nil
}

func (x *IteratorRequest) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *IteratorRequest) GetChunkSize() int64 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

type IteratorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variant int32             `protobuf:"varint,1,opt,name=variant,proto3" json:"variant,omitempty"`
	Command int32             `protobuf:"varint,2,opt,name=command,proto3" json:"command,omitempty"`
	Frame   *Frame            `protobuf:"bytes,3,opt,name=frame,proto3" json:"frame,omitempty"`
	NodeKey int32             `protobuf:"varint,43,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	Ack     bool              `protobuf:"varint,5,opt,name=ack,proto3" json:"ack,omitempty"`
	SeqNum  int32             `protobuf:"varint,6,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	Error   *errors.PBPayload `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *IteratorResponse) Reset() {
	*x = IteratorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IteratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IteratorResponse) ProtoMessage() {}

func (x *IteratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IteratorResponse.ProtoReflect.Descriptor instead.
func (*IteratorResponse) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{1}
}

func (x *IteratorResponse) GetVariant() int32 {
	if x != nil {
		return x.Variant
	}
	return 0
}

func (x *IteratorResponse) GetCommand() int32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *IteratorResponse) GetFrame() *Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *IteratorResponse) GetNodeKey() int32 {
	if x != nil {
		return x.NodeKey
	}
	return 0
}

func (x *IteratorResponse) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

func (x *IteratorResponse) GetSeqNum() int32 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

func (x *IteratorResponse) GetError() *errors.PBPayload {
	if x != nil {
		return x.Error
	}
	return nil
}

type RelayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []uint32 `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
}

func (x *RelayRequest) Reset() {
	*x = RelayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayRequest) ProtoMessage() {}

func (x *RelayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayRequest.ProtoReflect.Descriptor instead.
func (*RelayRequest) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{2}
}

func (x *RelayRequest) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

type RelayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame *Frame            `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	Error *errors.PBPayload `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RelayResponse) Reset() {
	*x = RelayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayResponse) ProtoMessage() {}

func (x *RelayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayResponse.ProtoReflect.Descriptor instead.
func (*RelayResponse) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{3}
}

func (x *RelayResponse) GetFrame() *Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *RelayResponse) GetError() *errors.PBPayload {
	if x != nil {
		return x.Error
	}
	return nil
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys   []uint32          `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
	Series []*telem.PBSeries `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{4}
}

func (x *Frame) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Frame) GetSeries() []*telem.PBSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

type WriterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command int32         `protobuf:"varint,1,opt,name=command,proto3" json:"command,omitempty"`
	Config  *WriterConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	Frame   *Frame        `protobuf:"bytes,3,opt,name=frame,proto3" json:"frame,omitempty"`
}

func (x *WriterRequest) Reset() {
	*x = WriterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriterRequest) ProtoMessage() {}

func (x *WriterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriterRequest.ProtoReflect.Descriptor instead.
func (*WriterRequest) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{5}
}

func (x *WriterRequest) GetCommand() int32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *WriterRequest) GetConfig() *WriterConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *WriterRequest) GetFrame() *Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

type WriterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys  []uint32 `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
	Start int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *WriterConfig) Reset() {
	*x = WriterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriterConfig) ProtoMessage() {}

func (x *WriterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriterConfig.ProtoReflect.Descriptor instead.
func (*WriterConfig) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{6}
}

func (x *WriterConfig) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *WriterConfig) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

type WriterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command   int32             `protobuf:"varint,1,opt,name=command,proto3" json:"command,omitempty"`
	Ack       bool              `protobuf:"varint,2,opt,name=ack,proto3" json:"ack,omitempty"`
	SeqNum    int32             `protobuf:"varint,3,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	NodeKey   int32             `protobuf:"varint,4,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	Error     *errors.PBPayload `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	TimeStamp int64             `protobuf:"varint,6,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
}

func (x *WriterResponse) Reset() {
	*x = WriterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriterResponse) ProtoMessage() {}

func (x *WriterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriterResponse.ProtoReflect.Descriptor instead.
func (*WriterResponse) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{7}
}

func (x *WriterResponse) GetCommand() int32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *WriterResponse) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

func (x *WriterResponse) GetSeqNum() int32 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

func (x *WriterResponse) GetNodeKey() int32 {
	if x != nil {
		return x.NodeKey
	}
	return 0
}

func (x *WriterResponse) GetError() *errors.PBPayload {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WriterResponse) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys   []uint32           `protobuf:"varint,1,rep,packed,name=keys,proto3" json:"keys,omitempty"`
	Names  []string           `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	Bounds *telem.PBTimeRange `protobuf:"bytes,3,opt,name=bounds,proto3" json:"bounds,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *DeleteRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *DeleteRequest) GetBounds() *telem.PBTimeRange {
	if x != nil {
		return x.Bounds
	}
	return nil
}

var File_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto protoreflect.FileDescriptor

var file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x79, 0x6e, 0x6e, 0x61, 0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x18, 0x78, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x78, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x2e, 0x50, 0x42, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x10, 0x49, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x05,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61,
	0x63, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x12, 0x27, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x50, 0x42, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x50, 0x42, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x50, 0x42, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x0d,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63,
	0x6b, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x50, 0x42,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x65, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x2e, 0x50, 0x42, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x73, 0x32, 0x53, 0x0a, 0x0f, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x49, 0x74, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x46, 0x0a, 0x0c, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x12, 0x13, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x32, 0x4b, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x47,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x14, 0x2e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x91, 0x01, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x07, 0x54, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e,
	0x6e, 0x61, 0x78, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x6e, 0x61, 0x78, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02,
	0x05, 0x54, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x05, 0x54, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x11, 0x54, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x06, 0x54, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescOnce sync.Once
	file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescData = file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDesc
)

func file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescGZIP() []byte {
	file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescOnce.Do(func() {
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescData = protoimpl.X.CompressGZIP(file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescData)
	})
	return file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDescData
}

var file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_goTypes = []any{
	(*IteratorRequest)(nil),   // 0: ts.v1.IteratorRequest
	(*IteratorResponse)(nil),  // 1: ts.v1.IteratorResponse
	(*RelayRequest)(nil),      // 2: ts.v1.RelayRequest
	(*RelayResponse)(nil),     // 3: ts.v1.RelayResponse
	(*Frame)(nil),             // 4: ts.v1.Frame
	(*WriterRequest)(nil),     // 5: ts.v1.WriterRequest
	(*WriterConfig)(nil),      // 6: ts.v1.WriterConfig
	(*WriterResponse)(nil),    // 7: ts.v1.WriterResponse
	(*DeleteRequest)(nil),     // 8: ts.v1.DeleteRequest
	(*telem.PBTimeRange)(nil), // 9: telem.PBTimeRange
	(*errors.PBPayload)(nil),  // 10: errors.PBPayload
	(*telem.PBSeries)(nil),    // 11: telem.PBSeries
	(*emptypb.Empty)(nil),     // 12: google.protobuf.Empty
}
var file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_depIdxs = []int32{
	9,  // 0: ts.v1.IteratorRequest.bounds:type_name -> telem.PBTimeRange
	4,  // 1: ts.v1.IteratorResponse.frame:type_name -> ts.v1.Frame
	10, // 2: ts.v1.IteratorResponse.error:type_name -> errors.PBPayload
	4,  // 3: ts.v1.RelayResponse.frame:type_name -> ts.v1.Frame
	10, // 4: ts.v1.RelayResponse.error:type_name -> errors.PBPayload
	11, // 5: ts.v1.Frame.series:type_name -> telem.PBSeries
	6,  // 6: ts.v1.WriterRequest.config:type_name -> ts.v1.WriterConfig
	4,  // 7: ts.v1.WriterRequest.frame:type_name -> ts.v1.Frame
	10, // 8: ts.v1.WriterResponse.error:type_name -> errors.PBPayload
	9,  // 9: ts.v1.DeleteRequest.bounds:type_name -> telem.PBTimeRange
	0,  // 10: ts.v1.IteratorService.Iterate:input_type -> ts.v1.IteratorRequest
	2,  // 11: ts.v1.RelayService.Relay:input_type -> ts.v1.RelayRequest
	5,  // 12: ts.v1.WriterService.Write:input_type -> ts.v1.WriterRequest
	8,  // 13: ts.v1.DeleteService.Exec:input_type -> ts.v1.DeleteRequest
	1,  // 14: ts.v1.IteratorService.Iterate:output_type -> ts.v1.IteratorResponse
	3,  // 15: ts.v1.RelayService.Relay:output_type -> ts.v1.RelayResponse
	7,  // 16: ts.v1.WriterService.Write:output_type -> ts.v1.WriterResponse
	12, // 17: ts.v1.DeleteService.Exec:output_type -> google.protobuf.Empty
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_init() }
func file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_init() {
	if File_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IteratorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IteratorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RelayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RelayResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Frame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WriterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*WriterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*WriterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_goTypes,
		DependencyIndexes: file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_depIdxs,
		MessageInfos:      file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_msgTypes,
	}.Build()
	File_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto = out.File
	file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_rawDesc = nil
	file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_goTypes = nil
	file_synnax_pkg_distribution_transport_grpc_framer_v1_ts_proto_depIdxs = nil
}
