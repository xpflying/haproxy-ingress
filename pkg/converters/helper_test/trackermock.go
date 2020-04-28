/*
Copyright 2020 The HAProxy Ingress Controller Authors.

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

package helper_test

import (
	convtypes "github.com/jcmoraisjr/haproxy-ingress/pkg/converters/types"
	hatypes "github.com/jcmoraisjr/haproxy-ingress/pkg/haproxy/types"
)

// TrackerMock ...
type TrackerMock struct{}

// NewTrackerMock ...
func NewTrackerMock() *TrackerMock {
	return &TrackerMock{}
}

// TrackHostname ...
func (t *TrackerMock) TrackHostname(rtype convtypes.ResourceType, name, hostname string) {
}

// TrackBackend ...
func (t *TrackerMock) TrackBackend(rtype convtypes.ResourceType, name string, backendID hatypes.BackendID) {
}

// TrackMissingOnHostname ...
func (t *TrackerMock) TrackMissingOnHostname(rtype convtypes.ResourceType, name, hostname string) {
}

// GetDirtyLinks ...
func (t *TrackerMock) GetDirtyLinks(oldIngList, oldServiceList, addServiceList, oldSecretList, addSecretList []string) (dirtyIngs, dirtyHosts []string, dirtyBacks []hatypes.BackendID) {
	return nil, nil, nil
}

// DeleteHostnames ...
func (t *TrackerMock) DeleteHostnames(hostnames []string) {
}

// DeleteBackends ...
func (t *TrackerMock) DeleteBackends(backends []hatypes.BackendID) {
}
