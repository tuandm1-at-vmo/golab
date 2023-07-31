package mod05

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ServerGetTemperatureRequest struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type ServerGetTemperatureResponse struct {
	Temperature float64 `json:"temp"`
}

type OpenWeatherMapResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

type WeatherService int64

func (ws *WeatherService) GetTemperature(req ServerGetTemperatureRequest, res *ServerGetTemperatureResponse) error {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?units=metric&appid=81804a170a781fff7807ddc3eb8fb016&lon=%f&lat=%f", req.Longitude, req.Latitude)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	parsed := OpenWeatherMapResponse{}
	err = json.Unmarshal(bytes, &parsed)
	if err != nil {
		return err
	}
	res = &ServerGetTemperatureResponse{
		Temperature: parsed.Main.Temperature,
	}
	return nil
}

func StartServer(port int) {
	rpc.Register(new(WeatherService))
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	} else {
		go http.Serve(listener, nil)
	}
}
