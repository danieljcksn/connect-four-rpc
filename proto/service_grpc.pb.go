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
	GameService_Connect_FullMethodName          = "/game.GameService/Connect"
	GameService_NotifyPlayerTurn_FullMethodName = "/game.GameService/NotifyPlayerTurn"
	GameService_ShowGameState_FullMethodName    = "/game.GameService/ShowGameState"
	GameService_MakeMove_FullMethodName         = "/game.GameService/MakeMove"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
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

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, GameService_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) NotifyPlayerTurn(ctx context.Context, in *NotifyPlayerTurnRequest, opts ...grpc.CallOption) (*NotifyPlayerTurnResponse, error) {
	out := new(NotifyPlayerTurnResponse)
	err := c.cc.Invoke(ctx, GameService_NotifyPlayerTurn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ShowGameState(ctx context.Context, in *ShowGameStateRequest, opts ...grpc.CallOption) (*ShowGameStateResponse, error) {
	out := new(ShowGameStateResponse)
	err := c.cc.Invoke(ctx, GameService_ShowGameState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) MakeMove(ctx context.Context, in *MakeMoveRequest, opts ...grpc.CallOption) (*MakeMoveResponse, error) {
	out := new(MakeMoveResponse)
	err := c.cc.Invoke(ctx, GameService_MakeMove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
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
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedGameServiceServer) NotifyPlayerTurn(context.Context, *NotifyPlayerTurnRequest) (*NotifyPlayerTurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyPlayerTurn not implemented")
}
func (UnimplementedGameServiceServer) ShowGameState(context.Context, *ShowGameStateRequest) (*ShowGameStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowGameState not implemented")
}
func (UnimplementedGameServiceServer) MakeMove(context.Context, *MakeMoveRequest) (*MakeMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeMove not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_NotifyPlayerTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyPlayerTurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).NotifyPlayerTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_NotifyPlayerTurn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).NotifyPlayerTurn(ctx, req.(*NotifyPlayerTurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ShowGameState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowGameStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ShowGameState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_ShowGameState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ShowGameState(ctx, req.(*ShowGameStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_MakeMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).MakeMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_MakeMove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).MakeMove(ctx, req.(*MakeMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _GameService_Connect_Handler,
		},
		{
			MethodName: "NotifyPlayerTurn",
			Handler:    _GameService_NotifyPlayerTurn_Handler,
		},
		{
			MethodName: "ShowGameState",
			Handler:    _GameService_ShowGameState_Handler,
		},
		{
			MethodName: "MakeMove",
			Handler:    _GameService_MakeMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
