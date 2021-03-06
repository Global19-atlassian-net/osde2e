/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	time "time"
)

// AddOnInstallationKind is the name of the type used to represent objects
// of type 'add_on_installation'.
const AddOnInstallationKind = "AddOnInstallation"

// AddOnInstallationLinkKind is the name of the type used to represent links
// to objects of type 'add_on_installation'.
const AddOnInstallationLinkKind = "AddOnInstallationLink"

// AddOnInstallationNilKind is the name of the type used to nil references
// to objects of type 'add_on_installation'.
const AddOnInstallationNilKind = "AddOnInstallationNil"

// AddOnInstallation represents the values of the 'add_on_installation' type.
//
// Representation of an add-on installation in a cluster.
type AddOnInstallation struct {
	id                *string
	href              *string
	link              bool
	addon             *AddOn
	cluster           *Cluster
	creationTimestamp *time.Time
	operatorVersion   *string
	parameters        *AddOnInstallationParameterList
	state             *AddOnInstallationState
	stateDescription  *string
	updatedTimestamp  *time.Time
}

// Kind returns the name of the type of the object.
func (o *AddOnInstallation) Kind() string {
	if o == nil {
		return AddOnInstallationNilKind
	}
	if o.link {
		return AddOnInstallationLinkKind
	}
	return AddOnInstallationKind
}

// ID returns the identifier of the object.
func (o *AddOnInstallation) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *AddOnInstallation) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *AddOnInstallation) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *AddOnInstallation) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *AddOnInstallation) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *AddOnInstallation) Empty() bool {
	return o == nil || (o.id == nil &&
		o.creationTimestamp == nil &&
		o.operatorVersion == nil &&
		o.parameters.Len() == 0 &&
		o.state == nil &&
		o.stateDescription == nil &&
		o.updatedTimestamp == nil &&
		true)
}

// Addon returns the value of the 'addon' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Link to add-on attached to this cluster.
func (o *AddOnInstallation) Addon() *AddOn {
	if o == nil {
		return nil
	}
	return o.addon
}

// GetAddon returns the value of the 'addon' attribute and
// a flag indicating if the attribute has a value.
//
// Link to add-on attached to this cluster.
func (o *AddOnInstallation) GetAddon() (value *AddOn, ok bool) {
	ok = o != nil && o.addon != nil
	if ok {
		value = o.addon
	}
	return
}

// Cluster returns the value of the 'cluster' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// ID used to identify the cluster that this add-on is attached to.
func (o *AddOnInstallation) Cluster() *Cluster {
	if o == nil {
		return nil
	}
	return o.cluster
}

// GetCluster returns the value of the 'cluster' attribute and
// a flag indicating if the attribute has a value.
//
// ID used to identify the cluster that this add-on is attached to.
func (o *AddOnInstallation) GetCluster() (value *Cluster, ok bool) {
	ok = o != nil && o.cluster != nil
	if ok {
		value = o.cluster
	}
	return
}

// CreationTimestamp returns the value of the 'creation_timestamp' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Date and time when the add-on was initially installed in the cluster.
func (o *AddOnInstallation) CreationTimestamp() time.Time {
	if o != nil && o.creationTimestamp != nil {
		return *o.creationTimestamp
	}
	return time.Time{}
}

// GetCreationTimestamp returns the value of the 'creation_timestamp' attribute and
// a flag indicating if the attribute has a value.
//
// Date and time when the add-on was initially installed in the cluster.
func (o *AddOnInstallation) GetCreationTimestamp() (value time.Time, ok bool) {
	ok = o != nil && o.creationTimestamp != nil
	if ok {
		value = *o.creationTimestamp
	}
	return
}

// OperatorVersion returns the value of the 'operator_version' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Version of the operator installed by the add-on.
func (o *AddOnInstallation) OperatorVersion() string {
	if o != nil && o.operatorVersion != nil {
		return *o.operatorVersion
	}
	return ""
}

// GetOperatorVersion returns the value of the 'operator_version' attribute and
// a flag indicating if the attribute has a value.
//
// Version of the operator installed by the add-on.
func (o *AddOnInstallation) GetOperatorVersion() (value string, ok bool) {
	ok = o != nil && o.operatorVersion != nil
	if ok {
		value = *o.operatorVersion
	}
	return
}

// Parameters returns the value of the 'parameters' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// List of add-on parameters for this add-on installation.
func (o *AddOnInstallation) Parameters() *AddOnInstallationParameterList {
	if o == nil {
		return nil
	}
	return o.parameters
}

// GetParameters returns the value of the 'parameters' attribute and
// a flag indicating if the attribute has a value.
//
// List of add-on parameters for this add-on installation.
func (o *AddOnInstallation) GetParameters() (value *AddOnInstallationParameterList, ok bool) {
	ok = o != nil && o.parameters != nil
	if ok {
		value = o.parameters
	}
	return
}

// State returns the value of the 'state' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Overall state of the add-on installation.
func (o *AddOnInstallation) State() AddOnInstallationState {
	if o != nil && o.state != nil {
		return *o.state
	}
	return AddOnInstallationState("")
}

// GetState returns the value of the 'state' attribute and
// a flag indicating if the attribute has a value.
//
// Overall state of the add-on installation.
func (o *AddOnInstallation) GetState() (value AddOnInstallationState, ok bool) {
	ok = o != nil && o.state != nil
	if ok {
		value = *o.state
	}
	return
}

// StateDescription returns the value of the 'state_description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Reason for the current State.
func (o *AddOnInstallation) StateDescription() string {
	if o != nil && o.stateDescription != nil {
		return *o.stateDescription
	}
	return ""
}

// GetStateDescription returns the value of the 'state_description' attribute and
// a flag indicating if the attribute has a value.
//
// Reason for the current State.
func (o *AddOnInstallation) GetStateDescription() (value string, ok bool) {
	ok = o != nil && o.stateDescription != nil
	if ok {
		value = *o.stateDescription
	}
	return
}

// UpdatedTimestamp returns the value of the 'updated_timestamp' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Date and time when the add-on installation information was last updated.
func (o *AddOnInstallation) UpdatedTimestamp() time.Time {
	if o != nil && o.updatedTimestamp != nil {
		return *o.updatedTimestamp
	}
	return time.Time{}
}

// GetUpdatedTimestamp returns the value of the 'updated_timestamp' attribute and
// a flag indicating if the attribute has a value.
//
// Date and time when the add-on installation information was last updated.
func (o *AddOnInstallation) GetUpdatedTimestamp() (value time.Time, ok bool) {
	ok = o != nil && o.updatedTimestamp != nil
	if ok {
		value = *o.updatedTimestamp
	}
	return
}

// AddOnInstallationListKind is the name of the type used to represent list of objects of
// type 'add_on_installation'.
const AddOnInstallationListKind = "AddOnInstallationList"

// AddOnInstallationListLinkKind is the name of the type used to represent links to list
// of objects of type 'add_on_installation'.
const AddOnInstallationListLinkKind = "AddOnInstallationListLink"

// AddOnInstallationNilKind is the name of the type used to nil lists of objects of
// type 'add_on_installation'.
const AddOnInstallationListNilKind = "AddOnInstallationListNil"

// AddOnInstallationList is a list of values of the 'add_on_installation' type.
type AddOnInstallationList struct {
	href  *string
	link  bool
	items []*AddOnInstallation
}

// Kind returns the name of the type of the object.
func (l *AddOnInstallationList) Kind() string {
	if l == nil {
		return AddOnInstallationListNilKind
	}
	if l.link {
		return AddOnInstallationListLinkKind
	}
	return AddOnInstallationListKind
}

// Link returns true iif this is a link.
func (l *AddOnInstallationList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *AddOnInstallationList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *AddOnInstallationList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *AddOnInstallationList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *AddOnInstallationList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AddOnInstallationList) Get(i int) *AddOnInstallation {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *AddOnInstallationList) Slice() []*AddOnInstallation {
	var slice []*AddOnInstallation
	if l == nil {
		slice = make([]*AddOnInstallation, 0)
	} else {
		slice = make([]*AddOnInstallation, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AddOnInstallationList) Each(f func(item *AddOnInstallation) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *AddOnInstallationList) Range(f func(index int, item *AddOnInstallation) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
