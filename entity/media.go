package entity

import (
	"mime/multipart"
)

type Media struct {
	File   *multipart.File
	Header *multipart.FileHeader
}
