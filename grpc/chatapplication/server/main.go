package main

import (
	pb "chatservice/protos"
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type User struct {
	Id   int
	Name string
}

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	mu           sync.Mutex
	users        map[int]User
	rooms        map[int]map[int]chan *pb.ChatResponse
	privateRooms map[int]chan *pb.ChatResponse
}

func NewChatServiceServer() *ChatServiceServer {
	return &ChatServiceServer{
		users:        make(map[int]User),
		rooms:        make(map[int]map[int]chan *pb.ChatResponse),
		privateRooms: make(map[int]chan *pb.ChatResponse),
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, NewChatServiceServer())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (cs *ChatServiceServer) SendPrivateMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	senderId := int(req.SenderId)
	receiverId := int(req.ReceiverId)
	text := req.Text

	cs.mu.Lock()
	defer cs.mu.Unlock()

	receiverCh, receiverExists := cs.privateRooms[receiverId]
	if !receiverExists {
		receiverCh = make(chan *pb.ChatResponse, 10)
		cs.privateRooms[receiverId] = receiverCh
	}
	receiverCh <- &pb.ChatResponse{SenderId: int32(senderId), Text: text}
	return &pb.MessageResponse{Status: "Success"}, fmt.Errorf("Message sent successfully to %d", receiverId)
}

func (cs *ChatServiceServer) ChatRoom(stream pb.ChatService_ChatRoomServer) error {
	msgChannel := make(chan *pb.ChatResponse, 10)
	privateMsgChannel := make(chan *pb.ChatResponse, 10)

	var wg sync.WaitGroup
	wg.Add(1)

	var roomId, userId int

	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-msgChannel:
				if !ok {
					return
				}
				if err := stream.Send(msg); err != nil {
					log.Printf("Send error to user %d (room): %v", userId, err)
					return
				}
			case msg, ok := <-privateMsgChannel:
				if !ok {
					return
				}
				if err := stream.Send(msg); err != nil {
					log.Printf("Send error to user %d (private): %v", userId, err)
					return
				}
			}
		}
	}()

	for {
		req, err := stream.Recv()
		if err != nil {
			log.Printf("Client disconnected from room %d", roomId)
			break
		}

		if roomId == 0 && userId == 0 {
			roomId = int(req.RoomId)
			userId = int(req.SenderId)
			cs.mu.Lock()
			if cs.rooms[roomId] == nil {
				cs.rooms[roomId] = make(map[int]chan *pb.ChatResponse)
			}
			cs.rooms[roomId][userId] = msgChannel

			privateMsgChannel = make(chan *pb.ChatResponse, 10)
			cs.privateRooms[userId] = privateMsgChannel
			cs.mu.Unlock()
		}

		cs.mu.Lock()
		for uid, ch := range cs.rooms[roomId] {
			if uid != userId {
				select {
				case ch <- &pb.ChatResponse{SenderId: req.SenderId, Text: req.Text}:
				default:
					log.Printf("Message dropped for user %d (channel full)", uid)
				}
			}
		}
		cs.mu.Unlock()
	}

	cs.mu.Lock()
	if users, ok := cs.rooms[roomId]; ok {
		delete(users, userId)
		if len(users) == 0 {
			delete(cs.rooms, roomId)
		}
	}
	cs.mu.Unlock()
	close(msgChannel)

	cs.mu.Lock()
	delete(cs.privateRooms, userId)
	cs.mu.Unlock()

	close(msgChannel)
	close(privateMsgChannel)
	wg.Wait()
	return nil
}

func (cs *ChatServiceServer) LeaveRoom(ctx context.Context, req *pb.LeaveRoomRequest) (*pb.LeaveRoomResponse, error) {
	userId := int(req.UserId)
	roomId := int(req.RoomId)

	cs.mu.Lock()
	defer cs.mu.Unlock()

	if users, ok := cs.rooms[roomId]; ok {
		if _, exists := users[userId]; exists {
			delete(users, userId)
			if len(users) == 0 {
				delete(cs.rooms, roomId)
			}
			for uid, ch := range cs.rooms[roomId] {
				if uid != userId {
					ch <- &pb.ChatResponse{SenderId: int32(userId), Text: "[left the room]"}
				}
			}
			return &pb.LeaveRoomResponse{Status: "Success"}, nil
		}
	}
	return &pb.LeaveRoomResponse{Status: "Error"}, fmt.Errorf("User %d not found in room %d", userId, roomId)
}
