package erik

import (
	bhttp "github.com/pinterest/bender/http"
	bthrift "github.com/pinterest/bender/thrift"
	qs "git.appannie.org/appannie/performance/thrift/query_service"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net/http"
	"bytes"
	"time"
)

type ITestCase interface {
	SetCaseName(string)
	GenerateRequest() interface{}
	Run()
}

type TestCase struct {
	Name          string
	Query         []byte
	Interval      float64
	Repeat        int
	ServerAddress string
}

func (t *TestCase) SetCaseName(name string) {
	t.Name = name
}

type HttpTestCase struct {
	TestCase
	URL           string
	Method        string
	BodyValidator bhttp.ResponseValidator
}

func NewHttpTestCase() *HttpTestCase {
	return &HttpTestCase{}
}

func (t *HttpTestCase) GenerateRequest() interface{} {
	req, err := http.NewRequest(t.Method, t.URL, bytes.NewBuffer(t.Query))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req
}

func (t *HttpTestCase) Run() {
	requests := SyntheticRequests(t.Repeat, t)
	exec := bhttp.CreateExecutor(nil, nil, t.BodyValidator)
	LaunchBender(t.Name, t.Interval, t.Repeat, exec, requests)
}

type ThriftTestCase struct {
	TestCase
	Auth string
	Timeout time.Duration
}

func NewThriftTestCase() *ThriftTestCase {
	return &ThriftTestCase{}
}

func (t *ThriftTestCase) GenerateRequest() interface{} {
	request := qs.NewQueryServiceQueryArgs()
	request.Auth = t.Auth
	request.Q = t.Query
	return *request
}

func (t *ThriftTestCase) Run() {
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	requests := SyntheticRequests(t.Repeat, t)
	exec := bthrift.NewThriftRequestExec(transportFactory, QueryServiceExecutor, t.Timeout, t.ServerAddress)
	LaunchBender(t.Name, t.Interval, t.Repeat, exec, requests)
}

func QueryServiceExecutor(request interface{}, transport thrift.TTransport) (interface{}, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := qs.NewQueryServiceClientFactory(transport, protocolFactory)
	return client.Query(request.(qs.QueryServiceQueryArgs))
}
