//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/math_server.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: proto/math_server.proto

package math_server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TwoInts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N1 int32 `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	N2 int32 `protobuf:"varint,2,opt,name=n2,proto3" json:"n2,omitempty"`
}

func (x *TwoInts) Reset() {
	*x = TwoInts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwoInts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwoInts) ProtoMessage() {}

func (x *TwoInts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwoInts.ProtoReflect.Descriptor instead.
func (*TwoInts) Descriptor() ([]byte, []int) {
	return file_proto_math_server_proto_rawDescGZIP(), []int{0}
}

func (x *TwoInts) GetN1() int32 {
	if x != nil {
		return x.N1
	}
	return 0
}

func (x *TwoInts) GetN2() int32 {
	if x != nil {
		return x.N2
	}
	return 0
}

type IntInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N1 int32 `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
}

func (x *IntInput) Reset() {
	*x = IntInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntInput) ProtoMessage() {}

func (x *IntInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntInput.ProtoReflect.Descriptor instead.
func (*IntInput) Descriptor() ([]byte, []int) {
	return file_proto_math_server_proto_rawDescGZIP(), []int{1}
}

func (x *IntInput) GetN1() int32 {
	if x != nil {
		return x.N1
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_math_server_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type CapabilitiesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability string `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *CapabilitiesResult) Reset() {
	*x = CapabilitiesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilitiesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilitiesResult) ProtoMessage() {}

func (x *CapabilitiesResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilitiesResult.ProtoReflect.Descriptor instead.
func (*CapabilitiesResult) Descriptor() ([]byte, []int) {
	return file_proto_math_server_proto_rawDescGZIP(), []int{3}
}

func (x *CapabilitiesResult) GetCapability() string {
	if x != nil {
		return x.Capability
	}
	return ""
}

type TransformationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transformed string `protobuf:"bytes,1,opt,name=transformed,proto3" json:"transformed,omitempty"`
}

func (x *TransformationResult) Reset() {
	*x = TransformationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformationResult) ProtoMessage() {}

func (x *TransformationResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformationResult.ProtoReflect.Descriptor instead.
func (*TransformationResult) Descriptor() ([]byte, []int) {
	return file_proto_math_server_proto_rawDescGZIP(), []int{4}
}

func (x *TransformationResult) GetTransformed() string {
	if x != nil {
		return x.Transformed
	}
	return ""
}

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_proto_math_server_proto_rawDescGZIP(), []int{5}
}

var File_proto_math_server_proto protoreflect.FileDescriptor

var file_proto_math_server_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x07, 0x54, 0x77, 0x6f, 0x49, 0x6e, 0x74,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e,
	0x31, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e,
	0x32, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e, 0x31, 0x22, 0x20, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x34, 0x0a, 0x12, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x22,
	0x07, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0x87, 0x03, 0x0a, 0x0b, 0x4d, 0x61, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x14, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x77,
	0x6f, 0x49, 0x6e, 0x74, 0x73, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x77, 0x6f, 0x49, 0x6e, 0x74, 0x73, 0x1a, 0x13,
	0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x4e, 0x12, 0x15, 0x2e,
	0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3b, 0x0a,
	0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x4e, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x6d,
	0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79,
	0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_math_server_proto_rawDescOnce sync.Once
	file_proto_math_server_proto_rawDescData = file_proto_math_server_proto_rawDesc
)

func file_proto_math_server_proto_rawDescGZIP() []byte {
	file_proto_math_server_proto_rawDescOnce.Do(func() {
		file_proto_math_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_math_server_proto_rawDescData)
	})
	return file_proto_math_server_proto_rawDescData
}

var file_proto_math_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_math_server_proto_goTypes = []interface{}{
	(*TwoInts)(nil),              // 0: math_server.TwoInts
	(*IntInput)(nil),             // 1: math_server.IntInput
	(*Result)(nil),               // 2: math_server.Result
	(*CapabilitiesResult)(nil),   // 3: math_server.CapabilitiesResult
	(*TransformationResult)(nil), // 4: math_server.TransformationResult
	(*Dummy)(nil),                // 5: math_server.Dummy
}
var file_proto_math_server_proto_depIdxs = []int32{
	0, // 0: math_server.MathService.Add:input_type -> math_server.TwoInts
	0, // 1: math_server.MathService.Multiply:input_type -> math_server.TwoInts
	1, // 2: math_server.MathService.AddN:input_type -> math_server.IntInput
	1, // 3: math_server.MathService.MultiplyN:input_type -> math_server.IntInput
	5, // 4: math_server.MathService.GetCapabilties:input_type -> math_server.Dummy
	1, // 5: math_server.MathService.Transform:input_type -> math_server.IntInput
	2, // 6: math_server.MathService.Add:output_type -> math_server.Result
	2, // 7: math_server.MathService.Multiply:output_type -> math_server.Result
	2, // 8: math_server.MathService.AddN:output_type -> math_server.Result
	2, // 9: math_server.MathService.MultiplyN:output_type -> math_server.Result
	3, // 10: math_server.MathService.GetCapabilties:output_type -> math_server.CapabilitiesResult
	4, // 11: math_server.MathService.Transform:output_type -> math_server.TransformationResult
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_math_server_proto_init() }
func file_proto_math_server_proto_init() {
	if File_proto_math_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_math_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwoInts); i {
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
		file_proto_math_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntInput); i {
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
		file_proto_math_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_proto_math_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilitiesResult); i {
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
		file_proto_math_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformationResult); i {
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
		file_proto_math_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dummy); i {
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
			RawDescriptor: file_proto_math_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_math_server_proto_goTypes,
		DependencyIndexes: file_proto_math_server_proto_depIdxs,
		MessageInfos:      file_proto_math_server_proto_msgTypes,
	}.Build()
	File_proto_math_server_proto = out.File
	file_proto_math_server_proto_rawDesc = nil
	file_proto_math_server_proto_goTypes = nil
	file_proto_math_server_proto_depIdxs = nil
}
