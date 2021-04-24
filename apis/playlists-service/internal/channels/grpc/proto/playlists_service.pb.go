// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protos/playlists_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Playlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId string            `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Tracks []*Playlist_Track `protobuf:"bytes,4,rep,name=tracks,proto3" json:"tracks,omitempty"`
}

func (x *Playlist) Reset() {
	*x = Playlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playlist) ProtoMessage() {}

func (x *Playlist) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playlist.ProtoReflect.Descriptor instead.
func (*Playlist) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{1}
}

func (x *Playlist) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Playlist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Playlist) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Playlist) GetTracks() []*Playlist_Track {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type PlaylistsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Playlists []*Playlist `protobuf:"bytes,1,rep,name=playlists,proto3" json:"playlists,omitempty"`
}

func (x *PlaylistsList) Reset() {
	*x = PlaylistsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistsList) ProtoMessage() {}

func (x *PlaylistsList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistsList.ProtoReflect.Descriptor instead.
func (*PlaylistsList) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{2}
}

func (x *PlaylistsList) GetPlaylists() []*Playlist {
	if x != nil {
		return x.Playlists
	}
	return nil
}

type CreatePlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreatePlaylistRequest) Reset() {
	*x = CreatePlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlaylistRequest) ProtoMessage() {}

func (x *CreatePlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlaylistRequest.ProtoReflect.Descriptor instead.
func (*CreatePlaylistRequest) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePlaylistRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePlaylistRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdatePlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdatePlaylistRequest) Reset() {
	*x = UpdatePlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlaylistRequest) ProtoMessage() {}

func (x *UpdatePlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlaylistRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlaylistRequest) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePlaylistRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePlaylistRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Music struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	DurationInSeconds int32    `protobuf:"varint,3,opt,name=durationInSeconds,proto3" json:"durationInSeconds,omitempty"`
	File              string   `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Composers         []string `protobuf:"bytes,5,rep,name=composers,proto3" json:"composers,omitempty"`
	Lyrics            string   `protobuf:"bytes,6,opt,name=lyrics,proto3" json:"lyrics,omitempty"`
	AlbumId           string   `protobuf:"bytes,7,opt,name=albumId,proto3" json:"albumId,omitempty"`
	Views             int32    `protobuf:"varint,8,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *Music) Reset() {
	*x = Music{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Music) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Music) ProtoMessage() {}

func (x *Music) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Music.ProtoReflect.Descriptor instead.
func (*Music) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{5}
}

func (x *Music) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Music) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Music) GetDurationInSeconds() int32 {
	if x != nil {
		return x.DurationInSeconds
	}
	return 0
}

func (x *Music) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Music) GetComposers() []string {
	if x != nil {
		return x.Composers
	}
	return nil
}

func (x *Music) GetLyrics() string {
	if x != nil {
		return x.Lyrics
	}
	return ""
}

func (x *Music) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

func (x *Music) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{6}
}

type Playlist_Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Music *Music `protobuf:"bytes,2,opt,name=music,proto3" json:"music,omitempty"`
}

func (x *Playlist_Track) Reset() {
	*x = Playlist_Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_playlists_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playlist_Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playlist_Track) ProtoMessage() {}

func (x *Playlist_Track) ProtoReflect() protoreflect.Message {
	mi := &file_protos_playlists_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playlist_Track.ProtoReflect.Descriptor instead.
func (*Playlist_Track) Descriptor() ([]byte, []int) {
	return file_protos_playlists_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Playlist_Track) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Playlist_Track) GetMusic() *Music {
	if x != nil {
		return x.Music
	}
	return nil
}

var File_protos_playlists_service_proto protoreflect.FileDescriptor

var file_protos_playlists_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb8, 0x01,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x1a, 0x41, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x73, 0x69,
	0x63, 0x52, 0x05, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x22, 0x3e, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x43, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x05, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x79,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x79, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x9e, 0x02, 0x0a, 0x09,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_playlists_service_proto_rawDescOnce sync.Once
	file_protos_playlists_service_proto_rawDescData = file_protos_playlists_service_proto_rawDesc
)

func file_protos_playlists_service_proto_rawDescGZIP() []byte {
	file_protos_playlists_service_proto_rawDescOnce.Do(func() {
		file_protos_playlists_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_playlists_service_proto_rawDescData)
	})
	return file_protos_playlists_service_proto_rawDescData
}

var file_protos_playlists_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_playlists_service_proto_goTypes = []interface{}{
	(*Id)(nil),                    // 0: proto.Id
	(*Playlist)(nil),              // 1: proto.Playlist
	(*PlaylistsList)(nil),         // 2: proto.PlaylistsList
	(*CreatePlaylistRequest)(nil), // 3: proto.CreatePlaylistRequest
	(*UpdatePlaylistRequest)(nil), // 4: proto.UpdatePlaylistRequest
	(*Music)(nil),                 // 5: proto.Music
	(*Empty)(nil),                 // 6: proto.Empty
	(*Playlist_Track)(nil),        // 7: proto.Playlist.Track
}
var file_protos_playlists_service_proto_depIdxs = []int32{
	7, // 0: proto.Playlist.tracks:type_name -> proto.Playlist.Track
	1, // 1: proto.PlaylistsList.playlists:type_name -> proto.Playlist
	5, // 2: proto.Playlist.Track.music:type_name -> proto.Music
	0, // 3: proto.Playlists.GetPlaylist:input_type -> proto.Id
	0, // 4: proto.Playlists.GetPlaylists:input_type -> proto.Id
	3, // 5: proto.Playlists.CreatePlaylist:input_type -> proto.CreatePlaylistRequest
	4, // 6: proto.Playlists.UpdatePlaylist:input_type -> proto.UpdatePlaylistRequest
	0, // 7: proto.Playlists.DeletePlaylist:input_type -> proto.Id
	1, // 8: proto.Playlists.GetPlaylist:output_type -> proto.Playlist
	2, // 9: proto.Playlists.GetPlaylists:output_type -> proto.PlaylistsList
	1, // 10: proto.Playlists.CreatePlaylist:output_type -> proto.Playlist
	1, // 11: proto.Playlists.UpdatePlaylist:output_type -> proto.Playlist
	6, // 12: proto.Playlists.DeletePlaylist:output_type -> proto.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_playlists_service_proto_init() }
func file_protos_playlists_service_proto_init() {
	if File_protos_playlists_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_playlists_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_protos_playlists_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playlist); i {
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
		file_protos_playlists_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistsList); i {
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
		file_protos_playlists_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlaylistRequest); i {
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
		file_protos_playlists_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlaylistRequest); i {
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
		file_protos_playlists_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Music); i {
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
		file_protos_playlists_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_playlists_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playlist_Track); i {
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
			RawDescriptor: file_protos_playlists_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_playlists_service_proto_goTypes,
		DependencyIndexes: file_protos_playlists_service_proto_depIdxs,
		MessageInfos:      file_protos_playlists_service_proto_msgTypes,
	}.Build()
	File_protos_playlists_service_proto = out.File
	file_protos_playlists_service_proto_rawDesc = nil
	file_protos_playlists_service_proto_goTypes = nil
	file_protos_playlists_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlaylistsClient is the client API for Playlists service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaylistsClient interface {
	GetPlaylist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Playlist, error)
	GetPlaylists(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PlaylistsList, error)
	CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	UpdatePlaylist(ctx context.Context, in *UpdatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error)
	DeletePlaylist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type playlistsClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistsClient(cc grpc.ClientConnInterface) PlaylistsClient {
	return &playlistsClient{cc}
}

func (c *playlistsClient) GetPlaylist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/proto.Playlists/GetPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistsClient) GetPlaylists(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PlaylistsList, error) {
	out := new(PlaylistsList)
	err := c.cc.Invoke(ctx, "/proto.Playlists/GetPlaylists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistsClient) CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/proto.Playlists/CreatePlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistsClient) UpdatePlaylist(ctx context.Context, in *UpdatePlaylistRequest, opts ...grpc.CallOption) (*Playlist, error) {
	out := new(Playlist)
	err := c.cc.Invoke(ctx, "/proto.Playlists/UpdatePlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistsClient) DeletePlaylist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Playlists/DeletePlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistsServer is the server API for Playlists service.
type PlaylistsServer interface {
	GetPlaylist(context.Context, *Id) (*Playlist, error)
	GetPlaylists(context.Context, *Id) (*PlaylistsList, error)
	CreatePlaylist(context.Context, *CreatePlaylistRequest) (*Playlist, error)
	UpdatePlaylist(context.Context, *UpdatePlaylistRequest) (*Playlist, error)
	DeletePlaylist(context.Context, *Id) (*Empty, error)
}

// UnimplementedPlaylistsServer can be embedded to have forward compatible implementations.
type UnimplementedPlaylistsServer struct {
}

func (*UnimplementedPlaylistsServer) GetPlaylist(context.Context, *Id) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylist not implemented")
}
func (*UnimplementedPlaylistsServer) GetPlaylists(context.Context, *Id) (*PlaylistsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylists not implemented")
}
func (*UnimplementedPlaylistsServer) CreatePlaylist(context.Context, *CreatePlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylist not implemented")
}
func (*UnimplementedPlaylistsServer) UpdatePlaylist(context.Context, *UpdatePlaylistRequest) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylist not implemented")
}
func (*UnimplementedPlaylistsServer) DeletePlaylist(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlaylist not implemented")
}

func RegisterPlaylistsServer(s *grpc.Server, srv PlaylistsServer) {
	s.RegisterService(&_Playlists_serviceDesc, srv)
}

func _Playlists_GetPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistsServer).GetPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlists/GetPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistsServer).GetPlaylist(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlists_GetPlaylists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistsServer).GetPlaylists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlists/GetPlaylists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistsServer).GetPlaylists(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlists_CreatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistsServer).CreatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlists/CreatePlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistsServer).CreatePlaylist(ctx, req.(*CreatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlists_UpdatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistsServer).UpdatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlists/UpdatePlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistsServer).UpdatePlaylist(ctx, req.(*UpdatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlists_DeletePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistsServer).DeletePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlists/DeletePlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistsServer).DeletePlaylist(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Playlists_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Playlists",
	HandlerType: (*PlaylistsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlaylist",
			Handler:    _Playlists_GetPlaylist_Handler,
		},
		{
			MethodName: "GetPlaylists",
			Handler:    _Playlists_GetPlaylists_Handler,
		},
		{
			MethodName: "CreatePlaylist",
			Handler:    _Playlists_CreatePlaylist_Handler,
		},
		{
			MethodName: "UpdatePlaylist",
			Handler:    _Playlists_UpdatePlaylist_Handler,
		},
		{
			MethodName: "DeletePlaylist",
			Handler:    _Playlists_DeletePlaylist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/playlists_service.proto",
}
