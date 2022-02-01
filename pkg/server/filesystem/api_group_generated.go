// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package filesystem

import (
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl/v1"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl/v1beta1"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl/v1beta2"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl/v2alpha1"
	srvtypes "github.com/kubernetes-csi/csi-proxy/pkg/server/types"
)

const name = "filesystem"

// ensure the server defines all the required methods
var _ impl.ServerInterface = &Server{}

func (s *Server) VersionedAPIs() []*srvtypes.VersionedAPI {
	v1alpha1Server := v1alpha1.NewVersionedServer(s)
	v1beta1Server := v1beta1.NewVersionedServer(s)
	v1beta2Server := v1beta2.NewVersionedServer(s)
	v1Server := v1.NewVersionedServer(s)
	v2alpha1Server := v2alpha1.NewVersionedServer(s)

	return []*srvtypes.VersionedAPI{
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha1"),
			Registrant: v1alpha1Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1beta1"),
			Registrant: v1beta1Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1beta2"),
			Registrant: v1beta2Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1"),
			Registrant: v1Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v2alpha1"),
			Registrant: v2alpha1Server.Register,
		},
	}
}
