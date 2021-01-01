package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/ditdittdittt/backend-sitpi/domain/response"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00"
)

func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

func ReadRequest(req *http.Request, response *response.Response) ([]byte, error) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.Data = err.Error()
	} else {
		// And now set a new body, which will simulate the same data we read:
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	return body, err
}

func SetResponse(res http.ResponseWriter, req *http.Request, response *response.Response) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_ = json.NewEncoder(res).Encode(response)

}

func ValidateRequest(input interface{}, response *response.Response) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		response.Code = "VE"
		response.Desc = "Validation Error"
		response.Data = err.Error()
	}
	return err
}
