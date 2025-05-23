
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "shared/messages"
)

type server struct {
    pb.UnimplementedCouponServiceServer
}

func main() {
    conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewCouponServiceClient(conn)
    res, err := client.ValidateCoupon(context.Background(), &pb.CouponValidationRequest{
        CouponCode: "ABC123",
    })

    if err != nil {
        log.Fatalf("Error when calling ValidateCoupon: %v", err)
    }

    log.Printf("Response from server: %v", res.Message)
}
