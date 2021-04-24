package yunzhanghu

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	httpClient *http.Client
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	s          = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	httpClient = &http.Client{}
}

type (
	CommonResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		ReqID   string `json:"request_id"`
	}
)

func randomString(length int) string {
	if length > len(s) || length < 0 {
		return ""
	}
	b := bytes.NewBuffer(nil)
	for i := 0; i < length; i++ {
		t := random.Intn(36)
		b.WriteByte(s[t])
	}
	return b.String()
}

func (y *Yunzhanghu) postJSON(uri, apiName string, obj interface{}) ([]byte, error) {
	var (
		now         = time.Now()
		r           = rand.New(rand.NewSource(now.UnixNano()))
		b, _        = json.Marshal(obj)
		data, err   = TripleDesEncrypt(b, []byte(y.DesKey))
		encodedData = base64.StdEncoding.EncodeToString(data)
		randInt     = r.Intn(99999)
		parms       = fmt.Sprintf(`data=%s&mess=%d&timestamp=%d&key=%s`, string(encodedData), randInt, now.Unix(), y.Appkey)
		requestId   = randomString(10)
	)
	if err != nil {
		return nil, err
	}
	hash := hmac.New(sha256.New, []byte(y.Appkey))
	hash.Write([]byte(parms))
	md := hash.Sum(nil)
	hashStr := hex.EncodeToString(md)

	params := url.Values{}
	params.Add("data", string(encodedData))
	params.Add("mess", strconv.Itoa(randInt))
	params.Add("timestamp", strconv.FormatInt(now.Unix(), 10))
	params.Add("sign", hashStr)
	params.Add("sign_type", "sha256")

	var (
		resp   *http.Response
		req, _ = http.NewRequest(http.MethodPost, uri, strings.NewReader(params.Encode()))
	)
	req.Header.Set("dealer-id", y.Dealer)
	req.Header.Set("request-id", requestId)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (y *Yunzhanghu) decodeWithError(responseBytes []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(responseBytes, obj)
	if err != nil {
		return fmt.Errorf("json.Unmarshal Error, error = %v", err)
	}
	responseObject := reflect.ValueOf(obj)
	if !responseObject.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonResponse := responseObject.Elem().FieldByName("CommonResponse")
	if !commonResponse.IsValid() || commonResponse.Kind() != reflect.Struct {
		return fmt.Errorf("commonResponse is invalid or not struct")
	}
	code := commonResponse.FieldByName("Code")
	msg := commonResponse.FieldByName("Message")
	if !code.IsValid() || !msg.IsValid() {
		return fmt.Errorf("code or message is invalid")
	}
	if code.String() != "0000" {
		return fmt.Errorf("%s Error, errcode=%v , errmsg=%v", apiName, code, msg)
	}
	return nil
}
