package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string      `json:"routeId"`
	ClientID  string      `json: "clientId"`
	Positions []Positions `json:"position`
}

type Positions struct {
	Lat  float64 `json:"Lat"`
	Long float64 `json:"Long"`
}

func NewRoute() *Route {
	return &Route{}
}

// Carregando Posições
func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Delimitador
		data := strings.Split(scanner.Text(), ",")
		// Converter dados da latitude para Float
		Lat, err := strconv.ParseFloat(data[0], 64)
		// Error
		if err != nil {
			return nil
		}
		Long, err := strconv.ParseFloat(data[1], 64)
		// Error
		if err != nil {
			return nil
		}
		// Adicionando uma nova posição na lista
		r.Positions = append(r.Positions, Positions{
			Lat:  Lat,
			Long: Long,
		})

	}
	return nil
}

// POsição Parcial
type PartialRoutePosition struct {
	ID       string    `json:"RouteId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"` // FInished = True = Viagem finalizada
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
