package models

import (
	"crypto/md5"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
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
	body, _ := ioutil.ReadAll(file)
	file.Seek(0, io.SeekStart)
	resp.Md5 = fmt.Sprintf("%x", md5.Sum(body))
}


// 上传文件
func OssFileUpload(objectKey string, reader multipart.File) (ossResp *OssResp, err error) {
	resp := &OssResp{}
	resp.FileName = objectKey
	resp.setFileMd5(reader)
	ossFile := resp.buildUrl(ok.EndPoint, ok.BucketName )
	if err := bucket.PutObject(ossFile, reader); err != nil {
		return nil, err
	}
	reader.Seek(0, 0)
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


