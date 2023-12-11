package rest

import (
	"encoding/json"
	"testing"
)

func TestResult(t *testing.T) {
	result := Success()
	JsonResult, _ := json.Marshal(result)
	t.Log(string(JsonResult))

	result = SuccessMessage("abc")
	JsonResult, _ = json.Marshal(result)
	t.Log(string(JsonResult))

	result = Error(400)
	JsonResult, _ = json.Marshal(result)
	t.Log(string(JsonResult))

	d := make(map[string]string)
	d["a"] = "b"
	d["c"] = "d"
	result = Data(d)
	JsonResult, _ = json.Marshal(result)
	t.Log(string(JsonResult))

	result = Data(nil)
	JsonResult, _ = json.Marshal(result)
	t.Log(string(JsonResult))
}
