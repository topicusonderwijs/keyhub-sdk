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
	"github.com/dghubble/sling"
)

type Version struct {
	KeyhubVersion    string `json:"keyHubVersion"`
	ContractVersions []int  `json:"contractVersions"`
}

type VersionService struct {
	sling *sling.Sling
}

func newVersionService(sling *sling.Sling) *VersionService {
	return &VersionService{
		sling: sling.Path("/keyhub/rest/v1/info"),
	}
}

func (s *VersionService) Get() (v *Version, err error) {
	v = new(Version)
	_, err = s.sling.New().Get("").ReceiveSuccess(v)
	return
}
