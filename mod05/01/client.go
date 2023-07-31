package mod05

import (
	"fmt"
	"net/rpc"
)

type ClientGetTemperatureRequest struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type ClientGetTemperatureResponse struct {
	Temperature float64 `json:"temp"`
}

func StartClient(longitude float64, latitude float64, port int) float64 {
	req := ClientGetTemperatureRequest{
		Longitude: longitude,
		Latitude:  latitude,
	}
	res := ClientGetTemperatureResponse{}
	client, _ := rpc.DialHTTP("tcp", fmt.Sprintf(":%d", port))
	client.Call("WeatherService.GetTemperature", req, &res)
	return res.Temperature
}
