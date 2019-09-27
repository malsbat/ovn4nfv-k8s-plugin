// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nfn.proto

package nfn

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SubscribeContext struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeContext) Reset()         { *m = SubscribeContext{} }
func (m *SubscribeContext) String() string { return proto.CompactTextString(m) }
func (*SubscribeContext) ProtoMessage()    {}
func (*SubscribeContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b809db4a7814953, []int{0}
}

func (m *SubscribeContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeContext.Unmarshal(m, b)
}
func (m *SubscribeContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeContext.Marshal(b, m, deterministic)
}
func (m *SubscribeContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeContext.Merge(m, src)
}
func (m *SubscribeContext) XXX_Size() int {
	return xxx_messageInfo_SubscribeContext.Size(m)
}
func (m *SubscribeContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeContext.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeContext proto.InternalMessageInfo

func (m *SubscribeContext) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Notification struct {
	CniType string `protobuf:"bytes,1,opt,name=cni_type,json=cniType,proto3" json:"cni_type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Notification_InSync
	//	*Notification_ProviderNwCreate
	//	*Notification_ProviderNwRemove
	Payload              isNotification_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b809db4a7814953, []int{1}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetCniType() string {
	if m != nil {
		return m.CniType
	}
	return ""
}

type isNotification_Payload interface {
	isNotification_Payload()
}

type Notification_InSync struct {
	InSync *InSync `protobuf:"bytes,2,opt,name=in_sync,json=inSync,proto3,oneof"`
}

type Notification_ProviderNwCreate struct {
	ProviderNwCreate *ProviderNetworkCreate `protobuf:"bytes,3,opt,name=provider_nw_create,json=providerNwCreate,proto3,oneof"`
}

type Notification_ProviderNwRemove struct {
	ProviderNwRemove *ProviderNetworkRemove `protobuf:"bytes,4,opt,name=provider_nw_remove,json=providerNwRemove,proto3,oneof"`
}

func (*Notification_InSync) isNotification_Payload() {}

func (*Notification_ProviderNwCreate) isNotification_Payload() {}

func (*Notification_ProviderNwRemove) isNotification_Payload() {}

func (m *Notification) GetPayload() isNotification_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Notification) GetInSync() *InSync {
	if x, ok := m.GetPayload().(*Notification_InSync); ok {
		return x.InSync
	}
	return nil
}

func (m *Notification) GetProviderNwCreate() *ProviderNetworkCreate {
	if x, ok := m.GetPayload().(*Notification_ProviderNwCreate); ok {
		return x.ProviderNwCreate
	}
	return nil
}

func (m *Notification) GetProviderNwRemove() *ProviderNetworkRemove {
	if x, ok := m.GetPayload().(*Notification_ProviderNwRemove); ok {
		return x.ProviderNwRemove
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Notification) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Notification_InSync)(nil),
		(*Notification_ProviderNwCreate)(nil),
		(*Notification_ProviderNwRemove)(nil),
	}
}

type ProviderNetworkCreate struct {
	ProviderNwName       string    `protobuf:"bytes,1,opt,name=provider_nw_name,json=providerNwName,proto3" json:"provider_nw_name,omitempty"`
	Vlan                 *VlanInfo `protobuf:"bytes,2,opt,name=vlan,proto3" json:"vlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProviderNetworkCreate) Reset()         { *m = ProviderNetworkCreate{} }
func (m *ProviderNetworkCreate) String() string { return proto.CompactTextString(m) }
func (*ProviderNetworkCreate) ProtoMessage()    {}
func (*ProviderNetworkCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b809db4a7814953, []int{2}
}

func (m *ProviderNetworkCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderNetworkCreate.Unmarshal(m, b)
}
func (m *ProviderNetworkCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderNetworkCreate.Marshal(b, m, deterministic)
}
func (m *ProviderNetworkCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderNetworkCreate.Merge(m, src)
}
func (m *ProviderNetworkCreate) XXX_Size() int {
	return xxx_messageInfo_ProviderNetworkCreate.Size(m)
}
func (m *ProviderNetworkCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderNetworkCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderNetworkCreate proto.InternalMessageInfo

func (m *ProviderNetworkCreate) GetProviderNwName() string {
	if m != nil {
		return m.ProviderNwName
	}
	return ""
}

func (m *ProviderNetworkCreate) GetVlan() *VlanInfo {
	if m != nil {
		return m.Vlan
	}
	return nil
}

type ProviderNetworkRemove struct {
	ProviderNwName       string   `protobuf:"bytes,1,opt,name=provider_nw_name,json=providerNwName,proto3" json:"provider_nw_name,omitempty"`
	VlanLogicalIntf      string   `protobuf:"bytes,2,opt,name=vlan_logical_intf,json=vlanLogicalIntf,proto3" json:"vlan_logical_intf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderNetworkRemove) Reset()         { *m = ProviderNetworkRemove{} }
func (m *ProviderNetworkRemove) String() string { return proto.CompactTextString(m) }
func (*ProviderNetworkRemove) ProtoMessage()    {}
func (*ProviderNetworkRemove) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b809db4a7814953, []int{3}
}

func (m *ProviderNetworkRemove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderNetworkRemove.Unmarshal(m, b)
}
func (m *ProviderNetworkRemove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderNetworkRemove.Marshal(b, m, deterministic)
}
func (m *ProviderNetworkRemove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderNetworkRemove.Merge(m, src)
}
func (m *ProviderNetworkRemove) XXX_Size() int {
	return xxx_messageInfo_ProviderNetworkRemove.Size(m)
}
func (m *ProviderNetworkRemove) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderNetworkRemove.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderNetworkRemove proto.InternalMessageInfo

func (m *ProviderNetworkRemove) GetProviderNwName() string {
	if m != nil {
		return m.ProviderNwName
	}
	return ""
}

func (m *ProviderNetworkRemove) GetVlanLogicalIntf() string {
	if m != nil {
		return m.VlanLogicalIntf
	}
	return ""
}

type VlanInfo struct {
	VlanId               string   `protobuf:"bytes,1,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	ProviderIntf         string   `protobuf:"bytes,2,opt,name=provider_intf,json=providerIntf,proto3" json:"provider_intf,omitempty"`
	LogicalIntf          string   `protobuf:"bytes,3,opt,name=logical_intf,json=logicalIntf,proto3" json:"logical_intf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlanInfo) Reset()         { *m = VlanInfo{} }
func (m *VlanInfo) String() string { return proto.CompactTextString(m) }
func (*VlanInfo) ProtoMessage()    {}
func (*VlanInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b809db4a7814953, []int{4}
}

func (m *VlanInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanInfo.Unmarshal(m, b)
}
func (m *VlanInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanInfo.Marshal(b, m, deterministic)
}
func (m *VlanInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanInfo.Merge(m, src)
}
func (m *VlanInfo) XXX_Size() int {
	return xxx_messageInfo_VlanInfo.Size(m)
}
func (m *VlanInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VlanInfo proto.InternalMessageInfo

func (m *VlanInfo) GetVlanId() string {
	if m != nil {
		return m.VlanId
	}
	return ""
}

func (m *VlanInfo) GetProviderIntf() string {
	if m != nil {
		return m.ProviderIntf
	}
	return ""
}

func (m *VlanInfo) GetLogicalIntf() string {
	if m != nil {
		return m.LogicalIntf
	}
	return ""
}

type InSync struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InSync) Reset()         { *m = InSync{} }
func (m *InSync) String() string { return proto.CompactTextString(m) }
func (*InSync) ProtoMessage()    {}
func (*InSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b809db4a7814953, []int{5}
}

func (m *InSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InSync.Unmarshal(m, b)
}
func (m *InSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InSync.Marshal(b, m, deterministic)
}
func (m *InSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InSync.Merge(m, src)
}
func (m *InSync) XXX_Size() int {
	return xxx_messageInfo_InSync.Size(m)
}
func (m *InSync) XXX_DiscardUnknown() {
	xxx_messageInfo_InSync.DiscardUnknown(m)
}

var xxx_messageInfo_InSync proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SubscribeContext)(nil), "SubscribeContext")
	proto.RegisterType((*Notification)(nil), "Notification")
	proto.RegisterType((*ProviderNetworkCreate)(nil), "ProviderNetworkCreate")
	proto.RegisterType((*ProviderNetworkRemove)(nil), "ProviderNetworkRemove")
	proto.RegisterType((*VlanInfo)(nil), "VlanInfo")
	proto.RegisterType((*InSync)(nil), "InSync")
}

func init() { proto.RegisterFile("nfn.proto", fileDescriptor_5b809db4a7814953) }

var fileDescriptor_5b809db4a7814953 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0xaf, 0xd2, 0x40,
	0x14, 0xc5, 0xa9, 0x90, 0xfe, 0xb9, 0x80, 0xc2, 0x24, 0x6a, 0xd5, 0x98, 0x60, 0xdd, 0x10, 0x17,
	0xc5, 0xe0, 0xd6, 0x95, 0x24, 0x86, 0x26, 0xa6, 0x31, 0xc5, 0xb8, 0xad, 0xc3, 0x74, 0x6a, 0x26,
	0xb4, 0x77, 0x9a, 0x32, 0x82, 0xfd, 0xc6, 0xef, 0x63, 0xbc, 0x74, 0x4a, 0x79, 0x85, 0xc7, 0xe6,
	0xed, 0xda, 0x73, 0xee, 0x9c, 0xdf, 0xf4, 0xf6, 0x80, 0x83, 0x29, 0xfa, 0x45, 0x29, 0x95, 0xf4,
	0x16, 0x30, 0xd9, 0xfc, 0xdb, 0xee, 0x59, 0x29, 0xb6, 0x7c, 0x25, 0x51, 0xf1, 0xff, 0x8a, 0xbc,
	0x03, 0x07, 0x65, 0xc2, 0x63, 0xa4, 0x39, 0x77, 0x8d, 0x99, 0x31, 0x77, 0x22, 0xbb, 0x16, 0x42,
	0x9a, 0x73, 0xef, 0xce, 0x80, 0x51, 0x28, 0x95, 0x48, 0x05, 0xa3, 0x4a, 0x48, 0x24, 0x6f, 0xc0,
	0x66, 0x28, 0x62, 0x55, 0x15, 0xed, 0xb0, 0xc5, 0x50, 0xfc, 0xaa, 0x0a, 0x4e, 0x3c, 0xb0, 0x04,
	0xc6, 0xfb, 0x0a, 0x99, 0xfb, 0x6c, 0x66, 0xcc, 0x87, 0x4b, 0xcb, 0x0f, 0x70, 0x53, 0x21, 0x5b,
	0xf7, 0x22, 0x53, 0xe8, 0x27, 0xf2, 0x1d, 0x48, 0x51, 0xca, 0x83, 0x48, 0x78, 0x19, 0xe3, 0x31,
	0x66, 0x25, 0xa7, 0x8a, 0xbb, 0x7d, 0x3d, 0xfe, 0xca, 0xff, 0x79, 0xb2, 0x42, 0xae, 0x8e, 0xb2,
	0xdc, 0xad, 0xb4, 0xbb, 0xee, 0x45, 0x93, 0xf6, 0x4c, 0x78, 0x6c, 0xb4, 0xeb, 0x9c, 0x92, 0xe7,
	0xf2, 0xc0, 0xdd, 0xc1, 0xed, 0x9c, 0x48, 0xbb, 0x97, 0x39, 0x8d, 0xf6, 0xcd, 0x01, 0xab, 0xa0,
	0x55, 0x26, 0x69, 0xe2, 0xfd, 0x81, 0x97, 0x37, 0xf9, 0x64, 0x0e, 0x93, 0x2e, 0xab, 0xb3, 0xa7,
	0xe7, 0x0f, 0x79, 0xf5, 0xb6, 0xc8, 0x7b, 0x18, 0x1c, 0x32, 0x8a, 0xa7, 0xcf, 0x77, 0xfc, 0xdf,
	0x19, 0xc5, 0x00, 0x53, 0x19, 0x69, 0xd9, 0xcb, 0x1f, 0x11, 0x9a, 0x5b, 0x3c, 0x81, 0xf0, 0x09,
	0xa6, 0x75, 0x54, 0x9c, 0xc9, 0xbf, 0x82, 0xd1, 0x2c, 0x16, 0xa8, 0x52, 0x8d, 0x73, 0xa2, 0x17,
	0xb5, 0xf1, 0xa3, 0xd1, 0x03, 0x54, 0xa9, 0xb7, 0x03, 0xbb, 0xbd, 0x00, 0x79, 0x0d, 0x96, 0x3e,
	0x27, 0x92, 0x53, 0xb0, 0x59, 0xbf, 0x06, 0x09, 0xf9, 0x08, 0xe3, 0x33, 0xba, 0x13, 0x36, 0x6a,
	0xc5, 0x3a, 0x89, 0x7c, 0x80, 0xd1, 0x05, 0xb0, 0xaf, 0x67, 0x86, 0x59, 0x07, 0x66, 0x83, 0xd9,
	0xfc, 0xec, 0xe5, 0x57, 0x5d, 0x38, 0x5d, 0x9a, 0x8a, 0x2c, 0xc0, 0x39, 0x17, 0x8e, 0x4c, 0xfd,
	0xeb, 0xf2, 0xbd, 0x1d, 0xfb, 0xdd, 0x76, 0x7d, 0x36, 0xb6, 0xa6, 0x2e, 0xea, 0x97, 0xfb, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa6, 0xd9, 0x3e, 0x9e, 0xb5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NfnNotifyClient is the client API for NfnNotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NfnNotifyClient interface {
	Subscribe(ctx context.Context, in *SubscribeContext, opts ...grpc.CallOption) (NfnNotify_SubscribeClient, error)
}

type nfnNotifyClient struct {
	cc *grpc.ClientConn
}

func NewNfnNotifyClient(cc *grpc.ClientConn) NfnNotifyClient {
	return &nfnNotifyClient{cc}
}

func (c *nfnNotifyClient) Subscribe(ctx context.Context, in *SubscribeContext, opts ...grpc.CallOption) (NfnNotify_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NfnNotify_serviceDesc.Streams[0], "/nfnNotify/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &nfnNotifySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NfnNotify_SubscribeClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type nfnNotifySubscribeClient struct {
	grpc.ClientStream
}

func (x *nfnNotifySubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NfnNotifyServer is the server API for NfnNotify service.
type NfnNotifyServer interface {
	Subscribe(*SubscribeContext, NfnNotify_SubscribeServer) error
}

// UnimplementedNfnNotifyServer can be embedded to have forward compatible implementations.
type UnimplementedNfnNotifyServer struct {
}

func (*UnimplementedNfnNotifyServer) Subscribe(req *SubscribeContext, srv NfnNotify_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterNfnNotifyServer(s *grpc.Server, srv NfnNotifyServer) {
	s.RegisterService(&_NfnNotify_serviceDesc, srv)
}

func _NfnNotify_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeContext)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NfnNotifyServer).Subscribe(m, &nfnNotifySubscribeServer{stream})
}

type NfnNotify_SubscribeServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type nfnNotifySubscribeServer struct {
	grpc.ServerStream
}

func (x *nfnNotifySubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

var _NfnNotify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nfnNotify",
	HandlerType: (*NfnNotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NfnNotify_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nfn.proto",
}
