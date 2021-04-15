///////////////////////////////////////////////////////////////////////////
// Copyright 2019 Roku, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//////////////////////////////////////////////////////////////////////////

package ecpClient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

var base, _ = url.Parse("some.api")
// RoundTripFunc .
type RoundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func GetMockedClient(resp *string) *EcpClient {
	httpClient := TestingHTTPClient(resp)

	cli := NewClient()
	cli.HttpClient = httpClient

	contentsClient := EcpClient{
		BaseClient: &BaseClient{
			BaseURL:        base,
			HttpClient:     cli,
		},
	}

	return &contentsClient
}


func TestingHTTPClient(resp *string) *http.Client {
	var function func(r *http.Request) (*http.Response, error)
	if resp != nil{
		function = func(req *http.Request) (*http.Response, error) {
			// Test request parameters
			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(*resp)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}, nil
		}
	} else {
		function = func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 408,
				// Send response to be tested
				Body: nil,
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}, http.ErrHandlerTimeout
		}
	}

	cli := NewTestClient(function)
	return cli
}
