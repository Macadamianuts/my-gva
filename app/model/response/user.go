package response

import "gva-lbx/app/model"

type UserLogin struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
	*model.User
}
