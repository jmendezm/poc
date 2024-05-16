package domain

import "time"

type Site struct {
	AccountId         int64
	SiteID            int64 `gorm:"<-:create;primaryKey"`
	SiteName          string
	CreateDate        time.Time `gorm:"<-:create;autoCreateTime"`
	LastUpdated       time.Time `gorm:"<-:create;autoUpdateTime:milli"`
	RecordStatus      RecordStatus
	Active            RecordStatus
	Des               string
	Description       string
	OperateBy         string
	Logo              string
	RulesDocuments    string
	ServicesAmenities string
	Type              string
	Email             string
	Phone             string
	Address           string
	Website           string
	Geolocation       string
}

func getSiteOrdenableFields() []string {
	return []string{
		"account_id", "site_id", "site_name", "create_date", "last_updated", "record_status",
		"active", "des", "description", "operated_by", "logo", "rules_documents", "services_amenities",
		"type", "email", "phone", "address", "website", "geolocation",
	}
}

func IsSiteOrdenableField(field string) bool {
	for _, f := range getSiteOrdenableFields() {
		if field == f {
			return true
		}
	}
	return false
}
