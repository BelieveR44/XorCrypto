package Xor2Base64

import "testing"
var(
	XB = Xor2Base64{Key:"My key to Xor2Base64"}
	str string = "5db705fd659a7e373061b15a" //测试使用 mongodb 的objectID
	en string = "eB1CXFVMRhBZFWEORVdxVkBVAAUvSBUK"
)
func TestXor2Base64_Encode(t *testing.T) {
	if XB.Encode(str) != en {
		t.Error("TestXor2Base64_Encode Error")
	}
}
func TestXor2Base64_Decode(t *testing.T) {
	if XB.Decode(en)!= str{
		t.Error("TestXor2Base64_Decode Error")
	}
}
func BenchmarkXor2Base64_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XB.Encode(str)
	}
}
func BenchmarkXor2Base64_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XB.Decode(en)
	}
}
