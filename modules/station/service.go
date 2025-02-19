package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aliffadillah/mrt-schedule-golang.git/common/client"
)

type Service interface {
	GetAllStations() (response []StationRespone, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (response []StationRespone, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	for _, item := range stations {
		response = append(response, StationRespone{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}
