package handlers

import "mime/multipart"

type ReleaseDTO struct {
	ID   uint    `json:"id"`
	Site SiteDTO `json:"site"`
	Tag  string  `json:"tag"`
}

type CreateReleaseReq struct {
	Tag  string                `json:"tag" binding:"required"`
	File *multipart.FileHeader `json:"file" binding:"required"`
}

type ReleaseIdReq struct {
	ID uint `json:"id" binding:"required"`
}
