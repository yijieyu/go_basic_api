package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yijieyu/go_basic_api/pkg/constant"
	"github.com/yijieyu/go_basic_api/pkg/crypto"
	aesAndBase64 "github.com/yijieyu/go_basic_api/pkg/crypto/aes_and_base64"
)

func Crypto(aes *aesAndBase64.AesAndBase64, rsa crypto.Decrypt, force bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.Request.Header.Get("Content-MD5") // 前1个字节为算法版本
		authKey := strings.TrimSpace(c.Request.Header.Get("X-Authorization"))
		header := c.Request.Header.Get("Content-Body")
		log := logrus.WithFields(logrus.Fields{
			"request_id":      c.GetString(constant.XRequestID),
			"query":           c.Request.URL.String(),
			"method":          c.Request.Method,
			"Content-MD5":     sign,
			"X-Authorization": authKey,
			"Content-Body":    header,
		})

		if authKey == "" && !force {
			log.Debug("not crypto")
			return
		}

		// 用于接口检查是否是加密请求
		defer func() {
			if !c.IsAborted() {
				c.Set(constant.Crypto, true)
			}
		}()

		if authKey == "" {
			log.Error("crypto key not found")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		originalSign, err := hex.DecodeString(sign)
		if err != nil {
			log.Errorf("Content-MD5 Parse fail %v", err)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if len(originalSign) < 1 {
			log.Errorf("Content-MD5 length fail")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if originalSign[0] != 1 {
			log.Errorf("crypto algorithm version won't support it %d", originalSign[0])
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Request.Header.Del("Content-MD5")
		c.Request.Header.Del("X-Authorization")
		c.Request.Header.Del("Content-Body")

		key, err := hex.DecodeString(authKey)
		if err != nil {
			log.Errorf("X-Authorization Parse fail %v", err)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		key, err = rsa.Decrypt(key)
		if err != nil {
			log.Errorf("rsa X-Authorization Decrypt fail %v", err)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		var iv, body, data []byte
		if header != "" {
			iv, data, err = aes.Decrypt(key, []byte(header))
			if err != nil {
				log.Errorf("aes decrypt request.Body fail %v", err)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			log.WithFields(logrus.Fields{
				"header_query": string(data),
			}).Debug("header query")

			query, err := url.ParseQuery(string(data))
			if err != nil {
				log.WithFields(logrus.Fields{
					"header_data": string(data),
				}).Errorf("url.ParseQuery common data fail %v", err)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			for k := range query {
				for _, v := range query[k] {
					c.Request.Header.Add(k, v)
				}
			}
		}

		defer func() {
			if len(iv) != 16 {
				iv = make([]byte, 16)
			}

			hmacKey := make([]byte, 32)
			xor(key[:16], iv, hmacKey[:16])
			xor(key[16:], iv, hmacKey[16:])
			sha := hmac.New(sha256.New, hmacKey)
			sha.Write([]byte(header))
			sha.Write(data)
			if hex.EncodeToString(sha.Sum([]byte{originalSign[0]})) != sign {
				log.Errorf("sign verify fail %x != %s", sha.Sum([]byte{originalSign[0]}), sign)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}()

		if c.Request.Body != nil {
			data, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Errorf("read request.Body fail %v", err)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			if len(data) < 1 {
				log.Debug("request.Body is empty")
				return
			}

			var biv []byte
			biv, body, err = aes.Decrypt(key, data)
			if err != nil {
				log.Errorf("aes decrypt request.Body fail %v", err)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

			if len(iv) < 1 {
				iv = biv
			}

			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}
}

func xor(b1, b2, buf []byte) []byte {
	for i := 0; i < len(buf); i++ {
		buf[i] = b1[i] ^ b2[i]
	}

	return buf
}
