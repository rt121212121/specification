// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assets.proto

package assets

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Message - Membership (MEM)
type Membership struct {
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,1,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	ValidFrom            uint64               `protobuf:"varint,2,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,3,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	ID                   string               `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
	MembershipClass      string               `protobuf:"bytes,5,opt,name=MembershipClass,proto3" json:"MembershipClass,omitempty"`
	RoleType             string               `protobuf:"bytes,6,opt,name=RoleType,proto3" json:"RoleType,omitempty"`
	MembershipType       string               `protobuf:"bytes,7,opt,name=MembershipType,proto3" json:"MembershipType,omitempty"`
	Description          string               `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Membership) Reset()         { *m = Membership{} }
func (m *Membership) String() string { return proto.CompactTextString(m) }
func (*Membership) ProtoMessage()    {}
func (*Membership) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{0}
}

func (m *Membership) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Membership.Unmarshal(m, b)
}
func (m *Membership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Membership.Marshal(b, m, deterministic)
}
func (m *Membership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Membership.Merge(m, src)
}
func (m *Membership) XXX_Size() int {
	return xxx_messageInfo_Membership.Size(m)
}
func (m *Membership) XXX_DiscardUnknown() {
	xxx_messageInfo_Membership.DiscardUnknown(m)
}

var xxx_messageInfo_Membership proto.InternalMessageInfo

func (m *Membership) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *Membership) GetValidFrom() uint64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *Membership) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *Membership) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Membership) GetMembershipClass() string {
	if m != nil {
		return m.MembershipClass
	}
	return ""
}

func (m *Membership) GetRoleType() string {
	if m != nil {
		return m.RoleType
	}
	return ""
}

func (m *Membership) GetMembershipType() string {
	if m != nil {
		return m.MembershipType
	}
	return ""
}

func (m *Membership) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Message - Currency (CUR)
type Currency struct {
	CurrencyCode         string   `protobuf:"bytes,1,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	MonetaryAuthority    string   `protobuf:"bytes,2,opt,name=MonetaryAuthority,proto3" json:"MonetaryAuthority,omitempty"`
	Precision            uint64   `protobuf:"varint,4,opt,name=Precision,proto3" json:"Precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{1}
}

func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Currency) GetMonetaryAuthority() string {
	if m != nil {
		return m.MonetaryAuthority
	}
	return ""
}

func (m *Currency) GetPrecision() uint64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// Message - Share - Common (SHC)
type ShareCommon struct {
	Ticker               string   `protobuf:"bytes,1,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	ISIN                 string   `protobuf:"bytes,2,opt,name=ISIN,proto3" json:"ISIN,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareCommon) Reset()         { *m = ShareCommon{} }
func (m *ShareCommon) String() string { return proto.CompactTextString(m) }
func (*ShareCommon) ProtoMessage()    {}
func (*ShareCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{2}
}

func (m *ShareCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareCommon.Unmarshal(m, b)
}
func (m *ShareCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareCommon.Marshal(b, m, deterministic)
}
func (m *ShareCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareCommon.Merge(m, src)
}
func (m *ShareCommon) XXX_Size() int {
	return xxx_messageInfo_ShareCommon.Size(m)
}
func (m *ShareCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareCommon.DiscardUnknown(m)
}

var xxx_messageInfo_ShareCommon proto.InternalMessageInfo

func (m *ShareCommon) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *ShareCommon) GetISIN() string {
	if m != nil {
		return m.ISIN
	}
	return ""
}

func (m *ShareCommon) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Message - Coupon (COU)
type Coupon struct {
	RedeemingEntity      string   `protobuf:"bytes,1,opt,name=RedeemingEntity,proto3" json:"RedeemingEntity,omitempty"`
	IssueDate            uint64   `protobuf:"varint,2,opt,name=IssueDate,proto3" json:"IssueDate,omitempty"`
	ExpiryDate           uint64   `protobuf:"varint,3,opt,name=ExpiryDate,proto3" json:"ExpiryDate,omitempty"`
	Value                uint64   `protobuf:"varint,4,opt,name=Value,proto3" json:"Value,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	Precision            uint64   `protobuf:"varint,7,opt,name=Precision,proto3" json:"Precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coupon) Reset()         { *m = Coupon{} }
func (m *Coupon) String() string { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()    {}
func (*Coupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{3}
}

func (m *Coupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coupon.Unmarshal(m, b)
}
func (m *Coupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coupon.Marshal(b, m, deterministic)
}
func (m *Coupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coupon.Merge(m, src)
}
func (m *Coupon) XXX_Size() int {
	return xxx_messageInfo_Coupon.Size(m)
}
func (m *Coupon) XXX_DiscardUnknown() {
	xxx_messageInfo_Coupon.DiscardUnknown(m)
}

var xxx_messageInfo_Coupon proto.InternalMessageInfo

func (m *Coupon) GetRedeemingEntity() string {
	if m != nil {
		return m.RedeemingEntity
	}
	return ""
}

func (m *Coupon) GetIssueDate() uint64 {
	if m != nil {
		return m.IssueDate
	}
	return 0
}

func (m *Coupon) GetExpiryDate() uint64 {
	if m != nil {
		return m.ExpiryDate
	}
	return 0
}

func (m *Coupon) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Coupon) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Coupon) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Coupon) GetPrecision() uint64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// Message - Loyalty Points (LOY)
type LoyaltyPoints struct {
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,1,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	OfferName            string               `protobuf:"bytes,2,opt,name=OfferName,proto3" json:"OfferName,omitempty"`
	ValidFrom            uint64               `protobuf:"varint,3,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,4,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoyaltyPoints) Reset()         { *m = LoyaltyPoints{} }
func (m *LoyaltyPoints) String() string { return proto.CompactTextString(m) }
func (*LoyaltyPoints) ProtoMessage()    {}
func (*LoyaltyPoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{4}
}

func (m *LoyaltyPoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoyaltyPoints.Unmarshal(m, b)
}
func (m *LoyaltyPoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoyaltyPoints.Marshal(b, m, deterministic)
}
func (m *LoyaltyPoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoyaltyPoints.Merge(m, src)
}
func (m *LoyaltyPoints) XXX_Size() int {
	return xxx_messageInfo_LoyaltyPoints.Size(m)
}
func (m *LoyaltyPoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LoyaltyPoints.DiscardUnknown(m)
}

var xxx_messageInfo_LoyaltyPoints proto.InternalMessageInfo

func (m *LoyaltyPoints) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *LoyaltyPoints) GetOfferName() string {
	if m != nil {
		return m.OfferName
	}
	return ""
}

func (m *LoyaltyPoints) GetValidFrom() uint64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *LoyaltyPoints) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *LoyaltyPoints) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Message - Ticket (Admission) (TIC)
type TicketAdmission struct {
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,1,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	AdmissionType        string               `protobuf:"bytes,2,opt,name=AdmissionType,proto3" json:"AdmissionType,omitempty"`
	Venue                string               `protobuf:"bytes,3,opt,name=Venue,proto3" json:"Venue,omitempty"`
	Class                string               `protobuf:"bytes,4,opt,name=Class,proto3" json:"Class,omitempty"`
	Area                 string               `protobuf:"bytes,5,opt,name=Area,proto3" json:"Area,omitempty"`
	Seat                 string               `protobuf:"bytes,6,opt,name=Seat,proto3" json:"Seat,omitempty"`
	StartTimeDate        uint64               `protobuf:"varint,7,opt,name=StartTimeDate,proto3" json:"StartTimeDate,omitempty"`
	ValidFrom            uint64               `protobuf:"varint,8,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,9,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	Description          string               `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TicketAdmission) Reset()         { *m = TicketAdmission{} }
func (m *TicketAdmission) String() string { return proto.CompactTextString(m) }
func (*TicketAdmission) ProtoMessage()    {}
func (*TicketAdmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{5}
}

func (m *TicketAdmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TicketAdmission.Unmarshal(m, b)
}
func (m *TicketAdmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TicketAdmission.Marshal(b, m, deterministic)
}
func (m *TicketAdmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketAdmission.Merge(m, src)
}
func (m *TicketAdmission) XXX_Size() int {
	return xxx_messageInfo_TicketAdmission.Size(m)
}
func (m *TicketAdmission) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketAdmission.DiscardUnknown(m)
}

var xxx_messageInfo_TicketAdmission proto.InternalMessageInfo

func (m *TicketAdmission) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *TicketAdmission) GetAdmissionType() string {
	if m != nil {
		return m.AdmissionType
	}
	return ""
}

func (m *TicketAdmission) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *TicketAdmission) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *TicketAdmission) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *TicketAdmission) GetSeat() string {
	if m != nil {
		return m.Seat
	}
	return ""
}

func (m *TicketAdmission) GetStartTimeDate() uint64 {
	if m != nil {
		return m.StartTimeDate
	}
	return 0
}

func (m *TicketAdmission) GetValidFrom() uint64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *TicketAdmission) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *TicketAdmission) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Message - Casino Chip (CHP)
type CasinoChip struct {
	CurrencyCode         string               `protobuf:"bytes,1,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	UseType              string               `protobuf:"bytes,2,opt,name=UseType,proto3" json:"UseType,omitempty"`
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,3,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	ValidFrom            uint64               `protobuf:"varint,4,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,5,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	Precision            uint64               `protobuf:"varint,6,opt,name=Precision,proto3" json:"Precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CasinoChip) Reset()         { *m = CasinoChip{} }
func (m *CasinoChip) String() string { return proto.CompactTextString(m) }
func (*CasinoChip) ProtoMessage()    {}
func (*CasinoChip) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{6}
}

func (m *CasinoChip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CasinoChip.Unmarshal(m, b)
}
func (m *CasinoChip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CasinoChip.Marshal(b, m, deterministic)
}
func (m *CasinoChip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CasinoChip.Merge(m, src)
}
func (m *CasinoChip) XXX_Size() int {
	return xxx_messageInfo_CasinoChip.Size(m)
}
func (m *CasinoChip) XXX_DiscardUnknown() {
	xxx_messageInfo_CasinoChip.DiscardUnknown(m)
}

var xxx_messageInfo_CasinoChip proto.InternalMessageInfo

func (m *CasinoChip) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *CasinoChip) GetUseType() string {
	if m != nil {
		return m.UseType
	}
	return ""
}

func (m *CasinoChip) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *CasinoChip) GetValidFrom() uint64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *CasinoChip) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *CasinoChip) GetPrecision() uint64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// Field - Age Restriction
type AgeRestrictionField struct {
	Lower                uint32   `protobuf:"varint,1,opt,name=Lower,proto3" json:"Lower,omitempty"`
	Upper                uint32   `protobuf:"varint,2,opt,name=Upper,proto3" json:"Upper,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgeRestrictionField) Reset()         { *m = AgeRestrictionField{} }
func (m *AgeRestrictionField) String() string { return proto.CompactTextString(m) }
func (*AgeRestrictionField) ProtoMessage()    {}
func (*AgeRestrictionField) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{7}
}

func (m *AgeRestrictionField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgeRestrictionField.Unmarshal(m, b)
}
func (m *AgeRestrictionField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgeRestrictionField.Marshal(b, m, deterministic)
}
func (m *AgeRestrictionField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgeRestrictionField.Merge(m, src)
}
func (m *AgeRestrictionField) XXX_Size() int {
	return xxx_messageInfo_AgeRestrictionField.Size(m)
}
func (m *AgeRestrictionField) XXX_DiscardUnknown() {
	xxx_messageInfo_AgeRestrictionField.DiscardUnknown(m)
}

var xxx_messageInfo_AgeRestrictionField proto.InternalMessageInfo

func (m *AgeRestrictionField) GetLower() uint32 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *AgeRestrictionField) GetUpper() uint32 {
	if m != nil {
		return m.Upper
	}
	return 0
}

func init() {
	proto.RegisterType((*Membership)(nil), "assets.Membership")
	proto.RegisterType((*Currency)(nil), "assets.Currency")
	proto.RegisterType((*ShareCommon)(nil), "assets.ShareCommon")
	proto.RegisterType((*Coupon)(nil), "assets.Coupon")
	proto.RegisterType((*LoyaltyPoints)(nil), "assets.LoyaltyPoints")
	proto.RegisterType((*TicketAdmission)(nil), "assets.TicketAdmission")
	proto.RegisterType((*CasinoChip)(nil), "assets.CasinoChip")
	proto.RegisterType((*AgeRestrictionField)(nil), "assets.AgeRestrictionField")
}

func init() { proto.RegisterFile("assets.proto", fileDescriptor_610ca40ce07a87fe) }

var fileDescriptor_610ca40ce07a87fe = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x51, 0x6b, 0xdb, 0x3a,
	0x14, 0xc7, 0x71, 0xe2, 0xa4, 0xc9, 0x69, 0xd3, 0xde, 0xab, 0x5e, 0x2e, 0x61, 0x1b, 0x23, 0x84,
	0x31, 0xf2, 0x30, 0x9a, 0xb1, 0x31, 0xf6, 0x9c, 0x39, 0x2d, 0x64, 0xb4, 0x5d, 0x71, 0xda, 0x3e,
	0x6c, 0x4f, 0xaa, 0x7d, 0x9a, 0x88, 0xda, 0x96, 0x91, 0x64, 0x36, 0xef, 0x69, 0x5f, 0x6d, 0x5f,
	0x65, 0xec, 0x61, 0x5f, 0x62, 0x30, 0x24, 0x39, 0x49, 0xed, 0xb4, 0x10, 0x58, 0xdf, 0x74, 0xfe,
	0x3a, 0x8e, 0x74, 0x7e, 0xff, 0x73, 0x14, 0xd8, 0xa1, 0x52, 0xa2, 0x92, 0x07, 0xa9, 0xe0, 0x8a,
	0x93, 0xa6, 0x8d, 0xfa, 0xdf, 0x6b, 0x00, 0x27, 0x18, 0x5f, 0xa1, 0x90, 0x73, 0x96, 0x12, 0x0f,
	0x76, 0x47, 0x33, 0xf4, 0x51, 0x2a, 0xc1, 0x02, 0xc5, 0x78, 0xd2, 0x75, 0x7a, 0xce, 0x60, 0xfb,
	0xd5, 0xe3, 0x83, 0xe2, 0xeb, 0xf2, 0xee, 0x11, 0xc3, 0x28, 0xf4, 0x2b, 0x9f, 0x90, 0x27, 0xd0,
	0xbe, 0xa4, 0x11, 0x0b, 0x8f, 0x04, 0x8f, 0xbb, 0xb5, 0x9e, 0x33, 0x70, 0xfd, 0x95, 0x40, 0x5e,
	0xc2, 0xfe, 0xe1, 0x97, 0x94, 0x09, 0xaa, 0x73, 0xcf, 0x59, 0x8c, 0x52, 0xd1, 0x38, 0xed, 0xd6,
	0x4d, 0xde, 0x5d, 0x5b, 0x64, 0x17, 0x6a, 0x93, 0x71, 0xd7, 0xed, 0x39, 0x83, 0xb6, 0x5f, 0x9b,
	0x8c, 0xc9, 0x00, 0xf6, 0x56, 0x57, 0xf6, 0x22, 0x2a, 0x65, 0xb7, 0x61, 0x36, 0xab, 0x32, 0x79,
	0x04, 0x2d, 0x9f, 0x47, 0x78, 0x9e, 0xa7, 0xd8, 0x6d, 0x9a, 0x94, 0x65, 0x4c, 0x9e, 0xc3, 0xee,
	0x2a, 0xdd, 0x64, 0x6c, 0x99, 0x8c, 0x8a, 0x4a, 0x7a, 0xb0, 0x3d, 0x46, 0x19, 0x08, 0x96, 0x1a,
	0x1e, 0x2d, 0x93, 0x74, 0x5b, 0xea, 0x7f, 0x73, 0xa0, 0xe5, 0x65, 0x42, 0x60, 0x12, 0xe4, 0xa4,
	0x0f, 0x3b, 0x8b, 0xb5, 0xc7, 0x43, 0x34, 0xfc, 0xda, 0x7e, 0x49, 0x23, 0x2f, 0xe0, 0xdf, 0x13,
	0x9e, 0xa0, 0xa2, 0x22, 0x1f, 0x65, 0x6a, 0xce, 0x05, 0x53, 0xb9, 0x01, 0xd5, 0xf6, 0xd7, 0x37,
	0x34, 0xce, 0x33, 0x81, 0x01, 0x93, 0xfa, 0x78, 0xd7, 0xe2, 0x5c, 0x0a, 0xef, 0xdd, 0x56, 0xfd,
	0x1f, 0xb7, 0xff, 0x09, 0xb6, 0xa7, 0x73, 0x2a, 0xd0, 0xe3, 0x71, 0xcc, 0x13, 0xf2, 0x3f, 0x34,
	0xcf, 0x59, 0x70, 0x83, 0xa2, 0x38, 0xbe, 0x88, 0x08, 0x01, 0x77, 0x32, 0x9d, 0x9c, 0x16, 0x67,
	0x99, 0x75, 0xb5, 0xbe, 0xfa, 0x7a, 0x7d, 0x3f, 0x1c, 0x68, 0x7a, 0x3c, 0x4b, 0x79, 0xa2, 0xd1,
	0xfb, 0x18, 0x22, 0xc6, 0x2c, 0x99, 0x1d, 0x26, 0x4a, 0xdf, 0xdb, 0x9e, 0x50, 0x95, 0xf5, 0xad,
	0x27, 0x52, 0x66, 0x38, 0xa6, 0x0a, 0x17, 0x4d, 0xb0, 0x14, 0xc8, 0x53, 0x00, 0xe3, 0x74, 0x6e,
	0xb6, 0xad, 0xf7, 0xb7, 0x14, 0xf2, 0x1f, 0x34, 0x2e, 0x69, 0x94, 0x61, 0x51, 0xaf, 0x0d, 0xb4,
	0x9d, 0x0b, 0x8e, 0x85, 0xe3, 0x2b, 0xee, 0x95, 0x32, 0x9a, 0x6b, 0x65, 0x94, 0x39, 0x6e, 0x55,
	0x38, 0xf6, 0x7f, 0x3a, 0xd0, 0x39, 0xe6, 0x39, 0x8d, 0x54, 0x7e, 0xc6, 0x59, 0xa2, 0xe4, 0x83,
	0xcd, 0xc2, 0x87, 0xeb, 0x6b, 0x14, 0xa7, 0x34, 0xc6, 0x02, 0xfb, 0x4a, 0x28, 0x4f, 0x4a, 0x7d,
	0xc3, 0x49, 0x71, 0xef, 0x9f, 0x94, 0x0a, 0x84, 0xc6, 0xba, 0x97, 0xbf, 0x6a, 0xb0, 0x67, 0x9a,
	0x41, 0x8d, 0xc2, 0x98, 0x49, 0x5d, 0xfa, 0xc3, 0x14, 0xfa, 0x0c, 0x3a, 0xcb, 0x5f, 0x34, 0xd3,
	0x64, 0x8b, 0x2d, 0x8b, 0xc6, 0x57, 0x4c, 0x32, 0x2c, 0xda, 0xcc, 0x06, 0x5a, 0xb5, 0x63, 0x6c,
	0x67, 0xdc, 0x06, 0xba, 0x59, 0x47, 0x02, 0x69, 0x51, 0x85, 0x59, 0x6b, 0x6d, 0x8a, 0x54, 0x15,
	0xf6, 0x9a, 0xb5, 0x3e, 0x79, 0xaa, 0xa8, 0x50, 0x1a, 0x83, 0x69, 0x27, 0xeb, 0x6d, 0x59, 0x2c,
	0xa3, 0x6e, 0x6d, 0x88, 0xba, 0xbd, 0x31, 0x6a, 0x58, 0x47, 0xfd, 0xdb, 0x01, 0xf0, 0xa8, 0x64,
	0x09, 0xf7, 0xf4, 0xd3, 0xba, 0xc9, 0xc3, 0xd0, 0x85, 0xad, 0x0b, 0x89, 0xb7, 0xf0, 0x2d, 0xc2,
	0x3b, 0x3c, 0xaa, 0xff, 0xe5, 0xc3, 0xec, 0x6e, 0xc8, 0xa0, 0x71, 0x3f, 0x83, 0xd2, 0x44, 0x35,
	0xab, 0x13, 0x35, 0x82, 0xfd, 0x3b, 0x2e, 0xa5, 0xcd, 0x3e, 0xe6, 0x9f, 0x8b, 0xa7, 0xa9, 0xe3,
	0xdb, 0x40, 0xab, 0x17, 0x69, 0x8a, 0xc2, 0xd4, 0xdd, 0xf1, 0x6d, 0xf0, 0xee, 0xed, 0xc7, 0x37,
	0x33, 0xa6, 0xe6, 0xd9, 0xd5, 0x41, 0xc0, 0xe3, 0xa1, 0xe2, 0x37, 0x98, 0xb0, 0xaf, 0x18, 0x0e,
	0x65, 0x8a, 0x01, 0xbb, 0x66, 0x81, 0xb9, 0xcf, 0x30, 0x64, 0x52, 0x0d, 0x67, 0x3c, 0xa2, 0xc9,
	0x6c, 0x68, 0x69, 0x5c, 0x35, 0xcd, 0xbf, 0xdc, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63,
	0xe1, 0x6f, 0xcb, 0xf5, 0x06, 0x00, 0x00,
}
