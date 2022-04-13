package pck

import (
	"encoding/json"
	"net/http"
)

type ResponseForm struct {
	Success bool
	Message string
	Data    []interface{}
}

type ResponseForm2 struct {
	Success bool
	Message string
	Data    interface{}
}

func APIResponseMoreData(w http.ResponseWriter, obj []interface{}) {
	res := ResponseForm{
		Success: false,
		Data:    obj,
	}
	jsonByte, err := json.Marshal(res)
	if err != nil {
		internalServerError(w)

	}
	w.Write(jsonByte)

}
func APIResponse(w http.ResponseWriter, obj interface{}) {
	res := ResponseForm2{
		Success: false,
		Data:    obj,
	}
	jsonByte, err := json.Marshal(res)
	if err != nil {
		internalServerError(w)
	}
	w.Write(jsonByte)

}
func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal api error"))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
func ParamsInvalid(w http.ResponseWriter, r *http.Request, params string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Truyền thiếu params : ..."))
}
func Authenticated(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("you cannot access , only admin access"))
}
