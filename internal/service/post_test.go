package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zfirdavs/another-rest/internal/entity"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(ctx context.Context, post *entity.Post) error {
	args := m.Called()

	return args.Error(0)
}

func (m *MockRepository) FindAll(ctx context.Context) ([]entity.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var id int64 = 1
	post := entity.Post{
		ID:    id,
		Title: "1",
		Text:  "2",
	}

	// setup the expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPost(mockRepo)

	result, _ := testService.FindAll(context.Background())

	// behavioral mock assertion
	mockRepo.AssertExpectations(t)

	// data assertion
	assert.Equal(t, "1", result[0].Title)
	assert.Equal(t, "2", result[0].Text)
	assert.Equal(t, id, result[0].ID)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPost(nil)

	err := testService.Validate(context.Background(), nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "the post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{
		ID:    1,
		Title: "",
		Text:  "A",
	}

	testService := NewPost(nil)

	err := testService.Validate(context.Background(), &post)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "the post title is empty")
}
