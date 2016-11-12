// This file was generated by counterfeiter
package converterfakes

import (
	"sync"

	"github.com/cbguder/v2e/converter"
	"github.com/cbguder/v2e/models"
)

type FakeExporter struct {
	ExportStub        func(string, []models.Note) error
	exportMutex       sync.RWMutex
	exportArgsForCall []struct {
		arg1 string
		arg2 []models.Note
	}
	exportReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExporter) Export(arg1 string, arg2 []models.Note) error {
	var arg2Copy []models.Note
	if arg2 != nil {
		arg2Copy = make([]models.Note, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.exportMutex.Lock()
	fake.exportArgsForCall = append(fake.exportArgsForCall, struct {
		arg1 string
		arg2 []models.Note
	}{arg1, arg2Copy})
	fake.recordInvocation("Export", []interface{}{arg1, arg2Copy})
	fake.exportMutex.Unlock()
	if fake.ExportStub != nil {
		return fake.ExportStub(arg1, arg2)
	} else {
		return fake.exportReturns.result1
	}
}

func (fake *FakeExporter) ExportCallCount() int {
	fake.exportMutex.RLock()
	defer fake.exportMutex.RUnlock()
	return len(fake.exportArgsForCall)
}

func (fake *FakeExporter) ExportArgsForCall(i int) (string, []models.Note) {
	fake.exportMutex.RLock()
	defer fake.exportMutex.RUnlock()
	return fake.exportArgsForCall[i].arg1, fake.exportArgsForCall[i].arg2
}

func (fake *FakeExporter) ExportReturns(result1 error) {
	fake.ExportStub = nil
	fake.exportReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exportMutex.RLock()
	defer fake.exportMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeExporter) recordInvocation(key string, args []interface{}) {
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

var _ converter.Exporter = new(FakeExporter)