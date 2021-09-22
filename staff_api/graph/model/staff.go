package model

import (
	"fmt"
	"staff_api/entity"
	"strconv"
)

type Staff struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Icon      string `json:"icon"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func StaffFromEntity(e *entity.Staff) *Staff {
	return &Staff{
		ID:        fmt.Sprintf("%d", e.ID),
		Name:      e.Name,
		Email:     e.Email,
		Icon:      e.Icon,
		CreatedAt: strconv.FormatInt(e.CreatedAt.Unix(), 10),
		UpdatedAt: strconv.FormatInt(e.UpdatedAt.Unix(), 10),
	}
}