package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/cahya-kargo/kargo-trucks/graph/generated"
	"github.com/cahya-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	truck := &model.Truck{
		ID:        fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo:   plateNo,
		IsDeleted: &f,
		CreatedAt: int(time.Now().UnixMicro()),
		UpdatedAt: int(time.Now().UnixMicro()),
	}
	r.Trucks = append(r.Trucks, truck)
	return truck, nil
}

func (r *mutationResolver) UpdateTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	truck := &model.Truck{
		PlateNo: plateNo,
	}
	if *id == "" {
		panic("Id cannot be null")
	}
	for _, v := range r.Trucks {
		if *id == v.ID {
			v.PlateNo = plateNo
			v.UpdatedAt = int(time.Now().UnixMicro())
		}
	}
	return truck, nil
}

func (r *mutationResolver) DeleteTruck(ctx context.Context, id *string) (*model.Response, error) {
	if *id == "" {
		panic("Id cannot be null")
	}
	for _, v := range r.Trucks {
		if *id == v.ID {
			v.IsDeleted = &t
			v.UpdatedAt = int(time.Now().UnixMicro())
		}
	}
	return &model.Response{
		Message: "Success Deleted 1 Data",
	}, nil
}

func (r *queryResolver) PaginatedTrucks(ctx context.Context) ([]*model.Truck, error) {
	var array []*model.Truck
	for _, v := range r.Trucks {
		if v.IsDeleted != &t {
			array = append(array, v)
		}
	}
	sort.SliceStable(array, func(i, j int) bool {
		return array[i].UpdatedAt > array[j].UpdatedAt
	})
	return array, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var t bool = true
var f bool = false
