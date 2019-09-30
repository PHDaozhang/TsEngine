package tsCrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
)

type Des struct {
	Strkey     string
	Iv         string // 密码
	PadType    int    // 数据补齐方式
	EncodeType int
}

func (this *Des) getKey() []byte {

	keyLen := len(this.Strkey)
	if keyLen < 8 {
		panic("des key 长度不能小于8")
	}
	arrKey := []byte(this.Strkey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	} else if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	} else if keyLen >= 16 {
		//取前16个字节
		return arrKey[:16]
	} else {
		return arrKey[:8]
	}
}

func (this *Des) getIv(size int) []byte {
	ivLen := len(this.Iv)
	if ivLen < 8 {
		panic("res iv 长度不能小于8")
	}
	arriv := []byte(this.Iv)
	if size == 32 && ivLen >= 32 {
		//取前32个字节
		return arriv[:32]
	} else if size == 24 && ivLen >= 24 {
		//取前24个字节
		return arriv[:24]
	} else if size == 16 && ivLen >= 16 {
		//取前16个字节
		return arriv[:16]
	} else {
		//取前8个字节
		return arriv[:8]
	}
}

//加密字符串
func (this *Des) Encrypt(origData []byte) ([]byte, error) {
	key := this.getKey()
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if this.PadType == PadByte {
		origData = ZeroPadding(origData, blockSize)
	} else {
		origData = PKCS5Padding(origData, blockSize)
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
func (this *Des) EncryptCBC(origData []byte) (string, error) {
	block, err := des.NewCipher(this.getKey())
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	if this.PadType == PadByte {
		origData = ZeroPadding(origData, blockSize)
	} else {
		origData = PKCS5Padding(origData, blockSize)
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
func (this *Des) Decrypt(data string) ([]byte, error) {

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
		origData = PKCS5UnPadding(origData)
	}
	return origData, nil
}

func (this *Des) DecryptCBC(data string) ([]byte, error) {

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

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, this.getIv(block.BlockSize()))
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	if this.PadType == PadByte {
		origData = ZeroUnPadding(origData)
	} else {
		origData = PKCS5UnPadding(origData)
	}

	return origData, nil
}
