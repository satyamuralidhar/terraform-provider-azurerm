package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ApplicationGatewayPrivateEndpointConnectionsClient is the network Client
type ApplicationGatewayPrivateEndpointConnectionsClient struct {
	BaseClient
}

// NewApplicationGatewayPrivateEndpointConnectionsClient creates an instance of the
// ApplicationGatewayPrivateEndpointConnectionsClient client.
func NewApplicationGatewayPrivateEndpointConnectionsClient(subscriptionID string) ApplicationGatewayPrivateEndpointConnectionsClient {
	return NewApplicationGatewayPrivateEndpointConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewayPrivateEndpointConnectionsClientWithBaseURI creates an instance of the
// ApplicationGatewayPrivateEndpointConnectionsClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewApplicationGatewayPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewayPrivateEndpointConnectionsClient {
	return ApplicationGatewayPrivateEndpointConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes the specified private endpoint connection on application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
// connectionName - the name of the application gateway private endpoint connection.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string) (result ApplicationGatewayPrivateEndpointConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateEndpointConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, applicationGatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"connectionName":         autorest.Encode("path", connectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) DeleteSender(req *http.Request) (future ApplicationGatewayPrivateEndpointConnectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified private endpoint connection on application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
// connectionName - the name of the application gateway private endpoint connection.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string) (result ApplicationGatewayPrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateEndpointConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, applicationGatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"connectionName":         autorest.Encode("path", connectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) GetResponder(resp *http.Response) (result ApplicationGatewayPrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all private endpoint connections on an application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) List(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewayPrivateEndpointConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateEndpointConnectionsClient.List")
		defer func() {
			sc := -1
			if result.agpeclr.Response.Response != nil {
				sc = result.agpeclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.agpeclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.agpeclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) ListResponder(resp *http.Response) (result ApplicationGatewayPrivateEndpointConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) listNextResults(ctx context.Context, lastResults ApplicationGatewayPrivateEndpointConnectionListResult) (result ApplicationGatewayPrivateEndpointConnectionListResult, err error) {
	req, err := lastResults.applicationGatewayPrivateEndpointConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewayPrivateEndpointConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateEndpointConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, applicationGatewayName)
	return
}

// Update updates the specified private endpoint connection on application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
// connectionName - the name of the application gateway private endpoint connection.
// parameters - parameters supplied to update application gateway private endpoint connection operation.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) Update(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection) (result ApplicationGatewayPrivateEndpointConnectionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateEndpointConnectionsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateEndpointConnectionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"connectionName":         autorest.Encode("path", connectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) UpdateSender(req *http.Request) (future ApplicationGatewayPrivateEndpointConnectionsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ApplicationGatewayPrivateEndpointConnectionsClient) UpdateResponder(resp *http.Response) (result ApplicationGatewayPrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
