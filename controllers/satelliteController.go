package controllers

import (
	"net/http"

	"github.com/asdine/storm"
	"github.com/labstack/echo/v4"
	"github.com/maurodanieldev/quasar-oper-fire/controllers/request"
	"github.com/maurodanieldev/quasar-oper-fire/interfaces"
)

type SatelliteHandler struct {
	satelliteService interfaces.ISatelliteService
}

func NewSatelliteHandler(service interfaces.ISatelliteService) *SatelliteHandler {
	return &SatelliteHandler{
		satelliteService: service,
	}
}

// SatellitesPostOne @Title SatellitesPostOne
// @Description Create a satellite
// @Accept json
// @Param satellite body request.Satellite true "satellite info"
// @Success 201 "Successfully create Satellite"
// @Failure 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /satellites [post]
func (h SatelliteHandler) SatellitesPostOne(c echo.Context) error {
	s := new(request.Satellite)
	err := c.Bind(s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	err = h.satelliteService.Save(s)
	if err != nil {
		if err == request.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	c.Response().Writer.Header().Set("Location", "/satellites/"+s.Name)
	return c.NoContent(http.StatusCreated)
}

// SatellitesDeleteOne @Title SatellitesDeleteOne
// @Description Delete satellite by name
// @Accept json
// @Param name path string true "name"
// @Success 200 {object} request.Satellite
// @Failure 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /satellites/{name} [delete]
func (h SatelliteHandler) SatellitesDeleteOne(c echo.Context) error {
	name := c.Param("name")
	err := h.satelliteService.Delete(name)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

// SatellitesPatchOne @Title SatellitesPatchOne
// @Description update a property satellite by name
// @Accept json
// @Param name path string true "name"
// @Param satellite body request.Satellite true "satellite info"
// @Success 200 {object} request.Satellite
// @Failure 400 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /satellites/{name} [patch]
func (h SatelliteHandler) SatellitesPatchOne(c echo.Context) error {
	name := c.Param("name")
	s, err := h.satelliteService.One(name)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = c.Bind(s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	s.Name = name
	err = h.satelliteService.Save(s)
	if err != nil {
		if err == request.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, s)
}

// SatellitesPutOne @Title SatellitesPutOne
// @Description update a satellite by name
// @Accept json
// @Param name formData string true "name"
// @Param satellite body request.Satellite true "satellite info"
// @Success 200 {object} request.Satellite
// @Failure 400 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /satellites/{name} [put]
func (h SatelliteHandler) SatellitesPutOne(c echo.Context) error {
	s := new(request.Satellite)
	err := c.Bind(s)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	s.Name = c.Param("name")
	err = h.satelliteService.Save(s)
	if err != nil {
		if err == request.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, s)
}

// SatellitesGetOne @Title SatellitesGetOne
// @Description get a satellite by name
// @Param name path string true "name"
// @Success 200 {object} request.Satellite
// @Failure 404 {object} responses.HTTPError
// @Failure 500 {object} responses.HTTPError
// @Router /satellites/{name} [get]
func (h SatelliteHandler) SatellitesGetOne(c echo.Context) error {
	name := c.Param("name")
	s, err := h.satelliteService.One(name)
	if err != nil {
		if err == storm.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, s)
}

// SatellitesGetAll @Title SatellitesGetAll
// @Description get all satellites
// @Success 200 {array} request.Satellite
// @Router /satellites [get]
// @Failure 500 {object} responses.HTTPError
func (h SatelliteHandler) SatellitesGetAll(c echo.Context) error {
	satellites, err := h.satelliteService.All()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, satellites)
}
