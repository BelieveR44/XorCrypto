 XorCrypto
=======
简单的异或加密string
### 使用
```go
go get github.com/BelieveR44/XorCrypto
```
### 简介
简单的Url参数加密
```go
var XB = Xor2Base64{Key:"My key to Xor2Base64"}
EnString := XB.Encode("Your string")
YourString := XB.Decode(EnString)
fmt.Printf("YouString:%s\nEnString:%s\nDeString:%s\nKey:%s\n","Your string",EnString,YourString,XB.Key)
/*
YouString:Your string
EnString:FBZVGUUKVAYGTj8=
DeString:Your string
Key:My key to Xor2Base64
*/
```
使用的base64.RawURL 在URL传参的时候是安全的
