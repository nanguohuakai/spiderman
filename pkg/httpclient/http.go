package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/nanguohuakai/spiderman/dto"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var httpClient *http.Client

func init() {
	// init the httpClient for later reuse
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 2
	retryClient.RetryWaitMin = 10 * time.Millisecond
	retryClient.RetryWaitMax = 500 * time.Millisecond
	retryClient.HTTPClient = &http.Client{
		Timeout:   5 * time.Second, //1s 超时
		Transport: t,
	}
	retryClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {

		if err != nil {
			var e *net.DNSError
			if errors.As(err, &e) {
				return true, err
			}
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}

	retryClient.ErrorHandler = retryablehttp.PassthroughErrorHandler

	httpClient = retryClient.StandardClient()
}

func Get(url string, conf dto.AppConf, s interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", getAuthorization(conf))
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		// close body stream to avoid memory leak
		if innerErr := resp.Body.Close(); innerErr != nil {
			err = innerErr
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &s)
	}

	return err
}

func Post(url string, conf dto.AppConf, params interface{}, rt interface{}) (err error) {
	postBody, _ := json.Marshal(params)
	responseBody := strings.NewReader(string(postBody))
	// create a request object
	req, _ := http.NewRequest(
		"POST",
		url,
		responseBody,
	)

	// add a request header
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", getAuthorization(conf))

	// send an HTTP using `req` object
	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}
	defer func() {
		// close body stream to avoid memory leak
		if innerErr := resp.Body.Close(); innerErr != nil {
			err = innerErr
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &rt)
	}
	return err
}

// get header authorization
func getAuthorization(conf dto.AppConf) string {
	return conf.ServiceName + " " + conf.Token
}

// check response status
func checkResponse(resp *http.Response) error {
	var err error
	if resp.StatusCode != 200 {
		if resp.StatusCode == 500 {
			err = errors.New("服务内部错误")
		} else if resp.StatusCode == 401 {
			err = errors.New("没有api权限")
		} else {
			err = errors.New("服务内部错误")
		}
	}
	return err
}

//HttpBuildQuery 拼接url ;uri host; path 请求路径; params 请求参数
//return url
func HttpBuildQuery(uri string, path string, params interface{}) string {
	var res map[string]interface{}
	//解析param struct
	j, _ := json.Marshal(params)
	_ = json.Unmarshal(j, &res)

	baseUri := uri + path
	//组合get参数拼接
	u, _ := url.Parse(baseUri)
	q := u.Query()
	for k, v := range res {
		q.Add(k, fmt.Sprintf("%v", v))
	}

	u.RawQuery = q.Encode()

	//get 请求
	return u.String()
}
