// Code generated by protoc-gen-go. DO NOT EDIT.
// source: omnitelk.proto

/*
Package omnitelk is a generated protocol buffer package.

It is generated from these files:
	omnitelk.proto

It has these top-level messages:
	ExportRequest
	EncodedRecord
	ExportResponse
	ShardDefinition
	ShardingConfig
	ConfigRequest
*/
package omnitelk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExportResponse_ResultCode int32

const (
	// Successfully received and accepted.
	ExportResponse_SUCCESS ExportResponse_ResultCode = 0
	// Failed to process the data. Sender MUST NOT retry this request.
	ExportResponse_FAILED_NOT_RETRYABLE ExportResponse_ResultCode = 1
	// Failed to process the data. Sender SHOULD retry this request.
	ExportResponse_FAILED_RETRYABLE ExportResponse_ResultCode = 2
	// Sharding configuration at receiver does not match
	// sharding performed by sender. Sender MUST re-shard the
	// data according to new sharding configuration specified
	// in shardingConfig field and send again.
	ExportResponse_SHARD_CONFIG_MISTMATCH ExportResponse_ResultCode = 3
)

var ExportResponse_ResultCode_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED_NOT_RETRYABLE",
	2: "FAILED_RETRYABLE",
	3: "SHARD_CONFIG_MISTMATCH",
}
var ExportResponse_ResultCode_value = map[string]int32{
	"SUCCESS":                0,
	"FAILED_NOT_RETRYABLE":   1,
	"FAILED_RETRYABLE":       2,
	"SHARD_CONFIG_MISTMATCH": 3,
}

func (x ExportResponse_ResultCode) String() string {
	return proto.EnumName(ExportResponse_ResultCode_name, int32(x))
}
func (ExportResponse_ResultCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ExportRequest struct {
	// Unique request id.
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Encoded record.
	Record *EncodedRecord `protobuf:"bytes,2,opt,name=record" json:"record,omitempty"`
	// The shard that the records were prepared for. The receiver will
	// accept the records if the current sharding configuration has a
	// shard with specified shardId and matching hash key range.
	Shard *ShardDefinition `protobuf:"bytes,3,opt,name=shard" json:"shard,omitempty"`
}

func (m *ExportRequest) Reset()                    { *m = ExportRequest{} }
func (m *ExportRequest) String() string            { return proto.CompactTextString(m) }
func (*ExportRequest) ProtoMessage()               {}
func (*ExportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExportRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportRequest) GetRecord() *EncodedRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *ExportRequest) GetShard() *ShardDefinition {
	if m != nil {
		return m.Shard
	}
	return nil
}

type EncodedRecord struct {
	// data is a byte sequence encoded as follows:
	// 1. github.com/omnition/opencensus-go-exporter-kinesis/SpanList
	//    encoded in Protocol Buffer format.
	// 2. Result of step 1 compressed using gzip.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Partition key defines the shard that this data is intended for.
	PartitionKey string `protobuf:"bytes,2,opt,name=partitionKey" json:"partitionKey,omitempty"`
	// Number of spans encoded in this record. Useful for recipient
	// to maintain span count stats without uncompressing the data.
	SpanCount uint64 `protobuf:"varint,3,opt,name=spanCount" json:"spanCount,omitempty"`
	// Size of encoded but uncompressed data in bytes. Useful for recipient
	// to calculate compression ratio stats without uncompressing the data.
	UncompressedBytes uint64 `protobuf:"varint,4,opt,name=uncompressedBytes" json:"uncompressedBytes,omitempty"`
}

func (m *EncodedRecord) Reset()                    { *m = EncodedRecord{} }
func (m *EncodedRecord) String() string            { return proto.CompactTextString(m) }
func (*EncodedRecord) ProtoMessage()               {}
func (*EncodedRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EncodedRecord) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EncodedRecord) GetPartitionKey() string {
	if m != nil {
		return m.PartitionKey
	}
	return ""
}

func (m *EncodedRecord) GetSpanCount() uint64 {
	if m != nil {
		return m.SpanCount
	}
	return 0
}

func (m *EncodedRecord) GetUncompressedBytes() uint64 {
	if m != nil {
		return m.UncompressedBytes
	}
	return 0
}

type ExportResponse struct {
	// Request id to which this response corresponds.
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Result of request processing.
	ResultCode ExportResponse_ResultCode `protobuf:"varint,2,opt,name=resultCode,enum=omnitelk.ExportResponse_ResultCode" json:"resultCode,omitempty"`
	// If resultCode=SHARD_CONFIG_MISTMATCH then this field
	// must contain the new sharding configuration otherwise
	// it should be omitted and ignored.
	ShardingConfig *ShardingConfig `protobuf:"bytes,3,opt,name=shardingConfig" json:"shardingConfig,omitempty"`
}

func (m *ExportResponse) Reset()                    { *m = ExportResponse{} }
func (m *ExportResponse) String() string            { return proto.CompactTextString(m) }
func (*ExportResponse) ProtoMessage()               {}
func (*ExportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExportResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportResponse) GetResultCode() ExportResponse_ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ExportResponse_SUCCESS
}

func (m *ExportResponse) GetShardingConfig() *ShardingConfig {
	if m != nil {
		return m.ShardingConfig
	}
	return nil
}

type ShardDefinition struct {
	// The id of the shard.
	ShardId string `protobuf:"bytes,1,opt,name=shardId" json:"shardId,omitempty"`
	// Shard starting and ending hash keys.
	// hasKey is byte sequence representation of big.Int hash key.
	StartingHashKey []byte `protobuf:"bytes,2,opt,name=startingHashKey,proto3" json:"startingHashKey,omitempty"`
	EndingHashKey   []byte `protobuf:"bytes,3,opt,name=endingHashKey,proto3" json:"endingHashKey,omitempty"`
}

func (m *ShardDefinition) Reset()                    { *m = ShardDefinition{} }
func (m *ShardDefinition) String() string            { return proto.CompactTextString(m) }
func (*ShardDefinition) ProtoMessage()               {}
func (*ShardDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ShardDefinition) GetShardId() string {
	if m != nil {
		return m.ShardId
	}
	return ""
}

func (m *ShardDefinition) GetStartingHashKey() []byte {
	if m != nil {
		return m.StartingHashKey
	}
	return nil
}

func (m *ShardDefinition) GetEndingHashKey() []byte {
	if m != nil {
		return m.EndingHashKey
	}
	return nil
}

type ShardingConfig struct {
	// Sharding configuration is a list of shard definitions.
	ShardDefinitions []*ShardDefinition `protobuf:"bytes,1,rep,name=shardDefinitions" json:"shardDefinitions,omitempty"`
}

func (m *ShardingConfig) Reset()                    { *m = ShardingConfig{} }
func (m *ShardingConfig) String() string            { return proto.CompactTextString(m) }
func (*ShardingConfig) ProtoMessage()               {}
func (*ShardingConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ShardingConfig) GetShardDefinitions() []*ShardDefinition {
	if m != nil {
		return m.ShardDefinitions
	}
	return nil
}

type ConfigRequest struct {
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*ExportRequest)(nil), "omnitelk.ExportRequest")
	proto.RegisterType((*EncodedRecord)(nil), "omnitelk.EncodedRecord")
	proto.RegisterType((*ExportResponse)(nil), "omnitelk.ExportResponse")
	proto.RegisterType((*ShardDefinition)(nil), "omnitelk.ShardDefinition")
	proto.RegisterType((*ShardingConfig)(nil), "omnitelk.ShardingConfig")
	proto.RegisterType((*ConfigRequest)(nil), "omnitelk.ConfigRequest")
	proto.RegisterEnum("omnitelk.ExportResponse_ResultCode", ExportResponse_ResultCode_name, ExportResponse_ResultCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OmnitelK service

type OmnitelKClient interface {
	// Get the sharding configuration. Typically called at the beginning of
	// communication, before calling Export.
	GetShardingConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ShardingConfig, error)
	// Export a single request.
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
}

type omnitelKClient struct {
	cc *grpc.ClientConn
}

func NewOmnitelKClient(cc *grpc.ClientConn) OmnitelKClient {
	return &omnitelKClient{cc}
}

func (c *omnitelKClient) GetShardingConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ShardingConfig, error) {
	out := new(ShardingConfig)
	err := grpc.Invoke(ctx, "/omnitelk.OmnitelK/GetShardingConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omnitelKClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := grpc.Invoke(ctx, "/omnitelk.OmnitelK/Export", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OmnitelK service

type OmnitelKServer interface {
	// Get the sharding configuration. Typically called at the beginning of
	// communication, before calling Export.
	GetShardingConfig(context.Context, *ConfigRequest) (*ShardingConfig, error)
	// Export a single request.
	Export(context.Context, *ExportRequest) (*ExportResponse, error)
}

func RegisterOmnitelKServer(s *grpc.Server, srv OmnitelKServer) {
	s.RegisterService(&_OmnitelK_serviceDesc, srv)
}

func _OmnitelK_GetShardingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmnitelKServer).GetShardingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omnitelk.OmnitelK/GetShardingConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmnitelKServer).GetShardingConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmnitelK_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmnitelKServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omnitelk.OmnitelK/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmnitelKServer).Export(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OmnitelK_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omnitelk.OmnitelK",
	HandlerType: (*OmnitelKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShardingConfig",
			Handler:    _OmnitelK_GetShardingConfig_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _OmnitelK_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omnitelk.proto",
}

func init() { proto.RegisterFile("omnitelk.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x86, 0x92, 0x30, 0x01, 0xe3, 0x8c, 0xa2, 0xd6, 0x8d, 0x7a, 0x40, 0xdb, 0x1e, 0x38,
	0x54, 0x44, 0xa2, 0xe7, 0x4a, 0x25, 0xc6, 0x09, 0x28, 0x1f, 0x48, 0x6b, 0xaa, 0xaa, 0x27, 0xe4,
	0xb2, 0x1b, 0x62, 0x35, 0xd9, 0x75, 0xbd, 0x8b, 0xd4, 0xa8, 0xa7, 0xfe, 0x81, 0x5e, 0xf2, 0x87,
	0x2b, 0xd6, 0x7c, 0xd8, 0x46, 0xed, 0xcd, 0x7e, 0xf3, 0xde, 0xec, 0xd3, 0x9b, 0x19, 0x70, 0xe4,
	0xa3, 0x88, 0x35, 0x7f, 0xf8, 0xde, 0x4b, 0x52, 0xa9, 0x25, 0x1e, 0x6e, 0xfe, 0xc9, 0x6f, 0x0b,
	0x5a, 0xc1, 0xcf, 0x44, 0xa6, 0x9a, 0xf2, 0x1f, 0x4b, 0xae, 0x34, 0x3a, 0x60, 0xc7, 0xcc, 0xb3,
	0x3a, 0x56, 0xb7, 0x46, 0xed, 0x98, 0xe1, 0x19, 0xd4, 0x53, 0x3e, 0x97, 0x29, 0xf3, 0xec, 0x8e,
	0xd5, 0x3d, 0xea, 0xbf, 0xea, 0x6d, 0x9b, 0x05, 0x62, 0x2e, 0x19, 0x67, 0xd4, 0x94, 0xe9, 0x9a,
	0x86, 0x67, 0xf0, 0x42, 0xdd, 0x47, 0x29, 0xf3, 0xaa, 0x86, 0xff, 0x7a, 0xc7, 0x0f, 0x57, 0xf0,
	0x90, 0xdf, 0xc5, 0x22, 0xd6, 0xb1, 0x14, 0x34, 0xe3, 0x91, 0xe7, 0x95, 0x87, 0x7c, 0x2b, 0x44,
	0xa8, 0xb1, 0x48, 0x47, 0xc6, 0x45, 0x93, 0x9a, 0x6f, 0x24, 0xd0, 0x4c, 0xa2, 0x54, 0x1b, 0xe5,
	0x15, 0x7f, 0x32, 0x6e, 0x1a, 0xb4, 0x80, 0xe1, 0x1b, 0x68, 0xa8, 0x24, 0x12, 0xbe, 0x5c, 0x0a,
	0x6d, 0x9e, 0xaf, 0xd1, 0x1d, 0x80, 0xef, 0xe1, 0x78, 0x29, 0xe6, 0xf2, 0x31, 0x49, 0xb9, 0x52,
	0x9c, 0x9d, 0x3f, 0x69, 0xae, 0xbc, 0x9a, 0x61, 0xed, 0x17, 0xc8, 0x1f, 0x1b, 0x9c, 0x4d, 0x32,
	0x2a, 0x91, 0x42, 0xf1, 0xbd, 0x68, 0x7c, 0x80, 0x94, 0xab, 0xe5, 0x83, 0xf6, 0x25, 0xe3, 0xc6,
	0x90, 0xd3, 0x7f, 0x9b, 0x8b, 0xa7, 0xa0, 0xee, 0xd1, 0x2d, 0x95, 0xe6, 0x64, 0xf8, 0x09, 0x1c,
	0x13, 0x43, 0x2c, 0x16, 0xbe, 0x14, 0x77, 0xf1, 0x62, 0x9d, 0x9b, 0x57, 0xca, 0x6d, 0x5b, 0xa7,
	0x25, 0x3e, 0xe1, 0x00, 0xbb, 0xde, 0x78, 0x04, 0x07, 0xe1, 0x67, 0xdf, 0x0f, 0xc2, 0xd0, 0xad,
	0xa0, 0x07, 0x27, 0x17, 0x83, 0xf1, 0x75, 0x30, 0x9c, 0xdd, 0x4e, 0xa6, 0x33, 0x1a, 0x4c, 0xe9,
	0xd7, 0xc1, 0xf9, 0x75, 0xe0, 0x5a, 0x78, 0x02, 0xee, 0xba, 0xb2, 0x43, 0x6d, 0x3c, 0x85, 0x97,
	0xe1, 0x68, 0x40, 0x87, 0x33, 0x7f, 0x72, 0x7b, 0x31, 0xbe, 0x9c, 0xdd, 0x8c, 0xc3, 0xe9, 0xcd,
	0x60, 0xea, 0x8f, 0xdc, 0x2a, 0xf9, 0x05, 0xed, 0xd2, 0x00, 0xd1, 0x83, 0x03, 0xe3, 0x65, 0x9c,
	0xa5, 0xd2, 0xa0, 0x9b, 0x5f, 0xec, 0x42, 0x5b, 0xe9, 0xd5, 0x68, 0xc4, 0x62, 0x14, 0xa9, 0xfb,
	0xcd, 0xc0, 0x9a, 0xb4, 0x0c, 0xe3, 0x3b, 0x68, 0x71, 0xc1, 0x72, 0xbc, 0xaa, 0xe1, 0x15, 0x41,
	0xf2, 0x05, 0x9c, 0x62, 0x0a, 0x18, 0x80, 0xab, 0x8a, 0x76, 0x94, 0x67, 0x75, 0xaa, 0xff, 0xdf,
	0xb8, 0x3d, 0x09, 0x69, 0x43, 0x6b, 0x1d, 0x6b, 0xb6, 0xff, 0xfd, 0x67, 0x0b, 0x0e, 0x27, 0x99,
	0xfe, 0x0a, 0x47, 0x70, 0x7c, 0xc9, 0x75, 0xe9, 0xe5, 0xdc, 0x05, 0x14, 0xa4, 0xa7, 0xff, 0x1c,
	0x19, 0xa9, 0xe0, 0x47, 0xa8, 0x67, 0xfb, 0x90, 0x97, 0x17, 0x2e, 0x2f, 0x2f, 0x2f, 0xae, 0x0e,
	0xa9, 0x7c, 0xab, 0x9b, 0xc3, 0xfd, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x44, 0xd5, 0x07, 0x5f,
	0xca, 0x03, 0x00, 0x00,
}
