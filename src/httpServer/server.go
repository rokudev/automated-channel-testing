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
	"github.com/gorilla/mux"
	ecp "ecpClient"
	"net/http"
	"time"
	"github.com/sirupsen/logrus"
)



type Server struct {
    router *mux.Router
	sessions map[string]*SessionInfo
} 

type SessionInfo struct {
	client *ecp.EcpClient
	plugin *ecp.PluginClient
	capability *Capability
	pressDelay time.Duration
}

func GetServerInstance() *Server {
	server := &Server{
		router: mux.NewRouter(),
		sessions: make(map[string]*SessionInfo),
	}
   
	return server
}

func (s *Server) Start(port string) {
	s.SetUpRoutes()
	err := http.ListenAndServe(":" + port, nil)
	if err != http.ErrServerClosed {
	   logrus.WithError(err).Error("Http Server stopped unexpected")
	} else {
	   logrus.WithError(err).Info("Http Server stopped")
	}
 }