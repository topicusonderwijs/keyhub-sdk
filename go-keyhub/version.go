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
