package services

import (
	"technical_test/domain"
	"technical_test/postgres_connection"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

var sites_service = &SitesService{}

type SitesService struct {
}

func GetSitesServiceInstance() *SitesService {
	if sites_service == nil {
		sites_service = &SitesService{}
	}
	return sites_service
}

func (ss *SitesService) GetSites(connectionID string, limit int, offset int, order string, sort string) ([]domain.Site, error) {
	pgConn := postgres_connection.Get()
	var sites = make([]domain.Site, 0)
	if order != "desc" && order != "asc" {
		order = "asc"
	}
	conn, err := GetAuthServiceInstance().GetConnectionByID(connectionID)
	if err != nil {
		return sites, err
	}
	if conn == nil {
		return sites, &domain.ErrNotLoggedIn
	}

	res := pgConn.Where("account_id = ?", conn.AccountId).
		Limit(limit).
		Offset(offset).
		Order(sort + " " + order).
		Find(&sites)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return sites, nil
		}
		return nil, &domain.ErrInternalServerError
	}
	return sites, nil
}

func (ss *SitesService) CreateSite(connectionID string, site *domain.Site) (*domain.Site, error) {
	pgConn := postgres_connection.Get()
	conn, err := GetAuthServiceInstance().GetConnectionByID(connectionID)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, &domain.ErrNotLoggedIn
	}
	if site.SiteName == "" {
		return nil, &domain.CustomError{
			Code:    400,
			Err:     "ERR_WRONG_SITE_NAME",
			Message: "Site name can't be empty",
		}
	}
	site.AccountId = conn.AccountId
	site.CreateDate = time.Now()
	site.LastUpdated = time.Now()
	site.RecordStatus = domain.RecordStatusActive
	site.Active = domain.RecordStatusActive
	res := pgConn.Save(&site)
	if res.Error != nil {
		return nil, &domain.ErrInternalServerError
	}
	return site, nil
}

func (ss *SitesService) GetSiteByID(connectionID string, siteID int64) (*domain.Site, error) {
	pgConn := postgres_connection.Get()
	conn, err := GetAuthServiceInstance().GetConnectionByID(connectionID)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, &domain.ErrNotLoggedIn
	}
	site := &domain.Site{SiteID: siteID}
	res := pgConn.Where("account_id = ?", conn.AccountId).First(&site)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, &domain.ErrSiteNotFound
		}
		return nil, &domain.ErrInternalServerError
	}
	return site, nil
}

func (ss *SitesService) DeleteSite(connectionID string, siteID int64) (*domain.Site, error) {
	pgConn := postgres_connection.Get()
	var site *domain.Site
	conn, err := GetAuthServiceInstance().GetConnectionByID(connectionID)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, &domain.ErrNotLoggedIn
	}
	err = pgConn.Transaction(func(tx *gorm.DB) error {
		site = &domain.Site{SiteID: siteID, AccountId: conn.AccountId}
		res := tx.Where("account_id = ?", conn.AccountId).First(&site)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return &domain.ErrSiteNotFound
			}
			return &domain.ErrInternalServerError
		}
		site.RecordStatus = domain.RecordStatusDeleted
		site.Active = domain.RecordStatusDeleted
		site.LastUpdated = time.Now()
		res = pgConn.Save(&site)
		if res.Error != nil {
			log.Error(res.Error.Error())
			return &domain.ErrInternalServerError
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return site, nil
}

func (ss *SitesService) UpdateSite(connectionID string, site *domain.Site) (*domain.Site, error) {
	pgConn := postgres_connection.Get()
	conn, err := GetAuthServiceInstance().GetConnectionByID(connectionID)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, &domain.ErrNotLoggedIn
	}
	if site.SiteName == "" {
		return nil, &domain.CustomError{
			Code:    400,
			Err:     "ERR_WRONG_SITE_NAME",
			Message: "Site name can't be empty",
		}
	}
	var siteToUpdate domain.Site
	err = pgConn.Transaction(func(tx *gorm.DB) error {
		siteToUpdate = domain.Site{AccountId: conn.AccountId, SiteID: site.SiteID}
		res := tx.First(&siteToUpdate)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return &domain.ErrSiteNotFound
			}
			log.Error(res.Error.Error())
			return &domain.ErrInternalServerError
		}
		siteToUpdate.LastUpdated = time.Now()
		siteToUpdate.SiteName = site.SiteName
		siteToUpdate.Address = site.Address
		siteToUpdate.Des = site.Des
		siteToUpdate.Description = site.Description
		siteToUpdate.Email = site.Email
		siteToUpdate.Geolocation = site.Geolocation
		siteToUpdate.Logo = site.Logo
		siteToUpdate.OperateBy = site.OperateBy
		siteToUpdate.Phone = site.Phone
		res = tx.Save(&siteToUpdate)
		if res.Error != nil {
			log.Error(res.Error.Error())
			return &domain.ErrInternalServerError
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &siteToUpdate, nil
}
