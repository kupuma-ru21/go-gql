package services

import (
	"context"
	"go-gql/graph/db"
	"go-gql/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
		// qm.Where("name = ?", name),
	).One(ctx, u.exec)

	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func convertUser(user *db.User) *model.User {
	return &model.User{ID: user.ID, Name: user.Name}
}
