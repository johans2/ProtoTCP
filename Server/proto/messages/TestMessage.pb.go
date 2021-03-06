// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: TestMessage.proto

package messages

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

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Number   int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// Types that are assignable to Payload:
	//	*TestMessage_Pl1
	//	*TestMessage_Pl2
	//	*TestMessage_Pl3
	Payload isTestMessage_Payload `protobuf_oneof:"payload"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_TestMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_TestMessage_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *TestMessage) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (m *TestMessage) GetPayload() isTestMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *TestMessage) GetPl1() *PayLoad1 {
	if x, ok := x.GetPayload().(*TestMessage_Pl1); ok {
		return x.Pl1
	}
	return nil
}

func (x *TestMessage) GetPl2() *PayLoad2 {
	if x, ok := x.GetPayload().(*TestMessage_Pl2); ok {
		return x.Pl2
	}
	return nil
}

func (x *TestMessage) GetPl3() *PayLoad3 {
	if x, ok := x.GetPayload().(*TestMessage_Pl3); ok {
		return x.Pl3
	}
	return nil
}

type isTestMessage_Payload interface {
	isTestMessage_Payload()
}

type TestMessage_Pl1 struct {
	Pl1 *PayLoad1 `protobuf:"bytes,3,opt,name=pl1,proto3,oneof"`
}

type TestMessage_Pl2 struct {
	Pl2 *PayLoad2 `protobuf:"bytes,4,opt,name=pl2,proto3,oneof"`
}

type TestMessage_Pl3 struct {
	Pl3 *PayLoad3 `protobuf:"bytes,5,opt,name=pl3,proto3,oneof"`
}

func (*TestMessage_Pl1) isTestMessage_Payload() {}

func (*TestMessage_Pl2) isTestMessage_Payload() {}

func (*TestMessage_Pl3) isTestMessage_Payload() {}

type PayLoad1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PayLoad1) Reset() {
	*x = PayLoad1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLoad1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLoad1) ProtoMessage() {}

func (x *PayLoad1) ProtoReflect() protoreflect.Message {
	mi := &file_TestMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLoad1.ProtoReflect.Descriptor instead.
func (*PayLoad1) Descriptor() ([]byte, []int) {
	return file_TestMessage_proto_rawDescGZIP(), []int{1}
}

func (x *PayLoad1) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PayLoad2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PayLoad2) Reset() {
	*x = PayLoad2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLoad2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLoad2) ProtoMessage() {}

func (x *PayLoad2) ProtoReflect() protoreflect.Message {
	mi := &file_TestMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLoad2.ProtoReflect.Descriptor instead.
func (*PayLoad2) Descriptor() ([]byte, []int) {
	return file_TestMessage_proto_rawDescGZIP(), []int{2}
}

func (x *PayLoad2) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PayLoad3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Percent float32 `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
}

func (x *PayLoad3) Reset() {
	*x = PayLoad3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestMessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLoad3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLoad3) ProtoMessage() {}

func (x *PayLoad3) ProtoReflect() protoreflect.Message {
	mi := &file_TestMessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLoad3.ProtoReflect.Descriptor instead.
func (*PayLoad3) Descriptor() ([]byte, []int) {
	return file_TestMessage_proto_rawDescGZIP(), []int{3}
}

func (x *PayLoad3) GetPercent() float32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

var File_TestMessage_proto protoreflect.FileDescriptor

var file_TestMessage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6c, 0x31, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x31, 0x48,
	0x00, 0x52, 0x03, 0x70, 0x6c, 0x31, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6c, 0x32, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x32, 0x48, 0x00,
	0x52, 0x03, 0x70, 0x6c, 0x32, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x6c, 0x33, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x33, 0x48, 0x00, 0x52,
	0x03, 0x70, 0x6c, 0x33, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x1c, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a,
	0x08, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x24, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x33, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x26, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0xaa, 0x02, 0x13, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TestMessage_proto_rawDescOnce sync.Once
	file_TestMessage_proto_rawDescData = file_TestMessage_proto_rawDesc
)

func file_TestMessage_proto_rawDescGZIP() []byte {
	file_TestMessage_proto_rawDescOnce.Do(func() {
		file_TestMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_TestMessage_proto_rawDescData)
	})
	return file_TestMessage_proto_rawDescData
}

var file_TestMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_TestMessage_proto_goTypes = []interface{}{
	(*TestMessage)(nil), // 0: TestMessage
	(*PayLoad1)(nil),    // 1: PayLoad1
	(*PayLoad2)(nil),    // 2: PayLoad2
	(*PayLoad3)(nil),    // 3: PayLoad3
}
var file_TestMessage_proto_depIdxs = []int32{
	1, // 0: TestMessage.pl1:type_name -> PayLoad1
	2, // 1: TestMessage.pl2:type_name -> PayLoad2
	3, // 2: TestMessage.pl3:type_name -> PayLoad3
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_TestMessage_proto_init() }
func file_TestMessage_proto_init() {
	if File_TestMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TestMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
		file_TestMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLoad1); i {
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
		file_TestMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLoad2); i {
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
		file_TestMessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLoad3); i {
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
	file_TestMessage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestMessage_Pl1)(nil),
		(*TestMessage_Pl2)(nil),
		(*TestMessage_Pl3)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TestMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TestMessage_proto_goTypes,
		DependencyIndexes: file_TestMessage_proto_depIdxs,
		MessageInfos:      file_TestMessage_proto_msgTypes,
	}.Build()
	File_TestMessage_proto = out.File
	file_TestMessage_proto_rawDesc = nil
	file_TestMessage_proto_goTypes = nil
	file_TestMessage_proto_depIdxs = nil
}
