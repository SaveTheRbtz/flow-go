// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package access

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

// AccessAPIClient is the client API for AccessAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessAPIClient interface {
	// Ping is used to check if the access node is alive and healthy.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// GetLatestBlockHeader gets the latest sealed or unsealed block header.
	GetLatestBlockHeader(ctx context.Context, in *GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	// GetBlockHeaderByID gets a block header by ID.
	GetBlockHeaderByID(ctx context.Context, in *GetBlockHeaderByIDRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	// GetBlockHeaderByHeight gets a block header by height.
	GetBlockHeaderByHeight(ctx context.Context, in *GetBlockHeaderByHeightRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	// GetLatestBlock gets the full payload of the latest sealed or unsealed
	// block.
	GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	// GetBlockByID gets a full block by ID.
	GetBlockByID(ctx context.Context, in *GetBlockByIDRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	// GetBlockByHeight gets a full block by height.
	GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	// GetCollectionByID gets a collection by ID.
	GetCollectionByID(ctx context.Context, in *GetCollectionByIDRequest, opts ...grpc.CallOption) (*CollectionResponse, error)
	// SendTransaction submits a transaction to the network.
	SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
	// GetTransaction gets a transaction by ID.
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	// GetTransactionResult gets the result of a transaction.
	GetTransactionResult(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResultResponse, error)
	// GetAccount is an alias for GetAccountAtLatestBlock.
	//
	// Warning: this function is deprecated. It behaves identically to
	// GetAccountAtLatestBlock and will be removed in a future version.
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	// GetAccountAtLatestBlock gets an account by address from the latest sealed
	// execution state.
	GetAccountAtLatestBlock(ctx context.Context, in *GetAccountAtLatestBlockRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	// GetAccountAtBlockHeight gets an account by address at the given block
	// height
	GetAccountAtBlockHeight(ctx context.Context, in *GetAccountAtBlockHeightRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	// ExecuteScriptAtLatestBlock executes a read-only Cadence script against the
	// latest sealed execution state.
	ExecuteScriptAtLatestBlock(ctx context.Context, in *ExecuteScriptAtLatestBlockRequest, opts ...grpc.CallOption) (*ExecuteScriptResponse, error)
	// ExecuteScriptAtBlockID executes a ready-only Cadence script against the
	// execution state at the block with the given ID.
	ExecuteScriptAtBlockID(ctx context.Context, in *ExecuteScriptAtBlockIDRequest, opts ...grpc.CallOption) (*ExecuteScriptResponse, error)
	// ExecuteScriptAtBlockHeight executes a ready-only Cadence script against the
	// execution state at the given block height.
	ExecuteScriptAtBlockHeight(ctx context.Context, in *ExecuteScriptAtBlockHeightRequest, opts ...grpc.CallOption) (*ExecuteScriptResponse, error)
	// GetEventsForHeightRange retrieves events emitted within the specified block
	// range.
	GetEventsForHeightRange(ctx context.Context, in *GetEventsForHeightRangeRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	// GetEventsForBlockIDs retrieves events for the specified block IDs and event
	// type.
	GetEventsForBlockIDs(ctx context.Context, in *GetEventsForBlockIDsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	// GetNetworkParameters retrieves the Flow network details
	GetNetworkParameters(ctx context.Context, in *GetNetworkParametersRequest, opts ...grpc.CallOption) (*GetNetworkParametersResponse, error)
	// GetLatestProtocolStateSnapshot retrieves the latest sealed protocol state
	// snapshot. Used by Flow nodes joining the network to bootstrap a
	// space-efficient local state.
	GetLatestProtocolStateSnapshot(ctx context.Context, in *GetLatestProtocolStateSnapshotRequest, opts ...grpc.CallOption) (*ProtocolStateSnapshotResponse, error)
}

type accessAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessAPIClient(cc grpc.ClientConnInterface) AccessAPIClient {
	return &accessAPIClient{cc}
}

func (c *accessAPIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetLatestBlockHeader(ctx context.Context, in *GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetLatestBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetBlockHeaderByID(ctx context.Context, in *GetBlockHeaderByIDRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetBlockHeaderByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetBlockHeaderByHeight(ctx context.Context, in *GetBlockHeaderByHeightRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetBlockHeaderByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetBlockByID(ctx context.Context, in *GetBlockByIDRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetBlockByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetBlockByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetCollectionByID(ctx context.Context, in *GetCollectionByIDRequest, opts ...grpc.CallOption) (*CollectionResponse, error) {
	out := new(CollectionResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetCollectionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetTransactionResult(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*TransactionResultResponse, error) {
	out := new(TransactionResultResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetTransactionResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetAccountAtLatestBlock(ctx context.Context, in *GetAccountAtLatestBlockRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetAccountAtLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetAccountAtBlockHeight(ctx context.Context, in *GetAccountAtBlockHeightRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetAccountAtBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) ExecuteScriptAtLatestBlock(ctx context.Context, in *ExecuteScriptAtLatestBlockRequest, opts ...grpc.CallOption) (*ExecuteScriptResponse, error) {
	out := new(ExecuteScriptResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/ExecuteScriptAtLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) ExecuteScriptAtBlockID(ctx context.Context, in *ExecuteScriptAtBlockIDRequest, opts ...grpc.CallOption) (*ExecuteScriptResponse, error) {
	out := new(ExecuteScriptResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/ExecuteScriptAtBlockID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) ExecuteScriptAtBlockHeight(ctx context.Context, in *ExecuteScriptAtBlockHeightRequest, opts ...grpc.CallOption) (*ExecuteScriptResponse, error) {
	out := new(ExecuteScriptResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/ExecuteScriptAtBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetEventsForHeightRange(ctx context.Context, in *GetEventsForHeightRangeRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetEventsForHeightRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetEventsForBlockIDs(ctx context.Context, in *GetEventsForBlockIDsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetEventsForBlockIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetNetworkParameters(ctx context.Context, in *GetNetworkParametersRequest, opts ...grpc.CallOption) (*GetNetworkParametersResponse, error) {
	out := new(GetNetworkParametersResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetNetworkParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessAPIClient) GetLatestProtocolStateSnapshot(ctx context.Context, in *GetLatestProtocolStateSnapshotRequest, opts ...grpc.CallOption) (*ProtocolStateSnapshotResponse, error) {
	out := new(ProtocolStateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/flow.access.AccessAPI/GetLatestProtocolStateSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessAPIServer is the server API for AccessAPI service.
// All implementations must embed UnimplementedAccessAPIServer
// for forward compatibility
type AccessAPIServer interface {
	// Ping is used to check if the access node is alive and healthy.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// GetLatestBlockHeader gets the latest sealed or unsealed block header.
	GetLatestBlockHeader(context.Context, *GetLatestBlockHeaderRequest) (*BlockHeaderResponse, error)
	// GetBlockHeaderByID gets a block header by ID.
	GetBlockHeaderByID(context.Context, *GetBlockHeaderByIDRequest) (*BlockHeaderResponse, error)
	// GetBlockHeaderByHeight gets a block header by height.
	GetBlockHeaderByHeight(context.Context, *GetBlockHeaderByHeightRequest) (*BlockHeaderResponse, error)
	// GetLatestBlock gets the full payload of the latest sealed or unsealed
	// block.
	GetLatestBlock(context.Context, *GetLatestBlockRequest) (*BlockResponse, error)
	// GetBlockByID gets a full block by ID.
	GetBlockByID(context.Context, *GetBlockByIDRequest) (*BlockResponse, error)
	// GetBlockByHeight gets a full block by height.
	GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*BlockResponse, error)
	// GetCollectionByID gets a collection by ID.
	GetCollectionByID(context.Context, *GetCollectionByIDRequest) (*CollectionResponse, error)
	// SendTransaction submits a transaction to the network.
	SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error)
	// GetTransaction gets a transaction by ID.
	GetTransaction(context.Context, *GetTransactionRequest) (*TransactionResponse, error)
	// GetTransactionResult gets the result of a transaction.
	GetTransactionResult(context.Context, *GetTransactionRequest) (*TransactionResultResponse, error)
	// GetAccount is an alias for GetAccountAtLatestBlock.
	//
	// Warning: this function is deprecated. It behaves identically to
	// GetAccountAtLatestBlock and will be removed in a future version.
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	// GetAccountAtLatestBlock gets an account by address from the latest sealed
	// execution state.
	GetAccountAtLatestBlock(context.Context, *GetAccountAtLatestBlockRequest) (*AccountResponse, error)
	// GetAccountAtBlockHeight gets an account by address at the given block
	// height
	GetAccountAtBlockHeight(context.Context, *GetAccountAtBlockHeightRequest) (*AccountResponse, error)
	// ExecuteScriptAtLatestBlock executes a read-only Cadence script against the
	// latest sealed execution state.
	ExecuteScriptAtLatestBlock(context.Context, *ExecuteScriptAtLatestBlockRequest) (*ExecuteScriptResponse, error)
	// ExecuteScriptAtBlockID executes a ready-only Cadence script against the
	// execution state at the block with the given ID.
	ExecuteScriptAtBlockID(context.Context, *ExecuteScriptAtBlockIDRequest) (*ExecuteScriptResponse, error)
	// ExecuteScriptAtBlockHeight executes a ready-only Cadence script against the
	// execution state at the given block height.
	ExecuteScriptAtBlockHeight(context.Context, *ExecuteScriptAtBlockHeightRequest) (*ExecuteScriptResponse, error)
	// GetEventsForHeightRange retrieves events emitted within the specified block
	// range.
	GetEventsForHeightRange(context.Context, *GetEventsForHeightRangeRequest) (*EventsResponse, error)
	// GetEventsForBlockIDs retrieves events for the specified block IDs and event
	// type.
	GetEventsForBlockIDs(context.Context, *GetEventsForBlockIDsRequest) (*EventsResponse, error)
	// GetNetworkParameters retrieves the Flow network details
	GetNetworkParameters(context.Context, *GetNetworkParametersRequest) (*GetNetworkParametersResponse, error)
	// GetLatestProtocolStateSnapshot retrieves the latest sealed protocol state
	// snapshot. Used by Flow nodes joining the network to bootstrap a
	// space-efficient local state.
	GetLatestProtocolStateSnapshot(context.Context, *GetLatestProtocolStateSnapshotRequest) (*ProtocolStateSnapshotResponse, error)
	mustEmbedUnimplementedAccessAPIServer()
}

// UnimplementedAccessAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAccessAPIServer struct {
}

func (UnimplementedAccessAPIServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAccessAPIServer) GetLatestBlockHeader(context.Context, *GetLatestBlockHeaderRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlockHeader not implemented")
}
func (UnimplementedAccessAPIServer) GetBlockHeaderByID(context.Context, *GetBlockHeaderByIDRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByID not implemented")
}
func (UnimplementedAccessAPIServer) GetBlockHeaderByHeight(context.Context, *GetBlockHeaderByHeightRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByHeight not implemented")
}
func (UnimplementedAccessAPIServer) GetLatestBlock(context.Context, *GetLatestBlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (UnimplementedAccessAPIServer) GetBlockByID(context.Context, *GetBlockByIDRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByID not implemented")
}
func (UnimplementedAccessAPIServer) GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHeight not implemented")
}
func (UnimplementedAccessAPIServer) GetCollectionByID(context.Context, *GetCollectionByIDRequest) (*CollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionByID not implemented")
}
func (UnimplementedAccessAPIServer) SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedAccessAPIServer) GetTransaction(context.Context, *GetTransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedAccessAPIServer) GetTransactionResult(context.Context, *GetTransactionRequest) (*TransactionResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionResult not implemented")
}
func (UnimplementedAccessAPIServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccessAPIServer) GetAccountAtLatestBlock(context.Context, *GetAccountAtLatestBlockRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountAtLatestBlock not implemented")
}
func (UnimplementedAccessAPIServer) GetAccountAtBlockHeight(context.Context, *GetAccountAtBlockHeightRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountAtBlockHeight not implemented")
}
func (UnimplementedAccessAPIServer) ExecuteScriptAtLatestBlock(context.Context, *ExecuteScriptAtLatestBlockRequest) (*ExecuteScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteScriptAtLatestBlock not implemented")
}
func (UnimplementedAccessAPIServer) ExecuteScriptAtBlockID(context.Context, *ExecuteScriptAtBlockIDRequest) (*ExecuteScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteScriptAtBlockID not implemented")
}
func (UnimplementedAccessAPIServer) ExecuteScriptAtBlockHeight(context.Context, *ExecuteScriptAtBlockHeightRequest) (*ExecuteScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteScriptAtBlockHeight not implemented")
}
func (UnimplementedAccessAPIServer) GetEventsForHeightRange(context.Context, *GetEventsForHeightRangeRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForHeightRange not implemented")
}
func (UnimplementedAccessAPIServer) GetEventsForBlockIDs(context.Context, *GetEventsForBlockIDsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForBlockIDs not implemented")
}
func (UnimplementedAccessAPIServer) GetNetworkParameters(context.Context, *GetNetworkParametersRequest) (*GetNetworkParametersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkParameters not implemented")
}
func (UnimplementedAccessAPIServer) GetLatestProtocolStateSnapshot(context.Context, *GetLatestProtocolStateSnapshotRequest) (*ProtocolStateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestProtocolStateSnapshot not implemented")
}
func (UnimplementedAccessAPIServer) mustEmbedUnimplementedAccessAPIServer() {}

// UnsafeAccessAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessAPIServer will
// result in compilation errors.
type UnsafeAccessAPIServer interface {
	mustEmbedUnimplementedAccessAPIServer()
}

func RegisterAccessAPIServer(s grpc.ServiceRegistrar, srv AccessAPIServer) {
	s.RegisterService(&AccessAPI_ServiceDesc, srv)
}

func _AccessAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetLatestBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetLatestBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetLatestBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetLatestBlockHeader(ctx, req.(*GetLatestBlockHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetBlockHeaderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeaderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetBlockHeaderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetBlockHeaderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetBlockHeaderByID(ctx, req.(*GetBlockHeaderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetBlockHeaderByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeaderByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetBlockHeaderByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetBlockHeaderByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetBlockHeaderByHeight(ctx, req.(*GetBlockHeaderByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetLatestBlock(ctx, req.(*GetLatestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetBlockByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetBlockByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetBlockByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetBlockByID(ctx, req.(*GetBlockByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetBlockByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetBlockByHeight(ctx, req.(*GetBlockByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetCollectionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetCollectionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetCollectionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetCollectionByID(ctx, req.(*GetCollectionByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).SendTransaction(ctx, req.(*SendTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetTransactionResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetTransactionResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetTransactionResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetTransactionResult(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetAccountAtLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountAtLatestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetAccountAtLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetAccountAtLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetAccountAtLatestBlock(ctx, req.(*GetAccountAtLatestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetAccountAtBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountAtBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetAccountAtBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetAccountAtBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetAccountAtBlockHeight(ctx, req.(*GetAccountAtBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_ExecuteScriptAtLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteScriptAtLatestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).ExecuteScriptAtLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/ExecuteScriptAtLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).ExecuteScriptAtLatestBlock(ctx, req.(*ExecuteScriptAtLatestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_ExecuteScriptAtBlockID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteScriptAtBlockIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).ExecuteScriptAtBlockID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/ExecuteScriptAtBlockID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).ExecuteScriptAtBlockID(ctx, req.(*ExecuteScriptAtBlockIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_ExecuteScriptAtBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteScriptAtBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).ExecuteScriptAtBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/ExecuteScriptAtBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).ExecuteScriptAtBlockHeight(ctx, req.(*ExecuteScriptAtBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetEventsForHeightRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsForHeightRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetEventsForHeightRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetEventsForHeightRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetEventsForHeightRange(ctx, req.(*GetEventsForHeightRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetEventsForBlockIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsForBlockIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetEventsForBlockIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetEventsForBlockIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetEventsForBlockIDs(ctx, req.(*GetEventsForBlockIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetNetworkParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetNetworkParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetNetworkParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetNetworkParameters(ctx, req.(*GetNetworkParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessAPI_GetLatestProtocolStateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestProtocolStateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessAPIServer).GetLatestProtocolStateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.access.AccessAPI/GetLatestProtocolStateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessAPIServer).GetLatestProtocolStateSnapshot(ctx, req.(*GetLatestProtocolStateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessAPI_ServiceDesc is the grpc.ServiceDesc for AccessAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.access.AccessAPI",
	HandlerType: (*AccessAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AccessAPI_Ping_Handler,
		},
		{
			MethodName: "GetLatestBlockHeader",
			Handler:    _AccessAPI_GetLatestBlockHeader_Handler,
		},
		{
			MethodName: "GetBlockHeaderByID",
			Handler:    _AccessAPI_GetBlockHeaderByID_Handler,
		},
		{
			MethodName: "GetBlockHeaderByHeight",
			Handler:    _AccessAPI_GetBlockHeaderByHeight_Handler,
		},
		{
			MethodName: "GetLatestBlock",
			Handler:    _AccessAPI_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlockByID",
			Handler:    _AccessAPI_GetBlockByID_Handler,
		},
		{
			MethodName: "GetBlockByHeight",
			Handler:    _AccessAPI_GetBlockByHeight_Handler,
		},
		{
			MethodName: "GetCollectionByID",
			Handler:    _AccessAPI_GetCollectionByID_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _AccessAPI_SendTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _AccessAPI_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionResult",
			Handler:    _AccessAPI_GetTransactionResult_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccessAPI_GetAccount_Handler,
		},
		{
			MethodName: "GetAccountAtLatestBlock",
			Handler:    _AccessAPI_GetAccountAtLatestBlock_Handler,
		},
		{
			MethodName: "GetAccountAtBlockHeight",
			Handler:    _AccessAPI_GetAccountAtBlockHeight_Handler,
		},
		{
			MethodName: "ExecuteScriptAtLatestBlock",
			Handler:    _AccessAPI_ExecuteScriptAtLatestBlock_Handler,
		},
		{
			MethodName: "ExecuteScriptAtBlockID",
			Handler:    _AccessAPI_ExecuteScriptAtBlockID_Handler,
		},
		{
			MethodName: "ExecuteScriptAtBlockHeight",
			Handler:    _AccessAPI_ExecuteScriptAtBlockHeight_Handler,
		},
		{
			MethodName: "GetEventsForHeightRange",
			Handler:    _AccessAPI_GetEventsForHeightRange_Handler,
		},
		{
			MethodName: "GetEventsForBlockIDs",
			Handler:    _AccessAPI_GetEventsForBlockIDs_Handler,
		},
		{
			MethodName: "GetNetworkParameters",
			Handler:    _AccessAPI_GetNetworkParameters_Handler,
		},
		{
			MethodName: "GetLatestProtocolStateSnapshot",
			Handler:    _AccessAPI_GetLatestProtocolStateSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flow/access/access.proto",
}
