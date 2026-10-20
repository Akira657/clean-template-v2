package domain

import "context"

type Profile struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID int64) (*Profile, error)
}
