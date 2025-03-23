package entity

type User struct {
	UserId               int64  `json:"-"`
	IdentityNumber       string `json:"nik" binding:"required" validate:"required"`
	FullName             string `json:"full_name" binding:"required" validate:"required"`
	LegalName            string `json:"legal_name" binding:"required" validate:"required"`
	PlaceOfBirth         string `json:"place_of_birth" binding:"required" validate:"required"`
	DateOfBirth          string `json:"date_of_birth" binding:"required" validate:"required"`
	Salary               int64  `json:"salary" binding:"required" validate:"required"`
	IdentityCardPhotoUrl string `json:"identity_card_photo_url"`
	SelfiePhotoUrl       string `json:"selfie_photo_url"`
}

type UserMedia struct {
	IdentityCardPhoto Media
	SelfiePhoto       Media
}
