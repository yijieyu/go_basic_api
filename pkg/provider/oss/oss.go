package oss

import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOSS struct {
	pool *sync.Pool
}

func New(endpoint, accessKeyID, accessKeySecret, bucket string, options ...oss.ClientOption) *AliyunOSS {
	return &AliyunOSS{
		pool: &sync.Pool{
			New: func() interface{} {
				client, err := oss.New(endpoint, accessKeyID, accessKeySecret, options...)
				if err != nil {
					return err
				}

				b, err := client.Bucket(bucket)
				if err != nil {
					return err
				}

				return b
			},
		},
	}
}

func (o *AliyunOSS) getBucket() (*oss.Bucket, error) {
	switch b := o.pool.Get(); b.(type) {
	case *oss.Bucket:
		return b.(*oss.Bucket), nil
	case error:
		return nil, b.(error)
	}

	return nil, errors.New("failure to establish connection")
}

func (o *AliyunOSS) PutObject(name string, r io.Reader) error {
	bucket, err := o.getBucket()
	if err != nil {
		return err
	}
	defer o.pool.Put(bucket)

	err = bucket.PutObject(name, r)
	return err
}

func (o *AliyunOSS) PutObjectFromFile(name, file string) error {
	bucket, err := o.getBucket()
	if err != nil {
		return err
	}
	defer o.pool.Put(bucket)

	return bucket.PutObjectFromFile(name, file)
}

func (o *AliyunOSS) GetObject(name string) (io.ReadCloser, error) {
	bucket, err := o.getBucket()
	if err != nil {
		return nil, err
	}
	defer o.pool.Put(bucket)

	return bucket.GetObject(name)
}

func (o *AliyunOSS) GetObjectToWriter(name string, w io.Writer) error {
	bucket, err := o.getBucket()
	if err != nil {
		return err
	}
	defer o.pool.Put(bucket)

	b, err := bucket.GetObject(name)
	if err != nil {
		return err
	}

	defer b.Close()

	_, err = io.Copy(w, b)
	return err
}

func (o *AliyunOSS) GetObjectToFile(name, file string) error {
	bucket, err := o.getBucket()
	if err != nil {
		return err
	}
	defer o.pool.Put(bucket)

	return bucket.GetObjectToFile(name, file)
}

func (o *AliyunOSS) IsObjectExist(name string) (bool, error) {
	bucket, err := o.getBucket()
	if err != nil {
		return false, err
	}
	defer o.pool.Put(bucket)

	return bucket.IsObjectExist(name)
}

func (o *AliyunOSS) ListObjectsForPrefix(prefix string) ([]string, error) {
	bucket, err := o.getBucket()
	if err != nil {
		return nil, err
	}
	defer o.pool.Put(bucket)

	res := make([]string, 0, 20)
	marker := oss.Marker("")
	for {
		lsRes, err := bucket.ListObjects(oss.MaxKeys(200), oss.Prefix(prefix), marker)
		if err != nil {
			return nil, err
		}

		marker = oss.Marker(lsRes.NextMarker)
		for _, object := range lsRes.Objects {
			res = append(res, object.Key)
		}

		if !lsRes.IsTruncated {
			break
		}
	}

	return res, nil
}

func (o *AliyunOSS) IsUpdate(name string, modTime time.Time) (bool, error) {
	bucket, err := o.getBucket()
	if err != nil {
		return false, err
	}
	defer o.pool.Put(bucket)

	header, err := bucket.GetObjectMeta(name)
	if err != nil {
		return true, err
	}

	laseDate, err := time.Parse(time.RFC1123, header.Get(oss.HTTPHeaderLastModified))
	if err != nil {
		return true, err
	}

	return laseDate.Before(modTime), nil
}

func (o *AliyunOSS) CopyObject(scrName, destName string) error {
	bucket, err := o.getBucket()
	if err != nil {
		return err
	}
	defer o.pool.Put(bucket)
	_, err = bucket.CopyObject(scrName, destName)

	return err
}
