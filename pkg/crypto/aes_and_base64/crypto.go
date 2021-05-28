package aesAndBase64

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

type AesAndBase64 struct{}

func NewCrypto() *AesAndBase64 {
	return &AesAndBase64{}
}

func (d *AesAndBase64) Encrypt(key, iv, src []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("NewCipher failure for key:%s, error:%v", key, err)
	}

	c, err := d.encrypt(block, iv, src)
	if err != nil {
		return []byte{}, fmt.Errorf("AesAndBase64 Encrypt failure, error:%v", err)
	}

	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(c)+len(iv)))
	base64.RawURLEncoding.Encode(buf, append(iv, c...))
	return buf, nil
}

func (d *AesAndBase64) Decrypt(key, src []byte) ([]byte, []byte, error) {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(src)))
	n, err := base64.RawURLEncoding.Decode(buf, src)
	if err != nil {
		return nil, nil, fmt.Errorf("AesAndBase64 base64decode failure, error:%v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, fmt.Errorf("NewCipher failure for key:%s, error:%v", key, err)
	}

	buf = buf[:n]
	res, err := d.decrypt(block, buf[:block.BlockSize()], buf[block.BlockSize():])
	if err != nil {
		return nil, nil, fmt.Errorf("AesAndBase64 Decrypt failure, error:%v", err)
	}

	return buf[:block.BlockSize()], res, nil
}

func (d *AesAndBase64) decrypt(block cipher.Block, iv, src []byte) (origData []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData = make([]byte, len(src))
	blockMode.CryptBlocks(origData, src)
	origData = d.pkcs5UnPadding(origData)

	return
}

func (d *AesAndBase64) encrypt(block cipher.Block, iv, src []byte) (c []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	origData := d.pkcs5Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	c = make([]byte, len(origData))
	blockMode.CryptBlocks(c, origData)
	return
}

func (d *AesAndBase64) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func (d *AesAndBase64) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
