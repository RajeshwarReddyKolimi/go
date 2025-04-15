package main

import (
	"bufio"
	"context"
	pb "fileservice/protos"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
)

type FileServer struct {
	pb.UnimplementedFileServiceServer
}

func main() {
	lis, er := net.Listen("tcp", ":50051")
	if er != nil {
		fmt.Println(er)
		return
	}

	server := grpc.NewServer()
	pb.RegisterFileServiceServer(server, &FileServer{})

	if er := server.Serve(lis); er != nil {
		fmt.Println(er)
		return
	}
}

func (s *FileServer) UploadFile(stream pb.FileService_UploadFileServer) error {
	var file *os.File
	var writer *bufio.Writer

	log.Println("Upload started")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			writer.Flush()
			file.Close()
			log.Println("Upload complete")
			return stream.SendAndClose(&pb.UploadFileResponse{Message: "Upload complete"})
		}
		if err != nil {
			return err
		}

		if file == nil {
			path := filepath.Join("uploads", req.Filename)
			file, err = os.Create(path)
			if err != nil {
				return err
			}
			log.Println("File created:", path)
			writer = bufio.NewWriter(file)
			continue
		}

		_, err = writer.Write(req.Data)
		if err != nil {
			return err
		}
		log.Println("Length: ", len(req.Data))
	}
}

func (s *FileServer) DownloadFile(req *pb.DownloadFileRequest, stream pb.FileService_DownloadFileServer) error {
	log.Println("Download started")
	path := filepath.Join("uploads", req.Filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("Reading file")
		stream.Send(&pb.DownloadFileResponse{Data: buffer[:n]})
	}

	log.Println("Download complete")
	return nil
}

func (s *FileServer) GetMetaData(ctx context.Context, req *pb.MetaDataRequest) (*pb.MetaDataResponse, error) {
	log.Println("Fetching metadata")
	path := filepath.Join("uploads", req.Filename)
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	mime := mime.TypeByExtension(filepath.Ext(req.Filename))
	log.Println("Fetching metadata complete")
	return &pb.MetaDataResponse{
		Filename:  req.Filename,
		Size:      stat.Size(),
		MimeType:  mime,
		CreatedAt: stat.ModTime().Format(time.RFC3339),
	}, nil
}
