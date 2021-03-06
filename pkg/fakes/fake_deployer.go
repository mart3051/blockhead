// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/blockhead/pkg/containermanager"
	"github.com/cloudfoundry-incubator/blockhead/pkg/deployer"
)

type FakeDeployer struct {
	DeployContractStub        func(contractInfo *deployer.ContractInfo, containerInfo *containermanager.ContainerInfo, nodePort string) (*deployer.NodeInfo, error)
	deployContractMutex       sync.RWMutex
	deployContractArgsForCall []struct {
		contractInfo  *deployer.ContractInfo
		containerInfo *containermanager.ContainerInfo
		nodePort      string
	}
	deployContractReturns struct {
		result1 *deployer.NodeInfo
		result2 error
	}
	deployContractReturnsOnCall map[int]struct {
		result1 *deployer.NodeInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeployer) DeployContract(contractInfo *deployer.ContractInfo, containerInfo *containermanager.ContainerInfo, nodePort string) (*deployer.NodeInfo, error) {
	fake.deployContractMutex.Lock()
	ret, specificReturn := fake.deployContractReturnsOnCall[len(fake.deployContractArgsForCall)]
	fake.deployContractArgsForCall = append(fake.deployContractArgsForCall, struct {
		contractInfo  *deployer.ContractInfo
		containerInfo *containermanager.ContainerInfo
		nodePort      string
	}{contractInfo, containerInfo, nodePort})
	fake.recordInvocation("DeployContract", []interface{}{contractInfo, containerInfo, nodePort})
	fake.deployContractMutex.Unlock()
	if fake.DeployContractStub != nil {
		return fake.DeployContractStub(contractInfo, containerInfo, nodePort)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deployContractReturns.result1, fake.deployContractReturns.result2
}

func (fake *FakeDeployer) DeployContractCallCount() int {
	fake.deployContractMutex.RLock()
	defer fake.deployContractMutex.RUnlock()
	return len(fake.deployContractArgsForCall)
}

func (fake *FakeDeployer) DeployContractArgsForCall(i int) (*deployer.ContractInfo, *containermanager.ContainerInfo, string) {
	fake.deployContractMutex.RLock()
	defer fake.deployContractMutex.RUnlock()
	return fake.deployContractArgsForCall[i].contractInfo, fake.deployContractArgsForCall[i].containerInfo, fake.deployContractArgsForCall[i].nodePort
}

func (fake *FakeDeployer) DeployContractReturns(result1 *deployer.NodeInfo, result2 error) {
	fake.DeployContractStub = nil
	fake.deployContractReturns = struct {
		result1 *deployer.NodeInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployer) DeployContractReturnsOnCall(i int, result1 *deployer.NodeInfo, result2 error) {
	fake.DeployContractStub = nil
	if fake.deployContractReturnsOnCall == nil {
		fake.deployContractReturnsOnCall = make(map[int]struct {
			result1 *deployer.NodeInfo
			result2 error
		})
	}
	fake.deployContractReturnsOnCall[i] = struct {
		result1 *deployer.NodeInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deployContractMutex.RLock()
	defer fake.deployContractMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeployer) recordInvocation(key string, args []interface{}) {
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

var _ deployer.Deployer = new(FakeDeployer)
