 XorCrypto
=======
可以简单的用异或加密string
### 简介
简单的Url参数加密
```go
var XB = Xor2Base64{Key:"My key to Xor2Base64"}
EnString := XB.Encode("Your string")
YourString := XB.Decode(EnString)
```
使用的base64.RawURL 在URL传参的时候是安全的