//Package service is where we handle every incoming services that are defined in router package
package service

// Elasthink, An alternative to elasticsearch engine written in Go for small set of documents that uses inverted index to build the index and utilizes redis to store the indexes.
// Copyright (C) 2020 Yuwono Bangun Nagoro (a.k.a SurgicalSteel)
//
// This file is part of Elasthink
//
// Elasthink is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Elasthink is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
import (
	"github.com/SurgicalSteel/elasthink/module"
	"net/http"
)

//HandlePing is the handler for a ping endpoint
func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
}

//ResponsePayload is the response payload struct for all API
type ResponsePayload struct {
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

func constructResponsePayload(rawModuleResponse module.Response) ResponsePayload {
	return ResponsePayload{
		ErrorMessage: rawModuleResponse.ErrorMessage,
		Data:         rawModuleResponse.Data,
	}
}
