// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type F2 struct {
	_tab flatbuffers.Table
}

func GetRootAsF2(buf []byte, offset flatbuffers.UOffsetT) *F2 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &F2{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *F2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *F2) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *F2) AllVals(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *F2) AllValsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func F2Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func F2AddAllVals(builder *flatbuffers.Builder, allVals flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(allVals), 0)
}
func F2StartAllValsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func F2End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
