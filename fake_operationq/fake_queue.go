// This file was generated by counterfeiter
package fake_operationq

import (
	"sync"

	"github.com/pivotal-golang/operationq"
)

type FakeQueue struct {
	PushStub        func(operationq.Operation)
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		arg1 operationq.Operation
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	WaitStub         func()
	waitMutex        sync.RWMutex
	waitArgsForCall  []struct{}
}

func (fake *FakeQueue) Push(arg1 operationq.Operation) {
	fake.pushMutex.Lock()
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		arg1 operationq.Operation
	}{arg1})
	fake.pushMutex.Unlock()
	if fake.PushStub != nil {
		fake.PushStub(arg1)
	}
}

func (fake *FakeQueue) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeQueue) PushArgsForCall(i int) operationq.Operation {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return fake.pushArgsForCall[i].arg1
}

func (fake *FakeQueue) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeQueue) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeQueue) Wait() {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		fake.WaitStub()
	}
}

func (fake *FakeQueue) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

var _ operationq.Queue = new(FakeQueue)
