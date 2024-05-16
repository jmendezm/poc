package sites

import (
	"encoding/json"
	"technical_test/api_server/dto"
	"technical_test/domain"
	"technical_test/services"

	"github.com/gofiber/fiber/v2"
)

func GetSites(ctx *fiber.Ctx) error {
	var (
		req   *dto.GetSitesRequest
		err   error
		sites []domain.Site
	)
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return ctx.Status(400).JSON(&dto.ErrorResponse{ErrorMessage: "Bad request"})
	}
	connID := ctx.Get("connection_id")
	if req.Page.Order != "desc" && req.Page.Order != "asc" {
		req.Page.Order = "asc"
	}
	if req.Page.Sort == "" {
		req.Page.Sort = "site_name"
	}
	if !domain.IsSiteOrdenableField(req.Page.Sort) {
		return ctx.Status(400).JSON(&dto.ErrorResponse{ErrorMessage: "not valid field for sort"})
	}
	sites, err = services.GetSitesServiceInstance().
		GetSites(connID, req.Page.Limit, req.Page.Offset, req.Page.Order, req.Page.Sort)
	if err != nil {
		ce := HandleError(err)
		resp := dto.ErrorResponse{ErrorMessage: ce.Message}
		return ctx.Status(ce.Code).JSON(resp)
	}
	resp := make([]dto.SiteResponse, 0)
	var temp dto.SiteResponse
	for _, s := range sites {
		temp = dto.SiteResponse{}
		temp.FromDomain(&s)
		resp = append(resp, temp)
	}
	fullResponse := dto.GetSitesResponse{
		Sites: resp,
		Page:  req.Page,
	}

	return ctx.Status(200).JSON(&fullResponse)
}
