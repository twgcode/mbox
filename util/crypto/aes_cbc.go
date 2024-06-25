/**
@Author: wei-g
@Date:   2021/8/11 5:13 下午
@Description:
*/

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// PKCS5Padding 填充明文
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}

// PKCS5UnPadding 去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// AesCBCEncrypt AES CBC模式 加密
func AesCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypt := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypt, origData)
	return encrypt, nil
}

// AesCBCEncryptEncodeToBase64String 加密后使用 base64编码
func AesCBCEncryptEncodeToBase64String(origData, key []byte) (encrypt string, err error) {
	var data []byte
	if data, err = AesCBCEncrypt(origData, key); err != nil {
		return
	}
	encrypt = base64.StdEncoding.EncodeToString(data)
	return
}

// AesCBCDecrypt AES CBC模式 解密
func AesCBCDecrypt(encrypt, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(encrypt))
	blockMode.CryptBlocks(origData, encrypt)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// AesCBCDecryptDecodeBase64String base64解码后在解密
func AesCBCDecryptDecodeBase64String(data string, key []byte) (decrypt []byte, err error) {
	var encrypt []byte
	if encrypt, err = base64.StdEncoding.DecodeString(data); err != nil {
		return
	}
	decrypt, err = AesCBCDecrypt(encrypt, key)
	return
}
