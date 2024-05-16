package dto

type UpdateSiteRequest struct {
	SiteName          string `json:"site_name"`
	Des               string `json:"des"`
	Description       string `json:"description"`
	OperateBy         string `json:"operated_by"`
	Logo              string `json:"logo"`
	RulesDocuments    string `json:"rules_document"`
	ServicesAmenities string `json:"services_amenities"`
	Type              string `json:"type"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	Website           string `json:"web_site"`
	Geolocation       string `json:"geolocation"`
}
