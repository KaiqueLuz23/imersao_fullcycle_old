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
	ID string 
	ClientID string
	Positions []Positions
}

type Positions struct {
	Lat float64
	Long float64
}


// Carregando Posições
func (r *Route) LoadPositions() error {
	if r.ID == ""{
		return errors.New(text: "route id not informed")
	}
	f, err := os.Open(name:"destinations/"  + r.ID + ".txt")
	if err != nil{
		return err
	}
	defer f.Close()

	scanner := bufior.NewScanner(f)
	for scanner.Scan(){
		// Delimitador
		data := string.Split(scanner.Text(), sep:",")
		// Converter dados da latitude para Float
		Lat, err := strconv.ParseFLoat(data[0],64)
		// Error
		if err != nil{
			return nil
		}
		Long, err := strconv.ParseFLoat(data[1],64)
		// Error
		if err != nil{
			return nil
		}
		// Adicionando uma nova posição na lista
		r.Positions = append (r.Positions, Positions{
			Lat: Lat,
			Long: Long,
		})

	}
	return nil
}


