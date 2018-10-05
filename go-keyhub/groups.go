/* Licensed to the Apache Software Foundation (ASF) under one or more
   contributor license agreements.  See the NOTICE file distributed with
   this work for additional information regarding copyright ownership.
   The ASF licenses this file to You under the Apache License, Version 2.0
   (the "License"); you may not use this file except in compliance with
   the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License. */

package keyhub

import (
	"errors"

	"github.com/dghubble/sling"
)

type groupList struct {
	Items []Group `json:"items"`
}

type Group struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name:`
	Links []struct {
		ID   int    `json:"id"`
		Rel  string `json:"rel"`
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"links"`
}

type GroupService struct {
	sling *sling.Sling
}

type GroupParams struct {
	UUID string `url:"uuid,omitempty"`
}

func newGroupService(sling *sling.Sling) *GroupService {
	return &GroupService{
		sling: sling.Path("/keyhub/rest/v1/group"),
	}
}

func (s *GroupService) List() (groups []Group, err error) {
	gl := new(groupList)
	_, err = s.sling.New().Get("").ReceiveSuccess(gl)
	groups = gl.Items
	return
}

func (s *GroupService) Get(uuid string) (g *Group, err error) {
	gl := new(groupList)
	params := &GroupParams{UUID: uuid}
	_, err = s.sling.New().Get("").QueryStruct(params).ReceiveSuccess(gl)
	if err == nil {
		if len(gl.Items) > 0 {
			g = &gl.Items[0]
		} else {
			err = errors.New("Group '" + uuid + "' not found!")
		}
	}
	return
}
