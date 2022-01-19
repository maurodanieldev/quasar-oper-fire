package util

import (
	"math"

	"github.com/maurodanieldev/quasar-oper-fire/controllers/request"
)

var Satellites = []string{"Kenobi", "Skywalker", "Sato"}

func GetMapDistances(request *request.SatellitesRequest) map[string]float64 {
	mapSatellites := map[string]float64{}
	for _, satellite := range request.Satellites {
		if SatelliteExist(satellite.Name) {
			mapSatellites[satellite.Name] = satellite.Distance
		}
	}
	return mapSatellites
}
func GetMapDistancesFromSatellitesRecord(satellites []*request.Satellite) map[string]float64 {
	mapSatellites := map[string]float64{}
	for _, satellite := range satellites {
		if SatelliteExist(satellite.Name) {
			mapSatellites[satellite.Name] = satellite.Distance
		}
	}
	return mapSatellites
}

func GetCoordinatesSatellites(satellites []*request.Satellite) map[string]Point {
	mapCoordinates := map[string]Point{}
	for _, satellite := range satellites {
		if SatelliteExist(satellite.Name) {
			mapCoordinates[satellite.Name] = Point{X: satellite.X, Y: satellite.Y}
		}
	}
	return mapCoordinates
}

func SatelliteExist(e string) bool {
	for _, a := range Satellites {
		if a == e {
			return true
		}
	}
	return false
}

func GetPart(message1, message2 []string) []string {
	tmp := make([]string, len(message1))
	for i := range message1 {
		tmp[i] = selectMessage(message1, message2, i)
	}
	return tmp
}

func selectMessage(message1 []string, message2 []string, index int) string {
	if message1[index] == "" {
		return message2[index]
	}
	return message1[index]
}

type Point struct {
	X float64
	Y float64
}

func norm(p Point) float64 {
	return math.Pow(math.Pow(p.X, 2)+math.Pow(p.Y, 2), .5)
}

func Trilateration(point1 Point, point2 Point, point3 Point, r1 float64, r2 float64, r3 float64) (float64, float64) {
	//unit vector in a direction from point1 to point 2
	p2p1Distance := math.Pow(math.Pow(point2.X-point1.X, 2)+math.Pow(point2.Y-point1.Y, 2), 0.5)
	ex := Point{(point2.X - point1.X) / p2p1Distance, (point2.Y - point1.Y) / p2p1Distance}
	aux := Point{point3.X - point1.X, point3.Y - point1.Y}
	//signed magnitude of the X component
	i := ex.X*aux.X + ex.Y*aux.Y
	//the unit vector in the y direction.
	aux2 := Point{point3.X - point1.X - i*ex.X, point3.Y - point1.Y - i*ex.Y}
	ey := Point{aux2.X / norm(aux2), aux2.Y / norm(aux2)}
	//the signed magnitude of the y component
	j := ey.X*aux.X + ey.Y*aux.Y
	//coordinates
	x := (math.Pow(r1, 2) - math.Pow(r2, 2) + math.Pow(p2p1Distance, 2)) / (2 * p2p1Distance)
	y := (math.Pow(r1, 2)-math.Pow(r3, 2)+math.Pow(i, 2)+math.Pow(j, 2))/(2*j) - i*x/j
	//result coordinates
	finalX := point1.X + x*ex.X + y*ey.X
	finalY := point1.Y + x*ex.Y + y*ey.Y

	return finalX, finalY
}
