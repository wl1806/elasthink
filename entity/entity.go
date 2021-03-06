//Package entity is the package where we store and manage elasthink's core entity
package entity

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

type entityData struct {
	documentTypes map[DocumentType]int
	stopwordData  StopwordData
}

var Entity entityData

func (e *entityData) Initialize(stopwordData StopwordData) {
	e.documentTypes = map[DocumentType]int{
		CampaignDocument:              1,
		AdvertisementCampaignDocument: 1,
	}
	e.stopwordData = stopwordData
}

func (e *entityData) GetDocumentTypes() map[DocumentType]int {
	return e.documentTypes
}

func (e *entityData) GetStopwordData() StopwordData {
	return e.stopwordData
}
