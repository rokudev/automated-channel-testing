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
    "net/http"
    "encoding/json"
    ecp "ecpClient"
    "io/ioutil"
    "io"
    "github.com/gorilla/mux"
    "strings"
    "time"
    "strconv"
    "runtime"
    "version"
    "bytes"
)

const defaultPressDelay = 1000

type appError struct {
    Message string
    Code    int
    InternalCode *int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (s *Server) GetStatusHandler() appHandler {
     return func(w http.ResponseWriter, r *http.Request) *appError {
            arch := runtime.GOARCH
            osName := runtime.GOOS
            vers := version.BuildVersion
            time := version.BuildTime
            response := &SessionResponse {       
            Status: 0,
            Value: &Status{
                Os: OsInfo{
                    Arch: arch,
                    Name: osName,
                },
                Build: BuildInfo{
                    Version: vers,
                    Time: time,
                },
            },
        }
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetStartSessionHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        b, err := ioutil.ReadAll(r.Body)
        var t Capability
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }

        if validIP4(t.Ip) != true {
            return &appError{ "Invalid IP", http.StatusBadRequest, nil}
        }
        client, err := ecp.GetEcpClient(t.Ip)
        if err != nil {
            status := responseStatuses["SessionNotCreatedException"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        plugin, err := ecp.GetPluginClient(t.Ip)
        if err != nil {
            status := responseStatuses["SessionNotCreatedException"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        info, err := client.GetDeviceInfo()
        if err != nil {
            status := responseStatuses["SessionNotCreatedException"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        id := info.AdvertisingID
        session := s.sessions[id]
        if session !=nil {
            status := responseStatuses["SessionNotCreatedException"]
            return &appError{ "Session already exist", http.StatusInternalServerError, &status}
        }

        timeout:= t.Timeout
        if timeout > 0  {
            client.SetTimeout(time.Duration(timeout))
        }

        pressDelay := defaultPressDelay
        delay := t.PressDelay
        if delay > 0 {
            pressDelay = delay
        }

        capability := &Capability {
            VendorName: info.VendorName,
            ModelName: info.ModelName,
            Language: info.Language,
            Country: info.Country,
            Ip: t.Ip,
            PressDelay: pressDelay,
            Timeout: client.GetTimeout(),
        }
        response := &SessionResponse{
            Id: id,
            Value: capability,
        }
        s.sessions[id] = &SessionInfo{
            client,
            plugin,
            capability,
            time.Duration(pressDelay),
        }

        return prepareResponse(w, response)
    }
 }

func (s *Server) GetSessionHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        vars := mux.Vars(r)
        id := vars["sessionId"]
        session := s.sessions[id]
        if session == nil {
            status := responseStatuses["NoSuchDriver"]
            return &appError{ "Invalid sessionId", http.StatusInternalServerError, &status}
        }
        response := &SessionResponse{
            Id: id,
            Status: responseStatuses["Success"],
            Value: session.capability,
        }
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetSessionsInfoHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        var sessions []*SessionResponse
        for key, element := range s.sessions {
            response := &SessionResponse{
                Id: key,
                Value: element.capability,
            }
            sessions = append(sessions, response)
        }
        
        return prepareResponse(w, sessions)
    }
}

func (s *Server) GetSessionDeleteHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        vars := mux.Vars(r)
        sessionId := vars["sessionId"]
        session := s.sessions[sessionId]
        if session != nil {
            delete(s.sessions, sessionId)
        } else {
            status := responseStatuses["NoSuchDriver"]
            return &appError{ "Invalid sessionId", http.StatusInternalServerError, &status}
        }
        
        response := &SessionResponse {
            Id: sessionId,
            Status: responseStatuses["Success"],
            Value: nil,
        }

        return prepareResponse(w, response)
    }
 }

func (s *Server) GetTimeoutsHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t TimeoutRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        timeout := t.Ms
        if timeout < 0 {
            return &appError{ "Timeout must be a positive number.", http.StatusBadRequest, nil}
        }
        switch t.Type {
        case "implicit":
            client.SetTimeout(time.Duration(timeout))
        case "pressDelay":
            s.sessions[id].pressDelay = time.Duration(timeout)
        default:
            return &appError{ "Invalid \"type\" value.", http.StatusBadRequest, nil}
        }
        response := &SessionResponse{
            Id: id,
            Status: 0,
            Value: nil,
        }
        return prepareResponse(w, response)
    }
 }

 func (s *Server) GetLoadHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        plugin, id, errorInfo := s.getPlugin(r)
        if  errorInfo != nil {
            return errorInfo
        }
        r.ParseMultipartForm(4096)
        file,_, _ := r.FormFile("channel")
        buf := bytes.NewBuffer(nil)
        _,err := io.Copy(buf, file);
        if err!= nil {
            return &appError{ "The \"channel\" field is required ", http.StatusBadRequest, nil}
        }
        user := r.FormValue("username")
        if  user == "" {
            return &appError{ "The \"username\" field is required ", http.StatusBadRequest, nil}
        }
        pass := r.FormValue("password")
        if  pass == "" {
            return &appError{ "The \"password\" field is required ", http.StatusBadRequest, nil}
        }
        res, err := plugin.Load(bytes.NewReader(buf.Bytes()), user, pass);
        if err !=nil || res == false {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        return prepareResponse(w, &SessionResponse{
            Id: id,
            Status: 0,
            Value: nil,
        })
    }
 }

 func (s *Server) GetImplicitTimeoutHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t TimeoutRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        timeout := t.Ms
        if timeout < 0 {
            return &appError{ "Timeout must be a positive number.", http.StatusBadRequest, nil}
        }
        urlPath := r.URL.Path
        pathParts := strings.Split(urlPath, "/")
        timeoutType := pathParts[len(pathParts) - 1]
        switch timeoutType {
        case "implicit_wait": 
            client.SetTimeout(time.Duration(timeout))
        case "press_wait":
            s.sessions[id].pressDelay = time.Duration(timeout)
        }
        response := &SessionResponse{
            Id: id,
            Status: 0,
            Value: nil,
        }
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetElementHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t ElementRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        if t.ElementData == nil {
            return &appError{ "The \"elementData\" is required ", http.StatusBadRequest, nil}
        }
        node, err := client.GetAppUi()
        if err != nil {
            status := responseStatuses["NoSuchDriver"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        result := node.Nodes
        var searchError *appError
        if t.ParentData != nil {
            result, searchError = getNodesByLocators(t.ParentData, result)
            if searchError != nil {
                return searchError
            }
        }
        result, searchError = getNodesByLocators(t.ElementData, result)
        if searchError != nil {
            return searchError
        }
        if result == nil {
            status := responseStatuses["NoSuchElement"]
            return &appError{ "An element could not be located on the screen using the given search parameters", http.StatusInternalServerError, &status}
        }
        response := &SessionResponse{
            Id: id,
            Status: 0,
            Value: result[0],
        }

        return prepareResponse(w, response)
    }
 }

 func getNodesByLocators(locators []Element, nodes []ecp.Node) ([]ecp.Node, *appError) {
    result := nodes
    var searchError *appError
    for i, element := range locators {
        checkChildNodes := false
        if i == 0 {
            checkChildNodes = true
        }
        result, searchError = searchMultipleResults(result, element, checkChildNodes)
        if searchError != nil {
            return nil, searchError
        }
    }
    return result, nil
 }

 func searchMultipleResults(nodes []ecp.Node, t Element, checkChildNodes bool) ( []ecp.Node, *appError ){
    var result []ecp.Node
    switch t.Using {
    case "tag":
        result = findMultipleNodes(nodes, t.Value, checkChildNodes)
    case "text":
        result = findMultipleNodesByText(nodes, t.Value, "text", checkChildNodes)
    case "attr":
        result = findMultipleNodesByText(nodes, t.Value, t.Attribute, checkChildNodes)
    default:
        return  nil, &appError{ "Invalid \"using\" value", http.StatusBadRequest, nil}
    }
    return result, nil
 }

 func (s *Server) GetInstallHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t ChannelRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        if len(t.ChannelId) == 0 {
            return  &appError{ "The \"channelId\" is required", http.StatusBadRequest, nil}
        }
        res, err := client.InstallChannel(t.ChannelId)
        if err !=nil || res == false {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        response := &SessionResponse {
            Id: id,
            Status: responseStatuses["Success"],
            Value: nil,
        }
        return prepareResponse(w, response)
    }
 }

 func (s *Server) GetLaunchHandler() appHandler { 
    return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t ChannelRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        if len(t.ChannelId) == 0 {
            return  &appError{ "The \"channelId\" is required", http.StatusBadRequest, nil}
        }
        res, err := client.LaunchChannel(t.ChannelId, t.ContentId, t.ContentType)
        if err !=nil || res == false {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        response := &SessionResponse {
            Id: id,
            Status: responseStatuses["Success"],
            Value: nil,
        }
        return prepareResponse(w, response)
    }
 }

 func (s *Server) GetInputHandler() appHandler { 
    return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t ChannelRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        if len(t.ChannelId) == 0 {
            return  &appError{ "The \"channelId\" is required", http.StatusBadRequest, nil}
        }
        res, err := client.InputChannel(t.ChannelId, t.ContentId, t.ContentType)
        if err !=nil || res == false {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        response := &SessionResponse {
            Id: id,
            Status: responseStatuses["Success"],
            Value: nil,
        }
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetElementsHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t ElementRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }
        if t.ElementData == nil {
            return &appError{ "The \"elementData\" is required", http.StatusBadRequest, nil}
        }
        node, err := client.GetAppUi()
        if err != nil {
            status := responseStatuses["NoSuchDriver"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        result := node.Nodes
        var searchError *appError
        if t.ParentData != nil {
            result, searchError = getNodesByLocators(t.ParentData, result)
            if searchError != nil {
                return searchError
            }
        }
        result, searchError = getNodesByLocators(t.ElementData, result)
        if searchError != nil {
            return searchError
        }

        if result == nil {
            status := responseStatuses["NoSuchElement"]
            return &appError{ "An element could not be located on the screen using the given search parameters", http.StatusInternalServerError, &status}
        }
        response := &SessionResponse{
            Id: id,
            Status: 0,
            Value: result,
        }

        return prepareResponse(w, response)
    }
 }

func (s *Server) GetActiveElementHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        node, err := client.GetAppUi()
        if err != nil {
            status := responseStatuses["NoSuchDriver"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        result := findFocusedNode(node.Nodes, false)
        
        if result == nil {
            status := responseStatuses["NoSuchElement"]
            return &appError{ "An element could not be located on the screen using the given search parameters", http.StatusInternalServerError, &status}
        }
        response := &SessionResponse{
            Id: id,
            Status: 0,
            Value: result,
        }
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetAppsHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        apps, err := client.GetApps()
        if err != nil {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }

        response := &SessionResponse{
            Status: responseStatuses["Success"],
            Id: id,
            Value: apps,
        }

        return prepareResponse(w, response)
    }
 }

 func (s *Server) GetPlayerHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        player, err := client.GetPlayer()
        if err != nil {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }

        response := &SessionResponse{
            Status: responseStatuses["Success"],
            Id: id,
            Value: player,
        }

        return prepareResponse(w, response)
    }
 }

func (s *Server) GetCurrentAppHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        app, err := client.GetActiveApp()
        if err != nil {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }

        response := &SessionResponse{
            Status: responseStatuses["Success"],
            Id: id,
            Value: app,
        }
        
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetPressButtonHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        b, err := ioutil.ReadAll(r.Body)
        var t ButtonRequest
        err = json.Unmarshal(b, &t)
        if err != nil {
            return &appError{ err.Error(), http.StatusBadRequest, nil}
        }

        if len([]rune(t.Button)) != 0 {
            result, err := client.KeyPress(t.Button)
            if err != nil || result == false {
                status := responseStatuses["UnknownError"]
                return &appError{ err.Error(), http.StatusInternalServerError, &status}
            }
        } else if len(t.Button_sequence) != 0 {
            delays := t.Button_delays
            delaysLength := len(delays)
            buttons := t.Button_sequence
            buttonsLength := len(buttons)
            var defaultDelay time.Duration
            if delaysLength > 0 {
                delayInt, error := strconv.Atoi(delays[delaysLength - 1])
                if error != nil {
                    status := responseStatuses["UnknownError"]
                    return &appError{ err.Error(), http.StatusInternalServerError, &status} 
                }
                defaultDelay = time.Duration(delayInt)
            } else {
                defaultDelay = s.sessions[id].pressDelay
            }
            for i, cmd := range buttons {
                result, err := client.KeyPress(cmd)
                if err != nil || result == false {
                    status := responseStatuses["UnknownError"]
                    return &appError{ err.Error(), http.StatusInternalServerError, &status}
                }
                if buttonsLength - 1 != i {
                    if delaysLength > i {
                        delayInt, error := strconv.Atoi(delays[i])
                        if error != nil {
                            status := responseStatuses["UnknownError"]
                            return &appError{ err.Error(), http.StatusInternalServerError, &status} 
                        }
                        time.Sleep(time.Duration(delayInt)*time.Millisecond)
                    } else {
                        time.Sleep(defaultDelay*time.Millisecond)
                    }
                }
            }
        } else {
            return &appError{ "button or button_Sequence is required.", http.StatusBadRequest, nil}
        }
        response := &SessionResponse{
            Status: responseStatuses["Success"],
            Id: id,
            Value: nil,
        }
        return prepareResponse(w, response)
    }
 }

func (s *Server) GetSourceHandler() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        client, id, errorInfo := s.getClient(r)
        if  errorInfo != nil {
            return errorInfo
        }
        result, err := client.GetSource()
        if err != nil  {
            status := responseStatuses["UnknownError"]
            return &appError{ err.Error(), http.StatusInternalServerError, &status}
        }
        response := &SessionResponse{
            Status: responseStatuses["Success"],
            Id: id,
            Value: result,
        }
        return prepareResponse(w, response)
    }
 }
 
 func (s *Server) notFound() appHandler {
	return func(w http.ResponseWriter, r *http.Request) *appError {
        return &appError{"Unimplemented Command", http.StatusNotImplemented, nil}
    }
 }

 //---------------------------Helpers--------------------------------------------

 func (s *Server) getClient(r *http.Request) (*ecp.EcpClient, string, *appError) {
    vars := mux.Vars(r)
    id := vars["sessionId"]
    session := s.sessions[id]
    if session == nil {
        status := responseStatuses["NoSuchDriver"]
        return nil, id, &appError{ "Invalid sessionId", http.StatusInternalServerError, &status}
    }
    client := session.client
    return client, id, nil
 }

 func (s *Server) getPlugin(r *http.Request) (*ecp.PluginClient, string, *appError) {
    vars := mux.Vars(r)
    id := vars["sessionId"]
    session := s.sessions[id]
    if session == nil {
        status := responseStatuses["NoSuchDriver"]
        return nil, id, &appError{ "Invalid sessionId", http.StatusInternalServerError, &status}
    }
    plugin := session.plugin
    return plugin, id, nil
 }

 func prepareResponse(w http.ResponseWriter, response interface{}) *appError {
    js, err := json.Marshal(response)
    if err != nil {
        status := responseStatuses["UnknownError"]
        return &appError{ err.Error(), http.StatusInternalServerError, &status}
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
    return nil
 }

 func findFocusedNode(nodes []ecp.Node, forceHandleFocusItem bool) *ecp.Node {
    for _, node := range nodes {
        isVisible := strings.ToLower(getAttributeValue(node, "visible"))
        isFocused := strings.ToLower(getAttributeValue(node, "focused"))
        if (isVisible == "" || isVisible == "true") && isFocused != "false" {
            focusItem := getAttributeValue(node, "focusItem")
            if focusItem != "" && (isFocused == "true" || forceHandleFocusItem) {
                result := handleFocusItem(node, focusItem)
                if result != nil {
                    return result
                }
            } else if node.Nodes != nil && focusItem == "" {
                childNode := findFocusedNode(node.Nodes, forceHandleFocusItem)
                if childNode != nil {
                    return childNode
                } else if isFocused == "true" {
                    return &node
                }
            } else if isFocused == "true" {
                return &node
            }
        }
    }
    return nil
 }

 func handleFocusItem(node ecp.Node, focusItem string) *ecp.Node {
    index, err := strconv.Atoi(focusItem)
    if err != nil {
        return nil
    }
    if (len(node.Nodes) <= index) || index < 0 {
        return nil
    }
    var childNode ecp.Node
    for _, childItem := range node.Nodes {
        nodeIndex := getAttributeValue(childItem, "index")
        if (focusItem == nodeIndex) {
            childNode = childItem
            break
        }
    }
    if childNode.Nodes != nil {
        childFocusedNode := findFocusedNode(childNode.Nodes, true)
        if childFocusedNode != nil {
            return childFocusedNode
        } else {
            return &childNode
        }
    } else {
        return &childNode
    }
 }

 func getAttributeValue(node ecp.Node, attribute string) string {
    for _, attrObj := range node.Attrs {
        if attrObj.Name.Local == attribute {
            return attrObj.Value
        }
    }
    return ""
 }

 func findMultipleNodes(nodes []ecp.Node, value string, checkChildNodes bool) []ecp.Node {
    var nodeArray []ecp.Node
    value = strings.ToLower(value)
    for _, node := range nodes {
        isVisible := strings.ToLower(getAttributeValue(node, "visible"))
        if (isVisible == "" || isVisible == "true") {
            if strings.ToLower(node.XMLName.Local) == value {
                nodeArray = append(nodeArray, node) 
            }
            if checkChildNodes && node.Nodes != nil {
                res := findMultipleNodes(node.Nodes, value, checkChildNodes)
                if res != nil {
                    nodeArray = append(nodeArray, res...)
                }
            }
        }
    }
    return nodeArray
 }

 func findMultipleNodesByText(nodes []ecp.Node, value string, attribute string, checkChildNodes bool) []ecp.Node  {
    var nodeArray []ecp.Node
    value = strings.ToLower(value)
    attribute = strings.ToLower(attribute)
    for _, node := range nodes {
        isVisible := strings.ToLower(getAttributeValue(node, "visible"))
        if (isVisible == "" || isVisible == "true") {
            if node.Attrs != nil {
                for _, attr := range node.Attrs {
                    attrValue := strings.ToLower(attr.Value)
                    if strings.ToLower(attr.Name.Local) == attribute && ((attribute == "text" && strings.Contains(attrValue, value) == true) || (attribute != "text" && attrValue == value)) {
                        nodeArray = append(nodeArray, node)
                    }
                }
            }
            if checkChildNodes && node.Nodes != nil {
                res := findMultipleNodesByText(node.Nodes, value, attribute, checkChildNodes)
                if res != nil {
                    nodeArray = append(nodeArray, res...)
                }
            }
        }
    }
    return nodeArray
 }