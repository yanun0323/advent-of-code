package main

import (
	"context"
	"io"
	"main/repository"
	"main/solution"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
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

func getPuzzleInput() ([]string, error) {
	url := viper.GetString("url.prefix") + os.Getenv("DAY") + viper.GetString("url.suffix")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", "session="+viper.GetString("session"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	buf, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(buf), "\n"), nil
}
