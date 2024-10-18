package server

import (
	"context"

	pb "github.com/RonexLemon/RPC/services/genproto/route"
)

type routeGuideServer struct{}

func(s *routeGuideServer) GetFeature(ctx context.Context,point *pb.Point)(*pb.Feature,error){
	//logic to get feature
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle,stream *pb.RouteGuide_ListFeaturesServer)error{

}

func (s *routeGuideServer) RecordRoute(stream *pb.RouteGuide_RecordRouteServer)error{

}
func(s *routeGuideServer)RouteChat(stream *pb.RouteGuide_RouteChatServer)error{
	
}