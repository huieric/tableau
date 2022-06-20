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

type XMLBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SheetMap maps sheet name to a XMLSheet struct
	// containing two trees which describe the XML structure
	SheetList []*XMLSheet      `protobuf:"bytes,1,rep,name=sheet_list,json=sheetList,proto3" json:"sheet_list,omitempty"`
	SheetMap  map[string]int32 `protobuf:"bytes,2,rep,name=sheet_map,json=sheetMap,proto3" json:"sheet_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // sheet name -> index
}

func (x *XMLBook) Reset() {
	*x = XMLBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLBook) ProtoMessage() {}

func (x *XMLBook) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use XMLBook.ProtoReflect.Descriptor instead.
func (*XMLBook) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{0}
}

func (x *XMLBook) GetSheetList() []*XMLSheet {
	if x != nil {
		return x.SheetList
	}
	return nil
}

func (x *XMLBook) GetSheetMap() map[string]int32 {
	if x != nil {
		return x.SheetMap
	}
	return nil
}

type XMLSheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// meta tree, describing the types of attributes of all nodes in XML
	Meta *XMLNode `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// data tree, describing the values of attributes of the explicitly specified nodes in XML
	Data *XMLNode `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// NodeMap maps path to a list of nodes matched the specified path in the tree
	MetaNodeMap map[string]*XMLNode           `protobuf:"bytes,3,rep,name=meta_node_map,json=metaNodeMap,proto3" json:"meta_node_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // path -> XMLNode
	DataNodeMap map[string]*XMLSheet_NodeList `protobuf:"bytes,4,rep,name=data_node_map,json=dataNodeMap,proto3" json:"data_node_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // path -> XMLNode list
}

func (x *XMLSheet) Reset() {
	*x = XMLSheet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLSheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLSheet) ProtoMessage() {}

func (x *XMLSheet) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use XMLSheet.ProtoReflect.Descriptor instead.
func (*XMLSheet) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{1}
}

func (x *XMLSheet) GetMeta() *XMLNode {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *XMLSheet) GetData() *XMLNode {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *XMLSheet) GetMetaNodeMap() map[string]*XMLNode {
	if x != nil {
		return x.MetaNodeMap
	}
	return nil
}

func (x *XMLSheet) GetDataNodeMap() map[string]*XMLSheet_NodeList {
	if x != nil {
		return x.DataNodeMap
	}
	return nil
}

type XMLNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// duplicated attributes in one node are forbidden
	AttrMap *XMLNode_AttrMap `protobuf:"bytes,2,opt,name=attr_map,json=attrMap,proto3" json:"attr_map,omitempty"`
	// - meta: Each child must be unique, so len(IndexList) must be 1
	// - data: Duplicated children are allowed, so len(IndexList) could be greater than 1
	ChildList []*XMLNode                    `protobuf:"bytes,3,rep,name=child_list,json=childList,proto3" json:"child_list,omitempty"`
	ChildMap  map[string]*XMLNode_IndexList `protobuf:"bytes,4,rep,name=child_map,json=childMap,proto3" json:"child_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // child name -> index list of child
	// record parent node so that we can trace back
	Parent *XMLNode `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// path that walks from root to self node, e.g.: Conf/Server/Toggle
	Path string `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *XMLNode) Reset() {
	*x = XMLNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLNode) ProtoMessage() {}

func (x *XMLNode) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use XMLNode.ProtoReflect.Descriptor instead.
func (*XMLNode) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{2}
}

func (x *XMLNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XMLNode) GetAttrMap() *XMLNode_AttrMap {
	if x != nil {
		return x.AttrMap
	}
	return nil
}

func (x *XMLNode) GetChildList() []*XMLNode {
	if x != nil {
		return x.ChildList
	}
	return nil
}

func (x *XMLNode) GetChildMap() map[string]*XMLNode_IndexList {
	if x != nil {
		return x.ChildMap
	}
	return nil
}

func (x *XMLNode) GetParent() *XMLNode {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *XMLNode) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type XMLSheet_NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*XMLNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *XMLSheet_NodeList) Reset() {
	*x = XMLSheet_NodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLSheet_NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLSheet_NodeList) ProtoMessage() {}

func (x *XMLSheet_NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMLSheet_NodeList.ProtoReflect.Descriptor instead.
func (*XMLSheet_NodeList) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{1, 2}
}

func (x *XMLSheet_NodeList) GetNodes() []*XMLNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type XMLNode_IndexList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indexes []int32 `protobuf:"varint,1,rep,packed,name=indexes,proto3" json:"indexes,omitempty"`
}

func (x *XMLNode_IndexList) Reset() {
	*x = XMLNode_IndexList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLNode_IndexList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLNode_IndexList) ProtoMessage() {}

func (x *XMLNode_IndexList) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMLNode_IndexList.ProtoReflect.Descriptor instead.
func (*XMLNode_IndexList) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{2, 1}
}

func (x *XMLNode_IndexList) GetIndexes() []int32 {
	if x != nil {
		return x.Indexes
	}
	return nil
}

// Combine list and map to get an ordered map
type XMLNode_AttrMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*XMLNode_AttrMap_Attr `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Map  map[string]int32        `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // attribute name -> index
}

func (x *XMLNode_AttrMap) Reset() {
	*x = XMLNode_AttrMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLNode_AttrMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLNode_AttrMap) ProtoMessage() {}

func (x *XMLNode_AttrMap) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMLNode_AttrMap.ProtoReflect.Descriptor instead.
func (*XMLNode_AttrMap) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{2, 2}
}

func (x *XMLNode_AttrMap) GetList() []*XMLNode_AttrMap_Attr {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *XMLNode_AttrMap) GetMap() map[string]int32 {
	if x != nil {
		return x.Map
	}
	return nil
}

type XMLNode_AttrMap_Attr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *XMLNode_AttrMap_Attr) Reset() {
	*x = XMLNode_AttrMap_Attr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_xml_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XMLNode_AttrMap_Attr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XMLNode_AttrMap_Attr) ProtoMessage() {}

func (x *XMLNode_AttrMap_Attr) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_xml_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XMLNode_AttrMap_Attr.ProtoReflect.Descriptor instead.
func (*XMLNode_AttrMap_Attr) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_xml_proto_rawDescGZIP(), []int{2, 2, 1}
}

func (x *XMLNode_AttrMap_Attr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XMLNode_AttrMap_Attr) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_tableau_protobuf_xml_proto protoreflect.FileDescriptor

var file_tableau_protobuf_xml_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x78, 0x6d, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x58, 0x4d, 0x4c, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e,
	0x58, 0x4d, 0x4c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75,
	0x2e, 0x58, 0x4d, 0x4c, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x70,
	0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc8, 0x03,
	0x0a, 0x08, 0x58, 0x4d, 0x4c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x46,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e,
	0x58, 0x4d, 0x4c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x50, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5a, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x32, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xdc, 0x04, 0x0a, 0x07, 0x58, 0x4d, 0x4c,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x4d, 0x61, 0x70, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x2f, 0x0a,
	0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x09, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x57, 0x0a, 0x0d, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x25, 0x0a, 0x09, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x1a, 0xdb, 0x01, 0x0a, 0x07, 0x41, 0x74,
	0x74, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x58, 0x4d,
	0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e,
	0x58, 0x4d, 0x4c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x2e,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a,
	0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x30, 0x0a, 0x04, 0x41, 0x74, 0x74, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x69, 0x6f, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x61, 0x75, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_tableau_protobuf_xml_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tableau_protobuf_xml_proto_goTypes = []interface{}{
	(*XMLBook)(nil),              // 0: tableau.XMLBook
	(*XMLSheet)(nil),             // 1: tableau.XMLSheet
	(*XMLNode)(nil),              // 2: tableau.XMLNode
	nil,                          // 3: tableau.XMLBook.SheetMapEntry
	nil,                          // 4: tableau.XMLSheet.MetaNodeMapEntry
	nil,                          // 5: tableau.XMLSheet.DataNodeMapEntry
	(*XMLSheet_NodeList)(nil),    // 6: tableau.XMLSheet.NodeList
	nil,                          // 7: tableau.XMLNode.ChildMapEntry
	(*XMLNode_IndexList)(nil),    // 8: tableau.XMLNode.IndexList
	(*XMLNode_AttrMap)(nil),      // 9: tableau.XMLNode.AttrMap
	nil,                          // 10: tableau.XMLNode.AttrMap.MapEntry
	(*XMLNode_AttrMap_Attr)(nil), // 11: tableau.XMLNode.AttrMap.Attr
}
var file_tableau_protobuf_xml_proto_depIdxs = []int32{
	1,  // 0: tableau.XMLBook.sheet_list:type_name -> tableau.XMLSheet
	3,  // 1: tableau.XMLBook.sheet_map:type_name -> tableau.XMLBook.SheetMapEntry
	2,  // 2: tableau.XMLSheet.meta:type_name -> tableau.XMLNode
	2,  // 3: tableau.XMLSheet.data:type_name -> tableau.XMLNode
	4,  // 4: tableau.XMLSheet.meta_node_map:type_name -> tableau.XMLSheet.MetaNodeMapEntry
	5,  // 5: tableau.XMLSheet.data_node_map:type_name -> tableau.XMLSheet.DataNodeMapEntry
	9,  // 6: tableau.XMLNode.attr_map:type_name -> tableau.XMLNode.AttrMap
	2,  // 7: tableau.XMLNode.child_list:type_name -> tableau.XMLNode
	7,  // 8: tableau.XMLNode.child_map:type_name -> tableau.XMLNode.ChildMapEntry
	2,  // 9: tableau.XMLNode.parent:type_name -> tableau.XMLNode
	2,  // 10: tableau.XMLSheet.MetaNodeMapEntry.value:type_name -> tableau.XMLNode
	6,  // 11: tableau.XMLSheet.DataNodeMapEntry.value:type_name -> tableau.XMLSheet.NodeList
	2,  // 12: tableau.XMLSheet.NodeList.nodes:type_name -> tableau.XMLNode
	8,  // 13: tableau.XMLNode.ChildMapEntry.value:type_name -> tableau.XMLNode.IndexList
	11, // 14: tableau.XMLNode.AttrMap.list:type_name -> tableau.XMLNode.AttrMap.Attr
	10, // 15: tableau.XMLNode.AttrMap.map:type_name -> tableau.XMLNode.AttrMap.MapEntry
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_tableau_protobuf_xml_proto_init() }
func file_tableau_protobuf_xml_proto_init() {
	if File_tableau_protobuf_xml_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tableau_protobuf_xml_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMLBook); i {
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
			switch v := v.(*XMLSheet); i {
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
			switch v := v.(*XMLNode); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMLSheet_NodeList); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMLNode_IndexList); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMLNode_AttrMap); i {
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
		file_tableau_protobuf_xml_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XMLNode_AttrMap_Attr); i {
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
			NumMessages:   12,
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
