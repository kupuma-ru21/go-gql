package services

import (
	"context"
	"go-gql/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	*repositoryService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByName(ctx context.Context, owner, name string) (*model.Repository, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
	}
}
