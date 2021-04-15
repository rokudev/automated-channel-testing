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

import(
	"net/http"
)

type Route struct {
    Method      string
    Pattern     string
    HandlerFunc appHandler
}

type Routes []Route

func (s *Server) SetUpRoutes() {
	routes := Routes{
		Route{
		   "GET",
		   "/v1/status",
		   s.GetStatusHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/load",
			s.GetLoadHandler(),
		 },
		Route{
			"POST",
			"/v1/session",
			s.GetStartSessionHandler(),
		},
		Route{
			"GET",
			"/v1/sessions",
			s.GetSessionsInfoHandler(),
		},
		Route{
			"GET",
			"/v1/session/{sessionId}",
			s.GetSessionHandler(),
		},
		Route{
			"DELETE",
			"/v1/session/{sessionId}",
			s.GetSessionDeleteHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/timeouts",
			s.GetTimeoutsHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/element",
			s.GetElementHandler(),
		},
		Route{
			"GET",
			"/v1/session/{sessionId}/player",
			s.GetPlayerHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/elements",
			s.GetElementsHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/element/active",
			s.GetActiveElementHandler(),
		},
		Route{
			"GET",
			"/v1/session/{sessionId}/apps",
			s.GetAppsHandler(),
		},
		Route{
			"GET",
			"/v1/session/{sessionId}/current_app",
			s.GetCurrentAppHandler(),
		},
		Route{
			"GET",
			"/v1/session/{sessionId}/source",
			s.GetSourceHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/press",
			s.GetPressButtonHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/launch",
			s.GetLaunchHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/input",
			s.GetInputHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/timeouts/implicit_wait",
			s.GetImplicitTimeoutHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/timeouts/press_wait",
			s.GetImplicitTimeoutHandler(),
		},
		Route{
			"POST",
			"/v1/session/{sessionId}/install",
			s.GetInstallHandler(),
		},
	}
		
	
	for _, route := range routes {
        s.router.
            Methods(route.Method).
            Path(route.Pattern).
            Handler(Middleware(route.HandlerFunc)).GetError()
    }
    s.router.NotFoundHandler = http.Handler(Middleware(s.notFound()))
	http.Handle("/", s.router)
}