package graph

import (
	"context"
	"errors"
	"fmt"
	"go-gql/graph/model"
	"go-gql/graph/services"

	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

type userBatcher struct {
	Srv services.Services
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}
	return &Loaders{
		// dataloader.Loader[string, *model.User]構造体型をセットするために、
		// dataloader.NewBatchedLoader関数を呼び出す
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	// 引数と戻り値のスライスlenは等しくする
	fmt.Println("IDs", IDs)
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexes := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexes[ID] = i
	}

	// サービス層のメソッドを使い、指定されたIDを持つユーザーを全て取得する
	// (ListUsersByIDメソッド内では、IN句を用いたselect文が実行されている)
	users, err := u.Srv.ListUsersByID(ctx, IDs)

	// 取得結果を、戻り値resultsの中の適切な場所に格納する
	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if err != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexes[user.ID]] = rsl
	}
	return results
}
