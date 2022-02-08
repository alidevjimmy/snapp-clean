package postgres

import (
	"context"
	"github.com/alidevjimmy/snapp-clean/internal/entity/model"
)

func (p *postgres) CreateUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (p *postgres) UpdateUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (p *postgres) DeleteUser(ctx context.Context, id int) error {
	panic("implement me")
}

func (p *postgres) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	panic("implement me")
}

func (p *postgres) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	panic("implement me")
}
