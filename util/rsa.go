package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	//"crypto/rand"
	//"crypto/rsa"
	//"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	//"github.com/wumansgy/goEncrypt/rsa"
	"os"
)

// 从文件中读取RSA key
func RSAReadKeyFromFile(filename string) []byte {
	f, err := os.Open(filename)
	var b []byte

	if err != nil {
		return b
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b = make([]byte, fileInfo.Size())
	f.Read(b)
	return b
}

// RSA加密
//func RSAEncrypt(data, publicBytes []byte) ([]byte, error) {
//	var res []byte
//	// 解析公钥
//	block, _ := pem.Decode(publicBytes)
//
//	if block == nil {
//		return res, fmt.Errorf("无法加密, 公钥可能不正确")
//	}
//
//	// 使用X509将解码之后的数据 解析出来
//	// x509.MarshalPKCS1PublicKey(block):解析之后无法用，所以采用以下方法：ParsePKIXPublicKey
//	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		return res, fmt.Errorf("无法加密, 公钥可能不正确, %v", err)
//	}
//	// 使用公钥加密数据
//	pubKey := keyInit.(*rsa.PublicKey)
//	res, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
//	if err != nil {
//		return res, fmt.Errorf("无法加密, 公钥可能不正确, %v", err)
//	}
//	// 将数据加密为base64格式
//	return []byte(EncodeStr2Base64(string(res))), nil
//}

// 对数据进行解密操作
func RSADecrypt(base64Data string, privateBytes []byte) ([]byte, error) {
	//var res []byte

	// 解析私钥
	block, _ := pem.Decode(privateBytes)
	//if block == nil {
	//	return res, fmt.Errorf("无法解密, 私钥可能不正确")
	//}
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	decodeString, err2 := base64.URLEncoding.DecodeString(base64Data)
	v15, err2 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeString)

	//decrypt, err := RsaDecryptByBase64(base64Data, base64.StdEncoding.EncodeToString(privateBytes))
	if err2 != nil {
		return nil, err2
	}
	return v15, nil

}

// 加密base64字符串
func EncodeStr2Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解密base64字符串
func DecodeStrFromBase64(str string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(decodeBytes)
}
