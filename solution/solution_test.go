package solution

import (
	"context"
	"main/repository"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SolutionSuite struct {
	suite.Suite
	Solution
}

func (su *SolutionSuite) SetupTest() {
	su.ctx = context.Background()
	repo := repository.New()
	su.Solution = New(su.ctx, repo)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SolutionSuite))
}
