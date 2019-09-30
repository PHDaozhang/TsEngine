package tsCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

/***
加密
  加密为Hex     tsCrypto.EncodeHex
  加密为Base64  tsCrypto.EncodeBase64
aesGb := tsCrypto.AesGB{Strkey: aesKey, Iv: aesKey, EncodeType: tsCrypto.EncodeHex, PadType: tsCrypto.PadString}
if encryptByte, err := aesGb.EncryptCBC([]byte(encryptValue)); err == nil {
	token = encryptByte
}

解密
  密文为Hex     tsCrypto.EncodeHex
  密文为Base64  tsCrypto.EncodeBase64
aesGb.DecryptCBC()

*/
//用户表模型
type AesGB struct {
	Strkey     string
	Iv         string // 密码
	PadType    int    // 数据补齐方式
	EncodeType int
}

func (this *AesGB) getKey() []byte {

	keyLen := len(this.Strkey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(this.Strkey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

func (this *AesGB) getIv(size int) []byte {
	ivLen := len(this.Iv)
	if ivLen < 16 {
		panic("res iv 长度不能小于16")
	}
	arriv := []byte(this.Iv)
	if size == 32 && ivLen >= 32 {
		//取前32个字节
		return arriv[:32]
	}
	if size == 24 && ivLen >= 24 {
		//取前24个字节
		return arriv[:24]
	}
	//取前16个字节
	return arriv[:16]
}

//加密字符串
func (this *AesGB) Encrypt(origData []byte) ([]byte, error) {
	key := this.getKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if this.PadType == PadByte {
		origData = ZeroPadding(origData, blockSize)
	} else {
		origData = PKCS7Padding(origData)
	}
	blockMode := NewECBEncrypter(block)

	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	if this.EncodeType == EncodeHex {
		return []byte(hex.EncodeToString(crypted)), nil
	} else {
		return []byte(base64.StdEncoding.EncodeToString(crypted)), nil
	}
}

/***
通用
*/
func (this *AesGB) EncryptCBC(origData []byte) (string, error) {
	block, err := aes.NewCipher(this.getKey())
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	if this.PadType == PadByte {
		origData = ZeroPadding(origData, blockSize)
	} else {
		origData = PKCS7Padding(origData)
	}

	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, len(origData))

	//block大小和初始向量大小一定要一致
	blockMode := cipher.NewCBCEncrypter(block, this.getIv(blockSize))
	blockMode.CryptBlocks(cipherText, origData)

	if this.EncodeType == EncodeHex {
		return hex.EncodeToString(cipherText), nil
	} else {
		return base64.StdEncoding.EncodeToString(cipherText), nil
	}
}

//解密字符串
func (this *AesGB) Decrypt(data string) ([]byte, error) {

	var crypted []byte
	var err error

	if this.EncodeType == EncodeHex {
		crypted, err = hex.DecodeString(data)
	} else {
		crypted, err = base64.StdEncoding.DecodeString(data)
	}

	if err != nil {
		return nil, err
	}
	key := this.getKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	if this.PadType == PadByte {
		origData = ZeroUnPadding(origData)
	} else {
		origData = PKCS7UnPadding(origData)
	}
	return origData, nil
}

func (this *AesGB) DecryptCBC(data string) ([]byte, error) {

	var crypted []byte
	var err error

	if this.EncodeType == EncodeHex {
		crypted, err = hex.DecodeString(data)
	} else {
		crypted, err = base64.StdEncoding.DecodeString(data)
	}

	if err != nil {
		return nil, err
	}
	key := this.getKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, this.getIv(block.BlockSize()))
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	if this.PadType == PadByte {
		origData = ZeroUnPadding(origData)
	} else {
		origData = PKCS7UnPadding(origData)
	}

	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if padding == blockSize {
		return ciphertext
	}
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {

	index := bytes.IndexByte(origData, 0)
	if index == -1 {
		return origData
	}
	rbyf_pn := origData[0:index]
	return rbyf_pn

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS7Padding(cipherText []byte) []byte {
	padding := aes.BlockSize - len(cipherText)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}
type ecbDecrypter ecb
type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {

	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}
