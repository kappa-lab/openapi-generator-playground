/*
 * A Todo list application
 *
 * From the todo list tutorial on goswagger.io
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// TodosApiService is a service that implements the logic for the TodosApiServicer
// This service should implement the business logic for every endpoint for the TodosApi API.
// Include any external packages or services that will be required by this service.
type TodosApiService struct {
}

// NewTodosApiService creates a default api service
func NewTodosApiService() TodosApiServicer {
	return &TodosApiService{}
}

// AddOne - 
func (s *TodosApiService) AddOne(ctx context.Context, item Item) (ImplResponse, error) {
	// TODO - update AddOne with the required logic for this service method.
	// Add api_todos_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Item{}) or use other options such as http.Ok ...
	//return Response(201, Item{}), nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddOne method not implemented")
}

// DestroyOne - 
func (s *TodosApiService) DestroyOne(ctx context.Context, id int64) (ImplResponse, error) {
	// TODO - update DestroyOne with the required logic for this service method.
	// Add api_todos_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DestroyOne method not implemented")
}

// FindTodos - 
func (s *TodosApiService) FindTodos(ctx context.Context, since int64, limit int32) (ImplResponse, error) {
	// TODO - update FindTodos with the required logic for this service method.
	// Add api_todos_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Item{}) or use other options such as http.Ok ...
	//return Response(200, []Item{}), nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("FindTodos method not implemented")
}

// UpdateOne - 
func (s *TodosApiService) UpdateOne(ctx context.Context, id int64, item Item) (ImplResponse, error) {
	// TODO - update UpdateOne with the required logic for this service method.
	// Add api_todos_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Item{}) or use other options such as http.Ok ...
	//return Response(200, Item{}), nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateOne method not implemented")
}