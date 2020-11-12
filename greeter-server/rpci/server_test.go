package rpci_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	pb "github.com/voprak/grpc-example/greeter-server/pb"
	"github.com/voprak/grpc-example/greeter-server/rpci"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterGreeterServer(server, rpci.NewServer())

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

var _ = Describe("Server", func() {

	Describe("Say Hello ", func() {
		Context("With value", func() {
			It("Should greet with the name provided", func() {
				ctx := context.Background()
				conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
				if err != nil {
					log.Fatal(err)
				}
				defer conn.Close()

				client := pb.NewGreeterClient(conn)

				request := &pb.HelloRequest{Name: "Go developers"}
				response, err := client.SayHello(ctx, request)
				m := response.GetMessage()
				fmt.Println(m)
				Expect(m).To(Equal("Hello Go developers"))

			})
		})
	})

})