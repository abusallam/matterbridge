// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UnifiedRoleDefinitionRequestBuilder is request builder for UnifiedRoleDefinition
type UnifiedRoleDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleDefinitionRequest
func (b *UnifiedRoleDefinitionRequestBuilder) Request() *UnifiedRoleDefinitionRequest {
	return &UnifiedRoleDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleDefinitionRequest is request for UnifiedRoleDefinition
type UnifiedRoleDefinitionRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Get(ctx context.Context) (resObj *UnifiedRoleDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Update(ctx context.Context, reqObj *UnifiedRoleDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}