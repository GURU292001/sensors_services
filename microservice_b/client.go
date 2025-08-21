package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"time"

	db "sensors/DB"
	pb "sensors/sensorpb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSensorServiceServer
	db *sql.DB
}

func (s *server) SendSensorData(ctx context.Context, data *pb.SensorData) (*pb.Ack, error) {
	log.Printf("Received data: %+v", data)

	query := "INSERT INTO readings_table (sensor_value, sensor_type, id1,time_stamp) VALUES (?, ?, ?,now())"
	_, err := s.db.Exec(query, data.Value, data.Type, data.Id1)
	if err != nil {
		return &pb.Ack{Status: "Failed"}, err
	}

	return &pb.Ack{Status: "OK"}, nil
}

func main() {
	// log file creation
	if err := os.MkdirAll("./log", 0755); err != nil {
		log.Fatalf("error creating log directory: %v", err)
	}
	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lFile.Close()
	log.SetOutput(lFile)

	// Global Database Connection
	db.Db_connection()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSensorServiceServer(grpcServer, &server{db: db.Gdb})

	log.Println("Microservice B (gRPC Server) running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
