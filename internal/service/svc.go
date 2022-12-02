package service

import (
	"context"
	"main/internal/repository"
	"os"
	"reflect"

	"github.com/yanun0323/pkg/logs"
)

type Service struct {
	l    *logs.Logger
	ctx  context.Context
	repo repository.Repo
}

func New(ctx context.Context, repo repository.Repo) Service {
	return Service{
		l:    logs.New("advent of code", 2),
		ctx:  ctx,
		repo: repo,
	}
}

func (svc Service) Run() {
	body, err := svc.repo.GetQuestion()
	if err != nil {
		svc.l.Errorf("get day failed, %+v", err)
		return
	}
	svc.invoke("Day"+os.Getenv("DAY")+"A", body)
	svc.invoke("Day"+os.Getenv("DAY")+"B", body)
}

func (svc Service) invoke(funcName string, args ...any) {
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
