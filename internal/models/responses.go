package models

type SuccessResponse struct {
	Status int
	Data   interface{}
}

type ErrResponse struct {
	Status int
	Msg    string
	MsgRus string
}
