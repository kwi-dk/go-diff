// Code generated by protoc-gen-gogo.
// source: diff.proto
// DO NOT EDIT!

/*
Package diff is a generated protocol buffer package.

It is generated from these files:
	diff.proto

It has these top-level messages:
	FileDiff
	Hunk
	Stat
*/
package diff

import proto "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"
import pbtypes "sourcegraph.com/sqs/pbtypes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// A FileDiff represents a unified diff for a single file.
//
// A file unified diff has a header that resembles the following:
//
//   --- oldname	2009-10-11 15:12:20.000000000 -0700
//   +++ newname	2009-10-11 15:12:30.000000000 -0700
type FileDiff struct {
	// the original name of the file
	OrigName string `protobuf:"bytes,1,opt,name=orig_name,proto3" json:"orig_name,omitempty"`
	// the original timestamp (nil if not present)
	OrigTime *pbtypes.Timestamp `protobuf:"bytes,2,opt,name=orig_time" json:"orig_time,omitempty"`
	// the new name of the file (often same as OrigName)
	NewName string `protobuf:"bytes,3,opt,name=new_name,proto3" json:"new_name,omitempty"`
	// the new timestamp (nil if not present)
	NewTime *pbtypes.Timestamp `protobuf:"bytes,4,opt,name=new_time" json:"new_time,omitempty"`
	// extended header lines (e.g., git's "new mode <mode>", "rename from <path>", etc.)
	Extended []string `protobuf:"bytes,5,rep,name=extended" json:"extended,omitempty"`
	// hunks that were changed from orig to new
	Hunks []*Hunk `protobuf:"bytes,6,rep,name=hunks" json:"hunks,omitempty"`
}

func (m *FileDiff) Reset()         { *m = FileDiff{} }
func (m *FileDiff) String() string { return proto.CompactTextString(m) }
func (*FileDiff) ProtoMessage()    {}

func (m *FileDiff) GetOrigTime() *pbtypes.Timestamp {
	if m != nil {
		return m.OrigTime
	}
	return nil
}

func (m *FileDiff) GetNewTime() *pbtypes.Timestamp {
	if m != nil {
		return m.NewTime
	}
	return nil
}

func (m *FileDiff) GetHunks() []*Hunk {
	if m != nil {
		return m.Hunks
	}
	return nil
}

// A Hunk represents a series of changes (additions or deletions) in a file's
// unified diff.
type Hunk struct {
	// starting line number in original file
	OrigStartLine int32 `protobuf:"varint,1,opt,name=orig_start_line,proto3" json:"orig_start_line,omitempty"`
	// number of lines the hunk applies to in the original file
	OrigLines int32 `protobuf:"varint,2,opt,name=orig_lines,proto3" json:"orig_lines,omitempty"`
	// if > 0, then the original file had a 'No newline at end of file' mark at this offset
	OrigNoNewlineAt int32 `protobuf:"varint,3,opt,name=orig_no_newline_at,proto3" json:"orig_no_newline_at,omitempty"`
	// starting line number in new file
	NewStartLine int32 `protobuf:"varint,4,opt,name=new_start_line,proto3" json:"new_start_line,omitempty"`
	// number of lines the hunk applies to in the new file
	NewLines int32 `protobuf:"varint,5,opt,name=new_lines,proto3" json:"new_lines,omitempty"`
	// optional section heading
	Section string `protobuf:"bytes,6,opt,name=section,proto3" json:"section,omitempty"`
	// 0-indexed line offset in unified file diff (including section headers); this is
	// only set when Hunks are read from entire file diff (i.e., when ReadAllHunks is
	// called) This accounts for hunk headers, too, so the StartPosition of the first
	// hunk will be 1.
	StartPosition int32 `protobuf:"varint,7,opt,name=start_position,proto3" json:"start_position,omitempty"`
	// hunk body (lines prefixed with '-', '+', or ' ')
	Body []byte `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Hunk) Reset()         { *m = Hunk{} }
func (m *Hunk) String() string { return proto.CompactTextString(m) }
func (*Hunk) ProtoMessage()    {}

// A Stat is a diff stat that represents the number of lines added/changed/deleted.
type Stat struct {
	// number of lines added
	Added int32 `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
	// number of lines changed
	Changed int32 `protobuf:"varint,2,opt,name=changed,proto3" json:"changed,omitempty"`
	// number of lines deleted
	Deleted int32 `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}

func init() {
}
