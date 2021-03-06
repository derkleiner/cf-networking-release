// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/store"
	"sync"
)

type DestinationRepo struct {
	CreateStub        func(store.Transaction, int, int, int, int, string) (int, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
		arg4 int
		arg5 int
		arg6 string
	}
	createReturns struct {
		result1 int
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	DeleteStub        func(store.Transaction, int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetIDStub        func(store.Transaction, int, int, int, int, string) (int, error)
	getIDMutex       sync.RWMutex
	getIDArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
		arg4 int
		arg5 int
		arg6 string
	}
	getIDReturns struct {
		result1 int
		result2 error
	}
	getIDReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	CountWhereGroupIDStub        func(store.Transaction, int) (int, error)
	countWhereGroupIDMutex       sync.RWMutex
	countWhereGroupIDArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
	}
	countWhereGroupIDReturns struct {
		result1 int
		result2 error
	}
	countWhereGroupIDReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DestinationRepo) Create(arg1 store.Transaction, arg2 int, arg3 int, arg4 int, arg5 int, arg6 string) (int, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
		arg4 int
		arg5 int
		arg6 string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *DestinationRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *DestinationRepo) CreateArgsForCall(i int) (store.Transaction, int, int, int, int, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3, fake.createArgsForCall[i].arg4, fake.createArgsForCall[i].arg5, fake.createArgsForCall[i].arg6
}

func (fake *DestinationRepo) CreateReturns(result1 int, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DestinationRepo) CreateReturnsOnCall(i int, result1 int, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DestinationRepo) Delete(arg1 store.Transaction, arg2 int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *DestinationRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *DestinationRepo) DeleteArgsForCall(i int) (store.Transaction, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1, fake.deleteArgsForCall[i].arg2
}

func (fake *DestinationRepo) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *DestinationRepo) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DestinationRepo) GetID(arg1 store.Transaction, arg2 int, arg3 int, arg4 int, arg5 int, arg6 string) (int, error) {
	fake.getIDMutex.Lock()
	ret, specificReturn := fake.getIDReturnsOnCall[len(fake.getIDArgsForCall)]
	fake.getIDArgsForCall = append(fake.getIDArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
		arg4 int
		arg5 int
		arg6 string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("GetID", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.getIDMutex.Unlock()
	if fake.GetIDStub != nil {
		return fake.GetIDStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getIDReturns.result1, fake.getIDReturns.result2
}

func (fake *DestinationRepo) GetIDCallCount() int {
	fake.getIDMutex.RLock()
	defer fake.getIDMutex.RUnlock()
	return len(fake.getIDArgsForCall)
}

func (fake *DestinationRepo) GetIDArgsForCall(i int) (store.Transaction, int, int, int, int, string) {
	fake.getIDMutex.RLock()
	defer fake.getIDMutex.RUnlock()
	return fake.getIDArgsForCall[i].arg1, fake.getIDArgsForCall[i].arg2, fake.getIDArgsForCall[i].arg3, fake.getIDArgsForCall[i].arg4, fake.getIDArgsForCall[i].arg5, fake.getIDArgsForCall[i].arg6
}

func (fake *DestinationRepo) GetIDReturns(result1 int, result2 error) {
	fake.GetIDStub = nil
	fake.getIDReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DestinationRepo) GetIDReturnsOnCall(i int, result1 int, result2 error) {
	fake.GetIDStub = nil
	if fake.getIDReturnsOnCall == nil {
		fake.getIDReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.getIDReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DestinationRepo) CountWhereGroupID(arg1 store.Transaction, arg2 int) (int, error) {
	fake.countWhereGroupIDMutex.Lock()
	ret, specificReturn := fake.countWhereGroupIDReturnsOnCall[len(fake.countWhereGroupIDArgsForCall)]
	fake.countWhereGroupIDArgsForCall = append(fake.countWhereGroupIDArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CountWhereGroupID", []interface{}{arg1, arg2})
	fake.countWhereGroupIDMutex.Unlock()
	if fake.CountWhereGroupIDStub != nil {
		return fake.CountWhereGroupIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.countWhereGroupIDReturns.result1, fake.countWhereGroupIDReturns.result2
}

func (fake *DestinationRepo) CountWhereGroupIDCallCount() int {
	fake.countWhereGroupIDMutex.RLock()
	defer fake.countWhereGroupIDMutex.RUnlock()
	return len(fake.countWhereGroupIDArgsForCall)
}

func (fake *DestinationRepo) CountWhereGroupIDArgsForCall(i int) (store.Transaction, int) {
	fake.countWhereGroupIDMutex.RLock()
	defer fake.countWhereGroupIDMutex.RUnlock()
	return fake.countWhereGroupIDArgsForCall[i].arg1, fake.countWhereGroupIDArgsForCall[i].arg2
}

func (fake *DestinationRepo) CountWhereGroupIDReturns(result1 int, result2 error) {
	fake.CountWhereGroupIDStub = nil
	fake.countWhereGroupIDReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DestinationRepo) CountWhereGroupIDReturnsOnCall(i int, result1 int, result2 error) {
	fake.CountWhereGroupIDStub = nil
	if fake.countWhereGroupIDReturnsOnCall == nil {
		fake.countWhereGroupIDReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countWhereGroupIDReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *DestinationRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getIDMutex.RLock()
	defer fake.getIDMutex.RUnlock()
	fake.countWhereGroupIDMutex.RLock()
	defer fake.countWhereGroupIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DestinationRepo) recordInvocation(key string, args []interface{}) {
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

var _ store.DestinationRepo = new(DestinationRepo)
