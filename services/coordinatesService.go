package services

import (
	"github.com/maurodanieldev/quasar-oper-fire/interfaces"
	"github.com/maurodanieldev/quasar-oper-fire/util"
)

type trilaterationService struct {
	satelliteService interfaces.ISatelliteService
}

func NewTrilaterationService(satelliteService interfaces.ISatelliteService) interfaces.ITrilaterationService {
	return &trilaterationService{
		satelliteService: satelliteService,
	}
}

func (s *trilaterationService) GetLocation(distances ...float64) (x, y float64) {
	satellites, _ := s.satelliteService.All()
	coordinates := util.GetCoordinatesSatellites(satellites)
	c1 := coordinates[util.Satellites[0]]
	c2 := coordinates[util.Satellites[1]]
	c3 := coordinates[util.Satellites[2]]
	return util.Trilateration(c1, c2, c3, distances[0], distances[1], distances[2])
}
