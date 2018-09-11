// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry-incubator/blockhead/pkg/containermanager"
)

type FakeContainerManager struct {
	ProvisionStub        func(ctx context.Context, cc *containermanager.ContainerConfig) error
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		ctx context.Context
		cc  *containermanager.ContainerConfig
	}
	provisionReturns struct {
		result1 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerManager) Provision(ctx context.Context, cc *containermanager.ContainerConfig) error {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		ctx context.Context
		cc  *containermanager.ContainerConfig
	}{ctx, cc})
	fake.recordInvocation("Provision", []interface{}{ctx, cc})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(ctx, cc)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.provisionReturns.result1
}

func (fake *FakeContainerManager) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeContainerManager) ProvisionArgsForCall(i int) (context.Context, *containermanager.ContainerConfig) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].ctx, fake.provisionArgsForCall[i].cc
}

func (fake *FakeContainerManager) ProvisionReturns(result1 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) ProvisionReturnsOnCall(i int, result1 error) {
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ containermanager.ContainerManager = new(FakeContainerManager)