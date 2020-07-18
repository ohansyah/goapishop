package register

// FormRegis as form data register
type FormRegis struct {
	Name    string `json:name`
	Address string `json:address`
	RoleID  string `json:role_id`
	Phone   string `json:phone`
	Email   string `json:email`
	Status  string `json:status`
}
