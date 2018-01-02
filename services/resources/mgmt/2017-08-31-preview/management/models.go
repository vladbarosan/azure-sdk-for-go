package management

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/go.uuid"
	"net/http"
)

// ChildType enumerates the values for child type.
type ChildType string

const (
	// Account ...
	Account ChildType = "Account"
	// Department ...
	Department ChildType = "Department"
	// Enrollment ...
	Enrollment ChildType = "Enrollment"
	// Subscription ...
	Subscription ChildType = "Subscription"
)

// ChildType1 enumerates the values for child type 1.
type ChildType1 string

const (
	// ChildType1Account ...
	ChildType1Account ChildType1 = "Account"
	// ChildType1Department ...
	ChildType1Department ChildType1 = "Department"
	// ChildType1Enrollment ...
	ChildType1Enrollment ChildType1 = "Enrollment"
	// ChildType1Subscription ...
	ChildType1Subscription ChildType1 = "Subscription"
)

// ManagementGroupType enumerates the values for management group type.
type ManagementGroupType string

const (
	// ManagementGroupTypeAccount ...
	ManagementGroupTypeAccount ManagementGroupType = "Account"
	// ManagementGroupTypeDepartment ...
	ManagementGroupTypeDepartment ManagementGroupType = "Department"
	// ManagementGroupTypeEnrollment ...
	ManagementGroupTypeEnrollment ManagementGroupType = "Enrollment"
	// ManagementGroupTypeSubscription ...
	ManagementGroupTypeSubscription ManagementGroupType = "Subscription"
)

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - One of a server-defined set of error codes.
	Code *string `json:"code,omitempty"`
	// Message - A human-readable representation of the error.
	Message *string `json:"message,omitempty"`
	// Target - (Optional) The target of the error.
	Target *string `json:"target,omitempty"`
}

// ErrorResponse the error object.
type ErrorResponse struct {
	Error *ErrorDetails `json:"error,omitempty"`
}

// Group the management group details.
type Group struct {
	// ID - The ID of the management group. E.g. /providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`
	// Type - The type of the resource. E.g. /providers/Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty"`
	// Name - The name of the management group. E.g. 20000000-0000-0000-0000-000000000000
	Name             *uuid.UUID `json:"name,omitempty"`
	*GroupProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for Group struct.
func (g *Group) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				g.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				g.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name uuid.UUID
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				g.Name = &name
			}
		case "properties":
			if v != nil {
				var groupProperties GroupProperties
				err = json.Unmarshal(*v, &groupProperties)
				if err != nil {
					return err
				}
				g.GroupProperties = &groupProperties
			}
		}
	}

	return nil
}

// GroupChildInfo the unique identifier (ID) of a management group.
type GroupChildInfo struct {
	// ChildType - Possible values include: 'Enrollment', 'Department', 'Account', 'Subscription'
	ChildType ChildType `json:"childType,omitempty"`
	// ChildID - The ID of the child resource (management group or subscription). E.g. /providers/Microsoft.Management/managementGroups/40000000-0000-0000-0000-000000000000
	ChildID *string `json:"childId,omitempty"`
	// DisplayName - The friendly name of the child resource.
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID - (Optional) The AAD Tenant ID associated with the child resource.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
}

// GroupDetailsProperties the details properties of a management group.
type GroupDetailsProperties struct {
	// Version - The version number of the object.
	Version *float64 `json:"version,omitempty"`
	// UpdatedTime - The date and time when this object was last updated.
	UpdatedTime *date.Time `json:"updatedTime,omitempty"`
	// UpdatedBy - The identity of the principal or process that updated the object.
	UpdatedBy *string          `json:"updatedBy,omitempty"`
	Parent    *ParentGroupInfo `json:"parent,omitempty"`
	// ManagementGroupType - Possible values include: 'ManagementGroupTypeEnrollment', 'ManagementGroupTypeDepartment', 'ManagementGroupTypeAccount', 'ManagementGroupTypeSubscription'
	ManagementGroupType ManagementGroupType `json:"managementGroupType,omitempty"`
}

// GroupInfo the management group.
type GroupInfo struct {
	// ID - The ID of the management group. E.g. /providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`
	// Type - The type of the resource. E.g. /providers/Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty"`
	// Name - The name of the management group. E.g. 20000000-0000-0000-0000-000000000000
	Name                 *uuid.UUID `json:"name,omitempty"`
	*GroupInfoProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for GroupInfo struct.
func (gi *GroupInfo) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				gi.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				gi.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name uuid.UUID
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				gi.Name = &name
			}
		case "properties":
			if v != nil {
				var groupInfoProperties GroupInfoProperties
				err = json.Unmarshal(*v, &groupInfoProperties)
				if err != nil {
					return err
				}
				gi.GroupInfoProperties = &groupInfoProperties
			}
		}
	}

	return nil
}

// GroupInfoProperties the generic properties of a management group.
type GroupInfoProperties struct {
	// TenantID - The AAD Tenant ID associated with the management group. E.g. 10000000-0000-0000-0000-000000000000
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// DisplayName - The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`
}

// GroupListResult the result of listing management groups.
type GroupListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of management groups.
	Value *[]GroupInfo `json:"value,omitempty"`
	// NextLink - The URL to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// GroupListResultIterator provides access to a complete listing of GroupInfo values.
type GroupListResultIterator struct {
	i    int
	page GroupListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *GroupListResultIterator) Next() error {
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
func (iter GroupListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter GroupListResultIterator) Response() GroupListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter GroupListResultIterator) Value() GroupInfo {
	if !iter.page.NotDone() {
		return GroupInfo{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (glr GroupListResult) IsEmpty() bool {
	return glr.Value == nil || len(*glr.Value) == 0
}

// groupListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (glr GroupListResult) groupListResultPreparer() (*http.Request, error) {
	if glr.NextLink == nil || len(to.String(glr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(glr.NextLink)))
}

// GroupListResultPage contains a page of GroupInfo values.
type GroupListResultPage struct {
	fn  func(GroupListResult) (GroupListResult, error)
	glr GroupListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *GroupListResultPage) Next() error {
	next, err := page.fn(page.glr)
	if err != nil {
		return err
	}
	page.glr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page GroupListResultPage) NotDone() bool {
	return !page.glr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page GroupListResultPage) Response() GroupListResult {
	return page.glr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page GroupListResultPage) Values() []GroupInfo {
	if page.glr.IsEmpty() {
		return nil
	}
	return *page.glr.Value
}

// GroupProperties the generic properties of a management group.
type GroupProperties struct {
	// TenantID - The AAD Tenant ID associated with the management group. E.g. 10000000-0000-0000-0000-000000000000
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// DisplayName - The friendly name of the management group.
	DisplayName *string                 `json:"displayName,omitempty"`
	Details     *GroupDetailsProperties `json:"details,omitempty"`
}

// GroupPropertiesWithChildren the generic properties of a management group.
type GroupPropertiesWithChildren struct {
	// TenantID - The AAD Tenant ID associated with the management group. E.g. 10000000-0000-0000-0000-000000000000
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// DisplayName - The friendly name of the management group.
	DisplayName *string                 `json:"displayName,omitempty"`
	Details     *GroupDetailsProperties `json:"details,omitempty"`
	// Children - The list of children.
	Children *[]GroupChildInfo `json:"children,omitempty"`
}

// GroupPropertiesWithHierarchy the generic properties of a management group.
type GroupPropertiesWithHierarchy struct {
	// TenantID - The AAD Tenant ID associated with the management group. E.g. 10000000-0000-0000-0000-000000000000
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// DisplayName - The friendly name of the management group.
	DisplayName *string                 `json:"displayName,omitempty"`
	Details     *GroupDetailsProperties `json:"details,omitempty"`
	// Children - The list of children.
	Children *[]GroupRecursiveChildInfo `json:"children,omitempty"`
}

// GroupRecursiveChildInfo the unique identifier (ID) of a management group.
type GroupRecursiveChildInfo struct {
	// ChildType - Possible values include: 'ChildType1Enrollment', 'ChildType1Department', 'ChildType1Account', 'ChildType1Subscription'
	ChildType ChildType `json:"childType,omitempty"`
	// ChildID - The ID of the child resource (management group or subscription). E.g. /providers/Microsoft.Management/managementGroups/40000000-0000-0000-0000-000000000000
	ChildID *string `json:"childId,omitempty"`
	// DisplayName - The friendly name of the child resource.
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID - (Optional) The AAD Tenant ID associated with the child resource.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// Children - The list of children.
	Children *[]GroupRecursiveChildInfo `json:"children,omitempty"`
}

// GroupWithChildren the management group details.
type GroupWithChildren struct {
	// ID - The ID of the management group. E.g. /providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`
	// Type - The type of the resource. E.g. /providers/Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty"`
	// Name - The name of the management group. E.g. 20000000-0000-0000-0000-000000000000
	Name                         *uuid.UUID `json:"name,omitempty"`
	*GroupPropertiesWithChildren `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for GroupWithChildren struct.
func (gwc *GroupWithChildren) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				gwc.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				gwc.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name uuid.UUID
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				gwc.Name = &name
			}
		case "properties":
			if v != nil {
				var groupPropertiesWithChildren GroupPropertiesWithChildren
				err = json.Unmarshal(*v, &groupPropertiesWithChildren)
				if err != nil {
					return err
				}
				gwc.GroupPropertiesWithChildren = &groupPropertiesWithChildren
			}
		}
	}

	return nil
}

// GroupWithHierarchy the management group details.
type GroupWithHierarchy struct {
	autorest.Response `json:"-"`
	// ID - The ID of the management group. E.g. /providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000
	ID *string `json:"id,omitempty"`
	// Type - The type of the resource. E.g. /providers/Microsoft.Management/managementGroups
	Type *string `json:"type,omitempty"`
	// Name - The name of the management group. E.g. 20000000-0000-0000-0000-000000000000
	Name                          *uuid.UUID `json:"name,omitempty"`
	*GroupPropertiesWithHierarchy `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for GroupWithHierarchy struct.
func (gwh *GroupWithHierarchy) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				gwh.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				gwh.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name uuid.UUID
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				gwh.Name = &name
			}
		case "properties":
			if v != nil {
				var groupPropertiesWithHierarchy GroupPropertiesWithHierarchy
				err = json.Unmarshal(*v, &groupPropertiesWithHierarchy)
				if err != nil {
					return err
				}
				gwh.GroupPropertiesWithHierarchy = &groupPropertiesWithHierarchy
			}
		}
	}

	return nil
}

// Operation a Management REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Management.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result listing  operations. It contains a list of operations and a URL link to get the next
// set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of management operations supported by the Microsoft.Management resource provider.
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

// ParentGroupInfo (Optional) The ID of the parent management group.
type ParentGroupInfo struct {
	// ParentID - The ID of the parent management group. E.g. /providers/Microsoft.Management/managementGroups/30000000-0000-0000-0000-000000000000
	ParentID *string `json:"parentId,omitempty"`
	// DisplayName - The friendly name of the management group.
	DisplayName *string `json:"displayName,omitempty"`
}
