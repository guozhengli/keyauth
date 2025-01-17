package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/app/system"
	"github.com/infraboard/keyauth/app/token"
	"github.com/infraboard/keyauth/app/token/issuer"
	"github.com/infraboard/keyauth/app/user"
	"github.com/infraboard/keyauth/app/verifycode"
	"github.com/infraboard/keyauth/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col    *mongo.Collection
	issuer issuer.Issuer
	system system.Service
	log    logger.Logger
	user   user.ServiceServer
	token  token.ServiceServer

	verifycode.UnimplementedServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("verify_code")

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

	s.system = app.GetInternalApp(system.AppName).(system.Service)
	s.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	s.token = app.GetGrpcApp(token.AppName).(token.ServiceServer)

	is, err := issuer.NewTokenIssuer()
	if err != nil {
		return err
	}
	s.issuer = is
	s.log = zap.L().Named("Verify Code")
	return nil
}

func (s *service) Name() string {
	return verifycode.AppName
}

func (s *service) Registry(server *grpc.Server) {
	verifycode.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
