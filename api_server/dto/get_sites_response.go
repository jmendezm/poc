package dto

type GetSitesResponse struct {
	Sites []SiteResponse `json:"sites"`
	Page  Pagination     `json:"page"`
}
