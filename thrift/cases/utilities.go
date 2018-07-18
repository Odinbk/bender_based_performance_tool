package cases

import (
	E "git.appannie.org/appannie/performance/erik"
	"github.com/ugorji/go/codec"
	"time"
)

func GenerateTestCase(query interface{}, data map[string]interface{}) E.ITestCase {
	ttc := E.NewThriftTestCase()
	ttc.Query = msgpackQuery(query)
	ttc.Interval = data["interval"].(float64)
	ttc.Repeat = data["repeat"].(int)
	ttc.ServerAddress = data["address"].(string)
	ttc.Auth = data["auth"].(string)
	ttc.Timeout = data["timeout"].(time.Duration)
	return ttc
}

func msgpackQuery(query interface{}) []byte {
	var (
		mh codec.MsgpackHandle
		h  = &mh
		b []byte = make([]byte, 0)
	)

	ec := codec.NewEncoderBytes(&b, h)
	if err := ec.Encode(query); err != nil {
		panic(err)
	}
	return b
}
