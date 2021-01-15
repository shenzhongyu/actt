package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/beego/beego/v2/client/httplib"
	"strings"
	"time"
)

type OssKey struct {
	EndPoint 			string
	AccessKeyId			string
	AccessKeySecret		string
	BucketName			string
	Sid 				string
	Token 				string
}

const api string = "http://www.medtmt.net/Cloud/AliyunAccess/GetKey.ashx"

func init() {

}

func (k *OssKey) buildToken() {
	md5Hash := md5.New()
	dt := time.Now().Format("200601021504")
	str := dt + ":" + "14574bb448f3f49b691c81c0d0c6c417" + ":" + strings.ToLower(k.BucketName)
	md5Hash.Write([]byte(str))
	code := md5Hash.Sum(nil)
	k.Token = hex.EncodeToString(code)
}

func (k *OssKey) getKey () {
	b := httplib.Post(api)
	b.Param("sid", k.Sid)
	b.Param("token", k.Token)
	b.ToJSON(k)
}