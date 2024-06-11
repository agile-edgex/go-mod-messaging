// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	types "github.com/agile-edgex/go-mod-messaging/v3/pkg/types"
)

// MessageClient is an autogenerated mock type for the MessageClient type
type MessageClient struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *MessageClient) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *MessageClient) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: message, topic
func (_m *MessageClient) Publish(message types.MessageEnvelope, topic string) error {
	ret := _m.Called(message, topic)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.MessageEnvelope, string) error); ok {
		r0 = rf(message, topic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Request provides a mock function with given fields: message, requestTopic, responseTopicPrefix, timeout
func (_m *MessageClient) Request(message types.MessageEnvelope, requestTopic string, responseTopicPrefix string, timeout time.Duration) (*types.MessageEnvelope, error) {
	ret := _m.Called(message, requestTopic, responseTopicPrefix, timeout)

	var r0 *types.MessageEnvelope
	if rf, ok := ret.Get(0).(func(types.MessageEnvelope, string, string, time.Duration) *types.MessageEnvelope); ok {
		r0 = rf(message, requestTopic, responseTopicPrefix, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MessageEnvelope)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.MessageEnvelope, string, string, time.Duration) error); ok {
		r1 = rf(message, requestTopic, responseTopicPrefix, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: topics, messageErrors
func (_m *MessageClient) Subscribe(topics []types.TopicChannel, messageErrors chan error) error {
	ret := _m.Called(topics, messageErrors)

	var r0 error
	if rf, ok := ret.Get(0).(func([]types.TopicChannel, chan error) error); ok {
		r0 = rf(topics, messageErrors)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: topics
func (_m *MessageClient) Unsubscribe(topics ...string) error {
	_va := make([]interface{}, len(topics))
	for _i := range topics {
		_va[_i] = topics[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(topics...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMessageClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageClient creates a new instance of MessageClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageClient(t mockConstructorTestingTNewMessageClient) *MessageClient {
	mock := &MessageClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
