package main_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	proto "github.com/bygui86/go-grpc-testing/domain"
	server "github.com/bygui86/go-grpc-testing/server"
)

func TestSayHello(t *testing.T) {
	grpcServer := server.GrpcServer{}

	// set up test cases
	tests := []struct {
		name string
		want string
	}{
		{
			name: "world",
			want: "Hello world",
		},
		{
			name: "123",
			want: "Hello 123",
		},
	}

	for _, tt := range tests {
		req := &proto.HelloRequest{Name: tt.name}
		resp, err := grpcServer.SayHello(context.Background(), req)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, tt.want, resp.Message)
		//
		// if err != nil {
		// 	t.Errorf("HelloTest(%s) got unexpected error", tt.name)
		// }
		// if resp == nil {
		// 	t.Errorf("HelloTest(%s) got unexpected empty response", tt.name)
		// }
		// if resp.Message != tt.want {
		// 	t.Errorf("HelloText(%s)=%v, wanted %s", tt.name, resp.Message, tt.want)
		// }
	}
}
