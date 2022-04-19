package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/zfirdavs/another-rest/internal/entity"
	"google.golang.org/api/iterator"
)

type repo struct {
	projectID      string
	collectionName string
}

func NewFirestoreRepository(projectID string, collectionName string) PostRepository {
	return &repo{projectID, collectionName}
}

func (r *repo) Save(ctx context.Context, post *entity.Post) error {
	client, err := firestore.NewClient(ctx, r.projectID)
	if err != nil {
		return fmt.Errorf("failed to create a firestore client: %w", err)
	}
	defer client.Close()

	_, _, err = client.Collection(r.collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		return fmt.Errorf("failed to add post to collection: %w", err)
	}
	return nil
}

func (r *repo) FindAll(ctx context.Context) ([]entity.Post, error) {
	client, err := firestore.NewClient(ctx, r.projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create a firestore client: %w", err)
	}
	defer client.Close()

	var posts []entity.Post

	iter := client.Collection(r.collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to iterate the lists of posts: %w", err)
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}
