package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"go_docker/domain/model"
	generated1 "go_docker/graphql/generated"
	"time"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "TODO-1",
			Text: "My Todo 1",
			User: &model.User{
				ID:   "User-1",
				Name: "hsaki",
			},
			Done: true,
		},
		{
			ID:   "TODO-2",
			Text: "My Todo 2",
			User: &model.User{
				ID:   "User-1",
				Name: "hsaki",
			},
			Done: false,
		},
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context) (string, error) {
	fmt.Println("wait 5 seconds for debug...")
	time.Sleep(time.Second * 5)
	fmt.Println("response")
	return "Hello World", nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
