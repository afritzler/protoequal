// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: test.proto

package test

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Foo message with Bar, Baz fields and a nested message Qux.
type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bar string `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"` // Bar field
	Baz string `protobuf:"bytes,2,opt,name=baz,proto3" json:"baz,omitempty"` // Baz field
	Qux *Qux   `protobuf:"bytes,3,opt,name=qux,proto3" json:"qux,omitempty"` // Qux is a nested message
}

func (x *Foo) Reset() {
	*x = Foo{}
	mi := &file_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

func (x *Foo) GetBaz() string {
	if x != nil {
		return x.Baz
	}
	return ""
}

func (x *Foo) GetQux() *Qux {
	if x != nil {
		return x.Qux
	}
	return nil
}

// The Qux message with Driver and Handle fields.
type Qux struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"` // Driver field
	Handle string `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"` // Handle field
}

func (x *Qux) Reset() {
	*x = Qux{}
	mi := &file_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Qux) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Qux) ProtoMessage() {}

func (x *Qux) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Qux.ProtoReflect.Descriptor instead.
func (*Qux) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Qux) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Qux) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x46, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x61, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x12, 0x1b, 0x0a,
	0x03, 0x71, 0x75, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x51, 0x75, 0x78, 0x52, 0x03, 0x71, 0x75, 0x78, 0x22, 0x35, 0x0a, 0x03, 0x51, 0x75,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x66, 0x72, 0x69, 0x74, 0x7a, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65,
	0x71, 0x75, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []any{
	(*Foo)(nil), // 0: test.Foo
	(*Qux)(nil), // 1: test.Qux
}
var file_test_proto_depIdxs = []int32{
	1, // 0: test.Foo.qux:type_name -> test.Qux
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
