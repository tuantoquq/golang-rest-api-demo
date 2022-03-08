package model

type Response struct {
	Code    int         `json:"code" binding:"required"`
	Message string      `json:"message" binding:"required"`
	Data    interface{} `json:"data"`
}
