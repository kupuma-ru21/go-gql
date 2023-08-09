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
	ProjectService
}

type services struct {
	*userService
	*repositoryService
	*issueService
	*projectService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByName(ctx context.Context, owner, name string) (*model.Repository, error)
	GetRepositoryByID(ctx context.Context, id string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type ProjectService interface {
	GetProjectByOwnerAndNumber(ctx context.Context, ownerID string, number int) (*model.ProjectV2, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService:      &issueService{exec: exec},
		projectService:    &projectService{exec: exec},
	}
}
