package helper

import (
	"bytes"
	"encoding/json"
	"github.com/ditdittdittt/backend-sitpi/domain/response"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
)

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
