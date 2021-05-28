package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"strconv"
)

// 生成字符串的md5值
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func MD5ToInt(str string, n int) (sum uint64) {
	md5 := MD5(str)
	text := md5[len(md5)-n:]
	size := n
	for i, c := range text {
		power := size - i - 1
		if val, err := strconv.Atoi(string(c)); err == nil {
			sum += uint64(val) << (4 * power) //数字
		} else {
			sum += uint64(byte(c)-byte('a')+10) << (4 * power)
		}
	}
	return sum
}

func MD5File(file string) (MD5Str string, err error) {
	c := md5.New()
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	_, err = io.Copy(c, f)
	if err != nil {
		return "", err
	}
	MD5Str = hex.EncodeToString(c.Sum(nil))
	return MD5Str, nil
}

//生成sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func SHA256String(s string) string {
	res := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", res)
}

// File SHA256
func SHA256File(filename string) (s string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	s = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func HmacSha256(data string, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

func Base64Encode(data []byte) (enc string) {
	enc = b64.StdEncoding.EncodeToString(data)
	return
}

func Base64Decode(enc string) (unDec []byte) {
	unDec, _ = b64.StdEncoding.DecodeString(enc)
	return

}
