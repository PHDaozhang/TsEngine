package tsCrypto

import (
	"crypto/cipher"
	"crypto/des"
)

//加密
func TripleDESEncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCSPadding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv[:8])
	sData := make([]byte, len(data))
	blockMode.CryptBlocks(sData, data)
	return sData, nil
}

//解密
func TripleDESDecrypt(sData, key, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv[:8])
	data := make([]byte, len(sData))
	blockMode.CryptBlocks(data, sData)
	data = PKCSUnPadding(data)
	return data, nil
}

//去除补码
func PKCSUnPadding(sData []byte) []byte {
	return PKCS7UnPadding(sData)
}

//补码
func PKCSPadding(data []byte, blockSize int) []byte {
	return PKCS7Padding(data)
}
