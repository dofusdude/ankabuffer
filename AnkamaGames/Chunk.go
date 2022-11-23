// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package AnkamaGames

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Chunk struct {
	_tab flatbuffers.Table
}

func GetRootAsChunk(buf []byte, offset flatbuffers.UOffsetT) *Chunk {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Chunk{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsChunk(buf []byte, offset flatbuffers.UOffsetT) *Chunk {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Chunk{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Chunk) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Chunk) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Chunk) Hash(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Chunk) HashLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Chunk) HashBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Chunk) MutateHash(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Chunk) Size() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Chunk) MutateSize(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *Chunk) Offset() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Chunk) MutateOffset(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *Chunk) Done() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Chunk) MutateDone(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func ChunkStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ChunkAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(hash), 0)
}
func ChunkStartHashVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ChunkAddSize(builder *flatbuffers.Builder, size int64) {
	builder.PrependInt64Slot(1, size, 0)
}
func ChunkAddOffset(builder *flatbuffers.Builder, offset int64) {
	builder.PrependInt64Slot(2, offset, 0)
}
func ChunkAddDone(builder *flatbuffers.Builder, done bool) {
	builder.PrependBoolSlot(3, done, false)
}
func ChunkEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
