// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tableau/protobuf/xml.proto

package tableaupb

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

type XMLProp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SheetPropMap map[string]*SheetProp `protobuf:"bytes,1,rep,name=sheet_prop_map,json=sheetPropMap,proto3" json:"sheet_prop_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *XMLProp) Reset() {
	*x = XMLProp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLProp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLProp) ProtoMessage() {}

func (x *XMLProp) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMLProp.ProtoReflect.Descriptor instead.
func (*XMLProp) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{0}
}

func (x *XMLProp) GetSheetPropMap() map[string]*SheetProp {
	if x != nil {
		return x.SheetPropMap
	}
	return nil
}

type SheetProp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *MetaProp `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data *DataProp `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SheetProp) Reset() {
	*x = SheetProp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetProp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetProp) ProtoMessage() {}

func (x *SheetProp) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetProp.ProtoReflect.Descriptor instead.
func (*SheetProp) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{1}
}

func (x *SheetProp) GetMeta() *MetaProp {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SheetProp) GetData() *DataProp {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderedAttrMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Attr          `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Map  map[string]int32 `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *OrderedAttrMap) Reset() {
	*x = OrderedAttrMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderedAttrMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderedAttrMap) ProtoMessage() {}

func (x *OrderedAttrMap) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderedAttrMap.ProtoReflect.Descriptor instead.
func (*OrderedAttrMap) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{2}
}

func (x *OrderedAttrMap) GetList() []*Attr {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *OrderedAttrMap) GetMap() map[string]int32 {
	if x != nil {
		return x.Map
	}
	return nil
}

type MetaProp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AttrMap *OrderedAttrMap `protobuf:"bytes,2,opt,name=attr_map,json=attrMap,proto3" json:"attr_map,omitempty"`
	// Combine list and map to get an ordered map
	ChildList []*MetaProp      `protobuf:"bytes,3,rep,name=child_list,json=childList,proto3" json:"child_list,omitempty"`
	ChildMap  map[string]int32 `protobuf:"bytes,4,rep,name=child_map,json=childMap,proto3" json:"child_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *MetaProp) Reset() {
	*x = MetaProp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaProp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaProp) ProtoMessage() {}

func (x *MetaProp) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaProp.ProtoReflect.Descriptor instead.
func (*MetaProp) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{3}
}

func (x *MetaProp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaProp) GetAttrMap() *OrderedAttrMap {
	if x != nil {
		return x.AttrMap
	}
	return nil
}

func (x *MetaProp) GetChildList() []*MetaProp {
	if x != nil {
		return x.ChildList
	}
	return nil
}

func (x *MetaProp) GetChildMap() map[string]int32 {
	if x != nil {
		return x.ChildMap
	}
	return nil
}

type DataProp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AttrMap   *OrderedAttrMap `protobuf:"bytes,2,opt,name=attr_map,json=attrMap,proto3" json:"attr_map,omitempty"`
	ChildList []*DataProp     `protobuf:"bytes,3,rep,name=child_list,json=childList,proto3" json:"child_list,omitempty"`
}

func (x *DataProp) Reset() {
	*x = DataProp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataProp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataProp) ProtoMessage() {}

func (x *DataProp) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataProp.ProtoReflect.Descriptor instead.
func (*DataProp) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{4}
}

func (x *DataProp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataProp) GetAttrMap() *OrderedAttrMap {
	if x != nil {
		return x.AttrMap
	}
	return nil
}

func (x *DataProp) GetChildList() []*DataProp {
	if x != nil {
		return x.ChildList
	}
	return nil
}

type Attr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Attr) Reset() {
	*x = Attr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attr) ProtoMessage() {}

func (x *Attr) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attr.ProtoReflect.Descriptor instead.
func (*Attr) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{5}
}

func (x *Attr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attr) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_tableau_protobuf_xml_proto protoreflect.FileDescriptor

var file_tableau_protobuf_xml_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x78, 0x6d, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x22, 0xa8, 0x01, 0x0a, 0x07, 0x58, 0x4d, 0x4c, 0x50, 0x72, 0x6f,
	0x70, 0x12, 0x48, 0x0a, 0x0e, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x50, 0x72, 0x6f, 0x70, 0x2e, 0x53, 0x68, 0x65, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x1a, 0x53, 0x0a, 0x11, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x53, 0x68, 0x65, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x59, 0x0a, 0x09, 0x53, 0x68, 0x65, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x12, 0x25, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x0e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x21,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xff, 0x01,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x4d,
	0x61, 0x70, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61,
	0x75, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4d,
	0x61, 0x70, 0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x84, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x32, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x07, 0x61, 0x74, 0x74,
	0x72, 0x4d, 0x61, 0x70, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x61, 0x75, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x09, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x04, 0x41, 0x74, 0x74, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x69, 0x6f,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tableau_protobuf_xml_proto_rawDescOnce sync.Once
	file_tableau_protobuf_xml_proto_rawDescData = file_tableau_protobuf_xml_proto_rawDesc
)

func file_tableau_protobuf_xml_proto_rawDescGZIP() []byte {
	file_tableau_protobuf_xml_proto_rawDescOnce.Do(func() {
		file_tableau_protobuf_xml_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableau_protobuf_xml_proto_rawDescData)
	})
	return file_tableau_protobuf_xml_proto_rawDescData
}

var file_tableau_protobuf_xml_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tableau_protobuf_xml_proto_goTypes = []interface{}{
	(*XMLProp)(nil),        // 0: tableau.XMLProp
	(*SheetProp)(nil),      // 1: tableau.SheetProp
	(*OrderedAttrMap)(nil), // 2: tableau.OrderedAttrMap
	(*MetaProp)(nil),       // 3: tableau.MetaProp
	(*DataProp)(nil),       // 4: tableau.DataProp
	(*Attr)(nil),           // 5: tableau.Attr
	nil,                    // 6: tableau.XMLProp.SheetPropMapEntry
	nil,                    // 7: tableau.OrderedAttrMap.MapEntry
	nil,                    // 8: tableau.MetaProp.ChildMapEntry
}
var file_tableau_protobuf_xml_proto_depIdxs = []int32{
	6,  // 0: tableau.XMLProp.sheet_prop_map:type_name -> tableau.XMLProp.SheetPropMapEntry
	3,  // 1: tableau.SheetProp.meta:type_name -> tableau.MetaProp
	4,  // 2: tableau.SheetProp.data:type_name -> tableau.DataProp
	5,  // 3: tableau.OrderedAttrMap.list:type_name -> tableau.Attr
	7,  // 4: tableau.OrderedAttrMap.map:type_name -> tableau.OrderedAttrMap.MapEntry
	2,  // 5: tableau.MetaProp.attr_map:type_name -> tableau.OrderedAttrMap
	3,  // 6: tableau.MetaProp.child_list:type_name -> tableau.MetaProp
	8,  // 7: tableau.MetaProp.child_map:type_name -> tableau.MetaProp.ChildMapEntry
	2,  // 8: tableau.DataProp.attr_map:type_name -> tableau.OrderedAttrMap
	4,  // 9: tableau.DataProp.child_list:type_name -> tableau.DataProp
	1,  // 10: tableau.XMLProp.SheetPropMapEntry.value:type_name -> tableau.SheetProp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tableau_protobuf_xml_proto_init() }
func file_tableau_protobuf_xml_proto_init() {
	if File_tableau_protobuf_xml_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tableau_protobuf_xml_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMLProp); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetProp); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderedAttrMap); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaProp); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataProp); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attr); i {
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
			RawDescriptor: file_tableau_protobuf_xml_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tableau_protobuf_xml_proto_goTypes,
		DependencyIndexes: file_tableau_protobuf_xml_proto_depIdxs,
		MessageInfos:      file_tableau_protobuf_xml_proto_msgTypes,
	}.Build()
	File_tableau_protobuf_xml_proto = out.File
	file_tableau_protobuf_xml_proto_rawDesc = nil
	file_tableau_protobuf_xml_proto_goTypes = nil
	file_tableau_protobuf_xml_proto_depIdxs = nil
}
