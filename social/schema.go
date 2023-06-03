package social

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/pulkitbhardwaj/alpha/social/internal"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (internal.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]internal.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *internal.User) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *userWhereInputResolver) ID(ctx context.Context, obj *internal.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *userWhereInputResolver) IDNeq(ctx context.Context, obj *internal.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *userWhereInputResolver) IDIn(ctx context.Context, obj *internal.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userWhereInputResolver) IDNotIn(ctx context.Context, obj *internal.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *userWhereInputResolver) IDGt(ctx context.Context, obj *internal.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *userWhereInputResolver) IDGte(ctx context.Context, obj *internal.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *userWhereInputResolver) IDLt(ctx context.Context, obj *internal.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *userWhereInputResolver) IDLte(ctx context.Context, obj *internal.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

// User returns internal.UserResolver implementation.
func (r *Resolver) User() internal.UserResolver { return &userResolver{r} }

// UserWhereInput returns internal.UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() internal.UserWhereInputResolver {
	return &userWhereInputResolver{r}
}

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
