package token

// ResToken as responds token data
type ResToken struct {
	ID           uint   `json:id`
	CreatedAt    string `json:created_at`
	ExpiredAt    string `json:expired_at`
	TokenCode    string `json:token_code`
	RefreshToken string `json:refresh_token`
	DeviceID     string `json:device_id`
	DeviceType   string `json:device_type`
}
