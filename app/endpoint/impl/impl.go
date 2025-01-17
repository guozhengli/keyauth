package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/app/micro"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string
	micro         micro.ServiceServer
	log           logger.Logger

	endpoint.UnimplementedServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("endpoint")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = col

	s.micro = app.GetGrpcApp(micro.AppName).(micro.ServiceServer)
	s.log = zap.L().Named("Endpoint")
	return nil
}

func (s *service) Name() string {
	return endpoint.AppName
}

func (s *service) Registry(server *grpc.Server) {
	endpoint.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
