package sites

import (
	"strconv"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func GetSiteByID(ctx *fiber.Ctx) error {
	var (
		err  error
		site *domain.Site
	)
	connID := ctx.Get("connection_id")
	siteIDParam := ctx.Params("site_id")
	siteID, err := strconv.ParseInt(siteIDParam, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{ErrorMessage: "Wrong site ID"})
	}
	site, err = services.GetSitesServiceInstance().GetSiteByID(connID, siteID)
	if err != nil {
		ce := HandleError(err)
		resp := dto.ErrorResponse{ErrorMessage: ce.Message}
		return ctx.Status(ce.Code).JSON(resp)
	}
	resp := dto.SiteResponse{}
	resp.FromDomain(site)
	return ctx.Status(200).JSON(&resp)
}
