package rsa

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDG0qpbOISvJNo2EXiqtFzFoqIz
Hq5m62dX+ifF5UfQwTcgQter5b9L4SS7EuQhQYDqTh07+8K3zcaNs4LEvdcyjDRC
tXjff2Gi/lpnsERCjJ/79YjBirs1MFdqmYueXYUYS9vuwYyRb3/L+0XFcK+0vc0D
KBWDtxJWYEoyrKCxqwIDAQAB
-----END PUBLIC KEY-----
`)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDG0qpbOISvJNo2EXiqtFzFoqIzHq5m62dX+ifF5UfQwTcgQter
5b9L4SS7EuQhQYDqTh07+8K3zcaNs4LEvdcyjDRCtXjff2Gi/lpnsERCjJ/79YjB
irs1MFdqmYueXYUYS9vuwYyRb3/L+0XFcK+0vc0DKBWDtxJWYEoyrKCxqwIDAQAB
AoGBAIdXeGfIPWoMKA6OLhnl1REj+E3cINQmgp68xr5LqWtNW302gyGcr2+zvxEp
dum8cUkOC5B9fLsp9HMZM3dB0258D4adBSPavzwmRQO8tpbfAkRwJpnRCIcLXYQC
gdrV5Vy6ZARctWVwpPy63I02yc8BV3NP+KxzhFIxW+psE6UhAkEA/WjN/ei3KObL
crw2dIIqVPRUNav0dbfArRu3ej9edfRZ2Zb1Qc5AZ2XXDPGurO8UOIveNhyX28Y1
UXkzfOtx7QJBAMjbAMalG+uuvv/e3RlcTGGrOOqPvCrf7DcvZnUiE49UxkCAheaz
3EFyu532BnLS+sFmwuQtadjcIR35/jMMHvcCQQCqRbXEGo6XyRyDJ7xfZOcU1meO
+M/2GVY5+xm55sKN93Y4IpsBcJkN4PwzEmIQiUTJChJxwZy4a+J/sDTMh1exAkA1
N3FZMF3ZcA9DE/VYxs+mAQuvei3M3k9jv7dRyZmNHFT6uhLAMV9mJ9P14j2LJsMx
gtyYQEAPKSPNIXNsawW5AkEApWuDqF2j5ugyLtd4XcNpQjhbR0M7DT6BQf0E1gFh
7VMgRuAMcXvqyPjUBuhF7Vu7XS40L8hjI3tsjkY7s2szjA==
-----END RSA PRIVATE KEY-----
`)

func TestNew(t *testing.T) {
	Convey("rsa crypto test", t, func() {
		rsa, err := New(publicKey, privateKey)
		So(err, ShouldBeNil)
		data := []byte("85452231sada8454212158asd")

		Convey("res encrypt", func() {
			dataEncrypt, err := rsa.Encrypt(data)
			So(err, ShouldBeNil)

			Convey("res decrypt", func() {
				actual, err := rsa.Decrypt(dataEncrypt)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, data)
			})
		})
	})
}
