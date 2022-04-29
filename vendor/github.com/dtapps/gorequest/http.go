package gorequest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Response 返回内容
type Response struct {
	RequestUrl            string      //【请求】链接
	RequestParams         url.Values  //【请求】参数
	RequestMethod         string      //【请求】方式
	RequestHeader         http.Header //【请求】头部
	ResponseHeader        http.Header //【返回】头部
	ResponseStatus        string      //【返回】状态
	ResponseStatusCode    int         //【返回】状态码
	ResponseBody          []byte      //【返回】内容
	ResponseContentLength int64       //【返回】大小
}

type app struct {
	httpUrl         string   // 请求地址
	httpMethod      string   // 请求方法
	httpHeader      Headers  // 请求头
	httpParams      Params   // 请求参数
	httpParamsMode  string   // 请求参数方式
	responseContent Response // 返回内容
}

var (
	httpParamsModeJson = "json"
	httpParamsModeForm = "form"
)

// NewHttp 实例化
func NewHttp() *app {
	return &app{
		httpHeader: NewHeaders(),
		httpParams: NewParams(),
	}
}

// SetUrl 设置请求地址
func (app *app) SetUrl(url string) {
	app.httpUrl = url
}

// SetMethod 设置请求方式地址
func (app *app) SetMethod(method string) {
	app.httpMethod = method
}

// SetHeader 设置请求头
func (app *app) SetHeader(key, value string) {
	if key == "" {
		panic("url is empty")
	}
	app.httpHeader.Set(key, value)
}

// SetHeaders 批量设置请求头
func (app *app) SetHeaders(headers Headers) {
	for key, value := range headers {
		app.httpHeader.Set(key, value)
	}
}

// SetAuthToken 设置身份验证令牌
func (app *app) SetAuthToken(token string) {
	app.httpHeader.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

// SetUserAgent 设置用户代理，空字符串就随机设置
func (app *app) SetUserAgent(ua string) {
	if ua == "" {
		ua = GetRandomUserAgent()
	}
	app.httpHeader.Set("User-Agent", ua)
}

// SetContentTypeJson 设置JSON格式
func (app *app) SetContentTypeJson() {
	app.httpParamsMode = httpParamsModeJson
	app.httpHeader.Set("Content-Type", "application/json")
}

// SetContentTypeForm 设置FORM格式
func (app *app) SetContentTypeForm() {
	app.httpParamsMode = httpParamsModeForm
	app.httpHeader.Set("Content-Type", "application/x-www-form-urlencoded")
}

// SetParam 设置请求参数
func (app *app) SetParam(key string, value interface{}) {
	app.httpParams.Set(key, value)
}

// SetParams 批量设置请求参数
func (app *app) SetParams(params Params) {
	for key, value := range params {
		app.httpParams.Set(key, value)
	}
}

// Get 发起GET请求
func (app *app) Get() (httpResponse Response, err error) {
	// 设置请求方法
	app.httpMethod = http.MethodGet
	return request(app)
}

// Post 发起POST请求
func (app *app) Post() (httpResponse Response, err error) {
	// 设置请求方法
	app.httpMethod = http.MethodPost
	return request(app)
}

// Request 发起请求
func (app *app) Request() (httpResponse Response, err error) {
	return request(app)
}

// 请求
func request(app *app) (httpResponse Response, err error) {

	// 创建 http 客户端
	client := &http.Client{}

	// 赋值
	httpResponse.RequestUrl = app.httpUrl
	httpResponse.RequestMethod = app.httpMethod

	// 携带 form 参数
	form := url.Values{}
	if app.httpMethod == http.MethodPost && app.httpParamsMode == httpParamsModeForm {
		if len(app.httpParams) > 0 {
			for k, v := range app.httpParams {
				form.Add(k, GetParamsString(v))
			}
			// 赋值
			httpResponse.RequestParams = form
		}
	}

	// 创建请求
	req, err := http.NewRequest(app.httpMethod, app.httpUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("创建请求出错 %s", err))
	}

	// GET 请求携带查询参数
	if app.httpMethod == http.MethodGet {
		if len(app.httpParams) > 0 {
			q := req.URL.Query()
			for k, v := range app.httpParams {
				q.Add(k, GetParamsString(v))
			}
			req.URL.RawQuery = q.Encode()
			// 赋值
			httpResponse.RequestParams = q
		}
	}

	// 设置请求头
	if len(app.httpHeader) > 0 {
		for key, value := range app.httpHeader {
			req.Header.Set(key, value)
		}
		// 赋值
		httpResponse.RequestHeader = req.Header
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}

	// 最后关闭连接
	defer resp.Body.Close()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}

	// 赋值
	httpResponse.ResponseStatus = resp.Status
	httpResponse.ResponseStatusCode = resp.StatusCode
	httpResponse.ResponseHeader = resp.Header
	httpResponse.ResponseBody = body
	httpResponse.ResponseContentLength = resp.ContentLength

	return httpResponse, err
}
