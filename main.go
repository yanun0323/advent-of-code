package main

import (
	"context"
	"main/infra"
	"main/internal/repository"
	"main/internal/service"

	"github.com/yanun0323/pkg/logs"
)

func main() {
	l := logs.New("advent of code", 2)
	ctx := context.Background()

	if err := infra.Init("config"); err != nil {
		l.Fatalf("init config failed, %+v", err)
	}

	svc := service.New(ctx, repository.New())
	svc.Run()
}
