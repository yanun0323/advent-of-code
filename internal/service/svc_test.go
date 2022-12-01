package service

import (
	"context"
	"main/infra"
	"main/internal/repository"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	Service
}

func (su *ServiceSuite) SetupSuit() {
	su.ctx = context.Background()
	su.Require().Nil(infra.Init("cocnfig"))
	repo := repository.New()
	su.Service = New(su.ctx, repo)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
