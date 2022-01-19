package services

import (
	"github.com/asdine/storm"
	"github.com/maurodanieldev/quasar-oper-fire/controllers/request"
	"github.com/maurodanieldev/quasar-oper-fire/interfaces"
	"github.com/maurodanieldev/quasar-oper-fire/util"
)

const (
	dbPath = "database.db"
)

type satelliteService struct{}

func NewSatelliteService() interfaces.ISatelliteService {
	return &satelliteService{}
}

func (r *satelliteService) All() ([]*request.Satellite, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var satellites []*request.Satellite
	err = db.All(&satellites)
	if err != nil {
		return nil, err
	}
	return satellites, err
}

func (r *satelliteService) One(id string) (*request.Satellite, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(request.Satellite)
	err = db.One("Name", id, u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *satelliteService) Delete(name string) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	s := new(request.Satellite)
	err = db.One("Name", name, s)
	if err != nil {
		return err
	}
	return db.DeleteStruct(s)
}
func (r *satelliteService) Save(s *request.Satellite) error {
	if util.SatelliteExist(s.Name) {
		if err := s.Validate(); err != nil {
			return err
		}
		db, err := storm.Open(dbPath)
		if err != nil {
			return err
		}
		defer db.Close()
		return db.Save(s)
	}
	return nil
}
