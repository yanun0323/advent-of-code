package service

import (
	"context"
	"main/internal/repository"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	Service
}

func (su *ServiceSuite) SetupTest() {
	su.ctx = context.Background()
	repo := repository.New()
	su.Service = New(su.ctx, repo)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
