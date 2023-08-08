package services

import (
	"context"
	"go-gql/graph/db"
	"go-gql/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func (u *userService) GetRepositoryByName(ctx context.Context, name string) (*model.Repository, error) {
	repository, err := db.Repositories(
		qm.Select(
			db.RepositoryTableColumns.ID,
			db.RepositoryTableColumns.Name,
			db.RepositoryTableColumns.CreatedAt,
		),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, u.exec)

	if err != nil {
		return nil, err
	}

	return convertRepository(repository), nil
}

func convertRepository(repository *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repository.ID,
		Name:      repository.Name,
		CreatedAt: repository.CreatedAt,
	}
}
