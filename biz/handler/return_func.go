package handler

import (
	"OceanShipBeidouMS/biz/common"
	"net/http"
)

type returnStruct struct {
	code    int         `json:"code"`
	message string      `json:"message"`
	value   interface{} `json:"value,omitempty"`
}

func returnSuccess(value interface{}, messages ...string) (int, interface{}) {
	res := returnStruct{value: value, code: 0, message: "success"}
	if len(messages) != 0 {
		res.message = ""
		for _, m := range messages {
			res.message += m + "\n"
		}
		return http.StatusOK, res
	}

	return http.StatusOK, res
}

func returnError(code int, messages ...string) (int, interface{}) {
	res := returnStruct{code: code, message: "error"}
	if len(messages) != 0 {
		res.message = ""
		for _, m := range messages {
			res.message += m + "\n"
		}
		return http.StatusInternalServerError, res
	}

	if message, ok := common.ErrMap[code]; ok {
		res.message = message
		return http.StatusInternalServerError, res
	}

	return http.StatusInternalServerError, res
}
