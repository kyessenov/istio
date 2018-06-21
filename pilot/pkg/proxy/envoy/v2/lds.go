// Copyright 2017 Istio Authors
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

package v2

import (
	"fmt"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/gogo/protobuf/types"
	"github.com/prometheus/client_golang/prometheus"

	"istio.io/istio/pilot/pkg/model"
)

func (s *DiscoveryServer) pushLds(con *XdsConnection) error {
	// TODO: Modify interface to take services, and config instead of making library query registry
	ls, err := s.ConfigGenerator.BuildListeners(s.env, *con.modelNode)
	if err != nil {
		adsLog.Warnf("LDS: Failed to generate listeners for node %s: %v", con.modelNode, err)
		pushes.With(prometheus.Labels{"type": "lds_builderr"}).Add(1)
		return err
	}

	for _, l := range ls {
		if err = l.Validate(); err != nil {
			retErr := fmt.Errorf("LDS: Generated invalid listener for node %s: %v", con.modelNode, err)
			adsLog.Errorf("LDS: Generated invalid listener for node %s: %v, %v", con.modelNode, err, l)
			pushes.With(prometheus.Labels{"type": "lds_builderr"}).Add(1)
			// Generating invalid listeners is a bug.
			// Panic instead of trying to recover from that, since we can't
			// assume anything about the state.
			panic(retErr.Error())
		}
	}

	con.HTTPListeners = ls
	response := ldsDiscoveryResponse(ls, *con.modelNode)
	err = con.send(response)
	if err != nil {
		adsLog.Warnf("LDS: Send failure, closing grpc %v", err)
		pushes.With(prometheus.Labels{"type": "lds_senderr"}).Add(1)
		return err
	}
	pushes.With(prometheus.Labels{"type": "lds"}).Add(1)

	adsLog.Infof("LDS: PUSH for node:%s addr:%q listeners:%d", con.modelNode, con.PeerAddr, len(ls))
	return nil
}

// LdsDiscoveryResponse returns a list of listeners for the given environment and source node.
func ldsDiscoveryResponse(ls []*xdsapi.Listener, node model.Proxy) *xdsapi.DiscoveryResponse {
	resp := &xdsapi.DiscoveryResponse{
		TypeUrl:     ListenerType,
		VersionInfo: versionInfo(),
		Nonce:       nonce(),
	}
	for _, ll := range ls {
		if ll == nil {
			adsLog.Errora("Nil listener ", ll)
			continue
		}
		lr, _ := types.MarshalAny(ll)
		resp.Resources = append(resp.Resources, *lr)
	}

	return resp
}
