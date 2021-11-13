package presenter

import "mime/multipart"

type ImportDataBody struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
