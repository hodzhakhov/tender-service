package orgresponsible

type OrganizationResponsible struct {
	ID             string `json:"id"`
	OrganizationId string `json:"organization_id"`
	UserId         string `json:"user_id"`
}
