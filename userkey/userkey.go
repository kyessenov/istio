package userkey

import (
	"github.com/gogo/googleapis/google/rpc"
	"golang.org/x/net/context"

	"istio.io/api/mixer/adapter/model/v1beta1"
	"istio.io/istio/userkey/config"
)

type Userkey struct{}

var externalTable = map[string]string{"foobar": "jason"}

func (Userkey) HandleUserkey(_ context.Context, req *HandleUserkeyRequest) (*HandleUserkeyResponse, error) {
	config := &config.Params{}
	if err := config.Unmarshal(req.AdapterConfig.Value); err != nil {
		return nil, err
	}

	user, ok := externalTable[req.Instance.Key]
	if ok {
		return &HandleUserkeyResponse{
			Result: &v1beta1.CheckResult{ValidDuration: config.ValidDuration},
			Output: &OutputMsg{User: user},
		}, nil
	}

	return &HandleUserkeyResponse{
		Result: &v1beta1.CheckResult{
			Status: rpc.Status{Code: int32(rpc.PERMISSION_DENIED)},
		},
	}, nil
}
