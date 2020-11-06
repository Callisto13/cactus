// Code generated by counterfeiter. DO NOT EDIT.
package runnerfakes

import (
	"sync"

	"github.com/callisto13/cactus/runner"
)

type FakeImageEditor struct {
	ClearStub        func()
	clearMutex       sync.RWMutex
	clearArgsForCall []struct {
	}
	CreateImageStub        func(int, int)
	createImageMutex       sync.RWMutex
	createImageArgsForCall []struct {
		arg1 int
		arg2 int
	}
	PrettyStub        func() string
	prettyMutex       sync.RWMutex
	prettyArgsForCall []struct {
	}
	prettyReturns struct {
		result1 string
	}
	prettyReturnsOnCall map[int]struct {
		result1 string
	}
	SetStub        func(int, int, string) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 int
		arg2 int
		arg3 string
	}
	setReturns struct {
		result1 error
	}
	setReturnsOnCall map[int]struct {
		result1 error
	}
	SetMultiXStub        func(int, int, int, string) error
	setMultiXMutex       sync.RWMutex
	setMultiXArgsForCall []struct {
		arg1 int
		arg2 int
		arg3 int
		arg4 string
	}
	setMultiXReturns struct {
		result1 error
	}
	setMultiXReturnsOnCall map[int]struct {
		result1 error
	}
	SetMultiYStub        func(int, int, int, string) error
	setMultiYMutex       sync.RWMutex
	setMultiYArgsForCall []struct {
		arg1 int
		arg2 int
		arg3 int
		arg4 string
	}
	setMultiYReturns struct {
		result1 error
	}
	setMultiYReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageEditor) Clear() {
	fake.clearMutex.Lock()
	fake.clearArgsForCall = append(fake.clearArgsForCall, struct {
	}{})
	fake.recordInvocation("Clear", []interface{}{})
	fake.clearMutex.Unlock()
	if fake.ClearStub != nil {
		fake.ClearStub()
	}
}

func (fake *FakeImageEditor) ClearCallCount() int {
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	return len(fake.clearArgsForCall)
}

func (fake *FakeImageEditor) ClearCalls(stub func()) {
	fake.clearMutex.Lock()
	defer fake.clearMutex.Unlock()
	fake.ClearStub = stub
}

func (fake *FakeImageEditor) CreateImage(arg1 int, arg2 int) {
	fake.createImageMutex.Lock()
	fake.createImageArgsForCall = append(fake.createImageArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CreateImage", []interface{}{arg1, arg2})
	fake.createImageMutex.Unlock()
	if fake.CreateImageStub != nil {
		fake.CreateImageStub(arg1, arg2)
	}
}

func (fake *FakeImageEditor) CreateImageCallCount() int {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	return len(fake.createImageArgsForCall)
}

func (fake *FakeImageEditor) CreateImageCalls(stub func(int, int)) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = stub
}

func (fake *FakeImageEditor) CreateImageArgsForCall(i int) (int, int) {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	argsForCall := fake.createImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageEditor) Pretty() string {
	fake.prettyMutex.Lock()
	ret, specificReturn := fake.prettyReturnsOnCall[len(fake.prettyArgsForCall)]
	fake.prettyArgsForCall = append(fake.prettyArgsForCall, struct {
	}{})
	fake.recordInvocation("Pretty", []interface{}{})
	fake.prettyMutex.Unlock()
	if fake.PrettyStub != nil {
		return fake.PrettyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.prettyReturns
	return fakeReturns.result1
}

func (fake *FakeImageEditor) PrettyCallCount() int {
	fake.prettyMutex.RLock()
	defer fake.prettyMutex.RUnlock()
	return len(fake.prettyArgsForCall)
}

func (fake *FakeImageEditor) PrettyCalls(stub func() string) {
	fake.prettyMutex.Lock()
	defer fake.prettyMutex.Unlock()
	fake.PrettyStub = stub
}

func (fake *FakeImageEditor) PrettyReturns(result1 string) {
	fake.prettyMutex.Lock()
	defer fake.prettyMutex.Unlock()
	fake.PrettyStub = nil
	fake.prettyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeImageEditor) PrettyReturnsOnCall(i int, result1 string) {
	fake.prettyMutex.Lock()
	defer fake.prettyMutex.Unlock()
	fake.PrettyStub = nil
	if fake.prettyReturnsOnCall == nil {
		fake.prettyReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.prettyReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeImageEditor) Set(arg1 int, arg2 int, arg3 string) error {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 int
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		return fake.SetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setReturns
	return fakeReturns.result1
}

func (fake *FakeImageEditor) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeImageEditor) SetCalls(stub func(int, int, string) error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *FakeImageEditor) SetArgsForCall(i int) (int, int, string) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageEditor) SetReturns(result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEditor) SetReturnsOnCall(i int, result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEditor) SetMultiX(arg1 int, arg2 int, arg3 int, arg4 string) error {
	fake.setMultiXMutex.Lock()
	ret, specificReturn := fake.setMultiXReturnsOnCall[len(fake.setMultiXArgsForCall)]
	fake.setMultiXArgsForCall = append(fake.setMultiXArgsForCall, struct {
		arg1 int
		arg2 int
		arg3 int
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetMultiX", []interface{}{arg1, arg2, arg3, arg4})
	fake.setMultiXMutex.Unlock()
	if fake.SetMultiXStub != nil {
		return fake.SetMultiXStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setMultiXReturns
	return fakeReturns.result1
}

func (fake *FakeImageEditor) SetMultiXCallCount() int {
	fake.setMultiXMutex.RLock()
	defer fake.setMultiXMutex.RUnlock()
	return len(fake.setMultiXArgsForCall)
}

func (fake *FakeImageEditor) SetMultiXCalls(stub func(int, int, int, string) error) {
	fake.setMultiXMutex.Lock()
	defer fake.setMultiXMutex.Unlock()
	fake.SetMultiXStub = stub
}

func (fake *FakeImageEditor) SetMultiXArgsForCall(i int) (int, int, int, string) {
	fake.setMultiXMutex.RLock()
	defer fake.setMultiXMutex.RUnlock()
	argsForCall := fake.setMultiXArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeImageEditor) SetMultiXReturns(result1 error) {
	fake.setMultiXMutex.Lock()
	defer fake.setMultiXMutex.Unlock()
	fake.SetMultiXStub = nil
	fake.setMultiXReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEditor) SetMultiXReturnsOnCall(i int, result1 error) {
	fake.setMultiXMutex.Lock()
	defer fake.setMultiXMutex.Unlock()
	fake.SetMultiXStub = nil
	if fake.setMultiXReturnsOnCall == nil {
		fake.setMultiXReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setMultiXReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEditor) SetMultiY(arg1 int, arg2 int, arg3 int, arg4 string) error {
	fake.setMultiYMutex.Lock()
	ret, specificReturn := fake.setMultiYReturnsOnCall[len(fake.setMultiYArgsForCall)]
	fake.setMultiYArgsForCall = append(fake.setMultiYArgsForCall, struct {
		arg1 int
		arg2 int
		arg3 int
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetMultiY", []interface{}{arg1, arg2, arg3, arg4})
	fake.setMultiYMutex.Unlock()
	if fake.SetMultiYStub != nil {
		return fake.SetMultiYStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setMultiYReturns
	return fakeReturns.result1
}

func (fake *FakeImageEditor) SetMultiYCallCount() int {
	fake.setMultiYMutex.RLock()
	defer fake.setMultiYMutex.RUnlock()
	return len(fake.setMultiYArgsForCall)
}

func (fake *FakeImageEditor) SetMultiYCalls(stub func(int, int, int, string) error) {
	fake.setMultiYMutex.Lock()
	defer fake.setMultiYMutex.Unlock()
	fake.SetMultiYStub = stub
}

func (fake *FakeImageEditor) SetMultiYArgsForCall(i int) (int, int, int, string) {
	fake.setMultiYMutex.RLock()
	defer fake.setMultiYMutex.RUnlock()
	argsForCall := fake.setMultiYArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeImageEditor) SetMultiYReturns(result1 error) {
	fake.setMultiYMutex.Lock()
	defer fake.setMultiYMutex.Unlock()
	fake.SetMultiYStub = nil
	fake.setMultiYReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEditor) SetMultiYReturnsOnCall(i int, result1 error) {
	fake.setMultiYMutex.Lock()
	defer fake.setMultiYMutex.Unlock()
	fake.SetMultiYStub = nil
	if fake.setMultiYReturnsOnCall == nil {
		fake.setMultiYReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setMultiYReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEditor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	fake.prettyMutex.RLock()
	defer fake.prettyMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.setMultiXMutex.RLock()
	defer fake.setMultiXMutex.RUnlock()
	fake.setMultiYMutex.RLock()
	defer fake.setMultiYMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageEditor) recordInvocation(key string, args []interface{}) {
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

var _ runner.ImageEditor = new(FakeImageEditor)
