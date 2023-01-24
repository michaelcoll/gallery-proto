// Photo related messages and service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: photo/v1/photo.proto

package photov1

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

// GetPhotosRequest represent a search query with pagination options to
// indicate which results to include in the response.
type GetPhotosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`   // the number of element to retrieve
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"` // the position of the first element to retrieve
}

func (x *GetPhotosRequest) Reset() {
	*x = GetPhotosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhotosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotosRequest) ProtoMessage() {}

func (x *GetPhotosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotosRequest.ProtoReflect.Descriptor instead.
func (*GetPhotosRequest) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{0}
}

func (x *GetPhotosRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPhotosRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetPhotosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photos []*Photo `protobuf:"bytes,1,rep,name=photos,proto3" json:"photos,omitempty"` // the photos
}

func (x *GetPhotosResponse) Reset() {
	*x = GetPhotosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhotosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhotosResponse) ProtoMessage() {}

func (x *GetPhotosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhotosResponse.ProtoReflect.Descriptor instead.
func (*GetPhotosResponse) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{1}
}

func (x *GetPhotosResponse) GetPhotos() []*Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

// GetByHashRequest represent a search query with a photo hash
type GetByHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"` // a photo hash
}

func (x *GetByHashRequest) Reset() {
	*x = GetByHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHashRequest) ProtoMessage() {}

func (x *GetByHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByHashRequest.ProtoReflect.Descriptor instead.
func (*GetByHashRequest) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{2}
}

func (x *GetByHashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetByHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photo *Photo `protobuf:"bytes,1,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *GetByHashResponse) Reset() {
	*x = GetByHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHashResponse) ProtoMessage() {}

func (x *GetByHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByHashResponse.ProtoReflect.Descriptor instead.
func (*GetByHashResponse) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{3}
}

func (x *GetByHashResponse) GetPhoto() *Photo {
	if x != nil {
		return x.Photo
	}
	return nil
}

// ExistsByHashRequest represent an existence search query with a photo hash
type ExistsByHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"` // a photo hash
}

func (x *ExistsByHashRequest) Reset() {
	*x = ExistsByHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsByHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsByHashRequest) ProtoMessage() {}

func (x *ExistsByHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsByHashRequest.ProtoReflect.Descriptor instead.
func (*ExistsByHashRequest) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{4}
}

func (x *ExistsByHashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type ExistsByHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *ExistsByHashResponse) Reset() {
	*x = ExistsByHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsByHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsByHashResponse) ProtoMessage() {}

func (x *ExistsByHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsByHashResponse.ProtoReflect.Descriptor instead.
func (*ExistsByHashResponse) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{5}
}

func (x *ExistsByHashResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

// ContentByHashRequest represent a search query with a photo hash
type ContentByHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"` // a photo hash
}

func (x *ContentByHashRequest) Reset() {
	*x = ContentByHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentByHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentByHashRequest) ProtoMessage() {}

func (x *ContentByHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentByHashRequest.ProtoReflect.Descriptor instead.
func (*ContentByHashRequest) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{6}
}

func (x *ContentByHashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type PhotoServiceContentByHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *PhotoServiceContentByHashResponse) Reset() {
	*x = PhotoServiceContentByHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoServiceContentByHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoServiceContentByHashResponse) ProtoMessage() {}

func (x *PhotoServiceContentByHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoServiceContentByHashResponse.ProtoReflect.Descriptor instead.
func (*PhotoServiceContentByHashResponse) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{7}
}

func (x *PhotoServiceContentByHashResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PhotoServiceContentByHashResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// ThumbnailByHashRequest represent a search query with a photo hash, a width or an height.
type ThumbnailByHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`      // a photo hash
	Width  uint32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`   // the requested width for the thumbnail
	Height uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"` // the requested height for the thumbnail
}

func (x *ThumbnailByHashRequest) Reset() {
	*x = ThumbnailByHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbnailByHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbnailByHashRequest) ProtoMessage() {}

func (x *ThumbnailByHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbnailByHashRequest.ProtoReflect.Descriptor instead.
func (*ThumbnailByHashRequest) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{8}
}

func (x *ThumbnailByHashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *ThumbnailByHashRequest) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ThumbnailByHashRequest) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type PhotoServiceThumbnailByHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *PhotoServiceThumbnailByHashResponse) Reset() {
	*x = PhotoServiceThumbnailByHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoServiceThumbnailByHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoServiceThumbnailByHashResponse) ProtoMessage() {}

func (x *PhotoServiceThumbnailByHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoServiceThumbnailByHashResponse.ProtoReflect.Descriptor instead.
func (*PhotoServiceThumbnailByHashResponse) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{9}
}

func (x *PhotoServiceThumbnailByHashResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PhotoServiceThumbnailByHashResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// Photo contains basic info of a photo and its metadata
type Photo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash         string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                                     // the photo hash
	Path         string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`                                     // the photo path on the daemon fs
	DateTime     string `protobuf:"bytes,3,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`             // the date
	Iso          uint32 `protobuf:"varint,4,opt,name=iso,proto3" json:"iso,omitempty"`                                      // the ISO of the image
	ExposureTime string `protobuf:"bytes,5,opt,name=exposure_time,json=exposureTime,proto3" json:"exposure_time,omitempty"` // the exposure as a fraction
	XDimension   uint32 `protobuf:"varint,6,opt,name=x_dimension,json=xDimension,proto3" json:"x_dimension,omitempty"`      // the x dimension of the image
	YDimension   uint32 `protobuf:"varint,7,opt,name=y_dimension,json=yDimension,proto3" json:"y_dimension,omitempty"`      // the y dimension of the image
	Model        string `protobuf:"bytes,8,opt,name=model,proto3" json:"model,omitempty"`                                   // the model of the camera that took the image
	FNumber      string `protobuf:"bytes,9,opt,name=f_number,json=fNumber,proto3" json:"f_number,omitempty"`                // the f Number
}

func (x *Photo) Reset() {
	*x = Photo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photo_v1_photo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Photo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Photo) ProtoMessage() {}

func (x *Photo) ProtoReflect() protoreflect.Message {
	mi := &file_photo_v1_photo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Photo.ProtoReflect.Descriptor instead.
func (*Photo) Descriptor() ([]byte, []int) {
	return file_photo_v1_photo_proto_rawDescGZIP(), []int{10}
}

func (x *Photo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Photo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Photo) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *Photo) GetIso() uint32 {
	if x != nil {
		return x.Iso
	}
	return 0
}

func (x *Photo) GetExposureTime() string {
	if x != nil {
		return x.ExposureTime
	}
	return ""
}

func (x *Photo) GetXDimension() uint32 {
	if x != nil {
		return x.XDimension
	}
	return 0
}

func (x *Photo) GetYDimension() uint32 {
	if x != nil {
		return x.YDimension
	}
	return 0
}

func (x *Photo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Photo) GetFNumber() string {
	if x != nil {
		return x.FNumber
	}
	return ""
}

var File_photo_v1_photo_proto protoreflect.FileDescriptor

var file_photo_v1_photo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22,
	0x2e, 0x0a, 0x14, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22,
	0x2a, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x5a, 0x0a, 0x21, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5a, 0x0a, 0x16, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x5c, 0x0a, 0x23, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xf6, 0x01, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x69,
	0x73, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x6f, 0x73,
	0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x78, 0x5f, 0x64, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x78, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x79, 0x5f, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x79,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xb9, 0x03, 0x0a, 0x0c, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1a, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x2e,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x66,
	0x0a, 0x0f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x20, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0xb2, 0x01, 0x0a, 0x1f, 0x66, 0x72, 0x2e, 0x6d, 0x69,
	0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x6f, 0x6c, 0x6c, 0x2e, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x6f, 0x6c, 0x6c,
	0x2f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x08, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_photo_v1_photo_proto_rawDescOnce sync.Once
	file_photo_v1_photo_proto_rawDescData = file_photo_v1_photo_proto_rawDesc
)

func file_photo_v1_photo_proto_rawDescGZIP() []byte {
	file_photo_v1_photo_proto_rawDescOnce.Do(func() {
		file_photo_v1_photo_proto_rawDescData = protoimpl.X.CompressGZIP(file_photo_v1_photo_proto_rawDescData)
	})
	return file_photo_v1_photo_proto_rawDescData
}

var file_photo_v1_photo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_photo_v1_photo_proto_goTypes = []interface{}{
	(*GetPhotosRequest)(nil),                    // 0: photo.v1.GetPhotosRequest
	(*GetPhotosResponse)(nil),                   // 1: photo.v1.GetPhotosResponse
	(*GetByHashRequest)(nil),                    // 2: photo.v1.GetByHashRequest
	(*GetByHashResponse)(nil),                   // 3: photo.v1.GetByHashResponse
	(*ExistsByHashRequest)(nil),                 // 4: photo.v1.ExistsByHashRequest
	(*ExistsByHashResponse)(nil),                // 5: photo.v1.ExistsByHashResponse
	(*ContentByHashRequest)(nil),                // 6: photo.v1.ContentByHashRequest
	(*PhotoServiceContentByHashResponse)(nil),   // 7: photo.v1.PhotoServiceContentByHashResponse
	(*ThumbnailByHashRequest)(nil),              // 8: photo.v1.ThumbnailByHashRequest
	(*PhotoServiceThumbnailByHashResponse)(nil), // 9: photo.v1.PhotoServiceThumbnailByHashResponse
	(*Photo)(nil),                               // 10: photo.v1.Photo
}
var file_photo_v1_photo_proto_depIdxs = []int32{
	10, // 0: photo.v1.GetPhotosResponse.photos:type_name -> photo.v1.Photo
	10, // 1: photo.v1.GetByHashResponse.photo:type_name -> photo.v1.Photo
	0,  // 2: photo.v1.PhotoService.GetPhotos:input_type -> photo.v1.GetPhotosRequest
	2,  // 3: photo.v1.PhotoService.GetByHash:input_type -> photo.v1.GetByHashRequest
	4,  // 4: photo.v1.PhotoService.ExistsByHash:input_type -> photo.v1.ExistsByHashRequest
	6,  // 5: photo.v1.PhotoService.ContentByHash:input_type -> photo.v1.ContentByHashRequest
	8,  // 6: photo.v1.PhotoService.ThumbnailByHash:input_type -> photo.v1.ThumbnailByHashRequest
	1,  // 7: photo.v1.PhotoService.GetPhotos:output_type -> photo.v1.GetPhotosResponse
	3,  // 8: photo.v1.PhotoService.GetByHash:output_type -> photo.v1.GetByHashResponse
	5,  // 9: photo.v1.PhotoService.ExistsByHash:output_type -> photo.v1.ExistsByHashResponse
	7,  // 10: photo.v1.PhotoService.ContentByHash:output_type -> photo.v1.PhotoServiceContentByHashResponse
	9,  // 11: photo.v1.PhotoService.ThumbnailByHash:output_type -> photo.v1.PhotoServiceThumbnailByHashResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_photo_v1_photo_proto_init() }
func file_photo_v1_photo_proto_init() {
	if File_photo_v1_photo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_photo_v1_photo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhotosRequest); i {
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
		file_photo_v1_photo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhotosResponse); i {
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
		file_photo_v1_photo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByHashRequest); i {
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
		file_photo_v1_photo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByHashResponse); i {
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
		file_photo_v1_photo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsByHashRequest); i {
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
		file_photo_v1_photo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsByHashResponse); i {
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
		file_photo_v1_photo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentByHashRequest); i {
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
		file_photo_v1_photo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoServiceContentByHashResponse); i {
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
		file_photo_v1_photo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThumbnailByHashRequest); i {
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
		file_photo_v1_photo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoServiceThumbnailByHashResponse); i {
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
		file_photo_v1_photo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Photo); i {
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
			RawDescriptor: file_photo_v1_photo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_photo_v1_photo_proto_goTypes,
		DependencyIndexes: file_photo_v1_photo_proto_depIdxs,
		MessageInfos:      file_photo_v1_photo_proto_msgTypes,
	}.Build()
	File_photo_v1_photo_proto = out.File
	file_photo_v1_photo_proto_rawDesc = nil
	file_photo_v1_photo_proto_goTypes = nil
	file_photo_v1_photo_proto_depIdxs = nil
}
