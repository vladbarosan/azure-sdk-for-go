package relay

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen ...
	Listen AccessRights = "Listen"
	// Manage ...
	Manage AccessRights = "Manage"
	// Send ...
	Send AccessRights = "Send"
)

// KeyType enumerates the values for key type.
type KeyType string

const (
	// PrimaryKey ...
	PrimaryKey KeyType = "PrimaryKey"
	// SecondaryKey ...
	SecondaryKey KeyType = "SecondaryKey"
)

// ProvisioningStateEnum enumerates the values for provisioning state enum.
type ProvisioningStateEnum string

const (
	// Created ...
	Created ProvisioningStateEnum = "Created"
	// Deleted ...
	Deleted ProvisioningStateEnum = "Deleted"
	// Failed ...
	Failed ProvisioningStateEnum = "Failed"
	// Succeeded ...
	Succeeded ProvisioningStateEnum = "Succeeded"
	// Unknown ...
	Unknown ProvisioningStateEnum = "Unknown"
	// Updating ...
	Updating ProvisioningStateEnum = "Updating"
)

// RelaytypeEnum enumerates the values for relaytype enum.
type RelaytypeEnum string

const (
	// HTTP ...
	HTTP RelaytypeEnum = "Http"
	// NetTCP ...
	NetTCP RelaytypeEnum = "NetTcp"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Standard ...
	Standard SkuTier = "Standard"
)

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// InvalidName ...
	InvalidName UnavailableReason = "InvalidName"
	// NameInLockdown ...
	NameInLockdown UnavailableReason = "NameInLockdown"
	// NameInUse ...
	NameInUse UnavailableReason = "NameInUse"
	// None ...
	None UnavailableReason = "None"
	// SubscriptionIsDisabled ...
	SubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// TooManyNamespaceInCurrentSubscription ...
	TooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// AccessKeys namespace/Relay Connection String
type AccessKeys struct {
	autorest.Response `json:"-"`
	// PrimaryConnectionString - Primary connection string of the created namespace authorization rule.
	PrimaryConnectionString *string `json:"primaryConnectionString,omitempty"`
	// SecondaryConnectionString - Secondary connection string of the created namespace authorization rule.
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
	// PrimaryKey - A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - A base64-encoded 256-bit secondary key for signing and validating the SAS token.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
	// KeyName - A string that describes the authorization rule.
	KeyName *string `json:"keyName,omitempty"`
}

// AuthorizationRule description of a namespace authorization rule.
type AuthorizationRule struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// AuthorizationRuleProperties - Authorization rule properties.
	*AuthorizationRuleProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for AuthorizationRule struct.
func (ar *AuthorizationRule) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var authorizationRuleProperties AuthorizationRuleProperties
				err = json.Unmarshal(*v, &authorizationRuleProperties)
				if err != nil {
					return err
				}
				ar.AuthorizationRuleProperties = &authorizationRuleProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ar.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ar.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ar.Type = &typeVar
			}
		}
	}

	return nil
}

// AuthorizationRuleListResult the response from the list namespace operation.
type AuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	// Value - Result of the list authorization rules operation.
	Value *[]AuthorizationRule `json:"value,omitempty"`
	// NextLink - Link to the next set of results. Not empty if value contains incomplete list of authorization rules.
	NextLink *string `json:"nextLink,omitempty"`
}

// AuthorizationRuleListResultIterator provides access to a complete listing of AuthorizationRule values.
type AuthorizationRuleListResultIterator struct {
	i    int
	page AuthorizationRuleListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AuthorizationRuleListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AuthorizationRuleListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AuthorizationRuleListResultIterator) Response() AuthorizationRuleListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AuthorizationRuleListResultIterator) Value() AuthorizationRule {
	if !iter.page.NotDone() {
		return AuthorizationRule{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (arlr AuthorizationRuleListResult) IsEmpty() bool {
	return arlr.Value == nil || len(*arlr.Value) == 0
}

// authorizationRuleListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (arlr AuthorizationRuleListResult) authorizationRuleListResultPreparer() (*http.Request, error) {
	if arlr.NextLink == nil || len(to.String(arlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(arlr.NextLink)))
}

// AuthorizationRuleListResultPage contains a page of AuthorizationRule values.
type AuthorizationRuleListResultPage struct {
	fn   func(AuthorizationRuleListResult) (AuthorizationRuleListResult, error)
	arlr AuthorizationRuleListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AuthorizationRuleListResultPage) Next() error {
	next, err := page.fn(page.arlr)
	if err != nil {
		return err
	}
	page.arlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AuthorizationRuleListResultPage) NotDone() bool {
	return !page.arlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AuthorizationRuleListResultPage) Response() AuthorizationRuleListResult {
	return page.arlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AuthorizationRuleListResultPage) Values() []AuthorizationRule {
	if page.arlr.IsEmpty() {
		return nil
	}
	return *page.arlr.Value
}

// AuthorizationRuleProperties authorization rule properties.
type AuthorizationRuleProperties struct {
	// Rights - The rights associated with the rule.
	Rights *[]AccessRights `json:"rights,omitempty"`
}

// CheckNameAvailability description of the check name availability request properties.
type CheckNameAvailability struct {
	// Name - The namespace name to check for availability. The namespace name can contain only letters, numbers, and hyphens. The namespace must start with a letter, and it must end with a letter or number.
	Name *string `json:"name,omitempty"`
}

// CheckNameAvailabilityResult description of the check name availability request properties.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// Message - The detailed info regarding the reason associated with the namespace.
	Message *string `json:"message,omitempty"`
	// NameAvailable - Value indicating namespace is available. Returns true if the namespace is available; otherwise, false.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason for unavailability of a namespace. Possible values include: 'None', 'InvalidName', 'SubscriptionIsDisabled', 'NameInUse', 'NameInLockdown', 'TooManyNamespaceInCurrentSubscription'
	Reason UnavailableReason `json:"reason,omitempty"`
}

// ErrorResponse error reponse indicates Relay service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponse struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// HybridConnection description of hybrid connection resource.
type HybridConnection struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// HybridConnectionProperties - Properties of the HybridConnection.
	*HybridConnectionProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for HybridConnection struct.
func (hc *HybridConnection) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var hybridConnectionProperties HybridConnectionProperties
				err = json.Unmarshal(*v, &hybridConnectionProperties)
				if err != nil {
					return err
				}
				hc.HybridConnectionProperties = &hybridConnectionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				hc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				hc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				hc.Type = &typeVar
			}
		}
	}

	return nil
}

// HybridConnectionListResult the response of the list hybrid connection operation.
type HybridConnectionListResult struct {
	autorest.Response `json:"-"`
	// Value - Result of the list hybrid connections.
	Value *[]HybridConnection `json:"value,omitempty"`
	// NextLink - Link to the next set of results. Not empty if value contains incomplete list hybrid connection operation.
	NextLink *string `json:"nextLink,omitempty"`
}

// HybridConnectionListResultIterator provides access to a complete listing of HybridConnection values.
type HybridConnectionListResultIterator struct {
	i    int
	page HybridConnectionListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *HybridConnectionListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter HybridConnectionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter HybridConnectionListResultIterator) Response() HybridConnectionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter HybridConnectionListResultIterator) Value() HybridConnection {
	if !iter.page.NotDone() {
		return HybridConnection{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (hclr HybridConnectionListResult) IsEmpty() bool {
	return hclr.Value == nil || len(*hclr.Value) == 0
}

// hybridConnectionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (hclr HybridConnectionListResult) hybridConnectionListResultPreparer() (*http.Request, error) {
	if hclr.NextLink == nil || len(to.String(hclr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(hclr.NextLink)))
}

// HybridConnectionListResultPage contains a page of HybridConnection values.
type HybridConnectionListResultPage struct {
	fn   func(HybridConnectionListResult) (HybridConnectionListResult, error)
	hclr HybridConnectionListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *HybridConnectionListResultPage) Next() error {
	next, err := page.fn(page.hclr)
	if err != nil {
		return err
	}
	page.hclr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page HybridConnectionListResultPage) NotDone() bool {
	return !page.hclr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page HybridConnectionListResultPage) Response() HybridConnectionListResult {
	return page.hclr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page HybridConnectionListResultPage) Values() []HybridConnection {
	if page.hclr.IsEmpty() {
		return nil
	}
	return *page.hclr.Value
}

// HybridConnectionProperties properties of the HybridConnection.
type HybridConnectionProperties struct {
	// CreatedAt - The time the hybrid connection was created.
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// UpdatedAt - The time the namespace was updated.
	UpdatedAt *date.Time `json:"updatedAt,omitempty"`
	// ListenerCount - The number of listeners for this hybrid connection. Note that min : 1 and max:25 are supported.
	ListenerCount *int32 `json:"listenerCount,omitempty"`
	// RequiresClientAuthorization - Returns true if client authorization is needed for this hybrid connection; otherwise, false.
	RequiresClientAuthorization *bool `json:"requiresClientAuthorization,omitempty"`
	// UserMetadata - The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}

// Namespace description of a namespace resource.
type Namespace struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Sku - SKU of the namespace.
	Sku *Sku `json:"sku,omitempty"`
	// NamespaceProperties - Description of Relay namespace
	*NamespaceProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Namespace.
func (n Namespace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if n.Sku != nil {
		objectMap["sku"] = n.Sku
	}
	if n.NamespaceProperties != nil {
		objectMap["properties"] = n.NamespaceProperties
	}
	if n.Location != nil {
		objectMap["location"] = n.Location
	}
	if n.Tags != nil {
		objectMap["tags"] = n.Tags
	}
	if n.ID != nil {
		objectMap["id"] = n.ID
	}
	if n.Name != nil {
		objectMap["name"] = n.Name
	}
	if n.Type != nil {
		objectMap["type"] = n.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Namespace struct.
func (n *Namespace) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				n.Sku = &sku
			}
		case "properties":
			if v != nil {
				var namespaceProperties NamespaceProperties
				err = json.Unmarshal(*v, &namespaceProperties)
				if err != nil {
					return err
				}
				n.NamespaceProperties = &namespaceProperties
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				n.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				n.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				n.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				n.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				n.Type = &typeVar
			}
		}
	}

	return nil
}

// NamespaceListResult the response from the list namespace operation.
type NamespaceListResult struct {
	autorest.Response `json:"-"`
	// Value - Result of the list namespace operation.
	Value *[]Namespace `json:"value,omitempty"`
	// NextLink - Link to the next set of results. Not empty if value contains incomplete list of namespaces.
	NextLink *string `json:"nextLink,omitempty"`
}

// NamespaceListResultIterator provides access to a complete listing of Namespace values.
type NamespaceListResultIterator struct {
	i    int
	page NamespaceListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *NamespaceListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter NamespaceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter NamespaceListResultIterator) Response() NamespaceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter NamespaceListResultIterator) Value() Namespace {
	if !iter.page.NotDone() {
		return Namespace{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (nlr NamespaceListResult) IsEmpty() bool {
	return nlr.Value == nil || len(*nlr.Value) == 0
}

// namespaceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (nlr NamespaceListResult) namespaceListResultPreparer() (*http.Request, error) {
	if nlr.NextLink == nil || len(to.String(nlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(nlr.NextLink)))
}

// NamespaceListResultPage contains a page of Namespace values.
type NamespaceListResultPage struct {
	fn  func(NamespaceListResult) (NamespaceListResult, error)
	nlr NamespaceListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *NamespaceListResultPage) Next() error {
	next, err := page.fn(page.nlr)
	if err != nil {
		return err
	}
	page.nlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page NamespaceListResultPage) NotDone() bool {
	return !page.nlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page NamespaceListResultPage) Response() NamespaceListResult {
	return page.nlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page NamespaceListResultPage) Values() []Namespace {
	if page.nlr.IsEmpty() {
		return nil
	}
	return *page.nlr.Value
}

// NamespaceProperties properties of the namespace.
type NamespaceProperties struct {
	// ProvisioningState - Possible values include: 'Created', 'Succeeded', 'Deleted', 'Failed', 'Updating', 'Unknown'
	ProvisioningState ProvisioningStateEnum `json:"provisioningState,omitempty"`
	// CreatedAt - The time the namespace was created.
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// UpdatedAt - The time the namespace was updated.
	UpdatedAt *date.Time `json:"updatedAt,omitempty"`
	// ServiceBusEndpoint - Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `json:"serviceBusEndpoint,omitempty"`
	// MetricID - Identifier for Azure Insights metrics.
	MetricID *string `json:"metricId,omitempty"`
}

// NamespacesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type NamespacesCreateOrUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future NamespacesCreateOrUpdateFuture) Result(client NamespacesClient) (n Namespace, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return n, autorest.NewError("relay.NamespacesCreateOrUpdateFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		n, err = client.CreateOrUpdateResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	n, err = client.CreateOrUpdateResponder(resp)
	return
}

// NamespacesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type NamespacesDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future NamespacesDeleteFuture) Result(client NamespacesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ar, autorest.NewError("relay.NamespacesDeleteFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ar, err = client.DeleteResponder(resp)
	return
}

// Operation a Relay REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Relay.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list Relay operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Relay operations supported by resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// RegenerateAccessKeyParameters parameters supplied to the regenerate authorization rule operation, specifies
// which key neeeds to be reset.
type RegenerateAccessKeyParameters struct {
	// KeyType - The access key to regenerate. Possible values include: 'PrimaryKey', 'SecondaryKey'
	KeyType KeyType `json:"keyType,omitempty"`
	// Key - Optional. If the key value is provided, this is set to key type, or autogenerated key value set for key type.
	Key *string `json:"key,omitempty"`
}

// Resource the resource definition.
type Resource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// ResourceNamespacePatch definition of resource.
type ResourceNamespacePatch struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ResourceNamespacePatch.
func (rnp ResourceNamespacePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rnp.Tags != nil {
		objectMap["tags"] = rnp.Tags
	}
	if rnp.ID != nil {
		objectMap["id"] = rnp.ID
	}
	if rnp.Name != nil {
		objectMap["name"] = rnp.Name
	}
	if rnp.Type != nil {
		objectMap["type"] = rnp.Type
	}
	return json.Marshal(objectMap)
}

// Sku SKU of the namespace.
type Sku struct {
	// Name - Name of this SKU.
	Name *string `json:"name,omitempty"`
	// Tier - The tier of this SKU. Possible values include: 'Standard'
	Tier SkuTier `json:"tier,omitempty"`
}

// TrackedResource definition of resource.
type TrackedResource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.ID != nil {
		objectMap["id"] = tr.ID
	}
	if tr.Name != nil {
		objectMap["name"] = tr.Name
	}
	if tr.Type != nil {
		objectMap["type"] = tr.Type
	}
	return json.Marshal(objectMap)
}

// UpdateParameters description of a namespace resource.
type UpdateParameters struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Sku - SKU of the namespace.
	Sku *Sku `json:"sku,omitempty"`
	// NamespaceProperties - Description of Relay namespace.
	*NamespaceProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for UpdateParameters.
func (up UpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if up.Sku != nil {
		objectMap["sku"] = up.Sku
	}
	if up.NamespaceProperties != nil {
		objectMap["properties"] = up.NamespaceProperties
	}
	if up.Tags != nil {
		objectMap["tags"] = up.Tags
	}
	if up.ID != nil {
		objectMap["id"] = up.ID
	}
	if up.Name != nil {
		objectMap["name"] = up.Name
	}
	if up.Type != nil {
		objectMap["type"] = up.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for UpdateParameters struct.
func (up *UpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				up.Sku = &sku
			}
		case "properties":
			if v != nil {
				var namespaceProperties NamespaceProperties
				err = json.Unmarshal(*v, &namespaceProperties)
				if err != nil {
					return err
				}
				up.NamespaceProperties = &namespaceProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				up.Tags = tags
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				up.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				up.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				up.Type = &typeVar
			}
		}
	}

	return nil
}

// WcfRelay description of the WCF relay resource.
type WcfRelay struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// WcfRelayProperties - Properties of the WCF relay.
	*WcfRelayProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for WcfRelay struct.
func (wr *WcfRelay) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var wcfRelayProperties WcfRelayProperties
				err = json.Unmarshal(*v, &wcfRelayProperties)
				if err != nil {
					return err
				}
				wr.WcfRelayProperties = &wcfRelayProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				wr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				wr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				wr.Type = &typeVar
			}
		}
	}

	return nil
}

// WcfRelayProperties properties of the WCF relay.
type WcfRelayProperties struct {
	// IsDynamic - Returns true if the relay is dynamic; otherwise, false.
	IsDynamic *bool `json:"isDynamic,omitempty"`
	// CreatedAt - The time the WCF relay was created.
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// UpdatedAt - The time the namespace was updated.
	UpdatedAt *date.Time `json:"updatedAt,omitempty"`
	// ListenerCount - The number of listeners for this relay. Note that min :1 and max:25 are supported.
	ListenerCount *int32 `json:"listenerCount,omitempty"`
	// RelayType - WCF relay type. Possible values include: 'NetTCP', 'HTTP'
	RelayType RelaytypeEnum `json:"relayType,omitempty"`
	// RequiresClientAuthorization - Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization *bool `json:"requiresClientAuthorization,omitempty"`
	// RequiresTransportSecurity - Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity *bool `json:"requiresTransportSecurity,omitempty"`
	// UserMetadata - The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}

// WcfRelaysListResult the response of the list WCF relay operation.
type WcfRelaysListResult struct {
	autorest.Response `json:"-"`
	// Value - Result of the list WCF relay operation.
	Value *[]WcfRelay `json:"value,omitempty"`
	// NextLink - Link to the next set of results. Not empty if value contains incomplete list of WCF relays.
	NextLink *string `json:"nextLink,omitempty"`
}

// WcfRelaysListResultIterator provides access to a complete listing of WcfRelay values.
type WcfRelaysListResultIterator struct {
	i    int
	page WcfRelaysListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *WcfRelaysListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter WcfRelaysListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter WcfRelaysListResultIterator) Response() WcfRelaysListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter WcfRelaysListResultIterator) Value() WcfRelay {
	if !iter.page.NotDone() {
		return WcfRelay{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (wrlr WcfRelaysListResult) IsEmpty() bool {
	return wrlr.Value == nil || len(*wrlr.Value) == 0
}

// wcfRelaysListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (wrlr WcfRelaysListResult) wcfRelaysListResultPreparer() (*http.Request, error) {
	if wrlr.NextLink == nil || len(to.String(wrlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(wrlr.NextLink)))
}

// WcfRelaysListResultPage contains a page of WcfRelay values.
type WcfRelaysListResultPage struct {
	fn   func(WcfRelaysListResult) (WcfRelaysListResult, error)
	wrlr WcfRelaysListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *WcfRelaysListResultPage) Next() error {
	next, err := page.fn(page.wrlr)
	if err != nil {
		return err
	}
	page.wrlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page WcfRelaysListResultPage) NotDone() bool {
	return !page.wrlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page WcfRelaysListResultPage) Response() WcfRelaysListResult {
	return page.wrlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page WcfRelaysListResultPage) Values() []WcfRelay {
	if page.wrlr.IsEmpty() {
		return nil
	}
	return *page.wrlr.Value
}
