// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package git

import (
	"bytes"
	"sync"
)

// Ensure, that ClientMock does implement Client.
// If this is not the case, regenerate this file with moq.
var _ Client = &ClientMock{}

// ClientMock is a mock implementation of Client.
//
//	func TestSomethingThatUsesClient(t *testing.T) {
//
//		// make and configure a mocked Client
//		mockedClient := &ClientMock{
//			ExecFunc: func(args ...string) (bytes.Buffer, bytes.Buffer, error) {
//				panic("mock out the Exec method")
//			},
//			HasCommitsFunc: func() bool {
//				panic("mock out the HasCommits method")
//			},
//			HasRemoteFunc: func(name string) bool {
//				panic("mock out the HasRemote method")
//			},
//			IsDirtyFunc: func() bool {
//				panic("mock out the IsDirty method")
//			},
//			IsInitializedFunc: func() bool {
//				panic("mock out the IsInitialized method")
//			},
//			IsInstalledFunc: func() bool {
//				panic("mock out the IsInstalled method")
//			},
//			StatusLinesFunc: func() ([]string, error) {
//				panic("mock out the StatusLines method")
//			},
//		}
//
//		// use mockedClient in code that requires Client
//		// and then make assertions.
//
//	}
type ClientMock struct {
	// ExecFunc mocks the Exec method.
	ExecFunc func(args ...string) (bytes.Buffer, bytes.Buffer, error)

	// HasCommitsFunc mocks the HasCommits method.
	HasCommitsFunc func() bool

	// HasRemoteFunc mocks the HasRemote method.
	HasRemoteFunc func(name string) bool

	// IsDirtyFunc mocks the IsDirty method.
	IsDirtyFunc func() bool

	// IsInitializedFunc mocks the IsInitialized method.
	IsInitializedFunc func() bool

	// IsInstalledFunc mocks the IsInstalled method.
	IsInstalledFunc func() bool

	// StatusLinesFunc mocks the StatusLines method.
	StatusLinesFunc func() ([]string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Args is the args argument value.
			Args []string
		}
		// HasCommits holds details about calls to the HasCommits method.
		HasCommits []struct {
		}
		// HasRemote holds details about calls to the HasRemote method.
		HasRemote []struct {
			// Name is the name argument value.
			Name string
		}
		// IsDirty holds details about calls to the IsDirty method.
		IsDirty []struct {
		}
		// IsInitialized holds details about calls to the IsInitialized method.
		IsInitialized []struct {
		}
		// IsInstalled holds details about calls to the IsInstalled method.
		IsInstalled []struct {
		}
		// StatusLines holds details about calls to the StatusLines method.
		StatusLines []struct {
		}
	}
	lockExec          sync.RWMutex
	lockHasCommits    sync.RWMutex
	lockHasRemote     sync.RWMutex
	lockIsDirty       sync.RWMutex
	lockIsInitialized sync.RWMutex
	lockIsInstalled   sync.RWMutex
	lockStatusLines   sync.RWMutex
}

// Exec calls ExecFunc.
func (mock *ClientMock) Exec(args ...string) (bytes.Buffer, bytes.Buffer, error) {
	if mock.ExecFunc == nil {
		panic("ClientMock.ExecFunc: method is nil but Client.Exec was just called")
	}
	callInfo := struct {
		Args []string
	}{
		Args: args,
	}
	mock.lockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	mock.lockExec.Unlock()
	return mock.ExecFunc(args...)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//
//	len(mockedClient.ExecCalls())
func (mock *ClientMock) ExecCalls() []struct {
	Args []string
} {
	var calls []struct {
		Args []string
	}
	mock.lockExec.RLock()
	calls = mock.calls.Exec
	mock.lockExec.RUnlock()
	return calls
}

// HasCommits calls HasCommitsFunc.
func (mock *ClientMock) HasCommits() bool {
	if mock.HasCommitsFunc == nil {
		panic("ClientMock.HasCommitsFunc: method is nil but Client.HasCommits was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHasCommits.Lock()
	mock.calls.HasCommits = append(mock.calls.HasCommits, callInfo)
	mock.lockHasCommits.Unlock()
	return mock.HasCommitsFunc()
}

// HasCommitsCalls gets all the calls that were made to HasCommits.
// Check the length with:
//
//	len(mockedClient.HasCommitsCalls())
func (mock *ClientMock) HasCommitsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHasCommits.RLock()
	calls = mock.calls.HasCommits
	mock.lockHasCommits.RUnlock()
	return calls
}

// HasRemote calls HasRemoteFunc.
func (mock *ClientMock) HasRemote(name string) bool {
	if mock.HasRemoteFunc == nil {
		panic("ClientMock.HasRemoteFunc: method is nil but Client.HasRemote was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockHasRemote.Lock()
	mock.calls.HasRemote = append(mock.calls.HasRemote, callInfo)
	mock.lockHasRemote.Unlock()
	return mock.HasRemoteFunc(name)
}

// HasRemoteCalls gets all the calls that were made to HasRemote.
// Check the length with:
//
//	len(mockedClient.HasRemoteCalls())
func (mock *ClientMock) HasRemoteCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockHasRemote.RLock()
	calls = mock.calls.HasRemote
	mock.lockHasRemote.RUnlock()
	return calls
}

// IsDirty calls IsDirtyFunc.
func (mock *ClientMock) IsDirty() bool {
	if mock.IsDirtyFunc == nil {
		panic("ClientMock.IsDirtyFunc: method is nil but Client.IsDirty was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsDirty.Lock()
	mock.calls.IsDirty = append(mock.calls.IsDirty, callInfo)
	mock.lockIsDirty.Unlock()
	return mock.IsDirtyFunc()
}

// IsDirtyCalls gets all the calls that were made to IsDirty.
// Check the length with:
//
//	len(mockedClient.IsDirtyCalls())
func (mock *ClientMock) IsDirtyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsDirty.RLock()
	calls = mock.calls.IsDirty
	mock.lockIsDirty.RUnlock()
	return calls
}

// IsInitialized calls IsInitializedFunc.
func (mock *ClientMock) IsInitialized() bool {
	if mock.IsInitializedFunc == nil {
		panic("ClientMock.IsInitializedFunc: method is nil but Client.IsInitialized was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsInitialized.Lock()
	mock.calls.IsInitialized = append(mock.calls.IsInitialized, callInfo)
	mock.lockIsInitialized.Unlock()
	return mock.IsInitializedFunc()
}

// IsInitializedCalls gets all the calls that were made to IsInitialized.
// Check the length with:
//
//	len(mockedClient.IsInitializedCalls())
func (mock *ClientMock) IsInitializedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsInitialized.RLock()
	calls = mock.calls.IsInitialized
	mock.lockIsInitialized.RUnlock()
	return calls
}

// IsInstalled calls IsInstalledFunc.
func (mock *ClientMock) IsInstalled() bool {
	if mock.IsInstalledFunc == nil {
		panic("ClientMock.IsInstalledFunc: method is nil but Client.IsInstalled was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsInstalled.Lock()
	mock.calls.IsInstalled = append(mock.calls.IsInstalled, callInfo)
	mock.lockIsInstalled.Unlock()
	return mock.IsInstalledFunc()
}

// IsInstalledCalls gets all the calls that were made to IsInstalled.
// Check the length with:
//
//	len(mockedClient.IsInstalledCalls())
func (mock *ClientMock) IsInstalledCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsInstalled.RLock()
	calls = mock.calls.IsInstalled
	mock.lockIsInstalled.RUnlock()
	return calls
}

// StatusLines calls StatusLinesFunc.
func (mock *ClientMock) StatusLines() ([]string, error) {
	if mock.StatusLinesFunc == nil {
		panic("ClientMock.StatusLinesFunc: method is nil but Client.StatusLines was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStatusLines.Lock()
	mock.calls.StatusLines = append(mock.calls.StatusLines, callInfo)
	mock.lockStatusLines.Unlock()
	return mock.StatusLinesFunc()
}

// StatusLinesCalls gets all the calls that were made to StatusLines.
// Check the length with:
//
//	len(mockedClient.StatusLinesCalls())
func (mock *ClientMock) StatusLinesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStatusLines.RLock()
	calls = mock.calls.StatusLines
	mock.lockStatusLines.RUnlock()
	return calls
}
