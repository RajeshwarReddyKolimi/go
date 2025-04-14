package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	pb "routeguide/protos"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error creating client: ", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getFeature(ctx, client, &pb.Point{Latitude: 200000000, Longitude: -300000000})
	listFeatures(ctx, client, &pb.Rectangle{Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000}, Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000}})
	runRecordRoute(ctx, client)
	runRouteChat(ctx, client)
}

func getFeature(ctx context.Context, client pb.RouteGuideServiceClient, point *pb.Point) {
	r1, e1 := client.GetFeature(ctx, point)
	fmt.Println(r1, e1)
}

func listFeatures(ctx context.Context, client pb.RouteGuideServiceClient, rect *pb.Rectangle) {
	stream, e := client.ListFeatures(ctx, &pb.Rectangle{Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000}, Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000}})
	if e != nil {
		fmt.Println(e)
	}
	for {
		feature, er := stream.Recv()
		if er == io.EOF {
			break
		}
		if er != nil {
			fmt.Println(er)
		}
		log.Println(feature)
	}
}

func runRecordRoute(ctx context.Context, client pb.RouteGuideServiceClient) {
	pointCount := int(rand.Int32N(100)) + 2
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint())
	}
	log.Printf("Traversing %d points.", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	log.Printf("Route summary: %v", reply)
}

func runRouteChat(ctx context.Context, client pb.RouteGuideServiceClient) {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("client.RouteChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.RouteChat failed: %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("client.RouteChat: stream.Send(%v) failed: %v", note, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func randomPoint() *pb.Point {
	lat := (rand.Int32N(180) - 90) * 1e7
	long := (rand.Int32N(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}
