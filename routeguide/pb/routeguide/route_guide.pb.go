// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routeguide/route_guide.proto

package routeguide

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude             int32    `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude            int32    `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_ccab5f943d2922ba, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi                   *Point   `protobuf:"bytes,2,opt,name=hi" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_ccab5f943d2922ba, []int{1}
}
func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (dst *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(dst, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Feature struct {
	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The point where the feature is detected.
	Location             *Point   `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_ccab5f943d2922ba, []int{2}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// The message to be sent.
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteNote) Reset()         { *m = RouteNote{} }
func (m *RouteNote) String() string { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()    {}
func (*RouteNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_ccab5f943d2922ba, []int{3}
}
func (m *RouteNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteNote.Unmarshal(m, b)
}
func (m *RouteNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteNote.Marshal(b, m, deterministic)
}
func (dst *RouteNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteNote.Merge(dst, src)
}
func (m *RouteNote) XXX_Size() int {
	return xxx_messageInfo_RouteNote.Size(m)
}
func (m *RouteNote) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteNote.DiscardUnknown(m)
}

var xxx_messageInfo_RouteNote proto.InternalMessageInfo

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime          int32    `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteSummary) Reset()         { *m = RouteSummary{} }
func (m *RouteSummary) String() string { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()    {}
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_ccab5f943d2922ba, []int{4}
}
func (m *RouteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSummary.Unmarshal(m, b)
}
func (m *RouteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSummary.Marshal(b, m, deterministic)
}
func (dst *RouteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSummary.Merge(dst, src)
}
func (m *RouteSummary) XXX_Size() int {
	return xxx_messageInfo_RouteSummary.Size(m)
}
func (m *RouteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSummary proto.InternalMessageInfo

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "routeguide.Point")
	proto.RegisterType((*Rectangle)(nil), "routeguide.Rectangle")
	proto.RegisterType((*Feature)(nil), "routeguide.Feature")
	proto.RegisterType((*RouteNote)(nil), "routeguide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "routeguide.RouteSummary")
}

func init() {
	proto.RegisterFile("routeguide/route_guide.proto", fileDescriptor_route_guide_ccab5f943d2922ba)
}

var fileDescriptor_route_guide_ccab5f943d2922ba = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0x89, 0xbb, 0x6e, 0x73, 0x13, 0x11, 0xaf, 0x08, 0x21, 0x2e, 0xe8, 0xc6, 0x97, 0xbe,
	0x18, 0x4b, 0x05, 0x1f, 0x2b, 0xb6, 0x60, 0x5f, 0x8a, 0xd4, 0xd8, 0xf7, 0x32, 0x26, 0xd7, 0x74,
	0x60, 0x92, 0x09, 0xc9, 0x04, 0xf4, 0x07, 0xf8, 0x0b, 0xfc, 0xc3, 0x92, 0x49, 0xd2, 0xa4, 0xda,
	0xb2, 0x6f, 0x73, 0xcf, 0x3d, 0xe7, 0x7e, 0x9c, 0xcb, 0xc0, 0x5d, 0xa9, 0x6a, 0x4d, 0x69, 0x2d,
	0x12, 0x7a, 0x67, 0x9e, 0x7b, 0xf3, 0x0e, 0x8b, 0x52, 0x69, 0x85, 0x30, 0x64, 0x83, 0x4f, 0x70,
	0xb3, 0x55, 0x22, 0xd7, 0xe8, 0xc3, 0x44, 0x72, 0x2d, 0x74, 0x9d, 0x90, 0xc7, 0x5e, 0xb3, 0xe9,
	0x4d, 0x74, 0x8c, 0xf1, 0x0e, 0x6c, 0xa9, 0xf2, 0xb4, 0x4d, 0x5a, 0x26, 0x39, 0x00, 0xc1, 0x57,
	0xb0, 0x23, 0x8a, 0x35, 0xcf, 0x53, 0x49, 0x78, 0x0f, 0x96, 0x54, 0xa6, 0x80, 0x33, 0x7f, 0x16,
	0x0e, 0x8d, 0x42, 0xd3, 0x25, 0xb2, 0xa4, 0x6a, 0x28, 0x07, 0x61, 0xca, 0x9c, 0xa7, 0x1c, 0x44,
	0xb0, 0x81, 0xdb, 0xcf, 0xc4, 0x75, 0x5d, 0x12, 0x22, 0x5c, 0xe7, 0x3c, 0x6b, 0x67, 0xb2, 0x23,
	0xf3, 0xc6, 0xb7, 0x30, 0x91, 0x2a, 0xe6, 0x5a, 0xa8, 0xfc, 0x72, 0x9d, 0x23, 0x25, 0xd8, 0x81,
	0x1d, 0x35, 0xd9, 0x2f, 0x4a, 0x9f, 0x6a, 0xd9, 0x83, 0x5a, 0xf4, 0xe0, 0x36, 0xa3, 0xaa, 0xe2,
	0x69, 0xbb, 0xb8, 0x1d, 0xf5, 0x61, 0xf0, 0x87, 0x81, 0x6b, 0xca, 0x7e, 0xab, 0xb3, 0x8c, 0x97,
	0xbf, 0xf0, 0x15, 0x38, 0x45, 0xa3, 0xde, 0xc7, 0xaa, 0xce, 0x75, 0x67, 0x22, 0x18, 0x68, 0xd5,
	0x20, 0xf8, 0x06, 0x9e, 0xfc, 0x68, 0xb7, 0xea, 0x28, 0xad, 0x95, 0x6e, 0x07, 0xb6, 0x24, 0x1f,
	0x26, 0x89, 0xa8, 0x34, 0xcf, 0x63, 0xf2, 0x1e, 0xb5, 0x77, 0xe8, 0x63, 0xbc, 0x07, 0x97, 0x24,
	0x2f, 0x2a, 0x4a, 0xf6, 0x5a, 0x64, 0xe4, 0x5d, 0x9b, 0xbc, 0xd3, 0x61, 0x3b, 0x91, 0xd1, 0xfc,
	0xb7, 0x05, 0x60, 0xa6, 0x5a, 0x37, 0xeb, 0xe0, 0x07, 0x80, 0x35, 0xe9, 0xde, 0xcb, 0xff, 0x37,
	0xf5, 0x9f, 0x8f, 0xa1, 0x8e, 0x17, 0x5c, 0xe1, 0x02, 0xdc, 0x8d, 0xa8, 0x7a, 0x61, 0x85, 0x2f,
	0xc6, 0xb4, 0xe3, 0xb5, 0x2f, 0xa8, 0x67, 0x0c, 0x17, 0xe0, 0x44, 0x14, 0xab, 0x32, 0x31, 0xb3,
	0x9c, 0x6b, 0xec, 0x9d, 0x54, 0x1c, 0xf9, 0x18, 0x5c, 0x4d, 0x19, 0x7e, 0xec, 0x4e, 0xb6, 0x3a,
	0x70, 0xfd, 0x4f, 0xf3, 0xfe, 0x92, 0xfe, 0x79, 0xb8, 0x91, 0xcf, 0xd8, 0x72, 0x06, 0x2f, 0x85,
	0x0a, 0xd3, 0xb2, 0x88, 0x43, 0xfa, 0xc9, 0xb3, 0x42, 0x52, 0x35, 0xa2, 0x2f, 0x9f, 0x0e, 0x1e,
	0x6d, 0x9b, 0x3f, 0xb1, 0x65, 0xdf, 0x1f, 0x9b, 0xcf, 0xf1, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0xb3, 0xf3, 0xe9, 0x3c, 0x03, 0x00, 0x00,
}
