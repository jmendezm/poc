package sites

import (
	"encoding/json"
	"strconv"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func UpdateSite(ctx *fiber.Ctx) error {
	var (
		req *dto.UpdateSiteRequest
		err error
	)
	siteIDParam := ctx.Params("site_id")
	siteID, err := strconv.ParseInt(siteIDParam, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{ErrorMessage: "Wrong site ID"})
	}
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return ctx.Status(400).JSON(&dto.ErrorResponse{ErrorMessage: "Bad request"})
	}
	connID := ctx.Get("connection_id")
	site := &domain.Site{
		SiteID:            siteID,
		SiteName:          req.SiteName,
		Des:               req.Des,
		Description:       req.Description,
		OperateBy:         req.OperateBy,
		Logo:              req.Logo,
		RulesDocuments:    req.RulesDocuments,
		ServicesAmenities: req.ServicesAmenities,
		Type:              req.Type,
		Email:             req.Email,
		Phone:             req.Phone,
		Address:           req.Address,
		Website:           req.Website,
		Geolocation:       req.Geolocation,
	}
	site, err = services.GetSitesServiceInstance().UpdateSite(connID, site)
	if err != nil {
		ce := HandleError(err)
		resp := dto.ErrorResponse{ErrorMessage: ce.Message}
		return ctx.Status(ce.Code).JSON(resp)
	}
	resp := dto.SiteResponse{}
	resp.FromDomain(site)

	return ctx.Status(200).JSON(&resp)
}
