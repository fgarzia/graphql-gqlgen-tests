package graphqls

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mercadolibre/graphql-tests/internal/model"
)

func (r *queryResolver) Employees(ctx context.Context, name *string) ([]*model.Employee, error) {
	var result []*model.Employee

	id1 := 22
	id2 := 23
	name1 := "Nicolas"
	name2 := "Oscar"
	supervisor := 33
	description := "description"

	result = append(result, &model.Employee{
		ID:         &id1,
		Name:       &name1,
		Supervisor: &supervisor,
		Role: &model.Role{
			ID:          &id1,
			Description: &description,
		},
	})

	result = append(result, &model.Employee{
		ID:         &id2,
		Name:       &name2,
		Supervisor: &supervisor,
		Role: &model.Role{
			ID:          &id2,
			Description: &description,
		},
	})

	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
