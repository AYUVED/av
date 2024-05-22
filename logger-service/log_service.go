package main

import (
	"context"
	"log"

	"github.com/ayuved/av/pb/logs"
	pb "github.com/ayuved/av/pb/logs"

	"loggerService/data"
)

type logServiceImpl struct {
	pb.UnimplementedLogService
	Models data.Models
}

// Log Log says Log.
func (s *logServiceImpl) Log(
	ctx context.Context,
	req *pb.LogRequest,
) (*pb.LogResponse, error) {
	log.Printf("Received Log request: %v", req)
	logEntry := data.LogEntry{
		Name: req.Name,
		Data: req.Data,
	}
	log.Printf("Log entry: %v", logEntry)

	err := s.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Msg: "Error inserting into logs: " + err.Error()}
		return res, err
	}

	// if err != nil {
	// 	log.Fatalf("Error inserting into logs: %v", err)
	// }

	result := "Log: " + req.Name + " " + req.Data
	rsp := &pb.LogResponse{Msg: req.App + req.Name + req.Data + result}
	return rsp, nil
}
