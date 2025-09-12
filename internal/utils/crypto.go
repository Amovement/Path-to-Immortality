package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var (
	// 密钥 - 密钥必须是32字节长度
	key = []byte("a123456789abcdefghijklmnopqrstuv")
)

// Encrypt 对字符串进行加密
// plaintext: 要加密的明文
// key: 加密密钥，必须是32字节(256位)
func Encrypt(plaintext string) (string, error) {

	// 检查密钥长度是否符合AES-256要求
	if len(key) != 32 {
		return "", errors.New("密钥必须是32字节长度")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 对明文进行PKCS#7填充
	paddedText := pkcs7Padding([]byte(plaintext), block.BlockSize())

	// 创建一个字节切片用于存储加密后的数据，长度为IV长度+明文长度
	ciphertext := make([]byte, aes.BlockSize+len(paddedText))

	// 生成随机的IV(初始向量)
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 执行加密操作
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedText)

	// 将加密结果转换为Base64字符串返回
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 对加密字符串进行解密
// ciphertext: 要解密的密文(Base64编码)
// key: 解密密钥，必须与加密时使用的密钥相同
func Decrypt(ciphertext string) (string, error) {
	// 检查密钥长度是否符合AES-256要求
	if len(key) != 32 {
		return "", errors.New("密钥必须是32字节长度")
	}

	// 解码Base64字符串
	data, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 检查数据长度是否足够
	if len(data) < aes.BlockSize {
		return "", errors.New("密文太短")
	}

	// 提取IV(初始向量)
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	// 执行解密操作
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	// 去除PKCS#7填充
	unpaddedText, err := pkcs7Unpadding(data)
	if err != nil {
		return "", err
	}

	return string(unpaddedText), nil
}

// pkcs7Padding 对数据进行PKCS#7填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

// pkcs7Unpadding 去除数据的PKCS#7填充
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据长度为0")
	}

	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("无效的填充")
	}

	return data[:length-padding], nil
}
