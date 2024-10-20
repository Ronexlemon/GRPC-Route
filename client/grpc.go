package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/RonexLemon/RPC/services/genproto/route"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)
var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:9090", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func printFeature(client pb.RouteGuideClient, point *pb.Point){
    ctx,cancel:= context.WithTimeout(context.Background(),10*time.Second)
    defer cancel()
    feature,err := client.GetFeature(ctx,point)
    if err!=nil{
        log.Fatal("Failed to get a feature",err)}

        fmt.Println("the feature from the point is",feature)

}

func main(){
    flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			// *caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
    conn, err := grpc.NewClient(*serverAddr, opts...)
    if err !=nil{
        log.Fatal("Failed to open a connection",err)
    }
    defer conn.Close()

    client := pb.NewRouteGuideClient(conn)
    printFeature(client,&pb.Point{Latitude:409146138 ,Longitude: -746188906})
    
}