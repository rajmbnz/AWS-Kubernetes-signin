// Copyright 2024 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the 'License');
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an 'AS IS' BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud/interface.go

// Package cloud is a generated GoMock package.
package cloud

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	gomock "github.com/golang/mock/gomock"
)

// MockCloud is a mock of Cloud interface.
type MockCloud struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMockRecorder
}

// MockCloudMockRecorder is the mock recorder for MockCloud.
type MockCloudMockRecorder struct {
	mock *MockCloud
}

// NewMockCloud creates a new mock instance.
func NewMockCloud(ctrl *gomock.Controller) *MockCloud {
	mock := &MockCloud{ctrl: ctrl}
	mock.recorder = &MockCloudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloud) EXPECT() *MockCloudMockRecorder {
	return m.recorder
}

// AttachDisk mocks base method.
func (m *MockCloud) AttachDisk(ctx context.Context, volumeID, nodeID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachDisk", ctx, volumeID, nodeID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachDisk indicates an expected call of AttachDisk.
func (mr *MockCloudMockRecorder) AttachDisk(ctx, volumeID, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachDisk", reflect.TypeOf((*MockCloud)(nil).AttachDisk), ctx, volumeID, nodeID)
}

// AvailabilityZones mocks base method.
func (m *MockCloud) AvailabilityZones(ctx context.Context) (map[string]struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityZones", ctx)
	ret0, _ := ret[0].(map[string]struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailabilityZones indicates an expected call of AvailabilityZones.
func (mr *MockCloudMockRecorder) AvailabilityZones(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityZones", reflect.TypeOf((*MockCloud)(nil).AvailabilityZones), ctx)
}

// CreateDisk mocks base method.
func (m *MockCloud) CreateDisk(ctx context.Context, volumeName string, diskOptions *DiskOptions) (*Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDisk", ctx, volumeName, diskOptions)
	ret0, _ := ret[0].(*Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDisk indicates an expected call of CreateDisk.
func (mr *MockCloudMockRecorder) CreateDisk(ctx, volumeName, diskOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDisk", reflect.TypeOf((*MockCloud)(nil).CreateDisk), ctx, volumeName, diskOptions)
}

// CreateSnapshot mocks base method.
func (m *MockCloud) CreateSnapshot(ctx context.Context, volumeID string, snapshotOptions *SnapshotOptions) (*Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", ctx, volumeID, snapshotOptions)
	ret0, _ := ret[0].(*Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockCloudMockRecorder) CreateSnapshot(ctx, volumeID, snapshotOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockCloud)(nil).CreateSnapshot), ctx, volumeID, snapshotOptions)
}

// DeleteDisk mocks base method.
func (m *MockCloud) DeleteDisk(ctx context.Context, volumeID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDisk", ctx, volumeID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDisk indicates an expected call of DeleteDisk.
func (mr *MockCloudMockRecorder) DeleteDisk(ctx, volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDisk", reflect.TypeOf((*MockCloud)(nil).DeleteDisk), ctx, volumeID)
}

// DeleteSnapshot mocks base method.
func (m *MockCloud) DeleteSnapshot(ctx context.Context, snapshotID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", ctx, snapshotID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockCloudMockRecorder) DeleteSnapshot(ctx, snapshotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockCloud)(nil).DeleteSnapshot), ctx, snapshotID)
}

// DetachDisk mocks base method.
func (m *MockCloud) DetachDisk(ctx context.Context, volumeID, nodeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachDisk", ctx, volumeID, nodeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachDisk indicates an expected call of DetachDisk.
func (mr *MockCloudMockRecorder) DetachDisk(ctx, volumeID, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachDisk", reflect.TypeOf((*MockCloud)(nil).DetachDisk), ctx, volumeID, nodeID)
}

// EnableFastSnapshotRestores mocks base method.
func (m *MockCloud) EnableFastSnapshotRestores(ctx context.Context, availabilityZones []string, snapshotID string) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableFastSnapshotRestores", ctx, availabilityZones, snapshotID)
	ret0, _ := ret[0].(*ec2.EnableFastSnapshotRestoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableFastSnapshotRestores indicates an expected call of EnableFastSnapshotRestores.
func (mr *MockCloudMockRecorder) EnableFastSnapshotRestores(ctx, availabilityZones, snapshotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableFastSnapshotRestores", reflect.TypeOf((*MockCloud)(nil).EnableFastSnapshotRestores), ctx, availabilityZones, snapshotID)
}

// GetDiskByID mocks base method.
func (m *MockCloud) GetDiskByID(ctx context.Context, volumeID string) (*Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskByID", ctx, volumeID)
	ret0, _ := ret[0].(*Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskByID indicates an expected call of GetDiskByID.
func (mr *MockCloudMockRecorder) GetDiskByID(ctx, volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskByID", reflect.TypeOf((*MockCloud)(nil).GetDiskByID), ctx, volumeID)
}

// GetDiskByName mocks base method.
func (m *MockCloud) GetDiskByName(ctx context.Context, name string, capacityBytes int64) (*Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskByName", ctx, name, capacityBytes)
	ret0, _ := ret[0].(*Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskByName indicates an expected call of GetDiskByName.
func (mr *MockCloudMockRecorder) GetDiskByName(ctx, name, capacityBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskByName", reflect.TypeOf((*MockCloud)(nil).GetDiskByName), ctx, name, capacityBytes)
}

// GetSnapshotByID mocks base method.
func (m *MockCloud) GetSnapshotByID(ctx context.Context, snapshotID string) (*Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotByID", ctx, snapshotID)
	ret0, _ := ret[0].(*Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotByID indicates an expected call of GetSnapshotByID.
func (mr *MockCloudMockRecorder) GetSnapshotByID(ctx, snapshotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotByID", reflect.TypeOf((*MockCloud)(nil).GetSnapshotByID), ctx, snapshotID)
}

// GetSnapshotByName mocks base method.
func (m *MockCloud) GetSnapshotByName(ctx context.Context, name string) (*Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotByName", ctx, name)
	ret0, _ := ret[0].(*Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotByName indicates an expected call of GetSnapshotByName.
func (mr *MockCloudMockRecorder) GetSnapshotByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotByName", reflect.TypeOf((*MockCloud)(nil).GetSnapshotByName), ctx, name)
}

// ListSnapshots mocks base method.
func (m *MockCloud) ListSnapshots(ctx context.Context, volumeID string, maxResults int32, nextToken string) (*ListSnapshotsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", ctx, volumeID, maxResults, nextToken)
	ret0, _ := ret[0].(*ListSnapshotsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots.
func (mr *MockCloudMockRecorder) ListSnapshots(ctx, volumeID, maxResults, nextToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockCloud)(nil).ListSnapshots), ctx, volumeID, maxResults, nextToken)
}

// ResizeOrModifyDisk mocks base method.
func (m *MockCloud) ResizeOrModifyDisk(ctx context.Context, volumeID string, newSizeBytes int64, options *ModifyDiskOptions) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeOrModifyDisk", ctx, volumeID, newSizeBytes, options)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResizeOrModifyDisk indicates an expected call of ResizeOrModifyDisk.
func (mr *MockCloudMockRecorder) ResizeOrModifyDisk(ctx, volumeID, newSizeBytes, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeOrModifyDisk", reflect.TypeOf((*MockCloud)(nil).ResizeOrModifyDisk), ctx, volumeID, newSizeBytes, options)
}

// WaitForAttachmentState mocks base method.
func (m *MockCloud) WaitForAttachmentState(ctx context.Context, volumeID, expectedState, expectedInstance, expectedDevice string, alreadyAssigned bool) (*types.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAttachmentState", ctx, volumeID, expectedState, expectedInstance, expectedDevice, alreadyAssigned)
	ret0, _ := ret[0].(*types.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForAttachmentState indicates an expected call of WaitForAttachmentState.
func (mr *MockCloudMockRecorder) WaitForAttachmentState(ctx, volumeID, expectedState, expectedInstance, expectedDevice, alreadyAssigned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAttachmentState", reflect.TypeOf((*MockCloud)(nil).WaitForAttachmentState), ctx, volumeID, expectedState, expectedInstance, expectedDevice, alreadyAssigned)
}
