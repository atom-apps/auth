package dto

type ResetPasswordForm struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type RoleUsersForm struct {
	Users []string `json:"users"`
	Role  string   `json:"role"`
}

type UserForm struct {
	DisplayName string
}
