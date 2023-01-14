package main

import (
	"context"
	pb "dataAnalyzerFile/frequentationPB"
	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type grpcServer struct {
	DataRepo DataRepo
	pb.UnimplementedFrequentationServer
}

type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type Data struct {
	UicCode string `csv:"code_uic"`
	Zipcode string `csv:"cp"`
	A2015   int32  `csv:"a2015"`
	A2016   int32  `csv:"a2016"`
	A2017   int32  `csv:"a2017"`
	A2018   int32  `csv:"a2018"`
	A2019   int32  `csv:"a2019"`
	A2020   int32  `csv:"a2020"`
	A2021   int32  `csv:"a2021"`
}

type DataRepo interface {
	DataList(ctx context.Context, offset int64, limit int64) ([]*Data, error)
}

func (s *grpcServer) DataList(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {

	stations := []*Data{}
	response := &pb.ListResponse{}
	rslt := []*pb.Data{}

	in, err := os.Open("data/frequentation-gares.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &stations); err != nil {
		panic(err)
	}

	for _, station := range stations {
		b := &pb.Data{
			UicCode: station.UicCode,
			Zipcode: station.Zipcode,
			A2015:   station.A2015,
			A2016:   station.A2016,
			A2017:   station.A2017,
			A2018:   station.A2018,
			A2019:   station.A2018,
			A2020:   station.A2020,
			A2021:   station.A2021,
		}
		rslt = append(rslt, b)
	}
	response.Data = rslt
	return response, nil
}

func main() {

	// load app.env file data to struct
	config, err := LoadConfig(".")

	lis, err := net.Listen("tcp", config.ServerAddress)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterFrequentationServer(s, &grpcServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
