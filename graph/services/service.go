package services

import (
	"context"
	"go-gql/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
}

type services struct {
	*userService
	*repositoryService
	*issueService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByName(ctx context.Context, owner, name string) (*model.Repository, error)
	GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService:      &issueService{exec: exec},
	}
}
