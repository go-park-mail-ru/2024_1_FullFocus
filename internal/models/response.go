package models

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	MsgRus string `json:"msgRus"`
}
