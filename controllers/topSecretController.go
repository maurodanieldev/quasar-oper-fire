package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/maurodanieldev/quasar-oper-fire/controllers/request"
	"github.com/maurodanieldev/quasar-oper-fire/controllers/responses"
	"github.com/maurodanieldev/quasar-oper-fire/interfaces"
	"github.com/maurodanieldev/quasar-oper-fire/util"
)

type TopSecretHandler struct {
	messageService       interfaces.IMessagesService
	trilaterationService interfaces.ITrilaterationService
}

func NewTopSecretHandler(messageService interfaces.IMessagesService, trilaterationService interfaces.ITrilaterationService) *TopSecretHandler {
	return &TopSecretHandler{
		messageService:       messageService,
		trilaterationService: trilaterationService,
	}
}

// GetMessages @Title GetMessages
// @Description get secret message and coordinates
// @Accept json
// @Param satellite body request.SatellitesRequest true "satellite info"
// @Success 200 {object} responses.MessageResponse
// @Failure 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /topsecret [post]
func (h TopSecretHandler) GetMessages(c echo.Context) error {
	req := new(request.SatellitesRequest)
	if err := c.Bind(&req); err != nil {
		log.Errorf("error parsing request: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(req.Satellites) != 3 {
		return echo.NewHTTPError(http.StatusBadRequest, "need message from 3 satellites")
	}
	m1 := req.Satellites[0].Message
	m2 := req.Satellites[1].Message
	m3 := req.Satellites[2].Message
	message := h.messageService.GetMessage(m1, m2, m3)
	if message == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	res := &responses.MessageResponse{
		Message: message,
	}
	x := 0.0
	y := 0.0
	mapSatellites := util.GetMapDistances(req)
	if len(mapSatellites) == 3 {
		x, y = h.trilaterationService.GetLocation(mapSatellites[util.Satellites[0]], mapSatellites[util.Satellites[1]], mapSatellites[util.Satellites[2]])
		res.Position.X = x
		res.Position.Y = y
	}
	return c.JSON(http.StatusOK, res)
}
