package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dummy-gql/graph/generated"
	"dummy-gql/model"
	"fmt"
)

func (r *queryResolver) GetRoot(ctx context.Context, someArg string) (*model.GetRootRes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *typeCompA1Resolver) TableA1(ctx context.Context, obj *model.TypeCompA1) ([]*model.TypeTableA1, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *typeCompA2Resolver) TableA2(ctx context.Context, obj *model.TypeCompA2) ([]*model.TypeTableA2, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *typeGroupAResolver) CompA1(ctx context.Context, obj *model.TypeGroupA) (*model.TypeCompA1, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *typeGroupAResolver) CompA2(ctx context.Context, obj *model.TypeGroupA) (*model.TypeCompA2, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *typeRootResolver) GroupA(ctx context.Context, obj *model.TypeRoot) ([]*model.TypeGroupA, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TypeCompA1 returns generated.TypeCompA1Resolver implementation.
func (r *Resolver) TypeCompA1() generated.TypeCompA1Resolver { return &typeCompA1Resolver{r} }

// TypeCompA2 returns generated.TypeCompA2Resolver implementation.
func (r *Resolver) TypeCompA2() generated.TypeCompA2Resolver { return &typeCompA2Resolver{r} }

// TypeGroupA returns generated.TypeGroupAResolver implementation.
func (r *Resolver) TypeGroupA() generated.TypeGroupAResolver { return &typeGroupAResolver{r} }

// TypeRoot returns generated.TypeRootResolver implementation.
func (r *Resolver) TypeRoot() generated.TypeRootResolver { return &typeRootResolver{r} }

type queryResolver struct{ *Resolver }
type typeCompA1Resolver struct{ *Resolver }
type typeCompA2Resolver struct{ *Resolver }
type typeGroupAResolver struct{ *Resolver }
type typeRootResolver struct{ *Resolver }
