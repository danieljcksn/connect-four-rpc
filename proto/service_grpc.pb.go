// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: service.proto

package proto

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

const (
	Connect4Game_GameSession_FullMethodName      = "/connect4.Connect4Game/GameSession"
	Connect4Game_Connect_FullMethodName          = "/connect4.Connect4Game/Connect"
	Connect4Game_NotifyPlayerTurn_FullMethodName = "/connect4.Connect4Game/NotifyPlayerTurn"
	Connect4Game_ShowGameState_FullMethodName    = "/connect4.Connect4Game/ShowGameState"
	Connect4Game_MakeMove_FullMethodName         = "/connect4.Connect4Game/MakeMove"
)

// Connect4GameClient is the client API for Connect4Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Connect4GameClient interface {
	GameSession(ctx context.Context, opts ...grpc.CallOption) (Connect4Game_GameSessionClient, error)
	// A client-to-server streaming RPC.
	//
	// Once the server has started the stream, the client can
	// send the first message in order to connect and start the game.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// A serve-to-client streaming RPC.
	//
	// The game will alternate between the local player (located in the server)
	// and the remote player (located in the client).
	// So the server needs to warn the client when it's their turn to play.
	NotifyPlayerTurn(ctx context.Context, in *NotifyPlayerTurnRequest, opts ...grpc.CallOption) (*NotifyPlayerTurnResponse, error)
	// A server-to-client streaming RPC.
	//
	// The client can request the current state of the game at any time.
	ShowGameState(ctx context.Context, in *ShowGameStateRequest, opts ...grpc.CallOption) (*ShowGameStateResponse, error)
	// A client-to-server streaming RPC.
	//
	// The client can send a message to the server to make a move.
	MakeMove(ctx context.Context, in *MakeMoveRequest, opts ...grpc.CallOption) (*MakeMoveResponse, error)
}

type connect4GameClient struct {
	cc grpc.ClientConnInterface
}

func NewConnect4GameClient(cc grpc.ClientConnInterface) Connect4GameClient {
	return &connect4GameClient{cc}
}

func (c *connect4GameClient) GameSession(ctx context.Context, opts ...grpc.CallOption) (Connect4Game_GameSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Connect4Game_ServiceDesc.Streams[0], Connect4Game_GameSession_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &connect4GameGameSessionClient{stream}
	return x, nil
}

type Connect4Game_GameSessionClient interface {
	Send(*GameCommand) error
	Recv() (*GameUpdate, error)
	grpc.ClientStream
}

type connect4GameGameSessionClient struct {
	grpc.ClientStream
}

func (x *connect4GameGameSessionClient) Send(m *GameCommand) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connect4GameGameSessionClient) Recv() (*GameUpdate, error) {
	m := new(GameUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connect4GameClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, Connect4Game_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connect4GameClient) NotifyPlayerTurn(ctx context.Context, in *NotifyPlayerTurnRequest, opts ...grpc.CallOption) (*NotifyPlayerTurnResponse, error) {
	out := new(NotifyPlayerTurnResponse)
	err := c.cc.Invoke(ctx, Connect4Game_NotifyPlayerTurn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connect4GameClient) ShowGameState(ctx context.Context, in *ShowGameStateRequest, opts ...grpc.CallOption) (*ShowGameStateResponse, error) {
	out := new(ShowGameStateResponse)
	err := c.cc.Invoke(ctx, Connect4Game_ShowGameState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connect4GameClient) MakeMove(ctx context.Context, in *MakeMoveRequest, opts ...grpc.CallOption) (*MakeMoveResponse, error) {
	out := new(MakeMoveResponse)
	err := c.cc.Invoke(ctx, Connect4Game_MakeMove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Connect4GameServer is the server API for Connect4Game service.
// All implementations must embed UnimplementedConnect4GameServer
// for forward compatibility
type Connect4GameServer interface {
	GameSession(Connect4Game_GameSessionServer) error
	// A client-to-server streaming RPC.
	//
	// Once the server has started the stream, the client can
	// send the first message in order to connect and start the game.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// A serve-to-client streaming RPC.
	//
	// The game will alternate between the local player (located in the server)
	// and the remote player (located in the client).
	// So the server needs to warn the client when it's their turn to play.
	NotifyPlayerTurn(context.Context, *NotifyPlayerTurnRequest) (*NotifyPlayerTurnResponse, error)
	// A server-to-client streaming RPC.
	//
	// The client can request the current state of the game at any time.
	ShowGameState(context.Context, *ShowGameStateRequest) (*ShowGameStateResponse, error)
	// A client-to-server streaming RPC.
	//
	// The client can send a message to the server to make a move.
	MakeMove(context.Context, *MakeMoveRequest) (*MakeMoveResponse, error)
	mustEmbedUnimplementedConnect4GameServer()
}

// UnimplementedConnect4GameServer must be embedded to have forward compatible implementations.
type UnimplementedConnect4GameServer struct {
}

func (UnimplementedConnect4GameServer) GameSession(Connect4Game_GameSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method GameSession not implemented")
}
func (UnimplementedConnect4GameServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedConnect4GameServer) NotifyPlayerTurn(context.Context, *NotifyPlayerTurnRequest) (*NotifyPlayerTurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyPlayerTurn not implemented")
}
func (UnimplementedConnect4GameServer) ShowGameState(context.Context, *ShowGameStateRequest) (*ShowGameStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowGameState not implemented")
}
func (UnimplementedConnect4GameServer) MakeMove(context.Context, *MakeMoveRequest) (*MakeMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMove not implemented")
}
func (UnimplementedConnect4GameServer) mustEmbedUnimplementedConnect4GameServer() {}

// UnsafeConnect4GameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Connect4GameServer will
// result in compilation errors.
type UnsafeConnect4GameServer interface {
	mustEmbedUnimplementedConnect4GameServer()
}

func RegisterConnect4GameServer(s grpc.ServiceRegistrar, srv Connect4GameServer) {
	s.RegisterService(&Connect4Game_ServiceDesc, srv)
}

func _Connect4Game_GameSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Connect4GameServer).GameSession(&connect4GameGameSessionServer{stream})
}

type Connect4Game_GameSessionServer interface {
	Send(*GameUpdate) error
	Recv() (*GameCommand, error)
	grpc.ServerStream
}

type connect4GameGameSessionServer struct {
	grpc.ServerStream
}

func (x *connect4GameGameSessionServer) Send(m *GameUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connect4GameGameSessionServer) Recv() (*GameCommand, error) {
	m := new(GameCommand)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Connect4Game_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Connect4GameServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connect4Game_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Connect4GameServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect4Game_NotifyPlayerTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyPlayerTurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Connect4GameServer).NotifyPlayerTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connect4Game_NotifyPlayerTurn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Connect4GameServer).NotifyPlayerTurn(ctx, req.(*NotifyPlayerTurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect4Game_ShowGameState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowGameStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Connect4GameServer).ShowGameState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connect4Game_ShowGameState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Connect4GameServer).ShowGameState(ctx, req.(*ShowGameStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect4Game_MakeMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Connect4GameServer).MakeMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connect4Game_MakeMove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Connect4GameServer).MakeMove(ctx, req.(*MakeMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Connect4Game_ServiceDesc is the grpc.ServiceDesc for Connect4Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connect4Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connect4.Connect4Game",
	HandlerType: (*Connect4GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Connect4Game_Connect_Handler,
		},
		{
			MethodName: "NotifyPlayerTurn",
			Handler:    _Connect4Game_NotifyPlayerTurn_Handler,
		},
		{
			MethodName: "ShowGameState",
			Handler:    _Connect4Game_ShowGameState_Handler,
		},
		{
			MethodName: "MakeMove",
			Handler:    _Connect4Game_MakeMove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GameSession",
			Handler:       _Connect4Game_GameSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
