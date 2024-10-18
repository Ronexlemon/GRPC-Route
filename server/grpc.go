package server

import (
	"context"
	"sync"

	pb "github.com/RonexLemon/RPC/services/genproto/route"
	"google.golang.org/protobuf/proto"
)

type routeGuideServer struct{
	pb.UnimplementedRouteGuideServer
	 savedFeatures  []*pb.Feature
	 mu  sync.Mutex
	 routeNotes map[string][]*pb.RouteNote

}

func(s *routeGuideServer) GetFeature(ctx context.Context,point *pb.Point)(*pb.Feature,error){
	//logic to get feature
	for _,feature := range s.savedFeatures{
		if proto.Equal(feature.Location,point){
			return feature,nil
		}
		
	}
	return &pb.Feature{Location: point},nil
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle,stream *pb.RouteGuide_ListFeaturesServer)error{

}

func (s *routeGuideServer) RecordRoute(stream *pb.RouteGuide_RecordRouteServer)error{

}
func(s *routeGuideServer)RouteChat(stream *pb.RouteGuide_RouteChatServer)error{

}