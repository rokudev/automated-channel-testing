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

package httpServer

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func Middleware(h appHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if e := h(w, r); e != nil {
			if e.InternalCode != nil {
				w.WriteHeader(e.Code)
				errorResponse := &errorResponse{
					Message: e.Message,
				}
				response := &SessionResponse{
					Status: *e.InternalCode,
					Value:  errorResponse,
				}
				js, _ := json.Marshal(response)
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			} else {
				http.Error(w, e.Message, e.Code)
			}

		}
	})
}
