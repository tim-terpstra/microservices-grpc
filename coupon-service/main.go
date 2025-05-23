
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
    usedCoupons map[string]bool
}

func (s *server) ValidateCoupon(ctx context.Context, req *pb.CouponValidationRequest) (*pb.CouponValidationResponse, error) {
    if s.usedCoupons[req.CouponCode] {
        return &pb.CouponValidationResponse{
            Valid:   false,
            Message: "Coupon already used",
        }, nil
    }
    s.usedCoupons[req.CouponCode] = true
    return &pb.CouponValidationResponse{
        Valid:   true,
        Message: "Coupon applied",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterCouponServiceServer(grpcServer, &server{usedCoupons: make(map[string]bool)})
    log.Println("Coupon service running on :50052")
    grpcServer.Serve(lis)
}
