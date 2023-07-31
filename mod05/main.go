package mod05

import (
	"fmt"

	lab01 "tuanm.dev/golab/mod05/01"
)

func Run(lab int) {
	switch lab {
	case 1:
		runLab01()
	default:
		println("no lab found")
	}
}

func runLab01() {
	port := 6789
	longitude := 51.956055
	latitude := 7.857940
	lab01.StartServer(port)
	temperature := lab01.StartClient(longitude, latitude, port)
	fmt.Printf("temp = %f\n", temperature)
}
