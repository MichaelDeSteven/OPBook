package model

import "mime/multipart"

type BookForm struct {
	Identify string                `form:"identify" binding:"required"`
	Zip      *multipart.FileHeader `form:"zipfile" binding:"required"`
}
