package pck

type HttpResponsesList struct {
	Success bool          `json:"isSuccess"`
	Data    []interface{} `json:"data"`
}

type HttpResponse struct {
	Success bool        `json:"isSuccess"`
	Data    interface{} `json:"data"`
}
type HttpErrorResponse struct {
	Success bool   `json:"isSuccess"`
	Message string `json:"message"`
}
