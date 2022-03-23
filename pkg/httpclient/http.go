package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/nanguohuakai/spiderman/dto"
	"io/ioutil"
	"net"
	"net/http"
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
		Timeout:   1 * time.Second, //1s 超时
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

func Get(uri string, conf dto.AppConf, s interface{}) error {
	e := checkAppConf(conf)
	if  e != nil{
		return e
	}
	url := conf.BaseUri + uri
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

func Post(uri string, conf dto.AppConf, params interface{}, rt interface{}) (err error) {
	e := checkAppConf(conf)
	if  e != nil{
		return e
	}
	url := conf.BaseUri + uri
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

func getAuthorization(conf dto.AppConf) string {
	return conf.ServiceName + " " + conf.Token
}

func checkAppConf(conf dto.AppConf) error {
	if conf.BaseUri =="" || conf.Token == "" || conf.ServiceName == "" {
		return errors.New("AppConf 配置信息缺失")
	}
	return nil
}