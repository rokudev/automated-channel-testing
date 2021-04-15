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
	"net/http"
	"time"
)

type Option func(*Client)

func SetHTTPClient(httpClient *http.Client) Option {
	return func(cli *Client) {
		cli.HttpClient = httpClient
	}
}

func SetRequestTimeout(timeout time.Duration) Option {
	return func(cli *Client) {
		cli.HttpClient.Timeout = timeout
	}
}

type Client struct {
	HttpClient *http.Client
}

func NewClient(options ...Option) *Client {
	cli := Client{
		HttpClient: &http.Client{},
	}

	for i := range options {
		options[i](&cli)
	}

	return &cli
}
