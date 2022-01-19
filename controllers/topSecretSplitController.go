package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maurodanieldev/quasar-oper-fire/controllers/request"
	"github.com/maurodanieldev/quasar-oper-fire/controllers/responses"
	"github.com/maurodanieldev/quasar-oper-fire/interfaces"
	"github.com/maurodanieldev/quasar-oper-fire/util"
)

type TopSecretSplitHandler struct {
	trilaterationService interfaces.ITrilaterationService
	satelliteService     interfaces.ISatelliteService
	messageService       interfaces.IMessagesService
}

func NewTopSecretSplitHandler(trilaterationService interfaces.ITrilaterationService, satelliteService interfaces.ISatelliteService, messageService interfaces.IMessagesService) *TopSecretSplitHandler {
	return &TopSecretSplitHandler{
		trilaterationService: trilaterationService,
		satelliteService:     satelliteService,
		messageService:       messageService,
	}
}

// GetMessage @Title GetMessage
// @Description get secret message and coordinates
// @Accept json
// @Success 200 {object} responses.MessageResponse
// @Failure 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /topsecret_split [get]
func (h TopSecretSplitHandler) GetMessage(c echo.Context) error {
	satellites, _ := h.satelliteService.All()
	m1 := satellites[0].Message
	m2 := satellites[1].Message
	m3 := satellites[2].Message
	message := h.messageService.GetMessage(m1, m2, m3)

	if message == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "insufficient information ")
	}
	mapSatellites := util.GetMapDistancesFromSatellitesRecord(satellites)

	res := &responses.MessageResponse{
		Message: message,
	}
	x := 0.0
	y := 0.0
	if len(mapSatellites) == 3 {
		x, y = h.trilaterationService.GetLocation(mapSatellites[util.Satellites[0]], mapSatellites[util.Satellites[1]], mapSatellites[util.Satellites[2]])
		res.Position.X = x
		res.Position.Y = y
	}
	return c.JSON(http.StatusOK, res)
}

// PutMessageOnSatellite @Title PutMessageOnSatellite
// @Description put message and distance on satellite
// @Accept json
// @Param name path string true "name"
// @Param satellite body request.Satellite true "satellite info"
// @Success 200 {object} request.Satellite
// @Failure 400 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /topsecret_split/{name} [post]
func (h TopSecretSplitHandler) PutMessageOnSatellite(c echo.Context) error {
	name := c.Param("name")
	s := new(request.Satellite)
	err := c.Bind(s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	entity, err2 := h.satelliteService.One(name)
	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	s.Name = entity.Name
	s.X = entity.X
	s.Y = entity.Y
	err = h.satelliteService.Save(s)
	if err != nil {
		if err == request.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, s)
}
