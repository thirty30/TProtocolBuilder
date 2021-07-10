package tprotocol
type CommonBool struct{
Value bool
}
func (pOwn *CommonBool) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeBOOL(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonBool) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeBOOL(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN8 struct{
Value int8
}
func (pOwn *CommonN8) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN8) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN16 struct{
Value int16
}
func (pOwn *CommonN16) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN16) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN32 struct{
Value int32
}
func (pOwn *CommonN32) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN32) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN64 struct{
Value int64
}
func (pOwn *CommonN64) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN64) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU8 struct{
Value uint8
}
func (pOwn *CommonU8) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU8) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU16 struct{
Value uint16
}
func (pOwn *CommonU16) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU16) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU32 struct{
Value uint32
}
func (pOwn *CommonU32) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU32) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU64 struct{
Value uint64
}
func (pOwn *CommonU64) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU64) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonF32 struct{
Value float32
}
func (pOwn *CommonF32) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonF32) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonF64 struct{
Value float64
}
func (pOwn *CommonF64) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeF64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonF64) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeF64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonStr struct{
Value string
}
func (pOwn *CommonStr) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonStr) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type T1 struct{
Value1 int32
Value2 float32
Value3 string
}
func (pOwn *T1) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value1)
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value2)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value3)
return nOffset
}
func (pOwn *T1) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value1, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value2, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value3, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type T2 struct{
Value1 int32
Value2 float32
Value3 string
Value4 int32
}
func (pOwn *T2) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value1)
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value2)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value3)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value4)
return nOffset
}
func (pOwn *T2) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value1, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value2, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value3, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value4, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type Test struct{
Value0 bool
Value1 int8
Value2 int16
Value3 int32
Value4 int64
Value5 uint8
Value6 uint16
Value7 uint32
Value8 uint64
Value9 float32
Value10 float64
Value11 string
Value12 T1
mValue13 []*T2
mValue14 []int32
mValue15 []string
}
func (pOwn *Test) GetValue13Count() int32 { return int32(len(pOwn.mValue13)) }
func (pOwn *Test) GetValue13At(aIdx int32) *T2 { return pOwn.mValue13[aIdx] }
func (pOwn *Test) AppendValue13(aData *T2) {
 if pOwn.mValue13 == nil { pOwn.mValue13 = make([]*T2, 0, 8) }
 pOwn.mValue13 = append(pOwn.mValue13, aData)
 }
func (pOwn *Test) GetValue14Count() int32 { return int32(len(pOwn.mValue14)) }
func (pOwn *Test) GetValue14At(aIdx int32) int32 { return pOwn.mValue14[aIdx] }
func (pOwn *Test) AppendValue14(aData int32) {
 if pOwn.mValue14 == nil { pOwn.mValue14 = make([]int32, 0, 8) }
 pOwn.mValue14 = append(pOwn.mValue14, aData)
 }
func (pOwn *Test) GetValue15Count() int32 { return int32(len(pOwn.mValue15)) }
func (pOwn *Test) GetValue15At(aIdx int32) string { return pOwn.mValue15[aIdx] }
func (pOwn *Test) AppendValue15(aData string) {
 if pOwn.mValue15 == nil { pOwn.mValue15 = make([]string, 0, 8) }
 pOwn.mValue15 = append(pOwn.mValue15, aData)
 }
func (pOwn *Test) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeBOOL(aBuffer[nOffset:], aSize-nOffset, pOwn.Value0)
nOffset += serializeN8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value1)
nOffset += serializeN16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value2)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value3)
nOffset += serializeN64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value4)
nOffset += serializeU8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value5)
nOffset += serializeU16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value6)
nOffset += serializeU32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value7)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value8)
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value9)
nOffset += serializeF64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value10)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value11)
nOffset += pOwn.Value12.Serialize(aBuffer[nOffset:], aSize-nOffset)
nValue13Count := len(pOwn.mValue13)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nValue13Count))
for i := 0; i < nValue13Count; i++ { nOffset += pOwn.mValue13[i].Serialize(aBuffer[nOffset:], aSize-nOffset) }
nValue14Count := len(pOwn.mValue14)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nValue14Count))
for i := 0; i < nValue14Count; i++ { nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.mValue14[i]) }
nValue15Count := len(pOwn.mValue15)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nValue15Count))
for i := 0; i < nValue15Count; i++ { nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.mValue15[i]) }
return nOffset
}
func (pOwn *Test) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value0, nTemp = deserializeBOOL(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value1, nTemp = deserializeN8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value2, nTemp = deserializeN16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value3, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value4, nTemp = deserializeN64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value5, nTemp = deserializeU8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value6, nTemp = deserializeU16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value7, nTemp = deserializeU32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value8, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value9, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value10, nTemp = deserializeF64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value11, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
nOffset += pOwn.Value12.Deserialize(aBuffer[nOffset:], aSize-nOffset)
nValue13Count, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nValue13Count); i++ {
var obj T2
nOffset += obj.Deserialize(aBuffer[nOffset:], aSize-nOffset)
if pOwn.mValue13 == nil { pOwn.mValue13 = make([]*T2, 0, 8) }
pOwn.mValue13 = append(pOwn.mValue13, &obj)
}
nValue14Count, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nValue14Count); i++ {
obj, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
if pOwn.mValue14 == nil { pOwn.mValue14 = make([]int32, 0, 8) }
pOwn.mValue14 = append(pOwn.mValue14, obj)
}
nValue15Count, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nValue15Count); i++ {
obj, nTemp := deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
if pOwn.mValue15 == nil { pOwn.mValue15 = make([]string, 0, 8) }
pOwn.mValue15 = append(pOwn.mValue15, obj)
}
return nOffset
}
const (
C2S_LOGIN = 1
S2C_LOGIN_RESP = 2
)
