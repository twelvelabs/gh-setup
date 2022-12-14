// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package gh

import (
	"io"
	"sync"
)

// Ensure, that RESTClientMock does implement RESTClient.
// If this is not the case, regenerate this file with moq.
var _ RESTClient = &RESTClientMock{}

// RESTClientMock is a mock implementation of RESTClient.
//
//	func TestSomethingThatUsesRESTClient(t *testing.T) {
//
//		// make and configure a mocked RESTClient
//		mockedRESTClient := &RESTClientMock{
//			DeleteFunc: func(path string, response interface{}) error {
//				panic("mock out the Delete method")
//			},
//			GetFunc: func(path string, response interface{}) error {
//				panic("mock out the Get method")
//			},
//			PatchFunc: func(path string, body io.Reader, response interface{}) error {
//				panic("mock out the Patch method")
//			},
//			PostFunc: func(path string, body io.Reader, response interface{}) error {
//				panic("mock out the Post method")
//			},
//			PutFunc: func(path string, body io.Reader, response interface{}) error {
//				panic("mock out the Put method")
//			},
//		}
//
//		// use mockedRESTClient in code that requires RESTClient
//		// and then make assertions.
//
//	}
type RESTClientMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(path string, response interface{}) error

	// GetFunc mocks the Get method.
	GetFunc func(path string, response interface{}) error

	// PatchFunc mocks the Patch method.
	PatchFunc func(path string, body io.Reader, response interface{}) error

	// PostFunc mocks the Post method.
	PostFunc func(path string, body io.Reader, response interface{}) error

	// PutFunc mocks the Put method.
	PutFunc func(path string, body io.Reader, response interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Path is the path argument value.
			Path string
			// Response is the response argument value.
			Response interface{}
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Path is the path argument value.
			Path string
			// Response is the response argument value.
			Response interface{}
		}
		// Patch holds details about calls to the Patch method.
		Patch []struct {
			// Path is the path argument value.
			Path string
			// Body is the body argument value.
			Body io.Reader
			// Response is the response argument value.
			Response interface{}
		}
		// Post holds details about calls to the Post method.
		Post []struct {
			// Path is the path argument value.
			Path string
			// Body is the body argument value.
			Body io.Reader
			// Response is the response argument value.
			Response interface{}
		}
		// Put holds details about calls to the Put method.
		Put []struct {
			// Path is the path argument value.
			Path string
			// Body is the body argument value.
			Body io.Reader
			// Response is the response argument value.
			Response interface{}
		}
	}
	lockDelete sync.RWMutex
	lockGet    sync.RWMutex
	lockPatch  sync.RWMutex
	lockPost   sync.RWMutex
	lockPut    sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *RESTClientMock) Delete(path string, response interface{}) error {
	if mock.DeleteFunc == nil {
		panic("RESTClientMock.DeleteFunc: method is nil but RESTClient.Delete was just called")
	}
	callInfo := struct {
		Path     string
		Response interface{}
	}{
		Path:     path,
		Response: response,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(path, response)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedRESTClient.DeleteCalls())
func (mock *RESTClientMock) DeleteCalls() []struct {
	Path     string
	Response interface{}
} {
	var calls []struct {
		Path     string
		Response interface{}
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *RESTClientMock) Get(path string, response interface{}) error {
	if mock.GetFunc == nil {
		panic("RESTClientMock.GetFunc: method is nil but RESTClient.Get was just called")
	}
	callInfo := struct {
		Path     string
		Response interface{}
	}{
		Path:     path,
		Response: response,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(path, response)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedRESTClient.GetCalls())
func (mock *RESTClientMock) GetCalls() []struct {
	Path     string
	Response interface{}
} {
	var calls []struct {
		Path     string
		Response interface{}
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Patch calls PatchFunc.
func (mock *RESTClientMock) Patch(path string, body io.Reader, response interface{}) error {
	if mock.PatchFunc == nil {
		panic("RESTClientMock.PatchFunc: method is nil but RESTClient.Patch was just called")
	}
	callInfo := struct {
		Path     string
		Body     io.Reader
		Response interface{}
	}{
		Path:     path,
		Body:     body,
		Response: response,
	}
	mock.lockPatch.Lock()
	mock.calls.Patch = append(mock.calls.Patch, callInfo)
	mock.lockPatch.Unlock()
	return mock.PatchFunc(path, body, response)
}

// PatchCalls gets all the calls that were made to Patch.
// Check the length with:
//
//	len(mockedRESTClient.PatchCalls())
func (mock *RESTClientMock) PatchCalls() []struct {
	Path     string
	Body     io.Reader
	Response interface{}
} {
	var calls []struct {
		Path     string
		Body     io.Reader
		Response interface{}
	}
	mock.lockPatch.RLock()
	calls = mock.calls.Patch
	mock.lockPatch.RUnlock()
	return calls
}

// Post calls PostFunc.
func (mock *RESTClientMock) Post(path string, body io.Reader, response interface{}) error {
	if mock.PostFunc == nil {
		panic("RESTClientMock.PostFunc: method is nil but RESTClient.Post was just called")
	}
	callInfo := struct {
		Path     string
		Body     io.Reader
		Response interface{}
	}{
		Path:     path,
		Body:     body,
		Response: response,
	}
	mock.lockPost.Lock()
	mock.calls.Post = append(mock.calls.Post, callInfo)
	mock.lockPost.Unlock()
	return mock.PostFunc(path, body, response)
}

// PostCalls gets all the calls that were made to Post.
// Check the length with:
//
//	len(mockedRESTClient.PostCalls())
func (mock *RESTClientMock) PostCalls() []struct {
	Path     string
	Body     io.Reader
	Response interface{}
} {
	var calls []struct {
		Path     string
		Body     io.Reader
		Response interface{}
	}
	mock.lockPost.RLock()
	calls = mock.calls.Post
	mock.lockPost.RUnlock()
	return calls
}

// Put calls PutFunc.
func (mock *RESTClientMock) Put(path string, body io.Reader, response interface{}) error {
	if mock.PutFunc == nil {
		panic("RESTClientMock.PutFunc: method is nil but RESTClient.Put was just called")
	}
	callInfo := struct {
		Path     string
		Body     io.Reader
		Response interface{}
	}{
		Path:     path,
		Body:     body,
		Response: response,
	}
	mock.lockPut.Lock()
	mock.calls.Put = append(mock.calls.Put, callInfo)
	mock.lockPut.Unlock()
	return mock.PutFunc(path, body, response)
}

// PutCalls gets all the calls that were made to Put.
// Check the length with:
//
//	len(mockedRESTClient.PutCalls())
func (mock *RESTClientMock) PutCalls() []struct {
	Path     string
	Body     io.Reader
	Response interface{}
} {
	var calls []struct {
		Path     string
		Body     io.Reader
		Response interface{}
	}
	mock.lockPut.RLock()
	calls = mock.calls.Put
	mock.lockPut.RUnlock()
	return calls
}
