package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"

	"github.com/ditdittdittt/backend-sitpi/config"
	"github.com/ditdittdittt/backend-sitpi/domain/response"
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

func IsAuthorized(res http.ResponseWriter, req *http.Request, response *response.Response) (err error) {
	if req.Header["Token"] != nil {
		token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return config.JwtSecret, nil
		})

		if err != nil {
			response.Code = "XX"
			response.Desc = "Login failed"
			response.Data = err.Error()
			return err
		}

		if !token.Valid {
			response.Code = "XX"
			response.Desc = "Token invalid"
			return errors.New("token invalid")
		}
	} else {
		response.Code = "XX"
		response.Desc = "Token not sent"
		return errors.New("token not send")
	}
	return nil
}

func parseJwt(tokenString string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSecret, nil
	})

	if err != nil {
		return claims, err
	}

	if !tkn.Valid {
		return claims, errors.New("token invalid")
	}

	return claims, nil
}
