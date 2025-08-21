package helpers

import (
	"encoding/json"
	"log"
)

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"statusCode"`
	ErrorMessage string `json:"msg"`
}

type MsgResponse struct {
	Status         string `json:"status"`
	Title          string `json:"title"`
	SuccessMessage string `json:"msg"`
}

func GetMsg_String(Msg_Title string, Msg_Description string) string {

	var Msg_Res MsgResponse

	Msg_Res.Status = "S"
	Msg_Res.Title = Msg_Title
	Msg_Res.SuccessMessage = Msg_Description

	result, err := json.Marshal(Msg_Res)

	if err != nil {
		log.Println(err)
	}

	return string(result)

}

func GetError_String(Err_Title string, Err_Description string) string {

	var Err_Response ErrorResponse

	Err_Response.Status = "E"
	Err_Response.ErrorCode = Err_Title
	Err_Response.ErrorMessage = Err_Description

	lResult, err := json.Marshal(Err_Response)

	if err != nil {
		log.Println(err)
	}

	return string(lResult)

}
