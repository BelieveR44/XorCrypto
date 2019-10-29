package Xor2Base64

import "encoding/base64"

type Xor2Base64 struct{
	Key string
}

func (xb *Xor2Base64)Decode(str string)string{
	strByte, _ := base64.RawURLEncoding.DecodeString(str)
	xorval:=xor(strByte,[]byte(xb.Key))
	return string(xorval)
}
func (xb *Xor2Base64)Encode(str string)string{
	xorval:=xor([]byte(str),[]byte(xb.Key))
	str = base64.RawURLEncoding.EncodeToString(xorval)
	return str
}
func xor(src []byte, key []byte) []byte {
	for i := 0; i < len(src); i++ {
		src[i] ^= key[i%len(key)]
	}
	return src
}
