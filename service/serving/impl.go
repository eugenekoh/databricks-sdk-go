// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just ServingEndpoints API methods
type servingEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *servingEndpointsImpl) BuildLogs(ctx context.Context, request BuildLogsRequest) (*BuildLogsResponse, error) {
	var buildLogsResponse BuildLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/build-logs", request.Name, request.ServedModelName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &buildLogsResponse)
	return &buildLogsResponse, err
}

func (a *servingEndpointsImpl) Create(ctx context.Context, request CreateServingEndpoint) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := "/api/2.0/serving-endpoints"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) Delete(ctx context.Context, request DeleteServingEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *servingEndpointsImpl) ExportMetrics(ctx context.Context, request ExportMetricsRequest) error {
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/metrics", request.Name)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, nil)
	return err
}

func (a *servingEndpointsImpl) Get(ctx context.Context, request GetServingEndpointRequest) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) GetPermissionLevels(ctx context.Context, request GetServingEndpointPermissionLevelsRequest) (*GetServingEndpointPermissionLevelsResponse, error) {
	var getServingEndpointPermissionLevelsResponse GetServingEndpointPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v/permissionLevels", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getServingEndpointPermissionLevelsResponse)
	return &getServingEndpointPermissionLevelsResponse, err
}

func (a *servingEndpointsImpl) GetPermissions(ctx context.Context, request GetServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}

func (a *servingEndpointsImpl) List(ctx context.Context) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := "/api/2.0/serving-endpoints"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

func (a *servingEndpointsImpl) Logs(ctx context.Context, request LogsRequest) (*ServerLogsResponse, error) {
	var serverLogsResponse ServerLogsResponse
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/served-models/%v/logs", request.Name, request.ServedModelName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &serverLogsResponse)
	return &serverLogsResponse, err
}

func (a *servingEndpointsImpl) Query(ctx context.Context, request QueryRequest) (*QueryEndpointResponse, error) {
	var queryEndpointResponse QueryEndpointResponse
	path := fmt.Sprintf("/serving-endpoints/%v/invocations", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &queryEndpointResponse)
	return &queryEndpointResponse, err
}

func (a *servingEndpointsImpl) SetPermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}

func (a *servingEndpointsImpl) UpdateConfig(ctx context.Context, request EndpointCoreConfigInput) (*ServingEndpointDetailed, error) {
	var servingEndpointDetailed ServingEndpointDetailed
	path := fmt.Sprintf("/api/2.0/serving-endpoints/%v/config", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &servingEndpointDetailed)
	return &servingEndpointDetailed, err
}

func (a *servingEndpointsImpl) UpdatePermissions(ctx context.Context, request ServingEndpointPermissionsRequest) (*ServingEndpointPermissions, error) {
	var servingEndpointPermissions ServingEndpointPermissions
	path := fmt.Sprintf("/api/2.0/permissions/serving-endpoints/%v", request.ServingEndpointId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &servingEndpointPermissions)
	return &servingEndpointPermissions, err
}
