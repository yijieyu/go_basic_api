package des3AndBase64

import (
	"bytes"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDes3AndBase64_Encrypt(t *testing.T) {
	Convey("Test data encryption", t, func() {
		key := []byte("038bcd079067797996c78200")
		vector := []byte("76540125")
		expected := []byte("o-U72kkZGsw7PWz-i2Wjl1FRwYvgJM-XFcw-nITSv-w")

		Convey("Routine test", func() {
			crypto, err := NewCrypto(key, vector)
			So(err, ShouldBeNil)

			Convey("encrypt result should resemble expected", func() {
				encryptByte := []byte("{requestId:5,\"adId\":\"2\"}")
				result, err := crypto.Encrypt(encryptByte)
				So(err, ShouldBeNil)

				So(result, ShouldResemble, expected)
			})
		})

		Convey("Smoke test", func() {
			crypto, err := NewCrypto(key, []byte("1"))
			So(err, ShouldBeNil)

			Convey("encrypt error not be nil", func() {
				encryptByte := []byte("{requestId:5,\"adId\":\"2\"}")
				_, err := crypto.Encrypt(encryptByte)
				So(err, ShouldNotBeNil)
			})
		})

	})
}

func TestDes3AndBase64_Decrypt(t *testing.T) {
	Convey("Test data encryption", t, func() {
		key := []byte("038bcd079067797996c78200")
		vector := []byte("76540125")
		expected := []byte("{requestId:5,\"adId\":\"2\"}")

		Convey("Routine test", func() {
			crypto, err := NewCrypto(key, vector)
			So(err, ShouldBeNil)

			Convey("decrypt result should resemble expected", func() {
				decryptByte := []byte("o-U72kkZGsw7PWz-i2Wjl1FRwYvgJM-XFcw-nITSv-w")
				result, err := crypto.Decrypt(decryptByte)
				So(err, ShouldBeNil)

				So(result, ShouldResemble, expected)
			})
		})

		Convey("Smoke test", func() {
			crypto, err := NewCrypto(key, []byte("1"))
			So(err, ShouldBeNil)

			Convey("decrypt error not be nil", func() {
				decryptByte := []byte("o-U72kkZGsw7PWz-i2Wjl1FRwYvgJM-XFcw-nITSv-w")
				_, err := crypto.Decrypt(decryptByte)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Be decrypted invalid characters", func() {
			crypto, err := NewCrypto(key, vector)
			So(err, ShouldBeNil)

			Convey("decrypt error not be nil", func() {
				decryptByte := []byte("o-U72kkZGsw7PWz-i2Wjl1FRwYvgJM+XFcw-nITSv-w")
				_, err := crypto.Decrypt(decryptByte)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestConcurrent_Security_Testing(t *testing.T) {
	randBytes := func(length int) []byte {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		buffer := bytes.NewBufferString("")
		for i := 0; i < length; i++ {
			isLetter := r.Intn(2)
			if isLetter > 0 {
				letter := r.Intn(52)
				if letter < 26 {
					letter += 97
				} else {
					letter += 65 - 26
				}
				buffer.WriteString(string(letter))
			} else {
				buffer.WriteString(strconv.Itoa(r.Intn(10)))
			}
		}
		return buffer.Bytes()
	}

	Convey("Concurrent security testing", t, func() {
		key := []byte("038bcd079067797996c78200")
		vector := []byte("76540125")

		Convey("Smoke test", func() {
			crypto, err := NewCrypto(key, vector)
			So(err, ShouldBeNil)

			checkFunc := func(w *sync.WaitGroup) {
				for i := 0; i < 100; i++ {

					s := randBytes(50)
					res, err := crypto.Encrypt(s)
					if err != nil {
						t.Error("encrypt failure")
					}

					result, err := crypto.Decrypt(res)
					if err != nil {
						t.Error("decrypt failure")
					}

					if !bytes.Equal(result, s) {
						t.Errorf("Before docking encrypted encrypted data inconsistencies %v != %s", s, result)
					}
				}
				w.Done()
			}

			Convey("decrypt result should resemble expected", func() {
				w := &sync.WaitGroup{}
				for i := 0; i < 1000; i++ {
					w.Add(1)
					go checkFunc(w)
				}
				w.Wait()
			})
		})
	})
}
