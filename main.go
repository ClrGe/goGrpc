package main

import (
	"context"
	pb "dataAnalyzerFile/frequentationPB"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

type Station struct {
	UicCode string `csv:"code_uic"`
	ZipCode string `csv:"cp"`
	A2015   int32  `csv:"a2015"`
	A2016   int32  `csv:"a2016"`
	A2017   int32  `csv:"a2017"`
	A2018   int32  `csv:"a2018"`
	A2019   int32  `csv:"a2019"`
	A2020   int32  `csv:"a2020"`
	A2021   int32  `csv:"a2021"`
}

type FrequentationServer struct {
	pb.UnimplementedFrequentationServer
}

//func (s *FrequentationServer) ReadStations(context.Context, *pb.FrequentationRequest) (*[]*pb.FrequentationResponse, error) {
//
//	in, err := os.Open("data/frequentation-gares.csv")
//	if err != nil {
//		panic(err)
//	}
//	defer in.Close()
//
//	stations := []*Station{}
//	if err := gocsv.UnmarshalFile(in, &stations); err != nil {
//		panic(err)
//	}
//	for _, station := range stations {
//		frequentation := pb.FrequentationList.GetValue(pb.FrequentationResponse{
//			UicCode: station.UicCode,
//			Zipcode: station.ZipCode,
//			A2015:   station.A2015,
//			A2016:   station.A2016,
//			A2017:   station.A2017,
//			A2018:   station.A2018,
//			A2019:   station.A2019,
//			A2020:   station.A2020,
//			A2021:   station.A2021})
//		return &frequentation, nil
//	}
//	return 200, nil
//}

func (srv *FrequentationServer) ReadStations(context.Context, *pb.FrequentationRequest) (*pb.FrequentationResponse, error) {
	bufferSize := 64 * 1024 //64KiB, tweak this as desired
	file, _ := os.Open("./data/output.json")
	resp := &pb.FrequentationResponse{}
	defer file.Close()
	buff := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		resp := &pb.FrequentationResponse{
			FileChunk: buff[:bytesRead],
		}
		return resp, nil
	}
	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFrequentationServer(s, &FrequentationServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
