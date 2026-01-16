package dto

type AddTodoReq struct {
	Title string `json:"title"  binding:"required"`
}

type BaseIdReq struct {
	Id int `json:"id"`
}
