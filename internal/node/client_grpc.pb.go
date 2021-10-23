// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node

import (
	context "context"
	api "github.com/sonr-io/core/internal/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	// Node Methods
	// Signing Method Request for Data
	Supply(ctx context.Context, in *api.SupplyRequest, opts ...grpc.CallOption) (*api.SupplyResponse, error)
	// Verification Method Request for Signed Data
	Edit(ctx context.Context, in *api.EditRequest, opts ...grpc.CallOption) (*api.EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(ctx context.Context, in *api.FetchRequest, opts ...grpc.CallOption) (*api.FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(ctx context.Context, in *api.ShareRequest, opts ...grpc.CallOption) (*api.ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(ctx context.Context, in *api.RespondRequest, opts ...grpc.CallOption) (*api.RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(ctx context.Context, in *api.SearchRequest, opts ...grpc.CallOption) (*api.SearchResponse, error)
	// Events Streams
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnLobbyRefreshClient, error)
	// Returns a stream of Mailbox Message Events
	OnMailboxMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnMailboxMessageClient, error)
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransferAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferAcceptedClient, error)
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransferDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferDeclinedClient, error)
	// Returns a stream of DecisionEvent's for Invites
	OnTransferInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferInviteClient, error)
	// Returns a stream of ProgressEvent's for Sessions
	OnTransferProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferProgressClient, error)
	// Returns a stream of Completed Transfers
	OnTransferComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferCompleteClient, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Supply(ctx context.Context, in *api.SupplyRequest, opts ...grpc.CallOption) (*api.SupplyResponse, error) {
	out := new(api.SupplyResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Edit(ctx context.Context, in *api.EditRequest, opts ...grpc.CallOption) (*api.EditResponse, error) {
	out := new(api.EditResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Fetch(ctx context.Context, in *api.FetchRequest, opts ...grpc.CallOption) (*api.FetchResponse, error) {
	out := new(api.FetchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Share(ctx context.Context, in *api.ShareRequest, opts ...grpc.CallOption) (*api.ShareResponse, error) {
	out := new(api.ShareResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Respond(ctx context.Context, in *api.RespondRequest, opts ...grpc.CallOption) (*api.RespondResponse, error) {
	out := new(api.RespondResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Respond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Search(ctx context.Context, in *api.SearchRequest, opts ...grpc.CallOption) (*api.SearchResponse, error) {
	out := new(api.SearchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.ClientService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnLobbyRefreshClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], "/sonr.node.ClientService/OnLobbyRefresh", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnLobbyRefreshClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnLobbyRefreshClient interface {
	Recv() (*api.RefreshEvent, error)
	grpc.ClientStream
}

type clientServiceOnLobbyRefreshClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnLobbyRefreshClient) Recv() (*api.RefreshEvent, error) {
	m := new(api.RefreshEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnMailboxMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnMailboxMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[1], "/sonr.node.ClientService/OnMailboxMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnMailboxMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnMailboxMessageClient interface {
	Recv() (*api.MailboxEvent, error)
	grpc.ClientStream
}

type clientServiceOnMailboxMessageClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnMailboxMessageClient) Recv() (*api.MailboxEvent, error) {
	m := new(api.MailboxEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferAcceptedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[2], "/sonr.node.ClientService/OnTransferAccepted", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferAcceptedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferAcceptedClient interface {
	Recv() (*api.DecisionEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferAcceptedClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferAcceptedClient) Recv() (*api.DecisionEvent, error) {
	m := new(api.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferDeclinedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[3], "/sonr.node.ClientService/OnTransferDeclined", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferDeclinedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferDeclinedClient interface {
	Recv() (*api.DecisionEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferDeclinedClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferDeclinedClient) Recv() (*api.DecisionEvent, error) {
	m := new(api.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferInviteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[4], "/sonr.node.ClientService/OnTransferInvite", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferInviteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferInviteClient interface {
	Recv() (*api.InviteEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferInviteClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferInviteClient) Recv() (*api.InviteEvent, error) {
	m := new(api.InviteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[5], "/sonr.node.ClientService/OnTransferProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferProgressClient interface {
	Recv() (*api.ProgressEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferProgressClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferProgressClient) Recv() (*api.ProgressEvent, error) {
	m := new(api.ProgressEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) OnTransferComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ClientService_OnTransferCompleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[6], "/sonr.node.ClientService/OnTransferComplete", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceOnTransferCompleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_OnTransferCompleteClient interface {
	Recv() (*api.CompleteEvent, error)
	grpc.ClientStream
}

type clientServiceOnTransferCompleteClient struct {
	grpc.ClientStream
}

func (x *clientServiceOnTransferCompleteClient) Recv() (*api.CompleteEvent, error) {
	m := new(api.CompleteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	// Node Methods
	// Signing Method Request for Data
	Supply(context.Context, *api.SupplyRequest) (*api.SupplyResponse, error)
	// Verification Method Request for Signed Data
	Edit(context.Context, *api.EditRequest) (*api.EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(context.Context, *api.FetchRequest) (*api.FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(context.Context, *api.ShareRequest) (*api.ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(context.Context, *api.RespondRequest) (*api.RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(context.Context, *api.SearchRequest) (*api.SearchResponse, error)
	// Events Streams
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(*Empty, ClientService_OnLobbyRefreshServer) error
	// Returns a stream of Mailbox Message Events
	OnMailboxMessage(*Empty, ClientService_OnMailboxMessageServer) error
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransferAccepted(*Empty, ClientService_OnTransferAcceptedServer) error
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransferDeclined(*Empty, ClientService_OnTransferDeclinedServer) error
	// Returns a stream of DecisionEvent's for Invites
	OnTransferInvite(*Empty, ClientService_OnTransferInviteServer) error
	// Returns a stream of ProgressEvent's for Sessions
	OnTransferProgress(*Empty, ClientService_OnTransferProgressServer) error
	// Returns a stream of Completed Transfers
	OnTransferComplete(*Empty, ClientService_OnTransferCompleteServer) error
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) Supply(context.Context, *api.SupplyRequest) (*api.SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedClientServiceServer) Edit(context.Context, *api.EditRequest) (*api.EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedClientServiceServer) Fetch(context.Context, *api.FetchRequest) (*api.FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedClientServiceServer) Share(context.Context, *api.ShareRequest) (*api.ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedClientServiceServer) Respond(context.Context, *api.RespondRequest) (*api.RespondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Respond not implemented")
}
func (UnimplementedClientServiceServer) Search(context.Context, *api.SearchRequest) (*api.SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedClientServiceServer) OnLobbyRefresh(*Empty, ClientService_OnLobbyRefreshServer) error {
	return status.Errorf(codes.Unimplemented, "method OnLobbyRefresh not implemented")
}
func (UnimplementedClientServiceServer) OnMailboxMessage(*Empty, ClientService_OnMailboxMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method OnMailboxMessage not implemented")
}
func (UnimplementedClientServiceServer) OnTransferAccepted(*Empty, ClientService_OnTransferAcceptedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferAccepted not implemented")
}
func (UnimplementedClientServiceServer) OnTransferDeclined(*Empty, ClientService_OnTransferDeclinedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferDeclined not implemented")
}
func (UnimplementedClientServiceServer) OnTransferInvite(*Empty, ClientService_OnTransferInviteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferInvite not implemented")
}
func (UnimplementedClientServiceServer) OnTransferProgress(*Empty, ClientService_OnTransferProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferProgress not implemented")
}
func (UnimplementedClientServiceServer) OnTransferComplete(*Empty, ClientService_OnTransferCompleteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferComplete not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Supply(ctx, req.(*api.SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Edit(ctx, req.(*api.EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Fetch(ctx, req.(*api.FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Share(ctx, req.(*api.ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Respond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.RespondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Respond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Respond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Respond(ctx, req.(*api.RespondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.ClientService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Search(ctx, req.(*api.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_OnLobbyRefresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnLobbyRefresh(m, &clientServiceOnLobbyRefreshServer{stream})
}

type ClientService_OnLobbyRefreshServer interface {
	Send(*api.RefreshEvent) error
	grpc.ServerStream
}

type clientServiceOnLobbyRefreshServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnLobbyRefreshServer) Send(m *api.RefreshEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnMailboxMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnMailboxMessage(m, &clientServiceOnMailboxMessageServer{stream})
}

type ClientService_OnMailboxMessageServer interface {
	Send(*api.MailboxEvent) error
	grpc.ServerStream
}

type clientServiceOnMailboxMessageServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnMailboxMessageServer) Send(m *api.MailboxEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferAccepted_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferAccepted(m, &clientServiceOnTransferAcceptedServer{stream})
}

type ClientService_OnTransferAcceptedServer interface {
	Send(*api.DecisionEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferAcceptedServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferAcceptedServer) Send(m *api.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferDeclined_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferDeclined(m, &clientServiceOnTransferDeclinedServer{stream})
}

type ClientService_OnTransferDeclinedServer interface {
	Send(*api.DecisionEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferDeclinedServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferDeclinedServer) Send(m *api.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferInvite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferInvite(m, &clientServiceOnTransferInviteServer{stream})
}

type ClientService_OnTransferInviteServer interface {
	Send(*api.InviteEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferInviteServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferInviteServer) Send(m *api.InviteEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferProgress(m, &clientServiceOnTransferProgressServer{stream})
}

type ClientService_OnTransferProgressServer interface {
	Send(*api.ProgressEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferProgressServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferProgressServer) Send(m *api.ProgressEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_OnTransferComplete_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).OnTransferComplete(m, &clientServiceOnTransferCompleteServer{stream})
}

type ClientService_OnTransferCompleteServer interface {
	Send(*api.CompleteEvent) error
	grpc.ServerStream
}

type clientServiceOnTransferCompleteServer struct {
	grpc.ServerStream
}

func (x *clientServiceOnTransferCompleteServer) Send(m *api.CompleteEvent) error {
	return x.ServerStream.SendMsg(m)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sonr.node.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Supply",
			Handler:    _ClientService_Supply_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _ClientService_Edit_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _ClientService_Fetch_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _ClientService_Share_Handler,
		},
		{
			MethodName: "Respond",
			Handler:    _ClientService_Respond_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ClientService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnLobbyRefresh",
			Handler:       _ClientService_OnLobbyRefresh_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnMailboxMessage",
			Handler:       _ClientService_OnMailboxMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferAccepted",
			Handler:       _ClientService_OnTransferAccepted_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferDeclined",
			Handler:       _ClientService_OnTransferDeclined_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferInvite",
			Handler:       _ClientService_OnTransferInvite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferProgress",
			Handler:       _ClientService_OnTransferProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferComplete",
			Handler:       _ClientService_OnTransferComplete_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/node/client.proto",
}