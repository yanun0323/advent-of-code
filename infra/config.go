package infra

import (
	"path/filepath"
	"runtime"
	"sync"

	"github.com/yanun0323/pkg/config"
)

var _once sync.Once

func Init(cfgName string) error {
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
