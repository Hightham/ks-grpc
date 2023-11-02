package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "ks-grpc/api/proto/api_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedChargingPointsServiceServer
}

func main() {

	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("error with listener")
	}
	s := grpc.NewServer()
	pb.RegisterChargingPointsServiceServer(s, &server{})

	reflection.Register(s)

	fmt.Println("Server is listening on port :4040")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("error starting server")
	}

	fmt.Print("Hello World")
}

func (s *server) GetChargingStation(ctx context.Context, req *pb.LocationRequest) (*pb.LocationResponse, error) {
	response := &pb.LocationResponse{
		Lat:       req.Lat,
		Long:      req.Long,
		Name:      "Charging Point",
		Uuid:      "xxxxx",
		Available: true,
	}
	return response, nil
}

func (s *server) StreamChargingStations(req *pb.LocationRequest, stream pb.ChargingPointsService_StreamChargingStationsServer) error {
	//querying logic here that get available charging stations within a certain region
	for {
		response := &pb.LocationResponse{
			Lat:       req.Lat,
			Long:      req.Long,
			Name:      "Charging Point",
			Uuid:      "xxxxx",
			Available: true,
		}
		err := stream.Send(response)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}
func (s *server) StreamLocation(stream pb.ChargingPointsService_StreamLocationServer) error {
	var responses []*pb.LocationResponse

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// The client has closed the stream, so you can send the list of responses here.
			locationResponseList := &pb.LocationResponseList{LocationResponse: responses}
			if err := stream.SendAndClose(locationResponseList); err != nil {
				log.Printf("Error sending response: %v", err)
			}
			return nil
		}
		if err != nil {
			log.Printf("Error receiving request: %v", err)
			return err
		}

		// Create a LocationResponse and add it to the list of responses
		response := &pb.LocationResponse{
			Lat:       request.Lat,
			Long:      request.Long,
			Name:      "ServerResponse",
			Uuid:      "xxx",
			Available: true,
		}
		responses = append(responses, response)
	}
}

func (s *server) TrackLocations(stream pb.ChargingPointsService_TrackLocationsServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		response := &pb.LocationResponse{
			Lat:       req.Lat,
			Long:      req.Long,
			Name:      "Charging Point",
			Uuid:      "xxxxx",
			Available: true,
		}

		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
}
