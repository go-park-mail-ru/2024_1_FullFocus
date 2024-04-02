package models

type SuccessResponse struct {
	Status int         `json:"Status"`
	Data   interface{} `json:"Data"`
}

type ErrResponse struct {
	Status int    `json:"Status"`
	Msg    string `json:"Msg"`
	MsgRus string `json:"MsgRus"`
}
