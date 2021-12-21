package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
response := BaseResponse{}
response.Meta.Status = http.StatusOK
response.Meta.Message = "success"
response.Data = data
return c.JSON(http.StatusOK,response)
}

func ErrorResponse(c echo.Context,status int,err string,errs error)error{
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Messages = []string{errs.Error()}
	response.Meta.Message = err

	return c.JSON(status,response)
}