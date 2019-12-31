// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatums

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///  request from tv box
type TerminalRequestT struct {
	UserID       int64
	SerialNumber string
	ActiveCode   string
	ApkType      string
	Operation    OperationType
}

func TerminalRequestPack(builder *flatbuffers.Builder, t *TerminalRequestT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	serialNumberOffset := builder.CreateString(t.SerialNumber)
	activeCodeOffset := builder.CreateString(t.ActiveCode)
	apkTypeOffset := builder.CreateString(t.ApkType)
	TerminalRequestStart(builder)
	TerminalRequestAddUserID(builder, t.UserID)
	TerminalRequestAddSerialNumber(builder, serialNumberOffset)
	TerminalRequestAddActiveCode(builder, activeCodeOffset)
	TerminalRequestAddApkType(builder, apkTypeOffset)
	TerminalRequestAddOperation(builder, t.Operation)
	return TerminalRequestEnd(builder)
}

func (rcv *TerminalRequest) UnPackTo(t *TerminalRequestT) {
	t.UserID = rcv.UserID()
	t.SerialNumber = string(rcv.SerialNumber())
	t.ActiveCode = string(rcv.ActiveCode())
	t.ApkType = string(rcv.ApkType())
	t.Operation = rcv.Operation()
}

func (rcv *TerminalRequest) UnPack() *TerminalRequestT {
	if rcv == nil {
		return nil
	}
	t := &TerminalRequestT{}
	rcv.UnPackTo(t)
	return t
}

type TerminalRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsTerminalRequest(buf []byte, offset flatbuffers.UOffsetT) *TerminalRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TerminalRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TerminalRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TerminalRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TerminalRequest) UserID() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TerminalRequest) MutateUserID(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *TerminalRequest) SerialNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TerminalRequest) ActiveCode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TerminalRequest) ApkType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TerminalRequest) Operation() OperationType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return OperationType(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *TerminalRequest) MutateOperation(n OperationType) bool {
	return rcv._tab.MutateInt32Slot(12, int32(n))
}

func TerminalRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func TerminalRequestAddUserID(builder *flatbuffers.Builder, userID int64) {
	builder.PrependInt64Slot(0, userID, 0)
}
func TerminalRequestAddSerialNumber(builder *flatbuffers.Builder, serialNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(serialNumber), 0)
}
func TerminalRequestAddActiveCode(builder *flatbuffers.Builder, activeCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(activeCode), 0)
}
func TerminalRequestAddApkType(builder *flatbuffers.Builder, apkType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(apkType), 0)
}
func TerminalRequestAddOperation(builder *flatbuffers.Builder, operation OperationType) {
	builder.PrependInt32Slot(4, int32(operation), 0)
}
func TerminalRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
