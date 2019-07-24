// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/squash/api/v1/debug_attachment.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DebugAttachment_State int32

const (
	// Newly created DebugAttachments have state RequestingAttachment
	DebugAttachment_RequestingAttachment DebugAttachment_State = 0
	// When the event loop begins fullfilling an attachment request it sets
	// DebugAttachments state to PendingAttachment
	DebugAttachment_PendingAttachment DebugAttachment_State = 1
	// When squash client successfully attaches, it sets state to Attached
	DebugAttachment_Attached DebugAttachment_State = 2
	// Indicates that user has requested an attachment be removed
	DebugAttachment_RequestingDelete DebugAttachment_State = 3
	// When the event loop begins fullfilling a delete request it sets this
	// status and triggers a cleanup routine
	// When the cleanup routine completes, it deletes the CRD
	DebugAttachment_PendingDelete DebugAttachment_State = 4
)

var DebugAttachment_State_name = map[int32]string{
	0: "RequestingAttachment",
	1: "PendingAttachment",
	2: "Attached",
	3: "RequestingDelete",
	4: "PendingDelete",
}

var DebugAttachment_State_value = map[string]int32{
	"RequestingAttachment": 0,
	"PendingAttachment":    1,
	"Attached":             2,
	"RequestingDelete":     3,
	"PendingDelete":        4,
}

func (x DebugAttachment_State) String() string {
	return proto.EnumName(DebugAttachment_State_name, int32(x))
}

func (DebugAttachment_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f76a2adbe78506d, []int{0, 0}
}

//
//Attachments store the information needed for squash to coordinate a debugging session
type DebugAttachment struct {
	Metadata             core.Metadata         `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	Status               core.Status           `protobuf:"bytes,2,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	PlankName            string                `protobuf:"bytes,3,opt,name=plank_name,json=plankName,proto3" json:"plank_name,omitempty"`
	Debugger             string                `protobuf:"bytes,4,opt,name=debugger,proto3" json:"debugger,omitempty"`
	Image                string                `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	ProcessName          string                `protobuf:"bytes,6,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Node                 string                `protobuf:"bytes,7,opt,name=node,proto3" json:"node,omitempty"`
	MatchRequest         bool                  `protobuf:"varint,8,opt,name=match_request,json=matchRequest,proto3" json:"match_request,omitempty"`
	DebugServerAddress   string                `protobuf:"bytes,9,opt,name=debug_server_address,json=debugServerAddress,proto3" json:"debug_server_address,omitempty"`
	Pod                  string                `protobuf:"bytes,11,opt,name=pod,proto3" json:"pod,omitempty"`
	Container            string                `protobuf:"bytes,12,opt,name=container,proto3" json:"container,omitempty"`
	DebugNamespace       string                `protobuf:"bytes,13,opt,name=debug_namespace,json=debugNamespace,proto3" json:"debug_namespace,omitempty"`
	State                DebugAttachment_State `protobuf:"varint,20,opt,name=state,proto3,enum=squash.solo.io.DebugAttachment_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DebugAttachment) Reset()         { *m = DebugAttachment{} }
func (m *DebugAttachment) String() string { return proto.CompactTextString(m) }
func (*DebugAttachment) ProtoMessage()    {}
func (*DebugAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f76a2adbe78506d, []int{0}
}
func (m *DebugAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugAttachment.Unmarshal(m, b)
}
func (m *DebugAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugAttachment.Marshal(b, m, deterministic)
}
func (m *DebugAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugAttachment.Merge(m, src)
}
func (m *DebugAttachment) XXX_Size() int {
	return xxx_messageInfo_DebugAttachment.Size(m)
}
func (m *DebugAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_DebugAttachment proto.InternalMessageInfo

func (m *DebugAttachment) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *DebugAttachment) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *DebugAttachment) GetPlankName() string {
	if m != nil {
		return m.PlankName
	}
	return ""
}

func (m *DebugAttachment) GetDebugger() string {
	if m != nil {
		return m.Debugger
	}
	return ""
}

func (m *DebugAttachment) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DebugAttachment) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *DebugAttachment) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *DebugAttachment) GetMatchRequest() bool {
	if m != nil {
		return m.MatchRequest
	}
	return false
}

func (m *DebugAttachment) GetDebugServerAddress() string {
	if m != nil {
		return m.DebugServerAddress
	}
	return ""
}

func (m *DebugAttachment) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *DebugAttachment) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *DebugAttachment) GetDebugNamespace() string {
	if m != nil {
		return m.DebugNamespace
	}
	return ""
}

func (m *DebugAttachment) GetState() DebugAttachment_State {
	if m != nil {
		return m.State
	}
	return DebugAttachment_RequestingAttachment
}

// Describes the user's debug intentions
type Intent struct {
	// type of debugger to use
	Debugger string `protobuf:"bytes,1,opt,name=debugger,proto3" json:"debugger,omitempty"`
	// pod to debug
	Pod *core.ResourceRef `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	// name of container to debug
	ContainerName string `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	// NOT YET IMPLEMENTED
	// if a container has multiple processes and you do not want to debug the first process, this string is used to select a specific process
	ProcessMatcher       string   `protobuf:"bytes,4,opt,name=process_matcher,json=processMatcher,proto3" json:"process_matcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Intent) Reset()         { *m = Intent{} }
func (m *Intent) String() string { return proto.CompactTextString(m) }
func (*Intent) ProtoMessage()    {}
func (*Intent) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f76a2adbe78506d, []int{1}
}
func (m *Intent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Intent.Unmarshal(m, b)
}
func (m *Intent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Intent.Marshal(b, m, deterministic)
}
func (m *Intent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Intent.Merge(m, src)
}
func (m *Intent) XXX_Size() int {
	return xxx_messageInfo_Intent.Size(m)
}
func (m *Intent) XXX_DiscardUnknown() {
	xxx_messageInfo_Intent.DiscardUnknown(m)
}

var xxx_messageInfo_Intent proto.InternalMessageInfo

func (m *Intent) GetDebugger() string {
	if m != nil {
		return m.Debugger
	}
	return ""
}

func (m *Intent) GetPod() *core.ResourceRef {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *Intent) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *Intent) GetProcessMatcher() string {
	if m != nil {
		return m.ProcessMatcher
	}
	return ""
}

// Describes the pod squash spawns for managing a particular debug session
type Plank struct {
	// plank pod reference
	Pod *core.ResourceRef `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	// indicates when plank has completed the debugger-specify preparation
	ReadyForConnect      bool     `protobuf:"varint,2,opt,name=ready_for_connect,json=readyForConnect,proto3" json:"ready_for_connect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plank) Reset()         { *m = Plank{} }
func (m *Plank) String() string { return proto.CompactTextString(m) }
func (*Plank) ProtoMessage()    {}
func (*Plank) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f76a2adbe78506d, []int{2}
}
func (m *Plank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plank.Unmarshal(m, b)
}
func (m *Plank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plank.Marshal(b, m, deterministic)
}
func (m *Plank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plank.Merge(m, src)
}
func (m *Plank) XXX_Size() int {
	return xxx_messageInfo_Plank.Size(m)
}
func (m *Plank) XXX_DiscardUnknown() {
	xxx_messageInfo_Plank.DiscardUnknown(m)
}

var xxx_messageInfo_Plank proto.InternalMessageInfo

func (m *Plank) GetPod() *core.ResourceRef {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *Plank) GetReadyForConnect() bool {
	if m != nil {
		return m.ReadyForConnect
	}
	return false
}

// Contains port information needed to connect or find a debugger
type PortSpec struct {
	// Types that are valid to be assigned to PortLocation:
	//	*PortSpec_Plank
	//	*PortSpec_Target
	PortLocation         isPortSpec_PortLocation `protobuf_oneof:"port_location"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PortSpec) Reset()         { *m = PortSpec{} }
func (m *PortSpec) String() string { return proto.CompactTextString(m) }
func (*PortSpec) ProtoMessage()    {}
func (*PortSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f76a2adbe78506d, []int{3}
}
func (m *PortSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortSpec.Unmarshal(m, b)
}
func (m *PortSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortSpec.Marshal(b, m, deterministic)
}
func (m *PortSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortSpec.Merge(m, src)
}
func (m *PortSpec) XXX_Size() int {
	return xxx_messageInfo_PortSpec.Size(m)
}
func (m *PortSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PortSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PortSpec proto.InternalMessageInfo

type isPortSpec_PortLocation interface {
	isPortSpec_PortLocation()
	Equal(interface{}) bool
}

type PortSpec_Plank struct {
	Plank string `protobuf:"bytes,1,opt,name=plank,proto3,oneof"`
}
type PortSpec_Target struct {
	Target string `protobuf:"bytes,2,opt,name=target,proto3,oneof"`
}

func (*PortSpec_Plank) isPortSpec_PortLocation()  {}
func (*PortSpec_Target) isPortSpec_PortLocation() {}

func (m *PortSpec) GetPortLocation() isPortSpec_PortLocation {
	if m != nil {
		return m.PortLocation
	}
	return nil
}

func (m *PortSpec) GetPlank() string {
	if x, ok := m.GetPortLocation().(*PortSpec_Plank); ok {
		return x.Plank
	}
	return ""
}

func (m *PortSpec) GetTarget() string {
	if x, ok := m.GetPortLocation().(*PortSpec_Target); ok {
		return x.Target
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PortSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PortSpec_OneofMarshaler, _PortSpec_OneofUnmarshaler, _PortSpec_OneofSizer, []interface{}{
		(*PortSpec_Plank)(nil),
		(*PortSpec_Target)(nil),
	}
}

func _PortSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PortSpec)
	// port_location
	switch x := m.PortLocation.(type) {
	case *PortSpec_Plank:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Plank)
	case *PortSpec_Target:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Target)
	case nil:
	default:
		return fmt.Errorf("PortSpec.PortLocation has unexpected type %T", x)
	}
	return nil
}

func _PortSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PortSpec)
	switch tag {
	case 1: // port_location.plank
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortLocation = &PortSpec_Plank{x}
		return true, err
	case 2: // port_location.target
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortLocation = &PortSpec_Target{x}
		return true, err
	default:
		return false, nil
	}
}

func _PortSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PortSpec)
	// port_location
	switch x := m.PortLocation.(type) {
	case *PortSpec_Plank:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Plank)))
		n += len(x.Plank)
	case *PortSpec_Target:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Target)))
		n += len(x.Target)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("squash.solo.io.DebugAttachment_State", DebugAttachment_State_name, DebugAttachment_State_value)
	proto.RegisterType((*DebugAttachment)(nil), "squash.solo.io.DebugAttachment")
	proto.RegisterType((*Intent)(nil), "squash.solo.io.Intent")
	proto.RegisterType((*Plank)(nil), "squash.solo.io.Plank")
	proto.RegisterType((*PortSpec)(nil), "squash.solo.io.PortSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/squash/api/v1/debug_attachment.proto", fileDescriptor_1f76a2adbe78506d)
}

var fileDescriptor_1f76a2adbe78506d = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x90, 0xe4, 0x3a, 0x87, 0xfc, 0x8e, 0x72, 0xd1, 0x80, 0xee, 0x2d, 0xa9, 0x5b, 0x44,
	0x44, 0x5b, 0xa7, 0xd0, 0x45, 0x2b, 0xba, 0x22, 0x45, 0xfd, 0x59, 0x80, 0x90, 0xd9, 0x75, 0x93,
	0x4e, 0xec, 0x13, 0xc7, 0x22, 0xf1, 0x98, 0x99, 0x09, 0x52, 0xb7, 0x3c, 0x48, 0xa5, 0xee, 0xfa,
	0x28, 0x7d, 0x0a, 0x16, 0x7d, 0x03, 0xfa, 0x04, 0x95, 0x67, 0x6c, 0x43, 0x50, 0x2b, 0xd1, 0x55,
	0xe6, 0x7c, 0xdf, 0xf9, 0x4e, 0xce, 0x9c, 0xf3, 0x79, 0xe0, 0x65, 0x18, 0xa9, 0xc9, 0x7c, 0xe4,
	0xfa, 0x7c, 0xd6, 0x97, 0x7c, 0xca, 0x9f, 0x45, 0xbc, 0x2f, 0xcf, 0xe7, 0x4c, 0x4e, 0xfa, 0x2c,
	0x89, 0xfa, 0x17, 0xbb, 0xfd, 0x00, 0x47, 0xf3, 0x70, 0xc8, 0x94, 0x62, 0xfe, 0x64, 0x86, 0xb1,
	0x72, 0x13, 0xc1, 0x15, 0x27, 0x0d, 0x93, 0xe5, 0xa6, 0x22, 0x37, 0xe2, 0x1b, 0x9d, 0x90, 0x87,
	0x5c, 0x53, 0xfd, 0xf4, 0x64, 0xb2, 0x36, 0x76, 0x7f, 0x57, 0x3e, 0xfd, 0x3d, 0x8b, 0x54, 0xfe,
	0x07, 0x33, 0x54, 0x2c, 0x60, 0x8a, 0x65, 0x92, 0xfe, 0x3d, 0x24, 0x52, 0x31, 0x35, 0x97, 0x99,
	0xe0, 0xe9, 0x3d, 0x04, 0x02, 0xc7, 0x7f, 0xd1, 0x51, 0x1e, 0x1b, 0x89, 0xf3, 0xa5, 0x0c, 0xcd,
	0xc3, 0x74, 0x0a, 0x07, 0xc5, 0x10, 0xc8, 0x2b, 0xb0, 0xf3, 0xbe, 0xa9, 0xd5, 0xb5, 0x7a, 0xab,
	0x7b, 0x6b, 0xae, 0xcf, 0x05, 0xe6, 0xf3, 0x70, 0x8f, 0x32, 0x76, 0x50, 0xfa, 0x7e, 0xb5, 0xb9,
	0xe4, 0x15, 0xd9, 0xe4, 0x1d, 0x54, 0x4c, 0xfb, 0x74, 0x59, 0xeb, 0x3a, 0x8b, 0xba, 0x53, 0xcd,
	0x0d, 0xd6, 0x53, 0xd5, 0xcf, 0xab, 0xcd, 0xb6, 0x42, 0xa9, 0x82, 0x68, 0x3c, 0xde, 0x77, 0xa2,
	0x30, 0xe6, 0x02, 0x1d, 0x2f, 0x93, 0x93, 0xff, 0x01, 0x92, 0x29, 0x8b, 0xcf, 0x86, 0x31, 0x9b,
	0x21, 0x5d, 0xe9, 0x5a, 0xbd, 0xaa, 0x57, 0xd5, 0xc8, 0x31, 0x9b, 0x21, 0xd9, 0x00, 0x5b, 0xaf,
	0x2e, 0x44, 0x41, 0x4b, 0x9a, 0x2c, 0x62, 0xd2, 0x81, 0x72, 0x34, 0x63, 0x21, 0xd2, 0xb2, 0x26,
	0x4c, 0x40, 0x1e, 0x42, 0x2d, 0x11, 0xdc, 0x47, 0x29, 0x4d, 0xc9, 0x8a, 0x26, 0x57, 0x33, 0x4c,
	0x17, 0x25, 0x50, 0x8a, 0x79, 0x80, 0xf4, 0x1f, 0x4d, 0xe9, 0x33, 0x79, 0x04, 0xf5, 0x19, 0x53,
	0xfe, 0x64, 0x28, 0xf0, 0x7c, 0x8e, 0x52, 0x51, 0xbb, 0x6b, 0xf5, 0x6c, 0xaf, 0xa6, 0x41, 0xcf,
	0x60, 0xe4, 0x39, 0x74, 0x8c, 0x91, 0x24, 0x8a, 0x0b, 0x14, 0x43, 0x16, 0x04, 0x02, 0xa5, 0xa4,
	0x55, 0x5d, 0x88, 0x68, 0xee, 0x54, 0x53, 0x07, 0x86, 0x21, 0x2d, 0x58, 0x49, 0x78, 0x40, 0x57,
	0x75, 0x42, 0x7a, 0x24, 0xff, 0x41, 0xd5, 0xe7, 0xb1, 0x62, 0x51, 0x8c, 0x82, 0xd6, 0xcc, 0x7d,
	0x0b, 0x80, 0x6c, 0x43, 0xd3, 0xfc, 0x43, 0xda, 0xbb, 0x4c, 0x98, 0x8f, 0xb4, 0xae, 0x73, 0x1a,
	0x1a, 0x3e, 0xce, 0x51, 0xf2, 0x1a, 0xca, 0xe9, 0x04, 0x91, 0x76, 0xba, 0x56, 0xaf, 0xb1, 0xb7,
	0xe5, 0x2e, 0x3a, 0xd9, 0xbd, 0xb3, 0x6a, 0xbd, 0x11, 0xf4, 0x8c, 0xc6, 0xe1, 0x50, 0xd6, 0x31,
	0xa1, 0xd0, 0xc9, 0xee, 0x16, 0xc5, 0xb7, 0xb2, 0x5b, 0x4b, 0xe4, 0x5f, 0x68, 0x9f, 0x60, 0x1c,
	0x2c, 0xc2, 0x16, 0xa9, 0x81, 0x6d, 0x62, 0x0c, 0x5a, 0xcb, 0xa4, 0x03, 0xad, 0x1b, 0xf9, 0x21,
	0x4e, 0x51, 0x61, 0x6b, 0x85, 0xb4, 0xa1, 0x9e, 0x49, 0x33, 0xa8, 0xb4, 0xef, 0x5c, 0x5e, 0x97,
	0x6c, 0xa8, 0x04, 0x38, 0x62, 0x4a, 0x5d, 0x5e, 0x97, 0x08, 0x69, 0xe9, 0xfb, 0xdc, 0x7c, 0x90,
	0xd2, 0xf9, 0x6a, 0x41, 0xe5, 0x43, 0xac, 0x52, 0x5f, 0xde, 0xde, 0xba, 0x75, 0x67, 0xeb, 0x4f,
	0xcc, 0x44, 0x8d, 0xed, 0xd6, 0x17, 0x6d, 0xe7, 0xa1, 0xe4, 0x73, 0xe1, 0xa3, 0x87, 0x63, 0x33,
	0xec, 0x2d, 0x68, 0x14, 0xb3, 0xbd, 0xed, 0xb0, 0x7a, 0x81, 0x6a, 0x43, 0x6c, 0x43, 0x33, 0xf7,
	0x8c, 0xde, 0x77, 0x61, 0xb6, 0x46, 0x06, 0x1f, 0x19, 0xd4, 0xf9, 0x04, 0xe5, 0x93, 0xd4, 0x9b,
	0x79, 0x17, 0xd6, 0xbd, 0xba, 0xd8, 0x81, 0xb6, 0x40, 0x16, 0x7c, 0x1e, 0x8e, 0xb9, 0x18, 0xfa,
	0x3c, 0x8e, 0xd1, 0x57, 0xfa, 0x02, 0xb6, 0xd7, 0xd4, 0xc4, 0x5b, 0x2e, 0xde, 0x18, 0xd8, 0x39,
	0x02, 0xfb, 0x84, 0x0b, 0x75, 0x9a, 0xa0, 0x4f, 0xd6, 0xa0, 0xac, 0xbf, 0x04, 0x33, 0x83, 0xf7,
	0x4b, 0x9e, 0x09, 0x09, 0x85, 0x8a, 0x62, 0x22, 0x44, 0x53, 0x24, 0x25, 0xb2, 0x78, 0xd0, 0x84,
	0x7a, 0xc2, 0x85, 0x1a, 0x4e, 0xb9, 0xcf, 0x54, 0xc4, 0xe3, 0xc1, 0xce, 0xb7, 0x1f, 0x0f, 0xac,
	0x8f, 0x8f, 0xff, 0xfc, 0x3e, 0x26, 0x67, 0x61, 0xf6, 0x60, 0x8c, 0x2a, 0xfa, 0xa1, 0x78, 0xf1,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x4a, 0x1e, 0xa4, 0x4e, 0x05, 0x00, 0x00,
}

func (this *DebugAttachment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DebugAttachment)
	if !ok {
		that2, ok := that.(DebugAttachment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if this.PlankName != that1.PlankName {
		return false
	}
	if this.Debugger != that1.Debugger {
		return false
	}
	if this.Image != that1.Image {
		return false
	}
	if this.ProcessName != that1.ProcessName {
		return false
	}
	if this.Node != that1.Node {
		return false
	}
	if this.MatchRequest != that1.MatchRequest {
		return false
	}
	if this.DebugServerAddress != that1.DebugServerAddress {
		return false
	}
	if this.Pod != that1.Pod {
		return false
	}
	if this.Container != that1.Container {
		return false
	}
	if this.DebugNamespace != that1.DebugNamespace {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Intent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Intent)
	if !ok {
		that2, ok := that.(Intent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Debugger != that1.Debugger {
		return false
	}
	if !this.Pod.Equal(that1.Pod) {
		return false
	}
	if this.ContainerName != that1.ContainerName {
		return false
	}
	if this.ProcessMatcher != that1.ProcessMatcher {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Plank) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plank)
	if !ok {
		that2, ok := that.(Plank)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Pod.Equal(that1.Pod) {
		return false
	}
	if this.ReadyForConnect != that1.ReadyForConnect {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PortSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PortSpec)
	if !ok {
		that2, ok := that.(PortSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.PortLocation == nil {
		if this.PortLocation != nil {
			return false
		}
	} else if this.PortLocation == nil {
		return false
	} else if !this.PortLocation.Equal(that1.PortLocation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PortSpec_Plank) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PortSpec_Plank)
	if !ok {
		that2, ok := that.(PortSpec_Plank)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Plank != that1.Plank {
		return false
	}
	return true
}
func (this *PortSpec_Target) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PortSpec_Target)
	if !ok {
		that2, ok := that.(PortSpec_Target)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Target != that1.Target {
		return false
	}
	return true
}
