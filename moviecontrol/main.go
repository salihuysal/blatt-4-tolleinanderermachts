package main

import (
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("moviectrl"))
	service.Init()
	err := proto.RegisterMovieControlHandler(service.Server(), new(MovieControl))

	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
