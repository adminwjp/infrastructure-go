package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

type  seurityUtl struct {
	BASE64Table string
	DesKey []byte //DES 中使用的密钥key只能为8位。
	AesKey []byte //AES 密钥key只能为16位、24位或32位，分别对应AES-128, AES-192和 AES-256
}
 const (
 	DesKey = "abcdefgh"
	BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

 )

var AesKey = []byte{0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
	0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
	0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
	0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
}

func NewSeurityUtl()seurityUtl{
	return seurityUtl{}
}




func (*seurityUtl) Base64Encode(data string) string  {
	content :=[]byte(data)
	return SeurityInstance.Base64EncodeByBytes(content)
}

func (*seurityUtl) Base64EncodeByBytes(buffer []byte) string  {
	coder :=base64.NewEncoding(SeurityInstance.BASE64Table)
	return coder.EncodeToString(buffer)
}

func (*seurityUtl) Base64EDecode(data string) string  {
	content :=[]byte(data)
	return SeurityInstance.Base64DecodeToString(content)
}

func (*seurityUtl) Base64DecodeToString(buffer []byte) string  {
	coder :=base64.NewEncoding(SeurityInstance.BASE64Table)
	result,err := coder.DecodeString(string(buffer[:]))
	if err!=nil{
		return ""
	}
	return  string(result)
}

func (*seurityUtl) Md5(str string) string  {
	return  SeurityInstance.Md5ByByte([]byte(str))
}

func (*seurityUtl) Md5ByByte(buffer []byte) string  {
	m :=md5.New()
	m.Write(buffer)
	return  hex.EncodeToString(m.Sum(nil))
}
func (*seurityUtl) Sha1(str string) string  {
	return  SeurityInstance.Sha1ByByte([]byte(str))
}

func (*seurityUtl) Sha1ByByte(buffer []byte) string  {
	m :=sha1.New()
	m.Write(buffer)
	return  hex.EncodeToString(m.Sum(nil))
}
func  zerPadding(cipherText []byte,blockSize int) []byte {
	padding :=blockSize-len(cipherText)%blockSize
	padText :=bytes.Repeat([]byte{0},padding)
	return  append(cipherText,padText...)
}

func  zeroUnPadding(origData []byte)[]byte  {
	return  bytes.TrimFunc(origData, func(r rune) bool {
		return  r==rune(0)
	})
}
func (*seurityUtl) DesEncrypt(text string)string  {
	return  SeurityInstance.DesEncryptByKey(text,SeurityInstance.DesKey)
}

func (seurityUtl) DesEncryptByKey(text string,key []byte)string  {
	block,err:=des.NewCipher(key)
	if err!=nil{
		return ""
	}
	bs :=block.BlockSize()
	src :=[]byte(text)
	temp :=zerPadding(src,bs)
	if len(temp)%bs!=0{
		return  ""
	}
	out :=make([]byte,len(temp))
	dst :=out
	for len(temp)>0 {
		block.Encrypt(dst,temp[:bs])
		temp=temp[bs:]
		dst=dst[bs:]
	}
	return  hex.EncodeToString(out)
}

func (*seurityUtl) DesDecrypt(hexStr string)string  {
	return  SeurityInstance.DesDecryptByKey(hexStr,SeurityInstance.DesKey)
}

func (*seurityUtl) DesDecryptByKey(hexStr string,key []byte)string  {
	src,err :=hex.DecodeString(hexStr)
	if err!=nil{
		return ""
	}
	block,err:=des.NewCipher(key)
	if err!=nil{
		return ""
	}
	bs :=block.BlockSize()
	//src :=[]byte(text)
	temp :=zerPadding(src,bs)
	if len(temp)%bs!=0{
		return  ""
	}
	out :=make([]byte,len(temp))
	dst :=out
	for len(temp)>0 {
		block.Decrypt(dst,temp[:bs])
		temp=temp[bs:]
		dst=dst[bs:]
	}
	out=zeroUnPadding(out)
	return  string(out)
}

func (*seurityUtl) AesEncrypt(text string)string  {
	return  SeurityInstance.AesEncryptByKey(text,SeurityInstance.AesKey)
}

func (*seurityUtl) AesEncryptByKey(text string,key []byte)string  {
	var iv=key[:aes.BlockSize]
	block,err:=aes.NewCipher(key)
	if err!=nil{
		return ""
	}
	encryptTd :=make([]byte,len(text))
	encryptor :=cipher.NewCFBEncrypter(block,iv)
	encryptor.XORKeyStream(encryptTd,[]byte(text))
	return  hex.EncodeToString(encryptTd)
}

func (*seurityUtl) AesDecrypt(hexStr string)string  {
	return  SeurityInstance.AesDecryptByKey(hexStr,SeurityInstance.AesKey)
}

func (*seurityUtl) AesDecryptByKey(hexStr string,key []byte)string  {
	text,err :=hex.DecodeString(hexStr)
	if err!=nil{
		return ""
	}
	var iv=key[:aes.BlockSize]
	block,err:=aes.NewCipher(key)
	if err!=nil{
		return ""
	}
	decryptTd :=make([]byte,len(text))
	encryptor :=cipher.NewCFBDecrypter(block,iv)
	encryptor.XORKeyStream(decryptTd,text)
	return  hex.EncodeToString(decryptTd)
}


