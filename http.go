package yunzhanghu

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	httpClient *http.Client
	s          = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	httpClient = &http.Client{}
}

type (
	CommonResponse struct {
		Code    StatusCode `json:"code"`
		Message string     `json:"message"`
		ReqID   string     `json:"request_id"`
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

func (y *Yunzhanghu) getJson(ctx context.Context, uri, apiName string, obj interface{}) ([]byte, error) {
	req, err := y.buildRequest(http.MethodGet, uri, apiName, obj)
	if err != nil {
		return nil, err
	}
	return y.doRequest(ctx, req)
}

func (y *Yunzhanghu) postJSON(ctx context.Context, uri, apiName string, obj interface{}) ([]byte, error) {
	req, err := y.buildRequest(http.MethodPost, uri, apiName, obj)
	if err != nil {
		return nil, err
	}
	return y.doRequest(ctx, req)
}

func (y *Yunzhanghu) postForm(ctx context.Context, uri, apiName string, obj interface{}, files map[string]io.Reader) ([]byte, error) {
	req, err := y.buildFormRequest(uri, apiName, obj, files)
	if err != nil {
		return nil, err
	}
	return y.doRequest(ctx, req)
}

func (y *Yunzhanghu) buildFormRequest(uri, apiName string, obj interface{}, files map[string]io.Reader) (*http.Request, error) {
	buf := bytes.NewBuffer(nil)
	mw := multipart.NewWriter(buf)
	for name, r := range files {
		var (
			fw  io.Writer
			err error
		)
		if c, ok := r.(io.Closer); ok {
			defer c.Close()
		}
		if f, ok := r.(*os.File); ok {
			fw, err = mw.CreateFormFile(name, f.Name())
		} else {
			fw, err = mw.CreateFormField(name)
		}
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}
	mw.Close()
	var (
		req, _    = http.NewRequest(http.MethodPost, y.ApiAddr+uri, buf)
		requestId string
		err       error
	)
	req.URL.RawQuery, requestId, err = y.buildParams(obj)
	if err != nil {
		return nil, err
	}
	req.Header.Set("dealer-id", y.Dealer)
	req.Header.Set("request-id", requestId)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	return req, nil
}

func (y *Yunzhanghu) buildRequest(method, uri, apiName string, obj interface{}) (*http.Request, error) {
	var (
		req, _    = http.NewRequest(method, y.ApiAddr+uri, nil)
		requestId string
		err       error
	)
	req.URL.RawQuery, requestId, err = y.buildParams(obj)
	if err != nil {
		return nil, err
	}
	req.Header.Set("dealer-id", y.Dealer)
	req.Header.Set("request-id", requestId)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (y *Yunzhanghu) buildParams(obj interface{}) (requestId string, query string, err error) {
	var (
		now     = time.Now()
		b, _    = json.Marshal(obj)
		randInt = random.Intn(99999)
		data    []byte
	)
	data, err = TripleDesEncrypt(b, []byte(y.DesKey))
	if err != nil {
		return
	}
	encodedData := base64.StdEncoding.EncodeToString(data)
	hash := hmac.New(sha256.New, []byte(y.Appkey))
	parms := fmt.Sprintf(`data=%s&mess=%d&timestamp=%d&key=%s`, string(encodedData), randInt, now.Unix(), y.Appkey)
	hash.Write([]byte(parms))
	md := hash.Sum(nil)
	hashStr := hex.EncodeToString(md)
	requestId = randomString(10)
	params := url.Values{}
	params.Add("data", string(encodedData))
	params.Add("mess", strconv.Itoa(randInt))
	params.Add("timestamp", strconv.FormatInt(now.Unix(), 10))
	params.Add("sign", hashStr)
	params.Add("sign_type", "sha256")
	query = params.Encode()
	return
}

func (y *Yunzhanghu) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	req = req.WithContext(ctx)
	var resp, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s %s %d", req.Method, req.URL.String(), resp.StatusCode)
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
