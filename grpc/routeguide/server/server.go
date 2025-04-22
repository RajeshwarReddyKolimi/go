package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	pb "routeguide/protos"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type RouteGuideServiceServer struct {
	pb.UnimplementedRouteGuideServiceServer
	savedFeatures []*pb.Feature
	mu            sync.Mutex
	routeNotes    map[string][]*pb.RouteNote
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterRouteGuideServiceServer(server, &RouteGuideServiceServer{})
	if err := server.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
	fmt.Println("Listening on port 50051...")
}

func (rgs *RouteGuideServiceServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	rgs.loadFeatures()
	rgs.routeNotes = make(map[string][]*pb.RouteNote)
	fmt.Println("Request for getFeature received")
	for _, feature := range rgs.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{Location: point}, nil
}

func (rgs *RouteGuideServiceServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuideService_ListFeaturesServer) error {
	fmt.Println("Request for listFeatures received")
	for _, feature := range rgs.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (rgs *RouteGuideServiceServer) RecordRoute(stream pb.RouteGuideService_RecordRouteServer) error {
	fmt.Println("Request for recordRoute received")
	var pointCount, featureCount, distance int32
	var lastPoint *pb.Point
	startTime := time.Now()

	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++
		for _, feature := range rgs.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		if lastPoint != nil {
			distance += calculateDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

func (rgs *RouteGuideServiceServer) RouteChat(stream pb.RouteGuideService_RouteChatServer) error {
	fmt.Println("Request for route chat")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)
		fmt.Println(key)
		rgs.mu.Lock()
		rgs.routeNotes[key] = append(rgs.routeNotes[key], in)

		rn := make([]*pb.RouteNote, len(rgs.routeNotes[key]))
		copy(rn, rgs.routeNotes[key])
		rgs.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

func serialize(point *pb.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}

func (s *RouteGuideServiceServer) loadFeatures() {
	if err := json.Unmarshal(exampleData, &s.savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

func calculateDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

var exampleData = []byte(`[{
    "location": {
        "latitude": 407838351,
        "longitude": -746143763
    },
    "name": "Patriots Path, Mendham, NJ 07945, USA"
}, {
    "location": {
        "latitude": 408122808,
        "longitude": -743999179
    },
    "name": "101 New Jersey 10, Whippany, NJ 07981, USA"
}, {
    "location": {
        "latitude": 413628156,
        "longitude": -749015468
    },
    "name": "U.S. 6, Shohola, PA 18458, USA"
}, {
    "location": {
        "latitude": 419999544,
        "longitude": -740371136
    },
    "name": "5 Conners Road, Kingston, NY 12401, USA"
}, {
    "location": {
        "latitude": 414008389,
        "longitude": -743951297
    },
    "name": "Mid Hudson Psychiatric Center, New Hampton, NY 10958, USA"
}, {
    "location": {
        "latitude": 419611318,
        "longitude": -746524769
    },
    "name": "287 Flugertown Road, Livingston Manor, NY 12758, USA"
}]`)
