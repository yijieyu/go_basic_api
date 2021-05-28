package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/crypto"
)

type Rsa struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func New(publicKey, privateKey []byte) (crypto.Crypto, error) {
	encrypt, err := NewEncrypt(publicKey)
	if err != nil {
		return nil, err
	}

	decrypt, err := NewDecrypt(privateKey)
	if err != nil {
		return nil, err
	}

	return &Rsa{
		publicKey:  encrypt.(*Rsa).publicKey,
		privateKey: decrypt.(*Rsa).privateKey,
	}, nil
}

func NewEncrypt(publicKey []byte) (crypto.Encrypt, error) {
	r := &Rsa{}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	r.publicKey = pubInterface.(*rsa.PublicKey)

	return r, nil
}

func NewDecrypt(privateKey []byte) (crypto.Decrypt, error) {
	r := &Rsa{}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	var err error
	r.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Rsa) Encrypt(src []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, src)
}

func (r *Rsa) Decrypt(src []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, src)
}
