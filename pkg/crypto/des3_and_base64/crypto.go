package des3AndBase64

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"

	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/crypto"
)

type des3AndBase64 struct {
	block  cipher.Block
	vector []byte
}

func NewCrypto(key, vector []byte) (crypto.Crypto, error) {
	c := &des3AndBase64{}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return c, fmt.Errorf("NewTripleDESCipher failure for key:%s, error:%v", key, err)
	}
	c.vector = vector
	c.block = block
	return c, nil
}

func (d *des3AndBase64) Encrypt(src []byte) ([]byte, error) {
	c, err := d.des3Encrypt(src)
	if err != nil {
		return []byte{}, fmt.Errorf("des3AndBase64 des3Encrypt failure, error:%v", err)
	}

	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(c)))
	base64.RawURLEncoding.Encode(buf, c)
	return buf, nil
}

func (d *des3AndBase64) Decrypt(src []byte) ([]byte, error) {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(src)))
	n, err := base64.RawURLEncoding.Decode(buf, src)
	if err != nil {
		return nil, fmt.Errorf("des3AndBase64 base64decode failure, error:%v", err)
	}

	res, err := d.des3Decrypt(buf[:n])
	if err != nil {
		return []byte{}, fmt.Errorf("des3AndBase64 des3Decrypt failure, error:%v", err)
	}

	return res, nil
}

func (d *des3AndBase64) des3Decrypt(c []byte) (origData []byte, err error) {
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

	blockMode := cipher.NewCBCDecrypter(d.block, d.vector)
	origData = make([]byte, len(c))
	blockMode.CryptBlocks(origData, c)
	origData = d.pkcs5UnPadding(origData)

	return
}

func (d *des3AndBase64) des3Encrypt(originData []byte) (c []byte, err error) {
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

	origData := d.pkcs5Padding(originData, d.block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(d.block, d.vector)
	c = make([]byte, len(origData))
	blockMode.CryptBlocks(c, origData)
	return
}

func (d *des3AndBase64) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func (d *des3AndBase64) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
