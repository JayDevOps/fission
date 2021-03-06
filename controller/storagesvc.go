/*
Copyright 2017 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (api *API) StorageServiceProxy(w http.ResponseWriter, r *http.Request) {
	u := api.storageServiceUrl
	ssUrl, err := url.Parse(u)
	if err != nil {
		msg := fmt.Sprintf("Error parsing url %v: %v", u, err)
		log.Println(msg)
		http.Error(w, msg, 500)
		return
	}
	director := func(req *http.Request) {
		req.URL.Scheme = ssUrl.Scheme
		req.URL.Host = ssUrl.Host
		req.URL.Path = "/v1/archive"
	}
	proxy := &httputil.ReverseProxy{
		Director: director,
	}
	proxy.ServeHTTP(w, r)
}
