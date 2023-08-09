package services

import (
	"context"
	"go-gql/graph/db"
	"go-gql/graph/model"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type projectService struct {
	exec boil.ContextExecutor
}

func (p *projectService) GetProjectByOwnerAndNumber(ctx context.Context, ownerID string, number int) (*model.ProjectV2, error) {
	project, err := db.Projects(
		qm.Select(
			db.ProjectColumns.ID,
			db.ProjectColumns.Title,
			db.ProjectColumns.Number,
			db.ProjectColumns.URL,
			db.ProjectColumns.Owner,
		),
		db.ProjectWhere.Owner.EQ(ownerID),
		db.ProjectWhere.Number.EQ(int64(number)),
	).One(ctx, p.exec)
	if err != nil {
		return nil, err
	}
	return convertProjectV2(project), nil
}

func convertProjectV2(project *db.Project) *model.ProjectV2 {
	projectURL, err := model.UnmarshalURI(project.URL)
	if err != nil {
		log.Println("invalid URI", project.URL)
	}

	return &model.ProjectV2{
		ID:     project.ID,
		Title:  project.Title,
		Number: int(project.Number),
		URL:    projectURL,
		Owner:  &model.User{ID: project.Owner},
	}
}
