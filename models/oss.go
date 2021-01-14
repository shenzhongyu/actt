package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/beego/beego/v2/server/web"
	"io"
	"log"
	"mime/multipart"
	"path"
	"time"
)

type OssResp struct {
	Url 		string			`json:"url"`
	Md5 		string			`json:"md5"`
	FileName	string			`json:"filename"`
}

var (
	client *oss.Client
	bucket *oss.Bucket
)

func init() {
	var err error
	EndPoint, _ := web.AppConfig.String("AliOss::Endpoint")
	AccessKeyId, _ := web.AppConfig.String("AliOss::AccessKeyId")
	AccessKeySecret, _ := web.AppConfig.String("AliOss::AccessKeySecret")
	Bucket, _ :=  web.AppConfig.String("AliOss::BucketName")
	client, err = oss.New("http://" + EndPoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}

	bucket, err = client.Bucket(Bucket)
	if err != nil {
		log.Fatal(err)
	}

}

func (resp *OssResp) buildUrl(endPoint string, bucketName string) string {
	ext := path.Ext(resp.FileName)
	md5Path := md5.Sum([]byte(resp.FileName))
	path := fmt.Sprintf("ATTACHED/%s/%x/%s%s", time.Now().Format("20060102"), md5Path, resp.Md5, ext)
	resp.Url = fmt.Sprintf("http://%s.%s/%s", bucketName, endPoint, path)
	return path
}

func (resp *OssResp) setFileMd5(file multipart.File) {
	f := md5.New()
	io.Copy(f, file)
	fMd5 := f.Sum([]byte(""))
	file.Seek(0, 0)
	resp.Md5 = hex.EncodeToString(fMd5[:])
}


// 上传文件
func OssFileUpload(objectKey string, reader multipart.File) (ossResp *OssResp, err error) {
	EndPoint, _ := web.AppConfig.String("AliOss::Endpoint")
	BucketName, _ :=  web.AppConfig.String("AliOss::BucketName")

	resp := &OssResp{}
	resp.FileName = objectKey
	resp.setFileMd5(reader)
	ossFile := resp.buildUrl(EndPoint, BucketName)

	if err := bucket.PutObject(ossFile, reader); err != nil {
		return nil, err
	}

	// 获取存储空间。
	return resp, nil
}

// 列举文件
func GetOssFileList() {

}

// 下载文件
func OssFileDownload() {

}

// 删除文件
func OssFileDelete() {

}


