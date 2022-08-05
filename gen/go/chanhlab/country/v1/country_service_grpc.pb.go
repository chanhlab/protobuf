// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: chanhlab/country/v1/country_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CountryServiceClient is the client API for CountryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountryServiceClient interface {
	// Country list.
	ListCountries(ctx context.Context, in *ListCountriesRequest, opts ...grpc.CallOption) (*ListCountriesResponse, error)
	// Get Country.
	GetCountry(ctx context.Context, in *GetCountryRequest, opts ...grpc.CallOption) (*GetCountryResponse, error)
	// Create Country.
	CreateCountry(ctx context.Context, in *CreateCountryRequest, opts ...grpc.CallOption) (*CreateCountryResponse, error)
	// Update Country.
	UpdateCountry(ctx context.Context, in *UpdateCountryRequest, opts ...grpc.CallOption) (*UpdateCountryResponse, error)
	// Delete Country.
	DeleteCountry(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountryResponse, error)
}

type countryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryServiceClient(cc grpc.ClientConnInterface) CountryServiceClient {
	return &countryServiceClient{cc}
}

func (c *countryServiceClient) ListCountries(ctx context.Context, in *ListCountriesRequest, opts ...grpc.CallOption) (*ListCountriesResponse, error) {
	out := new(ListCountriesResponse)
	err := c.cc.Invoke(ctx, "/chanhlab.country.v1.CountryService/ListCountries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) GetCountry(ctx context.Context, in *GetCountryRequest, opts ...grpc.CallOption) (*GetCountryResponse, error) {
	out := new(GetCountryResponse)
	err := c.cc.Invoke(ctx, "/chanhlab.country.v1.CountryService/GetCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) CreateCountry(ctx context.Context, in *CreateCountryRequest, opts ...grpc.CallOption) (*CreateCountryResponse, error) {
	out := new(CreateCountryResponse)
	err := c.cc.Invoke(ctx, "/chanhlab.country.v1.CountryService/CreateCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) UpdateCountry(ctx context.Context, in *UpdateCountryRequest, opts ...grpc.CallOption) (*UpdateCountryResponse, error) {
	out := new(UpdateCountryResponse)
	err := c.cc.Invoke(ctx, "/chanhlab.country.v1.CountryService/UpdateCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) DeleteCountry(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountryResponse, error) {
	out := new(DeleteCountryResponse)
	err := c.cc.Invoke(ctx, "/chanhlab.country.v1.CountryService/DeleteCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServiceServer is the server API for CountryService service.
// All implementations must embed UnimplementedCountryServiceServer
// for forward compatibility
type CountryServiceServer interface {
	// Country list.
	ListCountries(context.Context, *ListCountriesRequest) (*ListCountriesResponse, error)
	// Get Country.
	GetCountry(context.Context, *GetCountryRequest) (*GetCountryResponse, error)
	// Create Country.
	CreateCountry(context.Context, *CreateCountryRequest) (*CreateCountryResponse, error)
	// Update Country.
	UpdateCountry(context.Context, *UpdateCountryRequest) (*UpdateCountryResponse, error)
	// Delete Country.
	DeleteCountry(context.Context, *DeleteCountryRequest) (*DeleteCountryResponse, error)
	mustEmbedUnimplementedCountryServiceServer()
}

// UnimplementedCountryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCountryServiceServer struct {
}

func (UnimplementedCountryServiceServer) ListCountries(context.Context, *ListCountriesRequest) (*ListCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCountries not implemented")
}
func (UnimplementedCountryServiceServer) GetCountry(context.Context, *GetCountryRequest) (*GetCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountry not implemented")
}
func (UnimplementedCountryServiceServer) CreateCountry(context.Context, *CreateCountryRequest) (*CreateCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCountry not implemented")
}
func (UnimplementedCountryServiceServer) UpdateCountry(context.Context, *UpdateCountryRequest) (*UpdateCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCountry not implemented")
}
func (UnimplementedCountryServiceServer) DeleteCountry(context.Context, *DeleteCountryRequest) (*DeleteCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCountry not implemented")
}
func (UnimplementedCountryServiceServer) mustEmbedUnimplementedCountryServiceServer() {}

// UnsafeCountryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountryServiceServer will
// result in compilation errors.
type UnsafeCountryServiceServer interface {
	mustEmbedUnimplementedCountryServiceServer()
}

func RegisterCountryServiceServer(s grpc.ServiceRegistrar, srv CountryServiceServer) {
	s.RegisterService(&CountryService_ServiceDesc, srv)
}

func _CountryService_ListCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).ListCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chanhlab.country.v1.CountryService/ListCountries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).ListCountries(ctx, req.(*ListCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_GetCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).GetCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chanhlab.country.v1.CountryService/GetCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).GetCountry(ctx, req.(*GetCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_CreateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).CreateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chanhlab.country.v1.CountryService/CreateCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).CreateCountry(ctx, req.(*CreateCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_UpdateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).UpdateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chanhlab.country.v1.CountryService/UpdateCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).UpdateCountry(ctx, req.(*UpdateCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_DeleteCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).DeleteCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chanhlab.country.v1.CountryService/DeleteCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).DeleteCountry(ctx, req.(*DeleteCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CountryService_ServiceDesc is the grpc.ServiceDesc for CountryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chanhlab.country.v1.CountryService",
	HandlerType: (*CountryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCountries",
			Handler:    _CountryService_ListCountries_Handler,
		},
		{
			MethodName: "GetCountry",
			Handler:    _CountryService_GetCountry_Handler,
		},
		{
			MethodName: "CreateCountry",
			Handler:    _CountryService_CreateCountry_Handler,
		},
		{
			MethodName: "UpdateCountry",
			Handler:    _CountryService_UpdateCountry_Handler,
		},
		{
			MethodName: "DeleteCountry",
			Handler:    _CountryService_DeleteCountry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chanhlab/country/v1/country_service.proto",
}
