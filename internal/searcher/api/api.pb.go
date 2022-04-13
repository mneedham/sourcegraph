// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api.proto

package api

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

type SearchStructuralUnindexedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repo is the name of the repository to search. e.g. "github.com/gorilla/mux"
	Repo        string                 `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Commit      string                 `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	Limit       int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	PatternInfo *StructuralPatternInfo `protobuf:"bytes,4,opt,name=pattern_info,json=patternInfo,proto3" json:"pattern_info,omitempty"`
}

func (x *SearchStructuralUnindexedRequest) Reset() {
	*x = SearchStructuralUnindexedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchStructuralUnindexedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchStructuralUnindexedRequest) ProtoMessage() {}

func (x *SearchStructuralUnindexedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchStructuralUnindexedRequest.ProtoReflect.Descriptor instead.
func (*SearchStructuralUnindexedRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *SearchStructuralUnindexedRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *SearchStructuralUnindexedRequest) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *SearchStructuralUnindexedRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchStructuralUnindexedRequest) GetPatternInfo() *StructuralPatternInfo {
	if x != nil {
		return x.PatternInfo
	}
	return nil
}

type StructuralPatternInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern               string        `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	PatternMatchesContent bool          `protobuf:"varint,2,opt,name=pattern_matches_content,json=patternMatchesContent,proto3" json:"pattern_matches_content,omitempty"`
	PatternMatchesPath    bool          `protobuf:"varint,3,opt,name=pattern_matches_path,json=patternMatchesPath,proto3" json:"pattern_matches_path,omitempty"`
	PathPatterns          *PathPatterns `protobuf:"bytes,4,opt,name=path_patterns,json=pathPatterns,proto3" json:"path_patterns,omitempty"`
	Languages             []string      `protobuf:"bytes,5,rep,name=languages,proto3" json:"languages,omitempty"`
	CombyRule             string        `protobuf:"bytes,6,opt,name=comby_rule,json=combyRule,proto3" json:"comby_rule,omitempty"`
}

func (x *StructuralPatternInfo) Reset() {
	*x = StructuralPatternInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructuralPatternInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructuralPatternInfo) ProtoMessage() {}

func (x *StructuralPatternInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructuralPatternInfo.ProtoReflect.Descriptor instead.
func (*StructuralPatternInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *StructuralPatternInfo) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *StructuralPatternInfo) GetPatternMatchesContent() bool {
	if x != nil {
		return x.PatternMatchesContent
	}
	return false
}

func (x *StructuralPatternInfo) GetPatternMatchesPath() bool {
	if x != nil {
		return x.PatternMatchesPath
	}
	return false
}

func (x *StructuralPatternInfo) GetPathPatterns() *PathPatterns {
	if x != nil {
		return x.PathPatterns
	}
	return nil
}

func (x *StructuralPatternInfo) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *StructuralPatternInfo) GetCombyRule() string {
	if x != nil {
		return x.CombyRule
	}
	return ""
}

type PathPatterns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exclude         string   `protobuf:"bytes,1,opt,name=exclude,proto3" json:"exclude,omitempty"`
	Include         []string `protobuf:"bytes,2,rep,name=include,proto3" json:"include,omitempty"`
	IsRegexp        bool     `protobuf:"varint,3,opt,name=is_regexp,json=isRegexp,proto3" json:"is_regexp,omitempty"`
	IsCaseSensitive bool     `protobuf:"varint,4,opt,name=is_case_sensitive,json=isCaseSensitive,proto3" json:"is_case_sensitive,omitempty"`
}

func (x *PathPatterns) Reset() {
	*x = PathPatterns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathPatterns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathPatterns) ProtoMessage() {}

func (x *PathPatterns) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathPatterns.ProtoReflect.Descriptor instead.
func (*PathPatterns) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *PathPatterns) GetExclude() string {
	if x != nil {
		return x.Exclude
	}
	return ""
}

func (x *PathPatterns) GetInclude() []string {
	if x != nil {
		return x.Include
	}
	return nil
}

func (x *PathPatterns) GetIsRegexp() bool {
	if x != nil {
		return x.IsRegexp
	}
	return false
}

func (x *PathPatterns) GetIsCaseSensitive() bool {
	if x != nil {
		return x.IsCaseSensitive
	}
	return false
}

type SearchStructuralUnindexedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*SearchStructuralUnindexedResponse_Matches
	//	*SearchStructuralUnindexedResponse_Done
	Event isSearchStructuralUnindexedResponse_Event `protobuf_oneof:"event"`
}

func (x *SearchStructuralUnindexedResponse) Reset() {
	*x = SearchStructuralUnindexedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchStructuralUnindexedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchStructuralUnindexedResponse) ProtoMessage() {}

func (x *SearchStructuralUnindexedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchStructuralUnindexedResponse.ProtoReflect.Descriptor instead.
func (*SearchStructuralUnindexedResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (m *SearchStructuralUnindexedResponse) GetEvent() isSearchStructuralUnindexedResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *SearchStructuralUnindexedResponse) GetMatches() *EventMatches {
	if x, ok := x.GetEvent().(*SearchStructuralUnindexedResponse_Matches); ok {
		return x.Matches
	}
	return nil
}

func (x *SearchStructuralUnindexedResponse) GetDone() *EventDone {
	if x, ok := x.GetEvent().(*SearchStructuralUnindexedResponse_Done); ok {
		return x.Done
	}
	return nil
}

type isSearchStructuralUnindexedResponse_Event interface {
	isSearchStructuralUnindexedResponse_Event()
}

type SearchStructuralUnindexedResponse_Matches struct {
	Matches *EventMatches `protobuf:"bytes,1,opt,name=matches,proto3,oneof"`
}

type SearchStructuralUnindexedResponse_Done struct {
	Done *EventDone `protobuf:"bytes,2,opt,name=done,proto3,oneof"`
}

func (*SearchStructuralUnindexedResponse_Matches) isSearchStructuralUnindexedResponse_Event() {}

func (*SearchStructuralUnindexedResponse_Done) isSearchStructuralUnindexedResponse_Event() {}

type EventMatches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*FileMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
}

func (x *EventMatches) Reset() {
	*x = EventMatches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMatches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMatches) ProtoMessage() {}

func (x *EventMatches) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMatches.ProtoReflect.Descriptor instead.
func (*EventMatches) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *EventMatches) GetMatches() []*FileMatch {
	if x != nil {
		return x.Matches
	}
	return nil
}

type FileMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string       `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LineMatches []*LineMatch `protobuf:"bytes,2,rep,name=line_matches,json=lineMatches,proto3" json:"line_matches,omitempty"`
	// MatchCount is the number of matches.  Different from len(LineMatches), as multiple
	// lines may correspond to one logical match when doing a structural search
	MatchCount int64 `protobuf:"varint,3,opt,name=match_count,json=matchCount,proto3" json:"match_count,omitempty"`
	// LimitHit is true if LineMatches may not include all LineMatches.
	LimitHit bool `protobuf:"varint,4,opt,name=limit_hit,json=limitHit,proto3" json:"limit_hit,omitempty"`
}

func (x *FileMatch) Reset() {
	*x = FileMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMatch) ProtoMessage() {}

func (x *FileMatch) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMatch.ProtoReflect.Descriptor instead.
func (*FileMatch) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *FileMatch) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileMatch) GetLineMatches() []*LineMatch {
	if x != nil {
		return x.LineMatches
	}
	return nil
}

func (x *FileMatch) GetMatchCount() int64 {
	if x != nil {
		return x.MatchCount
	}
	return 0
}

func (x *FileMatch) GetLimitHit() bool {
	if x != nil {
		return x.LimitHit
	}
	return false
}

type LineMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Preview is the matched line.
	Preview string `protobuf:"bytes,1,opt,name=preview,proto3" json:"preview,omitempty"`
	// LineNumber is the 0-based line number.
	LineNumber       int64           `protobuf:"varint,2,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	OffsetAndLengths []*OffsetLength `protobuf:"bytes,3,rep,name=offset_and_lengths,json=offsetAndLengths,proto3" json:"offset_and_lengths,omitempty"`
}

func (x *LineMatch) Reset() {
	*x = LineMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineMatch) ProtoMessage() {}

func (x *LineMatch) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineMatch.ProtoReflect.Descriptor instead.
func (*LineMatch) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *LineMatch) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *LineMatch) GetLineNumber() int64 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *LineMatch) GetOffsetAndLengths() []*OffsetLength {
	if x != nil {
		return x.OffsetAndLengths
	}
	return nil
}

// Offsets and lengths are measured in characters, not bytes.
type OffsetLength struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Length int64 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *OffsetLength) Reset() {
	*x = OffsetLength{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffsetLength) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffsetLength) ProtoMessage() {}

func (x *OffsetLength) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffsetLength.ProtoReflect.Descriptor instead.
func (*OffsetLength) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *OffsetLength) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *OffsetLength) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type EventDone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LimitHit is true if the sent results may not be complete because a match limit was hit.
	LimitHit bool `protobuf:"varint,1,opt,name=limit_hit,json=limitHit,proto3" json:"limit_hit,omitempty"`
}

func (x *EventDone) Reset() {
	*x = EventDone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDone) ProtoMessage() {}

func (x *EventDone) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDone.ProtoReflect.Descriptor instead.
func (*EventDone) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *EventDone) GetLimitHit() bool {
	if x != nil {
		return x.LimitHit
	}
	return false
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0xa3, 0x01, 0x0a, 0x20, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x70, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x0d, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x52, 0x0c,
	0x70, 0x61, 0x74, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x62, 0x79, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x62, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x50, 0x61,
	0x74, 0x68, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x69,
	0x73, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x61, 0x73, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x21, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x0c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0b, 0x6c,
	0x69, 0x6e, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x69, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x6e,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x3f, 0x0a, 0x12, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x52, 0x10, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x73, 0x22, 0x3e, 0x0a, 0x0c, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x22, 0x28, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x69, 0x74, 0x32, 0x7a, 0x0a, 0x08,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x6e, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_goTypes = []interface{}{
	(*SearchStructuralUnindexedRequest)(nil),  // 0: api.SearchStructuralUnindexedRequest
	(*StructuralPatternInfo)(nil),             // 1: api.StructuralPatternInfo
	(*PathPatterns)(nil),                      // 2: api.PathPatterns
	(*SearchStructuralUnindexedResponse)(nil), // 3: api.SearchStructuralUnindexedResponse
	(*EventMatches)(nil),                      // 4: api.EventMatches
	(*FileMatch)(nil),                         // 5: api.FileMatch
	(*LineMatch)(nil),                         // 6: api.LineMatch
	(*OffsetLength)(nil),                      // 7: api.OffsetLength
	(*EventDone)(nil),                         // 8: api.EventDone
}
var file_api_proto_depIdxs = []int32{
	1, // 0: api.SearchStructuralUnindexedRequest.pattern_info:type_name -> api.StructuralPatternInfo
	2, // 1: api.StructuralPatternInfo.path_patterns:type_name -> api.PathPatterns
	4, // 2: api.SearchStructuralUnindexedResponse.matches:type_name -> api.EventMatches
	8, // 3: api.SearchStructuralUnindexedResponse.done:type_name -> api.EventDone
	5, // 4: api.EventMatches.matches:type_name -> api.FileMatch
	6, // 5: api.FileMatch.line_matches:type_name -> api.LineMatch
	7, // 6: api.LineMatch.offset_and_lengths:type_name -> api.OffsetLength
	0, // 7: api.Searcher.SearchStructuralUnindexed:input_type -> api.SearchStructuralUnindexedRequest
	3, // 8: api.Searcher.SearchStructuralUnindexed:output_type -> api.SearchStructuralUnindexedResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchStructuralUnindexedRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructuralPatternInfo); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathPatterns); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchStructuralUnindexedResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMatches); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMatch); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineMatch); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffsetLength); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDone); i {
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
	file_api_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SearchStructuralUnindexedResponse_Matches)(nil),
		(*SearchStructuralUnindexedResponse_Done)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
