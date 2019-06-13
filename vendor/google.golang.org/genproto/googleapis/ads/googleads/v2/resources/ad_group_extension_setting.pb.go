// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_extension_setting.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// An ad group extension setting.
type AdGroupExtensionSetting struct {
	// The resource name of the ad group extension setting.
	// AdGroupExtensionSetting resource names have the form:
	//
	// `customers/{customer_id}/adGroupExtensionSettings/{ad_group_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the ad group extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v2.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// The resource name of the ad group. The linked extension feed items will
	// serve under this ad group.
	// AdGroup resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The resource names of the extension feed items to serve under the ad group.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,4,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v2.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *AdGroupExtensionSetting) Reset()         { *m = AdGroupExtensionSetting{} }
func (m *AdGroupExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*AdGroupExtensionSetting) ProtoMessage()    {}
func (*AdGroupExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae21be15f4744e1, []int{0}
}

func (m *AdGroupExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupExtensionSetting.Unmarshal(m, b)
}
func (m *AdGroupExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupExtensionSetting.Marshal(b, m, deterministic)
}
func (m *AdGroupExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupExtensionSetting.Merge(m, src)
}
func (m *AdGroupExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_AdGroupExtensionSetting.Size(m)
}
func (m *AdGroupExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupExtensionSetting proto.InternalMessageInfo

func (m *AdGroupExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *AdGroupExtensionSetting) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupExtensionSetting)(nil), "google.ads.googleads.v2.resources.AdGroupExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_extension_setting.proto", fileDescriptor_aae21be15f4744e1)
}

var fileDescriptor_aae21be15f4744e1 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0xa5, 0x1d, 0x5d, 0xb5, 0xba, 0xfb, 0x50, 0x04, 0xcb, 0x32, 0xc8, 0xac, 0xb2, 0x30, 0x4f,
	0x29, 0xd4, 0x07, 0xa1, 0x8a, 0xd0, 0xc1, 0x75, 0xd0, 0x87, 0x65, 0xe8, 0xca, 0x3c, 0xc8, 0x40,
	0xc9, 0x4e, 0xee, 0xc6, 0xc0, 0x34, 0x09, 0x49, 0x3a, 0xba, 0x9f, 0xe0, 0x1f, 0xf8, 0xec, 0xa3,
	0x9f, 0xe2, 0xa7, 0xf8, 0x15, 0xd2, 0xa6, 0x89, 0x16, 0x99, 0xdd, 0x79, 0x3b, 0xcd, 0x3d, 0xe7,
	0xdc, 0x7b, 0x72, 0xd3, 0x68, 0x46, 0x85, 0xa0, 0x1b, 0x48, 0x31, 0xd1, 0xa9, 0x85, 0x2d, 0xda,
	0x66, 0xa9, 0x02, 0x2d, 0x1a, 0xb5, 0x06, 0x9d, 0x62, 0x52, 0x51, 0x25, 0x1a, 0x59, 0xc1, 0x57,
	0x03, 0x5c, 0x33, 0xc1, 0x2b, 0x0d, 0xc6, 0x30, 0x4e, 0x91, 0x54, 0xc2, 0x88, 0xf8, 0xc4, 0x0a,
	0x11, 0x26, 0x1a, 0x79, 0x0f, 0xb4, 0xcd, 0x90, 0xf7, 0x38, 0x7e, 0xbd, 0xab, 0x0d, 0xf0, 0xa6,
	0xd6, 0xe9, 0x7f, 0xce, 0x15, 0x81, 0x2d, 0x5b, 0x83, 0x6d, 0x70, 0x9c, 0xed, 0xab, 0x36, 0xd7,
	0xd2, 0x69, 0x9e, 0xf6, 0x9a, 0xee, 0xeb, 0xb2, 0xb9, 0x4a, 0xbf, 0x28, 0x2c, 0x25, 0x28, 0xdd,
	0xd7, 0xc7, 0xce, 0x53, 0xb2, 0x14, 0x73, 0x2e, 0x0c, 0x36, 0x4c, 0xf0, 0xbe, 0xfa, 0xec, 0xfb,
	0x28, 0x7a, 0x52, 0x90, 0x79, 0x1b, 0xfb, 0xcc, 0xb9, 0x5f, 0xd8, 0xd1, 0xe2, 0xe7, 0xd1, 0xa1,
	0x0b, 0x56, 0x71, 0x5c, 0x43, 0x12, 0x4c, 0x82, 0xe9, 0x83, 0xf2, 0x91, 0x3b, 0x3c, 0xc7, 0x35,
	0xc4, 0x10, 0x1d, 0x0d, 0xc7, 0x4a, 0xc2, 0x49, 0x30, 0x3d, 0xca, 0xde, 0xa0, 0x5d, 0x97, 0xd5,
	0x65, 0x41, 0xbe, 0xdb, 0xc7, 0x6b, 0x09, 0x67, 0xbc, 0xa9, 0x87, 0x27, 0xe5, 0x21, 0xfc, 0xfb,
	0x19, 0xbf, 0x8c, 0xee, 0xbb, 0xf5, 0x24, 0xa3, 0x49, 0x30, 0x7d, 0x98, 0x8d, 0x5d, 0x03, 0x17,
	0x1c, 0x5d, 0x18, 0xc5, 0x38, 0x5d, 0xe2, 0x4d, 0x03, 0xe5, 0x3d, 0x6c, 0x43, 0xc5, 0xe7, 0xd1,
	0xe3, 0xbf, 0xf3, 0x5d, 0x01, 0x90, 0x8a, 0x19, 0xa8, 0x75, 0x72, 0x67, 0x32, 0xba, 0xd5, 0x24,
	0xf6, 0xca, 0x77, 0x00, 0xe4, 0x7d, 0xab, 0x8b, 0x3f, 0x47, 0x07, 0x76, 0x65, 0xc9, 0xdd, 0x2e,
	0xe7, 0x62, 0xdf, 0x9c, 0xfd, 0xad, 0xbe, 0xed, 0xc4, 0xc3, 0xc0, 0x83, 0x52, 0xd9, 0xfb, 0xcf,
	0xbe, 0x85, 0xd1, 0xe9, 0x5a, 0xd4, 0xe8, 0xd6, 0x47, 0x37, 0x1b, 0xef, 0xd8, 0xe0, 0xa2, 0x0d,
	0xb5, 0x08, 0x3e, 0x7d, 0xe8, 0x2d, 0xa8, 0xd8, 0x60, 0x4e, 0x91, 0x50, 0x34, 0xa5, 0xc0, 0xbb,
	0xc8, 0xee, 0x99, 0x49, 0xa6, 0x6f, 0xf8, 0x35, 0x5e, 0x79, 0xf4, 0x23, 0x1c, 0xcd, 0x8b, 0xe2,
	0x67, 0x78, 0x32, 0xb7, 0x96, 0x05, 0xd1, 0xc8, 0xc2, 0x16, 0x2d, 0x33, 0x54, 0x3a, 0xe6, 0x2f,
	0xc7, 0x59, 0x15, 0x44, 0xaf, 0x3c, 0x67, 0xb5, 0xcc, 0x56, 0x9e, 0xf3, 0x3b, 0x3c, 0xb5, 0x85,
	0x3c, 0x2f, 0x88, 0xce, 0x73, 0xcf, 0xca, 0xf3, 0x65, 0x96, 0xe7, 0x9e, 0x77, 0x79, 0xd0, 0x0d,
	0xfb, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x6f, 0x78, 0xea, 0xc6, 0x03, 0x00, 0x00,
}
