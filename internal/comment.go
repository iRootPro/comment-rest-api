package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchFailed    = errors.New("failed to fetch comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Service struct {
	Store Store
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("recieve comment")
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err.Error())
		return Comment{}, ErrFetchFailed
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return comment, ErrNotImplemented
}
