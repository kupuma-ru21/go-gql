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

func (r *repositoryService) GetRepositoryByName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repository, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}

	return convertRepository(repository), nil
}

func convertRepository(repository *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repository.ID,
		Owner:     &model.User{ID: repository.Owner},
		Name:      repository.Name,
		CreatedAt: repository.CreatedAt,
	}
}
