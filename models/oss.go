package models

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
	contentMd5  string
}

var (
	client *oss.Client
	bucket *oss.Bucket
	ok *OssKey
)

func init() {
	var err error
	ok = &OssKey{
		Sid: "61be7993215c413345e8268b43e57d5f",
		EndPoint: "oss-cn-shanghai.aliyuncs.com",
		BucketName: "pain-resources",
	}
	ok.buildToken()
	ok.getKey()

	client, err = oss.New("https://" + ok.EndPoint, ok.AccessKeyId, ok.AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}

	bucket, err = client.Bucket(ok.BucketName)
	if err != nil {
		log.Fatal(err)
	}

}

func (resp *OssResp) buildUrl(endPoint string, bucketName string) string {
	ext := path.Ext(resp.FileName)
	md5Path := md5.Sum([]byte(resp.FileName))
	path := fmt.Sprintf("ATTACHED/%s/%x/%s%s", time.Now().Format("20060102"), md5Path, resp.Md5, ext)
	resp.Url = fmt.Sprintf("https://%s.%s/%s", bucketName, endPoint, path)
	return path
}

func (resp *OssResp) setFileMd5(file multipart.File) {
	h := md5.New()
	io.Copy(h, file)
	resp.contentMd5 = base64.StdEncoding.EncodeToString(h.Sum(nil))
	resp.Md5 = fmt.Sprintf("%x", h.Sum(nil))
	file.Seek(0, io.SeekStart)
}


// 上传文件
func OssFileUpload(file multipart.File, fileHeader *multipart.FileHeader) (ossResp *OssResp, err error) {
	resp := &OssResp{}
	resp.FileName = fileHeader.Filename
	resp.setFileMd5(file)
	resp.buildUrl(ok.EndPoint, ok.BucketName )

	ossFile := resp.buildUrl(ok.EndPoint, ok.BucketName )
	contentEncoding := oss.ContentEncoding("UTF-8")
	contentMd5 := oss.ContentMD5(resp.contentMd5)
	contentLength := oss.ContentLength(fileHeader.Size)
	if err := bucket.PutObject(ossFile, file, contentEncoding, contentMd5, contentLength); err != nil {
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
func OssFileDelete(url string) {

}


