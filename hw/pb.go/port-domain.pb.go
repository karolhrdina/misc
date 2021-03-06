// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: port-domain.proto

package pb_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Note:
// Based on provided .json example file
// * none of the port properties are guaranteed to be filled in (present) and we have to take that
//   into account => therefore the message's fields are wrappers, so that in case of update/snapshot,
//   we know whether the value is being changed or not
// * I selected key of root level properties as unique port identifier. There's also `unlocs` property,
//   that has the same value as far as the sample file goes, but we can not count on that since it's an array,
//   we have to consider that some ports might have more unlocs codes
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City        *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country     *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Province    *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Code        *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Alias       *RepeatedString         `protobuf:"bytes,8,opt,name=alias,proto3" json:"alias,omitempty"`
	Regions     *RepeatedString         `protobuf:"bytes,9,opt,name=regions,proto3" json:"regions,omitempty"`
	Coordinates *Coordinates            `protobuf:"bytes,10,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Unlocs      *RepeatedString         `protobuf:"bytes,11,opt,name=unlocs,proto3" json:"unlocs,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_port_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_port_domain_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Port) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Port) GetCity() *wrapperspb.StringValue {
	if x != nil {
		return x.City
	}
	return nil
}

func (x *Port) GetCountry() *wrapperspb.StringValue {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *Port) GetProvince() *wrapperspb.StringValue {
	if x != nil {
		return x.Province
	}
	return nil
}

func (x *Port) GetTimezone() *wrapperspb.StringValue {
	if x != nil {
		return x.Timezone
	}
	return nil
}

func (x *Port) GetCode() *wrapperspb.StringValue {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Port) GetAlias() *RepeatedString {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() *RepeatedString {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() *Coordinates {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetUnlocs() *RepeatedString {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

// wrapper for array of strings
type RepeatedString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *RepeatedString) Reset() {
	*x = RepeatedString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedString) ProtoMessage() {}

func (x *RepeatedString) ProtoReflect() protoreflect.Message {
	mi := &file_port_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedString.ProtoReflect.Descriptor instead.
func (*RepeatedString) Descriptor() ([]byte, []int) {
	return file_port_domain_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatedString) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

// wrapper for Coordinates
type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_port_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_port_domain_proto_rawDescGZIP(), []int{2}
}

func (x *Coordinates) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinates) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_port_domain_proto protoreflect.FileDescriptor

var file_port_domain_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x04, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x39, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0b, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x22, 0x26,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x79, 0x32, 0x76, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x34, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x72, 0x6f, 0x6c, 0x68, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2f, 0x68, 0x77, 0x2f, 0x70, 0x62, 0x2e, 0x67,
	0x6f, 0x3b, 0x70, 0x62, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_port_domain_proto_rawDescOnce sync.Once
	file_port_domain_proto_rawDescData = file_port_domain_proto_rawDesc
)

func file_port_domain_proto_rawDescGZIP() []byte {
	file_port_domain_proto_rawDescOnce.Do(func() {
		file_port_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_port_domain_proto_rawDescData)
	})
	return file_port_domain_proto_rawDescData
}

var file_port_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_port_domain_proto_goTypes = []interface{}{
	(*Port)(nil),                   // 0: portdomain.Port
	(*RepeatedString)(nil),         // 1: portdomain.RepeatedString
	(*Coordinates)(nil),            // 2: portdomain.Coordinates
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_port_domain_proto_depIdxs = []int32{
	3,  // 0: portdomain.Port.name:type_name -> google.protobuf.StringValue
	3,  // 1: portdomain.Port.city:type_name -> google.protobuf.StringValue
	3,  // 2: portdomain.Port.country:type_name -> google.protobuf.StringValue
	3,  // 3: portdomain.Port.province:type_name -> google.protobuf.StringValue
	3,  // 4: portdomain.Port.timezone:type_name -> google.protobuf.StringValue
	3,  // 5: portdomain.Port.code:type_name -> google.protobuf.StringValue
	1,  // 6: portdomain.Port.alias:type_name -> portdomain.RepeatedString
	1,  // 7: portdomain.Port.regions:type_name -> portdomain.RepeatedString
	2,  // 8: portdomain.Port.coordinates:type_name -> portdomain.Coordinates
	1,  // 9: portdomain.Port.unlocs:type_name -> portdomain.RepeatedString
	0,  // 10: portdomain.Portdomain.Snapshot:input_type -> portdomain.Port
	4,  // 11: portdomain.Portdomain.List:input_type -> google.protobuf.Empty
	4,  // 12: portdomain.Portdomain.Snapshot:output_type -> google.protobuf.Empty
	0,  // 13: portdomain.Portdomain.List:output_type -> portdomain.Port
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_port_domain_proto_init() }
func file_port_domain_proto_init() {
	if File_port_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_port_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_port_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedString); i {
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
		file_port_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
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
			RawDescriptor: file_port_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_port_domain_proto_goTypes,
		DependencyIndexes: file_port_domain_proto_depIdxs,
		MessageInfos:      file_port_domain_proto_msgTypes,
	}.Build()
	File_port_domain_proto = out.File
	file_port_domain_proto_rawDesc = nil
	file_port_domain_proto_goTypes = nil
	file_port_domain_proto_depIdxs = nil
}
