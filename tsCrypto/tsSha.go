package tsCrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func GetSha1(data []byte) (result string) {
	sha := sha1.New()
	sha.Write(data)
	result = hex.EncodeToString(sha.Sum(nil))
	return
}

func GetSha256(data []byte) (result string) {
	sha := sha256.New()
	sha.Write(data)
	result = hex.EncodeToString(sha.Sum(nil))
	return
}

func RsaSha1(originalData []byte, priKey *rsa.PrivateKey) (encryptedData []byte, err error) {
	hash := sha1.New()
	hash.Write(originalData)
	encryptedData, err = rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA1, hash.Sum(nil))
	return
}

func RsaSha256(originalData []byte, priKey *rsa.PrivateKey) (encryptedData []byte, err error) {
	hash := sha256.New()
	hash.Write(originalData)
	encryptedData, err = rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash.Sum(nil))
	return
}
