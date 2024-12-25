// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: inventory.proto

package inventory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Inventory_GetStock_FullMethodName         = "/inventory.Inventory/GetStock"
	Inventory_UpdateStock_FullMethodName      = "/inventory.Inventory/UpdateStock"
	Inventory_BatchGetStock_FullMethodName    = "/inventory.Inventory/BatchGetStock"
	Inventory_LockStock_FullMethodName        = "/inventory.Inventory/LockStock"
	Inventory_UnlockStock_FullMethodName      = "/inventory.Inventory/UnlockStock"
	Inventory_DeductStock_FullMethodName      = "/inventory.Inventory/DeductStock"
	Inventory_CreateWarehouse_FullMethodName  = "/inventory.Inventory/CreateWarehouse"
	Inventory_UpdateWarehouse_FullMethodName  = "/inventory.Inventory/UpdateWarehouse"
	Inventory_ListWarehouses_FullMethodName   = "/inventory.Inventory/ListWarehouses"
	Inventory_CreateStockIn_FullMethodName    = "/inventory.Inventory/CreateStockIn"
	Inventory_CreateStockOut_FullMethodName   = "/inventory.Inventory/CreateStockOut"
	Inventory_ListStockRecords_FullMethodName = "/inventory.Inventory/ListStockRecords"
)

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	// 库存管理
	GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error)
	BatchGetStock(ctx context.Context, in *BatchGetStockRequest, opts ...grpc.CallOption) (*BatchGetStockResponse, error)
	// 库存锁定/解锁
	LockStock(ctx context.Context, in *LockStockRequest, opts ...grpc.CallOption) (*LockStockResponse, error)
	UnlockStock(ctx context.Context, in *UnlockStockRequest, opts ...grpc.CallOption) (*UnlockStockResponse, error)
	DeductStock(ctx context.Context, in *DeductStockRequest, opts ...grpc.CallOption) (*DeductStockResponse, error)
	// 仓储管理
	CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error)
	UpdateWarehouse(ctx context.Context, in *UpdateWarehouseRequest, opts ...grpc.CallOption) (*UpdateWarehouseResponse, error)
	ListWarehouses(ctx context.Context, in *ListWarehousesRequest, opts ...grpc.CallOption) (*ListWarehousesResponse, error)
	// 入库/出库
	CreateStockIn(ctx context.Context, in *CreateStockInRequest, opts ...grpc.CallOption) (*CreateStockInResponse, error)
	CreateStockOut(ctx context.Context, in *CreateStockOutRequest, opts ...grpc.CallOption) (*CreateStockOutResponse, error)
	ListStockRecords(ctx context.Context, in *ListStockRecordsRequest, opts ...grpc.CallOption) (*ListStockRecordsResponse, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStockResponse)
	err := c.cc.Invoke(ctx, Inventory_GetStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStockResponse)
	err := c.cc.Invoke(ctx, Inventory_UpdateStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) BatchGetStock(ctx context.Context, in *BatchGetStockRequest, opts ...grpc.CallOption) (*BatchGetStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetStockResponse)
	err := c.cc.Invoke(ctx, Inventory_BatchGetStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) LockStock(ctx context.Context, in *LockStockRequest, opts ...grpc.CallOption) (*LockStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LockStockResponse)
	err := c.cc.Invoke(ctx, Inventory_LockStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) UnlockStock(ctx context.Context, in *UnlockStockRequest, opts ...grpc.CallOption) (*UnlockStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnlockStockResponse)
	err := c.cc.Invoke(ctx, Inventory_UnlockStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) DeductStock(ctx context.Context, in *DeductStockRequest, opts ...grpc.CallOption) (*DeductStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeductStockResponse)
	err := c.cc.Invoke(ctx, Inventory_DeductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) CreateWarehouse(ctx context.Context, in *CreateWarehouseRequest, opts ...grpc.CallOption) (*CreateWarehouseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWarehouseResponse)
	err := c.cc.Invoke(ctx, Inventory_CreateWarehouse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) UpdateWarehouse(ctx context.Context, in *UpdateWarehouseRequest, opts ...grpc.CallOption) (*UpdateWarehouseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateWarehouseResponse)
	err := c.cc.Invoke(ctx, Inventory_UpdateWarehouse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) ListWarehouses(ctx context.Context, in *ListWarehousesRequest, opts ...grpc.CallOption) (*ListWarehousesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListWarehousesResponse)
	err := c.cc.Invoke(ctx, Inventory_ListWarehouses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) CreateStockIn(ctx context.Context, in *CreateStockInRequest, opts ...grpc.CallOption) (*CreateStockInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStockInResponse)
	err := c.cc.Invoke(ctx, Inventory_CreateStockIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) CreateStockOut(ctx context.Context, in *CreateStockOutRequest, opts ...grpc.CallOption) (*CreateStockOutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStockOutResponse)
	err := c.cc.Invoke(ctx, Inventory_CreateStockOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) ListStockRecords(ctx context.Context, in *ListStockRecordsRequest, opts ...grpc.CallOption) (*ListStockRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStockRecordsResponse)
	err := c.cc.Invoke(ctx, Inventory_ListStockRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations must embed UnimplementedInventoryServer
// for forward compatibility.
type InventoryServer interface {
	// 库存管理
	GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error)
	UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error)
	BatchGetStock(context.Context, *BatchGetStockRequest) (*BatchGetStockResponse, error)
	// 库存锁定/解锁
	LockStock(context.Context, *LockStockRequest) (*LockStockResponse, error)
	UnlockStock(context.Context, *UnlockStockRequest) (*UnlockStockResponse, error)
	DeductStock(context.Context, *DeductStockRequest) (*DeductStockResponse, error)
	// 仓储管理
	CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	UpdateWarehouse(context.Context, *UpdateWarehouseRequest) (*UpdateWarehouseResponse, error)
	ListWarehouses(context.Context, *ListWarehousesRequest) (*ListWarehousesResponse, error)
	// 入库/出库
	CreateStockIn(context.Context, *CreateStockInRequest) (*CreateStockInResponse, error)
	CreateStockOut(context.Context, *CreateStockOutRequest) (*CreateStockOutResponse, error)
	ListStockRecords(context.Context, *ListStockRecordsRequest) (*ListStockRecordsResponse, error)
	mustEmbedUnimplementedInventoryServer()
}

// UnimplementedInventoryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInventoryServer struct{}

func (UnimplementedInventoryServer) GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedInventoryServer) UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedInventoryServer) BatchGetStock(context.Context, *BatchGetStockRequest) (*BatchGetStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetStock not implemented")
}
func (UnimplementedInventoryServer) LockStock(context.Context, *LockStockRequest) (*LockStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockStock not implemented")
}
func (UnimplementedInventoryServer) UnlockStock(context.Context, *UnlockStockRequest) (*UnlockStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockStock not implemented")
}
func (UnimplementedInventoryServer) DeductStock(context.Context, *DeductStockRequest) (*DeductStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStock not implemented")
}
func (UnimplementedInventoryServer) CreateWarehouse(context.Context, *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedInventoryServer) UpdateWarehouse(context.Context, *UpdateWarehouseRequest) (*UpdateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWarehouse not implemented")
}
func (UnimplementedInventoryServer) ListWarehouses(context.Context, *ListWarehousesRequest) (*ListWarehousesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWarehouses not implemented")
}
func (UnimplementedInventoryServer) CreateStockIn(context.Context, *CreateStockInRequest) (*CreateStockInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStockIn not implemented")
}
func (UnimplementedInventoryServer) CreateStockOut(context.Context, *CreateStockOutRequest) (*CreateStockOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStockOut not implemented")
}
func (UnimplementedInventoryServer) ListStockRecords(context.Context, *ListStockRecordsRequest) (*ListStockRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStockRecords not implemented")
}
func (UnimplementedInventoryServer) mustEmbedUnimplementedInventoryServer() {}
func (UnimplementedInventoryServer) testEmbeddedByValue()                   {}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	// If the following call pancis, it indicates UnimplementedInventoryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_GetStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetStock(ctx, req.(*GetStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_BatchGetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).BatchGetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_BatchGetStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).BatchGetStock(ctx, req.(*BatchGetStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_LockStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).LockStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_LockStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).LockStock(ctx, req.(*LockStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_UnlockStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).UnlockStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_UnlockStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).UnlockStock(ctx, req.(*UnlockStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_DeductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).DeductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_DeductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).DeductStock(ctx, req.(*DeductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_CreateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).CreateWarehouse(ctx, req.(*CreateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_UpdateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).UpdateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_UpdateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).UpdateWarehouse(ctx, req.(*UpdateWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_ListWarehouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWarehousesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).ListWarehouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_ListWarehouses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).ListWarehouses(ctx, req.(*ListWarehousesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_CreateStockIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStockInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).CreateStockIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_CreateStockIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).CreateStockIn(ctx, req.(*CreateStockInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_CreateStockOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStockOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).CreateStockOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_CreateStockOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).CreateStockOut(ctx, req.(*CreateStockOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_ListStockRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStockRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).ListStockRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_ListStockRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).ListStockRecords(ctx, req.(*ListStockRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStock",
			Handler:    _Inventory_GetStock_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _Inventory_UpdateStock_Handler,
		},
		{
			MethodName: "BatchGetStock",
			Handler:    _Inventory_BatchGetStock_Handler,
		},
		{
			MethodName: "LockStock",
			Handler:    _Inventory_LockStock_Handler,
		},
		{
			MethodName: "UnlockStock",
			Handler:    _Inventory_UnlockStock_Handler,
		},
		{
			MethodName: "DeductStock",
			Handler:    _Inventory_DeductStock_Handler,
		},
		{
			MethodName: "CreateWarehouse",
			Handler:    _Inventory_CreateWarehouse_Handler,
		},
		{
			MethodName: "UpdateWarehouse",
			Handler:    _Inventory_UpdateWarehouse_Handler,
		},
		{
			MethodName: "ListWarehouses",
			Handler:    _Inventory_ListWarehouses_Handler,
		},
		{
			MethodName: "CreateStockIn",
			Handler:    _Inventory_CreateStockIn_Handler,
		},
		{
			MethodName: "CreateStockOut",
			Handler:    _Inventory_CreateStockOut_Handler,
		},
		{
			MethodName: "ListStockRecords",
			Handler:    _Inventory_ListStockRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}