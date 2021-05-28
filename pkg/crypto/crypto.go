package crypto

// 密码
type Crypto interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

// 加密
type Encrypt interface {
	Encrypt([]byte) ([]byte, error)
}

// 解密
type Decrypt interface {
	Decrypt([]byte) ([]byte, error)
}
