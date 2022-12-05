package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
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
		Code      StatusCode `json:"code"`
		Message   string     `json:"message"`
		RequestId string     `json:"request_id"`
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

func GetJson(ctx *Context, uri string, obj interface{}) ([]byte, error) {
	req, err := buildRequest(ctx, http.MethodGet, uri, obj)
	if err != nil {
		return nil, err
	}
	return doRequest(ctx, req)
}

func PostJSON(ctx *Context, uri string, obj interface{}) ([]byte, error) {
	req, err := buildRequest(ctx, http.MethodPost, uri, obj)
	if err != nil {
		return nil, err
	}
	return doRequest(ctx, req)
}

func PostForm(ctx *Context, uri, apiName string, obj interface{}, files map[string]io.Reader) ([]byte, error) {
	req, err := buildFormRequest(ctx, uri, apiName, obj, files)
	if err != nil {
		return nil, err
	}
	return doRequest(ctx, req)
}

func buildFormRequest(ctx *Context, uri, apiName string, obj interface{}, files map[string]io.Reader) (*http.Request, error) {
	var (
		buf       = bytes.NewBuffer(nil)
		mw        = multipart.NewWriter(buf)
		req       *http.Request
		requestId string
		err       error
		params    url.Values
	)
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
	requestId, params, err = buildParams(ctx, obj)
	if err != nil {
		return nil, err
	}
	if _, err = mw.CreatePart(textproto.MIMEHeader(params)); err != nil {
		return nil, err
	}
	if err = mw.Close(); err != nil {
		return nil, err
	}
	u, _ := url.Parse(ctx.ApiAddr)
	u.Path = uri
	u.RawQuery = params.Encode()
	req, _ = http.NewRequest(http.MethodPost, u.String(), buf)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("dealer-id", ctx.Dealer)
	req.Header.Set("request-id", requestId)
	return req, nil
}

func buildRequest(ctx *Context, method, uri string, obj interface{}) (*http.Request, error) {
	var (
		req       *http.Request
		requestId string
		err       error
		params    url.Values
	)
	requestId, params, err = buildParams(ctx, obj)
	if err != nil {
		return nil, err
	}
	u, _ := url.Parse(ctx.ApiAddr)
	u.Path = uri
	u.RawQuery = params.Encode()
	req, _ = http.NewRequest(method, u.String(), nil)
	req.Header.Set("dealer-id", ctx.Dealer)
	req.Header.Set("request-id", requestId)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func buildParams(ctx *Context, obj interface{}) (requestId string, params url.Values, err error) {
	var (
		now     = time.Now().Unix()
		b, _    = json.Marshal(obj)
		randInt = random.Intn(99999)
		data    []byte
		hashStr string
	)
	data, err = TripleDesEncrypt(b, []byte(ctx.DesKey))
	if err != nil {
		return
	}
	encodedData := base64.StdEncoding.EncodeToString(data)
	parms := fmt.Sprintf(`data=%s&mess=%d&timestamp=%d&key=%s`, string(encodedData), randInt, now, ctx.Appkey)
	if hashStr, err = ctx.Signer.Sign(parms); err != nil {
		return
	}
	requestId = randomString(10)
	params = make(url.Values)
	params.Add("data", string(encodedData))
	params.Add("mess", strconv.Itoa(randInt))
	params.Add("timestamp", strconv.FormatInt(now, 10))
	params.Add("sign", hashStr)
	params.Add("sign_type", "rsa")
	return
}

func doRequest(ctx *Context, req *http.Request) ([]byte, error) {
	req = req.WithContext(ctx)
	var resp, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s %s %d", req.Method, req.URL.String(), resp.StatusCode)
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func DecodeWithError(responseBytes []byte, obj interface{}, apiName string) error {
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
	id := commonResponse.FieldByName("RequestId")
	if !code.IsValid() || !msg.IsValid() {
		return fmt.Errorf("code or message is invalid")
	}
	if code.String() != "0000" {
		return &Error{StatusCode(code.String()), msg.String(), id.String(), apiName}
	}
	return nil
}
