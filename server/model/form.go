package model

import "mime/multipart"

type AvatarForm struct {
	X      string                `form:"x" binding:"required"`
	Y      string                `form:"y" binding:"required"`
	Width  string                `form:"width" binding:"required"`
	Height string                `form:"height" binding:"required"`
	Avatar *multipart.FileHeader `form:"image-file" binding:"required"`
}

type CoverForm struct {
	X      string                `form:"x" binding:"required"`
	Y      string                `form:"y" binding:"required"`
	Width  string                `form:"width" binding:"required"`
	Height string                `form:"height" binding:"required"`
	Cover  *multipart.FileHeader `form:"image-file" binding:"required"`
	BookId int                   `form:"book_id" binding:"required"`
}
