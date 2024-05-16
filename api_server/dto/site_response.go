package dto

import (
	"technical_test/domain"

	"github.com/nleeper/goment"
)

type SiteResponse struct {
	AccountId         int64  `json:"account_id"`
	SiteID            int64  `json:"site_id"`
	SiteName          string `json:"site_name"`
	CreateDate        string `json:"create_date"`
	LastUpdated       string `json:"last_updated"`
	RecordStatus      string `json:"record_status"`
	Active            string `json:"active"`
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

func (sr *SiteResponse) FromDomain(s *domain.Site) {
	created, _ := goment.New(s.CreateDate)
	updated, _ := goment.New(s.LastUpdated)
	sr.AccountId = s.AccountId
	sr.SiteID = s.SiteID
	sr.SiteName = s.SiteName
	sr.CreateDate = created.Format("L LTS Z")
	sr.LastUpdated = updated.Format("L LTS Z")
	sr.RecordStatus = string(s.RecordStatus)
	sr.Active = string(s.Active)
	sr.Des = s.Des
	sr.Description = s.Description
	sr.OperateBy = s.OperateBy
	sr.Logo = s.Logo
	sr.RulesDocuments = s.RulesDocuments
	sr.ServicesAmenities = s.ServicesAmenities
	sr.Type = s.Type
	sr.Email = s.Email
	sr.Phone = s.Phone
	sr.Address = s.Address
	sr.Website = s.Website
	sr.Geolocation = s.Geolocation
}
