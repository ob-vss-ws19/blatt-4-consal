package cli

import (
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("client"))
	//TODO
	service.Init()
	switch nil {
	case "cinema":
	case "movie":
		//TODO: Add more cases and implementations
	default:
		return
	}

}
