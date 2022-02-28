package model

import "mime/multipart"

type BookForm struct {
	book *Book                 `form:"book" binding:"required"`
	zip  *multipart.FileHeader `form:"zipfile" binding:"required"`
}
