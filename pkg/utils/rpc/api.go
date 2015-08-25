package rpc

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/wandoulabs/codis/pkg/utils/atomic2"
	"github.com/wandoulabs/codis/pkg/utils/errors"
	"github.com/wandoulabs/codis/pkg/utils/log"
	"github.com/wandoulabs/codis/pkg/utils/trace"
)

const (
	MethodGet = "GET"
	MethodPut = "PUT"
)

var client *http.Client

func init() {
	var dials atomic2.Int64
	tr := &http.Transport{}
	tr.Dial = func(network, addr string) (net.Conn, error) {
		c, err := net.DialTimeout(network, addr, time.Second*10)
		if err == nil {
			log.Debugf("rpc: dial new connection to [%d] %s - %s",
				dials.Incr()-1, network, addr)
		}
		return c, err
	}
	tr.MaxIdleConnsPerHost = 4
	client = &http.Client{
		Transport: tr,
		Timeout:   time.Minute,
	}
}

type rpcError struct {
	Cause string
	Stack trace.Stack
}

func responseBodyAsBytes(rsp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return b, nil
}

func responseBodyAsString(rsp *http.Response) (string, error) {
	b, err := responseBodyAsBytes(rsp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func responseBodyAsError(rsp *http.Response) (error, error) {
	b, err := responseBodyAsBytes(rsp)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return nil, errors.Errorf("[Remote Error] Unknown Error")
	}
	e := &rpcError{}
	if err := json.Unmarshal(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	return &errors.TracedError{
		Cause: errors.New("[Remote Error] " + e.Cause),
		Stack: e.Stack,
	}, nil
}

func ApiGet(url string) ([]byte, error) {
	return apiRequest(MethodGet, url)
}

func ApiPut(url string) ([]byte, error) {
	return apiRequest(MethodPut, url)
}

func apiRequest(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.Trace(err)
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer func() {
		io.Copy(ioutil.Discard, rsp.Body)
		rsp.Body.Close()
	}()

	switch rsp.StatusCode {
	case 200:
		return responseBodyAsBytes(rsp)
	case 500:
		e, err := responseBodyAsError(rsp)
		if err != nil {
			return nil, err
		}
		return nil, e
	default:
		return nil, errors.Errorf("[%d] %s", rsp.StatusCode, http.StatusText(rsp.StatusCode))
	}
}

func ApiGetAsJson(url string, reply interface{}) error {
	return apiRequestAsJson(MethodGet, url, reply)
}

func ApiPutAsJson(url string, reply interface{}) error {
	return apiRequestAsJson(MethodPut, url, reply)
}

func apiRequestAsJson(method string, url string, reply interface{}) error {
	b, err := apiRequest(method, url)
	if err != nil {
		return err
	}
	if reply == nil {
		return nil
	}
	if err := json.Unmarshal(b, reply); err != nil {
		return errors.Trace(err)
	} else {
		return nil
	}
}

func ApiResponseString(s string) (int, string) {
	return 200, s
}

func ApiResponseError(err error) (int, string) {
	if err == nil {
		return 500, ""
	}
	e := &rpcError{
		Cause: err.Error(),
		Stack: errors.Stack(err),
	}
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return 500, ""
	} else {
		return 500, string(b)
	}
}

func ApiResponseJson(v interface{}) (int, string) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ApiResponseError(errors.Trace(err))
	} else {
		return ApiResponseString(string(b))
	}
}
