package resources


type UserResource struct {
	Email string `deepcopier:"field:email" json:"email"`
	AccessToken string `json:"access_token" deepcopier:"field:AccessToken"`
	RefreshToken string `json:"refresh_token" deepcopier:"field:refresh_token"`
	AtExpiredAt int64 `json:"at_expired_at" deepcopier:"field:at_expired_at"`
	RfExpiredAt int64	`json:"rf_expired_at" deepcopier:"field:rf_expired_at"`
	AccessUid string	`json:"access_uid" deepcopier:"field:access_uid"`
	RefreshUid string	`json:"refresh_uid" deepcopier:"field:refresh_uid"`
}