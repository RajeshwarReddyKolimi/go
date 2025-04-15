package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	pb "fileservice/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const chunkSize = 1024

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	err = uploadFile(client, "example.txt")
	if err != nil {
		fmt.Println("Upload failed:", err)
	}

	err = downloadFile(client, "example.txt", "downloads")
	if err != nil {
		fmt.Println("Download failed:", err)
	}

	err = getMetadata(client, "example.txt")
	if err != nil {
		fmt.Println("Metadata failed:", err)
	}
}

func uploadFile(client pb.FileServiceClient, filename string) error {
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		fmt.Println("1OK")

		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("2OK")

		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, chunkSize)

	err = stream.Send(&pb.UploadFileRequest{Filename: filepath.Base(filename)})
	if err != nil {
		fmt.Println("3OK")

		return err
	}

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("4OK")
			break
		}
		if err != nil {
			fmt.Println("5OK")

			return err
		}

		err = stream.Send(&pb.UploadFileRequest{Data: buffer[:n]})
		if err != nil {
			fmt.Println("6OK")

			return err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Println("Upload response:", res.Message)
	return nil
}

func downloadFile(client pb.FileServiceClient, filename, outputDir string) error {
	req := &pb.DownloadFileRequest{Filename: filename}
	stream, err := client.DownloadFile(context.Background(), req)
	if err != nil {
		return err
	}

	outPath := filepath.Join(outputDir, filename)
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		_, err = writer.Write(res.Data)
		if err != nil {
			return err
		}
	}

	fmt.Println("Download complete:", outPath)
	return nil
}

func getMetadata(client pb.FileServiceClient, filename string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.MetaDataRequest{Filename: filename}
	res, err := client.GetMetaData(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("Metadata for %s:\n", res.Filename)
	fmt.Printf("- Size: %d bytes\n", res.Size)
	fmt.Printf("- Type: %s\n", res.MimeType)
	fmt.Printf("- Created: %s\n", res.CreatedAt)
	return nil
}
