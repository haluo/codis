package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
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

func apiRequestJson(method string, url string, args, reply interface{}) error {
	var body []byte
	if args != nil {
		b, err := json.MarshalIndent(args, "", "    ")
		if err != nil {
			return errors.Trace(err)
		}
		body = b
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return errors.Trace(err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	rsp, err := client.Do(req)
	if err != nil {
		return errors.Trace(err)
	}
	defer func() {
		io.Copy(ioutil.Discard, rsp.Body)
		rsp.Body.Close()
	}()

	switch rsp.StatusCode {
	case 200:
		b, err := responseBodyAsBytes(rsp)
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
	case 1500:
		e, err := responseBodyAsError(rsp)
		if err != nil {
			return err
		} else {
			return e
		}
	default:
		return errors.Errorf("[%d] %s", rsp.StatusCode, http.StatusText(rsp.StatusCode))
	}
}

func ApiGetJson(url string, reply interface{}) error {
	return apiRequestJson(MethodGet, url, nil, reply)
}

func ApiPutJson(url string, args, reply interface{}) error {
	return apiRequestJson(MethodPut, url, args, reply)
}

func ApiResponseError(err error) (int, string) {
	if err == nil {
		return 1500, ""
	}
	e := &rpcError{
		Cause: err.Error(),
		Stack: errors.Stack(err),
	}
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return 1500, ""
	} else {
		return 1500, string(b)
	}
}

func ApiResponseJson(v interface{}) (int, string) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ApiResponseError(errors.Trace(err))
	} else {
		return 200, string(b)
	}
}

func EncodeURL(host string, format string, args ...interface{}) string {
	return "http://" + host + fmt.Sprintf(format, args...)
}
