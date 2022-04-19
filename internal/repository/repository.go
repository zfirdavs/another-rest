package repository

import (
	"context"

	"github.com/zfirdavs/another-rest/internal/entity"
)

type PostRepository interface {
	Save(ctx context.Context, post *entity.Post) error
	FindAll(ctx context.Context) ([]entity.Post, error)
}
