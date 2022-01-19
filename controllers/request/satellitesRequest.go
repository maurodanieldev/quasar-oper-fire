package request

import "errors"

type SatellitesRequest struct {
	Satellites []SatelliteRequest `json:"satellites" validate:"required"`
}

type SatelliteRequest struct {
	Name     string   `json:"name" validate:"required"`
	Distance float64  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required"`
}

type Satellite struct {
	Name     string   `json:"name" storm:"id"`
	X        float64  `json:"x"`
	Y        float64  `json:"y"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

func (s *Satellite) Validate() error {
	if s.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}

var ErrRecordInvalid = errors.New("record is invalid")
