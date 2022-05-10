package service

import (
	"context"
	"errors"
	"math/rand"

	"github.com/zfirdavs/another-rest/internal/entity"
	"github.com/zfirdavs/another-rest/internal/repository"
)

type PostService interface {
	Validate(ctx context.Context, post *entity.Post) error
	Create(ctx context.Context, post *entity.Post) error
	FindAll(ctx context.Context) ([]entity.Post, error)
}

type service struct {
	repo repository.PostRepository
}

func NewPost(repo repository.PostRepository) PostService {
	return &service{repo}
}

func (s *service) Validate(ctx context.Context, post *entity.Post) (err error) {
	if post == nil {
		err = errors.New("the post is empty")
	}

	if post.Title == "" {
		err = errors.New("the post title is empty")
	}
	return
}

func (s *service) Create(ctx context.Context, post *entity.Post) error {
	post.ID = rand.Int63()
	return s.repo.Save(ctx, post)
}

func (s *service) FindAll(ctx context.Context) ([]entity.Post, error) {
	return s.repo.FindAll(ctx)
}
