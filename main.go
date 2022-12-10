package main

import (
	"context"
	"main/repository"
	"main/solution"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/yanun0323/pkg/config"
	"github.com/yanun0323/pkg/logs"
)

func main() {
	l := logs.New("advent of code", 2)
	ctx := context.Background()

	if err := initConfig("config"); err != nil {
		l.Fatalf("init config failed, %+v", err)
	}

	svc := solution.New(ctx, repository.New())
	svc.Run()
}

var _once sync.Once

func initConfig(cfgName string) error {
	var err error
	_once.Do(
		func() {
			_, f, _, _ := runtime.Caller(0)
			cfgPath := filepath.Join(filepath.Dir(f), "../../config")
			if err = config.Init(cfgPath, cfgName, false); err != nil {
				return
			}
		},
	)
	return err
}
