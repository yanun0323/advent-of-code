package service

import (
	"context"
	"main/internal/repository"

	"github.com/yanun0323/pkg/logs"
)

type Service struct {
	l    *logs.Logger
	ctx  context.Context
	repo repository.Repo
}

func New(ctx context.Context, repo repository.Repo) Service {
	return Service{
		l:    logs.Get(ctx),
		ctx:  ctx,
		repo: repo,
	}
}

func (svc Service) Run() {
	body, err := svc.repo.GetQuestion()
	if err != nil {
		svc.l.Errorf("get day failed, %+v", err)
	}
	ans := svc.Day1B(body)
	if ans != nil {
		svc.l.Info("answer is ", ans)
		return
	}
	svc.l.Warn("no answer here")
}
