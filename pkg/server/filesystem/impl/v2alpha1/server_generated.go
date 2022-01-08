// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/api/filesystem/v2alpha1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/filesystem/impl"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v2alpha1")

type versionedAPI struct {
	apiGroupServer impl.ServerInterface
}

func NewVersionedServer(apiGroupServer impl.ServerInterface) impl.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v2alpha1.RegisterFilesystemServer(grpcServer, s)
}

func (s *versionedAPI) CreateSymlink(context context.Context, versionedRequest *v2alpha1.CreateSymlinkRequest) (*v2alpha1.CreateSymlinkResponse, error) {
	request := &impl.CreateSymlinkRequest{}
	if err := Convert_v2alpha1_CreateSymlinkRequest_To_impl_CreateSymlinkRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.CreateSymlink(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v2alpha1.CreateSymlinkResponse{}
	if err := Convert_impl_CreateSymlinkResponse_To_v2alpha1_CreateSymlinkResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) IsSymlink(context context.Context, versionedRequest *v2alpha1.IsSymlinkRequest) (*v2alpha1.IsSymlinkResponse, error) {
	request := &impl.IsSymlinkRequest{}
	if err := Convert_v2alpha1_IsSymlinkRequest_To_impl_IsSymlinkRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.IsSymlink(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v2alpha1.IsSymlinkResponse{}
	if err := Convert_impl_IsSymlinkResponse_To_v2alpha1_IsSymlinkResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) Mkdir(context context.Context, versionedRequest *v2alpha1.MkdirRequest) (*v2alpha1.MkdirResponse, error) {
	request := &impl.MkdirRequest{}
	if err := Convert_v2alpha1_MkdirRequest_To_impl_MkdirRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.Mkdir(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v2alpha1.MkdirResponse{}
	if err := Convert_impl_MkdirResponse_To_v2alpha1_MkdirResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) PathExists(context context.Context, versionedRequest *v2alpha1.PathExistsRequest) (*v2alpha1.PathExistsResponse, error) {
	request := &impl.PathExistsRequest{}
	if err := Convert_v2alpha1_PathExistsRequest_To_impl_PathExistsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.PathExists(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v2alpha1.PathExistsResponse{}
	if err := Convert_impl_PathExistsResponse_To_v2alpha1_PathExistsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) Rmdir(context context.Context, versionedRequest *v2alpha1.RmdirRequest) (*v2alpha1.RmdirResponse, error) {
	request := &impl.RmdirRequest{}
	if err := Convert_v2alpha1_RmdirRequest_To_impl_RmdirRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.Rmdir(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v2alpha1.RmdirResponse{}
	if err := Convert_impl_RmdirResponse_To_v2alpha1_RmdirResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
