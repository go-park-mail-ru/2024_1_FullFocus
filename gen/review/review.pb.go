// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.12.4
// source: review.proto

package gen

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CreateProductReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID  uint32             `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	ProfileID  uint32             `protobuf:"varint,2,opt,name=ProfileID,proto3" json:"ProfileID,omitempty"`
	ReviewData *ProductReviewData `protobuf:"bytes,3,opt,name=reviewData,proto3" json:"reviewData,omitempty"`
}

func (x *CreateProductReviewRequest) Reset() {
	*x = CreateProductReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReviewRequest) ProtoMessage() {}

func (x *CreateProductReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateProductReviewRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductReviewRequest) GetProductID() uint32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *CreateProductReviewRequest) GetProfileID() uint32 {
	if x != nil {
		return x.ProfileID
	}
	return 0
}

func (x *CreateProductReviewRequest) GetReviewData() *ProductReviewData {
	if x != nil {
		return x.ReviewData
	}
	return nil
}

type GetProductReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID    uint32 `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	LastReviewID uint32 `protobuf:"varint,2,opt,name=LastReviewID,proto3" json:"LastReviewID,omitempty"`
	Limit        uint32 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	SortingQuery string `protobuf:"bytes,4,opt,name=SortingQuery,proto3" json:"SortingQuery,omitempty"`
}

func (x *GetProductReviewsRequest) Reset() {
	*x = GetProductReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReviewsRequest) ProtoMessage() {}

func (x *GetProductReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetProductReviewsRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductReviewsRequest) GetProductID() uint32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *GetProductReviewsRequest) GetLastReviewID() uint32 {
	if x != nil {
		return x.LastReviewID
	}
	return 0
}

func (x *GetProductReviewsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetProductReviewsRequest) GetSortingQuery() string {
	if x != nil {
		return x.SortingQuery
	}
	return ""
}

type GetProductReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*ProductReview `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *GetProductReviewsResponse) Reset() {
	*x = GetProductReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReviewsResponse) ProtoMessage() {}

func (x *GetProductReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetProductReviewsResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductReviewsResponse) GetReviews() []*ProductReview {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type ProductReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewID   uint32             `protobuf:"varint,1,opt,name=ReviewID,proto3" json:"ReviewID,omitempty"`
	ProfileID  uint32             `protobuf:"varint,2,opt,name=ProfileID,proto3" json:"ProfileID,omitempty"`
	ReviewData *ProductReviewData `protobuf:"bytes,3,opt,name=reviewData,proto3" json:"reviewData,omitempty"`
	CreatedAt  string             `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *ProductReview) Reset() {
	*x = ProductReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReview) ProtoMessage() {}

func (x *ProductReview) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReview.ProtoReflect.Descriptor instead.
func (*ProductReview) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{3}
}

func (x *ProductReview) GetReviewID() uint32 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *ProductReview) GetProfileID() uint32 {
	if x != nil {
		return x.ProfileID
	}
	return 0
}

func (x *ProductReview) GetReviewData() *ProductReviewData {
	if x != nil {
		return x.ReviewData
	}
	return nil
}

func (x *ProductReview) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ProductReviewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating        uint32 `protobuf:"varint,1,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Advantages    string `protobuf:"bytes,2,opt,name=Advantages,proto3" json:"Advantages,omitempty"`
	Disadvantages string `protobuf:"bytes,3,opt,name=Disadvantages,proto3" json:"Disadvantages,omitempty"`
	Comment       string `protobuf:"bytes,4,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *ProductReviewData) Reset() {
	*x = ProductReviewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductReviewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductReviewData) ProtoMessage() {}

func (x *ProductReviewData) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductReviewData.ProtoReflect.Descriptor instead.
func (*ProductReviewData) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{4}
}

func (x *ProductReviewData) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *ProductReviewData) GetAdvantages() string {
	if x != nil {
		return x.Advantages
	}
	return ""
}

func (x *ProductReviewData) GetDisadvantages() string {
	if x != nil {
		return x.Disadvantages
	}
	return ""
}

func (x *ProductReviewData) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_review_proto protoreflect.FileDescriptor

var file_review_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x39,
	0x0a, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x96, 0x01, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x22, 0x4c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0a,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x69, 0x73, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x32, 0xb9, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x53,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32, 0x30,
	0x32, 0x34, 0x5f, 0x31, 0x5f, 0x46, 0x75, 0x6c, 0x6c, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x6e,
	0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_proto_rawDescOnce sync.Once
	file_review_proto_rawDescData = file_review_proto_rawDesc
)

func file_review_proto_rawDescGZIP() []byte {
	file_review_proto_rawDescOnce.Do(func() {
		file_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_proto_rawDescData)
	})
	return file_review_proto_rawDescData
}

var file_review_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_review_proto_goTypes = []interface{}{
	(*CreateProductReviewRequest)(nil), // 0: review.CreateProductReviewRequest
	(*GetProductReviewsRequest)(nil),   // 1: review.GetProductReviewsRequest
	(*GetProductReviewsResponse)(nil),  // 2: review.GetProductReviewsResponse
	(*ProductReview)(nil),              // 3: review.ProductReview
	(*ProductReviewData)(nil),          // 4: review.ProductReviewData
	(*empty.Empty)(nil),                // 5: google.protobuf.Empty
}
var file_review_proto_depIdxs = []int32{
	4, // 0: review.CreateProductReviewRequest.reviewData:type_name -> review.ProductReviewData
	3, // 1: review.GetProductReviewsResponse.reviews:type_name -> review.ProductReview
	4, // 2: review.ProductReview.reviewData:type_name -> review.ProductReviewData
	0, // 3: review.Review.CreateProductReview:input_type -> review.CreateProductReviewRequest
	1, // 4: review.Review.GetProductReviews:input_type -> review.GetProductReviewsRequest
	5, // 5: review.Review.CreateProductReview:output_type -> google.protobuf.Empty
	2, // 6: review.Review.GetProductReviews:output_type -> review.GetProductReviewsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_review_proto_init() }
func file_review_proto_init() {
	if File_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductReviewRequest); i {
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
		file_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReviewsRequest); i {
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
		file_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReviewsResponse); i {
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
		file_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReview); i {
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
		file_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductReviewData); i {
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
			RawDescriptor: file_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_proto_goTypes,
		DependencyIndexes: file_review_proto_depIdxs,
		MessageInfos:      file_review_proto_msgTypes,
	}.Build()
	File_review_proto = out.File
	file_review_proto_rawDesc = nil
	file_review_proto_goTypes = nil
	file_review_proto_depIdxs = nil
}
