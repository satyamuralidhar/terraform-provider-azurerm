package siterecovery

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

// ReplicationNetworkMappingsClient is the client for the ReplicationNetworkMappings methods of the Siterecovery
// service.
type ReplicationNetworkMappingsClient struct {
	BaseClient
}

// NewReplicationNetworkMappingsClient creates an instance of the ReplicationNetworkMappingsClient client.
func NewReplicationNetworkMappingsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationNetworkMappingsClient {
	return NewReplicationNetworkMappingsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationNetworkMappingsClientWithBaseURI creates an instance of the ReplicationNetworkMappingsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewReplicationNetworkMappingsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationNetworkMappingsClient {
	return ReplicationNetworkMappingsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to create an ASR network mapping.
// Parameters:
// fabricName - primary fabric name.
// networkName - primary network name.
// networkMappingName - network mapping name.
// input - create network mapping input.
func (client ReplicationNetworkMappingsClient) Create(ctx context.Context, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput) (result ReplicationNetworkMappingsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, fabricName, networkName, networkMappingName, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationNetworkMappingsClient) CreatePreparer(ctx context.Context, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":         autorest.Encode("path", fabricName),
		"networkMappingName": autorest.Encode("path", networkMappingName),
		"networkName":        autorest.Encode("path", networkName),
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"resourceName":       autorest.Encode("path", client.ResourceName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationNetworkMappingsClient) CreateSender(req *http.Request) (future ReplicationNetworkMappingsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationNetworkMappingsClient) CreateResponder(resp *http.Response) (result NetworkMapping, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a network mapping.
// Parameters:
// fabricName - primary fabric name.
// networkName - primary network name.
// networkMappingName - ARM Resource Name for network mapping.
func (client ReplicationNetworkMappingsClient) Delete(ctx context.Context, fabricName string, networkName string, networkMappingName string) (result ReplicationNetworkMappingsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, fabricName, networkName, networkMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationNetworkMappingsClient) DeletePreparer(ctx context.Context, fabricName string, networkName string, networkMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":         autorest.Encode("path", fabricName),
		"networkMappingName": autorest.Encode("path", networkMappingName),
		"networkName":        autorest.Encode("path", networkName),
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"resourceName":       autorest.Encode("path", client.ResourceName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationNetworkMappingsClient) DeleteSender(req *http.Request) (future ReplicationNetworkMappingsDeleteFuture, err error) {
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
func (client ReplicationNetworkMappingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of an ASR network mapping
// Parameters:
// fabricName - primary fabric name.
// networkName - primary network name.
// networkMappingName - network mapping name.
func (client ReplicationNetworkMappingsClient) Get(ctx context.Context, fabricName string, networkName string, networkMappingName string) (result NetworkMapping, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, networkName, networkMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationNetworkMappingsClient) GetPreparer(ctx context.Context, fabricName string, networkName string, networkMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":         autorest.Encode("path", fabricName),
		"networkMappingName": autorest.Encode("path", networkMappingName),
		"networkName":        autorest.Encode("path", networkName),
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"resourceName":       autorest.Encode("path", client.ResourceName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationNetworkMappingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationNetworkMappingsClient) GetResponder(resp *http.Response) (result NetworkMapping, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all ASR network mappings in the vault.
func (client ReplicationNetworkMappingsClient) List(ctx context.Context) (result NetworkMappingCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.List")
		defer func() {
			sc := -1
			if result.nmc.Response.Response != nil {
				sc = result.nmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.nmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "List", resp, "Failure sending request")
		return
	}

	result.nmc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationNetworkMappingsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationNetworkMappingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationNetworkMappingsClient) ListResponder(resp *http.Response) (result NetworkMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationNetworkMappingsClient) listNextResults(ctx context.Context, lastResults NetworkMappingCollection) (result NetworkMappingCollection, err error) {
	req, err := lastResults.networkMappingCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationNetworkMappingsClient) ListComplete(ctx context.Context) (result NetworkMappingCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationNetworks lists all ASR network mappings for the specified network.
// Parameters:
// fabricName - primary fabric name.
// networkName - primary network name.
func (client ReplicationNetworkMappingsClient) ListByReplicationNetworks(ctx context.Context, fabricName string, networkName string) (result NetworkMappingCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.ListByReplicationNetworks")
		defer func() {
			sc := -1
			if result.nmc.Response.Response != nil {
				sc = result.nmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationNetworksNextResults
	req, err := client.ListByReplicationNetworksPreparer(ctx, fabricName, networkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "ListByReplicationNetworks", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationNetworksSender(req)
	if err != nil {
		result.nmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "ListByReplicationNetworks", resp, "Failure sending request")
		return
	}

	result.nmc, err = client.ListByReplicationNetworksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "ListByReplicationNetworks", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationNetworksPreparer prepares the ListByReplicationNetworks request.
func (client ReplicationNetworkMappingsClient) ListByReplicationNetworksPreparer(ctx context.Context, fabricName string, networkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"networkName":       autorest.Encode("path", networkName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationNetworksSender sends the ListByReplicationNetworks request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationNetworkMappingsClient) ListByReplicationNetworksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationNetworksResponder handles the response to the ListByReplicationNetworks request. The method always
// closes the http.Response Body.
func (client ReplicationNetworkMappingsClient) ListByReplicationNetworksResponder(resp *http.Response) (result NetworkMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationNetworksNextResults retrieves the next set of results, if any.
func (client ReplicationNetworkMappingsClient) listByReplicationNetworksNextResults(ctx context.Context, lastResults NetworkMappingCollection) (result NetworkMappingCollection, err error) {
	req, err := lastResults.networkMappingCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "listByReplicationNetworksNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationNetworksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "listByReplicationNetworksNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationNetworksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "listByReplicationNetworksNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationNetworksComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationNetworkMappingsClient) ListByReplicationNetworksComplete(ctx context.Context, fabricName string, networkName string) (result NetworkMappingCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.ListByReplicationNetworks")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationNetworks(ctx, fabricName, networkName)
	return
}

// Update the operation to update an ASR network mapping.
// Parameters:
// fabricName - primary fabric name.
// networkName - primary network name.
// networkMappingName - network mapping name.
// input - update network mapping input.
func (client ReplicationNetworkMappingsClient) Update(ctx context.Context, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput) (result ReplicationNetworkMappingsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationNetworkMappingsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, fabricName, networkName, networkMappingName, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationNetworkMappingsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ReplicationNetworkMappingsClient) UpdatePreparer(ctx context.Context, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":         autorest.Encode("path", fabricName),
		"networkMappingName": autorest.Encode("path", networkMappingName),
		"networkName":        autorest.Encode("path", networkName),
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"resourceName":       autorest.Encode("path", client.ResourceName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationNetworkMappingsClient) UpdateSender(req *http.Request) (future ReplicationNetworkMappingsUpdateFuture, err error) {
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
func (client ReplicationNetworkMappingsClient) UpdateResponder(resp *http.Response) (result NetworkMapping, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
