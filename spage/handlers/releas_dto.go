package handlers

import (
	"github.com/LiteyukiStudio/spage/spage/models"
	"mime/multipart"
)

type ReleaseDTO struct {
	ID   uint        `json:"id"`
	Site SiteDTO     `json:"site"`
	Tag  string      `json:"tag"`
	File models.File `json:"file"`
}

type CreateReleaseReq struct {
	Tag  string                `json:"tag" binding:"required"`
	File *multipart.FileHeader `json:"file" binding:"required"`
}

type ReleaseIdReq struct {
	ID uint `json:"id" binding:"required"`
}
