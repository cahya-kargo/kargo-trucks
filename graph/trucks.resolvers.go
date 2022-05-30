package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/cahya-kargo/kargo-trucks/graph/generated"
	"github.com/cahya-kargo/kargo-trucks/graph/model"
	"github.com/segmentio/ksuid"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	isValidatedPlateNumber := validatePlateNumber(plateNo)
	if isValidatedPlateNumber == false {
		return nil, errors.New("Plate number is invalid")
	}
	truck := &model.Truck{
		ID:        ksuid.New().String(),
		PlateNo:   plateNo,
		IsDeleted: &f,
		CreatedAt: int(time.Now().UnixMicro()),
		UpdatedAt: int(time.Now().UnixMicro()),
	}
	r.Trucks = append(r.Trucks, truck)
	return truck, nil
}

func (r *mutationResolver) UpdateTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	isValidatedPlateNumber := validatePlateNumber(plateNo)
	truck := &model.Truck{
		PlateNo: plateNo,
	}
	if isValidatedPlateNumber == false {
		return nil, errors.New("Plate number is invalid")
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

func (r *queryResolver) PaginatedTrucks(ctx context.Context, first *int, page *int, id *string, plateNo *string) (*model.ResponsePagination, error) {
	var array []*model.Truck
	var ids string
	if id != nil {
		ids = *id
	}
	for _, v := range r.Trucks {
		if v.ID == ids && v.IsDeleted != &t {
			array = append(array, v)
		}
		if id == nil {
			array = append(array, v)

		}
	}
	var pages int
	if *page == 0 {
		pages = *page + 1
	} else {
		pages = *page
	}
	size := *first*pages + 1
	if size > len(array) {
		size = len(array)
	}
	fmt.Println(size, pages)
	result := array[*page:size]
	sort.SliceStable(array, func(i, j int) bool {
		return array[i].UpdatedAt > array[j].UpdatedAt
	})
	return &model.ResponsePagination{
		Data: result,
		Meta: &model.Pagination{
			Page:      *page,
			First:     *first,
			TotalData: len(array),
		},
	}, nil
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
func validatePlateNumber(plateNumber string) bool {
	if plateNumber == "" {
		return false
	} else {
		trimmedStr := strings.TrimSpace(plateNumber)
		arrayStr := strings.Split(trimmedStr, " ")
		matchFirst, _ := regexp.MatchString(`/^[a-zA-Z]{2}$`, arrayStr[0])
		matchSecond, _ := regexp.MatchString(`/^[0-9]{4}$`, arrayStr[1])
		matchThird, _ := regexp.MatchString(`/^[a-zA-Z]{1,3}$`, arrayStr[2])

		if matchFirst != false && matchSecond != false && matchThird != false {
			return false
		}
		return true
	}

}

var t bool = true
var f bool = false
