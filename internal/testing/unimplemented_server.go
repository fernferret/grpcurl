package testing

//go:generate protoc --go_out=. --go-grpc_out=. test.proto
//go:generate protoc --descriptor_set_out=./test.protoset test.proto
//go:generate protoc --descriptor_set_out=./example.protoset --include_imports example.proto

// UnimplementedServer implements the TestService interface defined in example.proto.
type UnimplementedServer struct {
	UnimplementedUnimplementedServiceServer
}

// This server contains unimplemented methods.
// var _ UnimplementedTestServiceServer = UnimplementedServer{}
