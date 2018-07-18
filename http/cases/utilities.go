package cases

import (
	E "git.appannie.org/appannie/performance/erik"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func GenerateTestCase(query string, data map[string]interface{}) E.ITestCase {
	query = fmt.Sprintf(query, data["user_id"], data["product_ids"])
	url := data["address"].(string)
	htc := E.NewHttpTestCase()
	htc.URL = url
	htc.BodyValidator = BodyValidator
	htc.Query = []byte(query)
	htc.Interval = data["interval"].(float64)
	htc.Repeat = data["repeat"].(int)
	htc.Method = data["method"].(string)
	return htc
}

type response struct {
	Data      interface{}   `json:"data"`
	DebugInfo string        `json:"debug_info"`
	Errors    []interface{} `json:"errors"`
}

func BodyValidator(request interface{}, resp *http.Response) error {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r := new(response)
	if err:= json.Unmarshal(data, r); err != nil {
		return err
	} else {
		if len(r.DebugInfo) > 0 {
			return errors.New(fmt.Sprintf("debug Info: %s.\n", r.DebugInfo))
		} else {
			return nil
		}
	}
}
