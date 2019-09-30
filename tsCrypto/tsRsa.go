package tsCrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/astaxie/beego/logs"
)

//获取密钥对
func GetRsaKeyPair() (pubPem, priPem []byte, err error) {
	priKey, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		logs.Error(err)
		return
	}
	priKeyDer := x509.MarshalPKCS1PrivateKey(priKey)
	priKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   priKeyDer,
	}
	priKeyPem := pem.EncodeToMemory(&priKeyBlock)
	pubKey := priKey.PublicKey
	pubKeyDer, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return nil, nil, err
	}
	pubKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyDer,
	}
	pubKeyPem := pem.EncodeToMemory(&pubKeyBlock)
	logs.Debug(string(priKeyPem))
	logs.Debug(string(pubKeyPem))
	return pubKeyPem, priKeyPem, nil
}

//公钥加密
func RsaEncryptByPub(msg, pubKey []byte) (encryptedData []byte, err error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("[Encrypt]Public Key Error")
	}
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := publicInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
}

//私钥解密
func RsaDecryptByPri(dMsg, priKey []byte) (data []byte, err error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("[Decrypt]Private Key Error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, dMsg)
}

//签名
func RsaSign(data, priKey []byte) (sign []byte, err error) {
	sha := sha256.New()
	sha.Write(data)
	sumSha := sha.Sum(nil)
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("[Sign]Private Key Error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, sumSha)
}

//验签
func RsaVerify(data, pubKey, sign []byte) (err error) {
	sunSha := sha256.Sum256(data)
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return errors.New("[VerifySign]Public Key Error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, sunSha[:], sign)
}
