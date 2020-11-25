// Copyright © 2018 Alex Kolbasov
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

const (
	hmsPortDefault = 9083

	jsonEncoding = "application/json; charset=UTF-8"

	paramMetadataUri = "metadataUri"
	paramDbName      = "dbName"
	paramTblName     = "tableName"
	paramPartName    = "partName"
)

var (
	webPort int
	hmsPort int
)

func main() {
	flag.IntVar(&hmsPort, "hmsport", hmsPortDefault, "HMS Thrift port")
	flag.IntVar(&webPort, "port", 8080, "web service port")
	flag.Parse()

	router := mux.NewRouter()

	// Show all routes as top-level index
	router.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%s\n", t)
			return nil
		})
	})

	router.HandleFunc("/", showHelp)
	router.HandleFunc("/{metadataUri}/databases", databaseList)
	router.HandleFunc("/{metadataUri}", databaseList)
	router.HandleFunc("/{metadataUri}/databases/{dbName}", databaseShow).Methods("GET")
	router.HandleFunc("/{metadataUri}/{dbName}", databaseShow).Methods("GET")
	router.HandleFunc("/{metadataUri}/databases/{dbName}", databaseCreate).Methods("POST")
	router.HandleFunc("/{metadataUri}/databases/{dbName}", databaseDrop).Methods("DELETE")
	router.HandleFunc("/{metadataUri}/{dbName}", databaseDrop).Methods("DELETE")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/", tablesList).Methods("GET")
	router.HandleFunc("/{metadataUri}/{dbName}/", tablesList).Methods("GET")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}", tablesShow).Methods("GET")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}", tablesShow).Methods("GET")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}", tableCreate).Methods("POST")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}", tableDrop).Methods("DELETE")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}", tableCreate).Methods("POST")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}", tableDrop).Methods("DELETE")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}/", partitionsList).Methods("GET")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}/", partitionsList).Methods("GET")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}/{partName}", partitionShow).Methods("GET")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}/{partName}", partitionShow).Methods("GET")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}/", partitionAdd).Methods("POST")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}/", partitionAdd).Methods("POST")
	router.HandleFunc("/{metadataUri}/{dbName}/{tableName}/{partName}", partitionDrop).Methods("DELETE")
	router.HandleFunc("/{metadataUri}/databases/{dbName}/{tableName}/{partName}", partitionDrop).Methods("DELETE")

	log.Printf("Start to listen at :%d\n", webPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), router))
}
