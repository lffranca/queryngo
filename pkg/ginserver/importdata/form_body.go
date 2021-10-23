package importdata

import "mime/multipart"

type formBody struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
