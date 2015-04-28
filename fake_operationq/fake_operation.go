// This file was generated by counterfeiter
package fake_operationq

import (
	"sync"

	"github.com/pivotal-golang/operationq"
)

type FakeOperation struct {
	KeyStub        func() string
	keyMutex       sync.RWMutex
	keyArgsForCall []struct{}
	keyReturns     struct {
		result1 string
	}
	ExecuteStub        func()
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
}

func (fake *FakeOperation) Key() string {
	fake.keyMutex.Lock()
	fake.keyArgsForCall = append(fake.keyArgsForCall, struct{}{})
	fake.keyMutex.Unlock()
	if fake.KeyStub != nil {
		return fake.KeyStub()
	} else {
		return fake.keyReturns.result1
	}
}

func (fake *FakeOperation) KeyCallCount() int {
	fake.keyMutex.RLock()
	defer fake.keyMutex.RUnlock()
	return len(fake.keyArgsForCall)
}

func (fake *FakeOperation) KeyReturns(result1 string) {
	fake.KeyStub = nil
	fake.keyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeOperation) Execute() {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		fake.ExecuteStub()
	}
}

func (fake *FakeOperation) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

var _ operationq.Operation = new(FakeOperation)
