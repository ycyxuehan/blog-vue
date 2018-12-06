package controllers

import (

)

type Result struct{
	ResultCode int `json:"ResultCode"`
	ResultString string `json:"ResultString"`
	ResultData interface{} `json:"ResultData"`
}