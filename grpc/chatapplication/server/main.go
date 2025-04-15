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

type Room struct {
	Id          int
	Name        string
	Description string
	Users       []User
}

type MessageType int

const (
	Private MessageType = iota
	Public
)

type Message struct {
	Id       int
	Sender   User
	Receiver User
	Text     string
	Type     MessageType
}

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	mu            sync.Mutex
	users         []User
	rooms         []Room
	messages      []Message
	lastUserId    int
	lastRoomId    int
	lastMessageId int

	roomStreams map[int]map[int]chan *pb.ChatResponse
}

func NewChatServiceServer() *ChatServiceServer {
	return &ChatServiceServer{
		roomStreams: make(map[int]map[int]chan *pb.ChatResponse),
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

func (cs *ChatServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.lastUserId++
	user := User{
		Id:   int(cs.lastUserId),
		Name: req.Name,
	}
	cs.users = append(cs.users, user)
	fmt.Println(cs.lastUserId)
	return &pb.RegisterResponse{UserId: int32(cs.lastUserId)}, nil
}

func (cs *ChatServiceServer) getUser(userId int) (User, bool) {
	for _, user := range cs.users {
		if user.Id == userId {
			return user, true
		}
	}
	return User{}, false
}
func (cs *ChatServiceServer) SendMessage(cts context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	if req.SenderId == req.ReceiverId {
		return &pb.MessageResponse{}, fmt.Errorf("Sender and receiver cannot be the same")
	}
	sender, senderExists := cs.getUser(int(req.SenderId))
	receiver, receiverExists := cs.getUser(int(req.ReceiverId))
	if !senderExists || !receiverExists {
		return &pb.MessageResponse{}, fmt.Errorf("Invalid user details")
	}
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.lastMessageId++
	message := Message{
		Id:       cs.lastMessageId,
		Sender:   sender,
		Receiver: receiver,
		Text:     req.Text,
		Type:     Private,
	}
	cs.messages = append(cs.messages, message)
	return &pb.MessageResponse{Message: "Message sent"}, nil
}

func (cs *ChatServiceServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.lastRoomId++
	room := Room{
		Id:          cs.lastRoomId,
		Name:        req.Name,
		Description: req.Description,
		Users:       []User{},
	}
	cs.rooms = append(cs.rooms, room)
	return &pb.CreateRoomResponse{RoomId: int32(cs.lastRoomId)}, nil
}

func (cs *ChatServiceServer) getRoom(roomId int) (Room, bool) {
	for _, room := range cs.rooms {
		if room.Id == roomId {
			return room, true
		}
	}
	return Room{}, false
}
func (cs *ChatServiceServer) JoinRoom(req *pb.JoinRoomRequest, stream pb.ChatService_JoinRoomServer) error {
	cs.mu.Lock()

	room, roomExists := cs.getRoom(int(req.RoomId))
	user, userExists := cs.getUser(int(req.UserId))
	if !roomExists || !userExists {
		cs.mu.Unlock()
		return fmt.Errorf("invalid room or user details")
	}

	room.Users = append(room.Users, user)

	if cs.roomStreams[room.Id] == nil {
		cs.roomStreams[room.Id] = make(map[int]chan *pb.ChatResponse)
	}
	msgCh := make(chan *pb.ChatResponse, 100)
	cs.roomStreams[room.Id][user.Id] = msgCh

	cs.mu.Unlock()

	for msg := range msgCh {
		if err := stream.Send(msg); err != nil {
			log.Printf("error sending to user %d: %v", user.Id, err)
			break
		}
	}

	cs.mu.Lock()
	delete(cs.roomStreams[room.Id], user.Id)
	cs.mu.Unlock()

	return nil
}

func (cs *ChatServiceServer) Chat(stream pb.ChatService_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Println("stream receive error:", err)
			return err
		}

		cs.mu.Lock()

		sender, senderExists := cs.getUser(int(req.SenderId))
		room, roomExists := cs.getRoom(int(req.RoomId))
		if !senderExists || !roomExists {
			cs.mu.Unlock()
			continue
		}

		cs.lastMessageId++
		msg := Message{
			Id:     cs.lastMessageId,
			Sender: sender,
			Text:   req.Text,
			Type:   Public,
		}
		cs.messages = append(cs.messages, msg)

		for userId, ch := range cs.roomStreams[room.Id] {
			if userId != sender.Id {
				select {
				case ch <- &pb.ChatResponse{
					Text:     req.Text,
					SenderId: int32(sender.Id),
				}:
				default:
					log.Printf("dropping message for user %d (channel full)", userId)
				}
			}
		}

		cs.mu.Unlock()
	}
}
