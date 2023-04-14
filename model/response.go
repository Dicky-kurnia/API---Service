package model

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"errors"`
}

type OKResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type GeneralError struct {
	General string `json:"general"`
}

type NotValidImage struct {
	Image string `json:"image"`
}
