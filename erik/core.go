package erik

import (
	"os"
	"encoding/json"
	"reflect"
	"strings"
	"fmt"
	"github.com/pinterest/bender"
	"github.com/pinterest/bender/hist"
	"time"
	"log"
)

func SyntheticRequests(repeat int, t ITestCase) chan interface{} {
	c := make(chan interface{}, repeat)
	go func() {
		for i := 0; i < repeat; i++ {
			req := t.GenerateRequest()
			c <- req
		}
		close(c)
	}()
	return c
}

func LoadConfig(filename string, ret interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&ret)
	if err != nil {
		return err
	}

	return nil
}

func LoadCases(tg interface{}) []ITestCase {
	var testCases []ITestCase
	var testAffix string = "Test"

	sType := reflect.TypeOf(tg)
	sValue := reflect.ValueOf(tg)
	for i := 0; i < sType.NumMethod(); i++ {
		method := sType.Method(i)
		sv := sValue.Method(i)
		if strings.HasPrefix(method.Name, testAffix) || strings.HasSuffix(method.Name, testAffix) {
			v := sv.Call([]reflect.Value{})
			if len(v) > 0 {
				t := v[0].Interface().(ITestCase)
				t.SetCaseName(method.Name)
				testCases = append(testCases, t)
			}
		}
	}
	return testCases
}

func LaunchBender(name string, interval float64, repeat int, exec bender.RequestExecutor, requests chan interface{}) {
	intervals := bender.ExponentialIntervalGenerator(interval)
	recorder := make(chan interface{}, repeat)

	l := log.New(os.Stdout, "", log.LstdFlags)
	h := hist.NewHistogram(60000, int(time.Millisecond))

	bender.LoadTestThroughput(intervals, requests, exec, recorder)
	bender.Record(recorder, bender.NewLoggingRecorder(l), bender.NewHistogramRecorder(h))
	fmt.Println(name)
	fmt.Println(h)
}
