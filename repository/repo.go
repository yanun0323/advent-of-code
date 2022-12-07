package repository

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type repo struct {
	client *http.Client
}

func New() Repo {
	return &repo{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (repo *repo) GetPuzzleInput() []string {
	url := viper.GetString("url.prefix") + os.Getenv("DAY") + viper.GetString("url.suffix")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Cookie", "session="+viper.GetString("session"))
	req.Header.Add("Content-Type", "application/json")

	response, err := repo.client.Do(req)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	buf, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	return strings.Split(string(buf), "\n")
}
