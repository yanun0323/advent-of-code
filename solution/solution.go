package solution

import (
	"context"
	"main/repository"
	"os"
	"reflect"

	"github.com/yanun0323/pkg/logs"
)

type Solution struct {
	l    *logs.Logger
	ctx  context.Context
	repo repository.Repo
}

func New(ctx context.Context, repo repository.Repo) Solution {
	return Solution{
		l:    logs.New("advent of code", 2),
		ctx:  ctx,
		repo: repo,
	}
}

func (svc Solution) Run() {
	body := svc.repo.GetPuzzleInput()
	if len(body) == 0 {
		svc.l.Errorf("empty puzzle input")
		return
	}
	svc.invoke("Day"+os.Getenv("DAY"), body)
}

func (svc Solution) invoke(funcName string, args ...any) {
	method, ok := reflect.TypeOf(svc).MethodByName(funcName)
	if !ok {
		svc.l.Warnf("can't find function %s", funcName)
		return
	}

	inputs := make([]reflect.Value, 0, len(args)+1)
	inputs = append(inputs, reflect.ValueOf(svc))
	for i := range args {
		inputs = append(inputs, reflect.ValueOf(args[i]))
	}
	ans := method.Func.Call(inputs)
	if len(ans) == 0 || ans[0].IsNil() {
		svc.l.Warn("nil answer from function %s", funcName)
		return
	}
	svc.l.Info("function ", funcName, " executed, answer is ", ans[0].Interface())
}
