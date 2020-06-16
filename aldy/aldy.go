/*
	目前使用阿里大于短信服务，本接口用于阿里大于迁移到阿里云通信后使用。
*/
package aldy

import (
	"crypto/hmac"
	"encoding/base64"

	"crypto/sha1"
	"encoding/json"

	"io/ioutil"

	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

// sendSmsResponse
type sendSmsResponse struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}

const (
	dySmsApiUrl = "http://dysmsapi.aliyuncs.com"
)

// signHMAC 获取签名
func signHMAC(params url.Values, appSecret string) (signature string) {
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	str := ""
	sort.Strings(keys)
	for _, k := range keys {
		str += "&" + url.QueryEscape(k) + "=" + url.QueryEscape(params.Get(k))
	}
	signstr := "GET&%2F&" + url.QueryEscape(str[1:])
	mac := hmac.New(sha1.New, []byte(appSecret+"&"))
	mac.Write([]byte(signstr))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// SendSMS
func SendSMS(mobileNo, signName, templateCode, paramString, appKey, appSecret string) (bool, string, error) {
	params := url.Values{}

	params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	params.Set("SignatureMethod", "HMAC-SHA1")
	params.Set("SignatureVersion", "1.0")
	uId := uuid.NewV4()
	params.Set("SignatureNonce", strings.ToLower(uId.String()))
	params.Set("AccessKeyId", appKey)
	params.Add("Format", "JSON")
	params.Set("RegionId", "cn-hangzhou")

	params.Set("SignName", signName)
	params.Set("TemplateCode", templateCode)
	params.Set("TemplateParam", paramString)
	params.Set("OutId", "")
	params.Set("Action", "SendSms")
	params.Set("PhoneNumbers", mobileNo)
	params.Set("Version", "2017-05-25")

	signstr := signHMAC(params, appSecret)
	params.Set("Signature", signstr)
	req, err := http.NewRequest(http.MethodGet, dySmsApiUrl+"/?"+params.Encode(), nil)
	if err != nil {
		return false, "", err
	}

	req.Header.Set("x-sdk-client", "wl520")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "wl520")

	c := new(http.Client)
	resp, err := c.Do(req)

	if err != nil {
		return false, "", err
	}

	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "", err
	}
	var result sendSmsResponse
	err = json.Unmarshal(bs, &result)
	if err != nil {
		return false, "", err
	}
	return result.Code == "OK", result.Message, nil
}
