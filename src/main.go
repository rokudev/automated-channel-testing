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

package main

import (
    "fmt"
	httpServer "httpServer"
	"os"
	"regexp"
)

func main() {
    defaultPort := "9000"
	validPort := regexp.MustCompile(`^[0-9]+$`)
	server := httpServer.GetServerInstance()
	if len(os.Args) > 1 && validPort.MatchString(os.Args[1]) {
		fmt.Println("Starting driver on port: " + os.Args[1])
		server.Start(os.Args[1])
	} else {
		fmt.Println("Starting driver on port: " + defaultPort)
		server.Start(defaultPort)
	}
}