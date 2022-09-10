package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-account-service/model/pb"
)

type UserMongoDBTransactionCreator interface {
	CreateNoSQLSession() (session mongo.Session, err error)
	StartTransaction(session mongo.Session) (err error)
	AbortTransaction(session mongo.Session, ctx context.Context) (err error)
	CommitTransaction(session mongo.Session, ctx context.Context) (err error)
}

type UserMongoDBRepositoryCommand interface {
	CreateUser(ctx context.Context, session mongo.Session, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error)
}

type UserMongoDBRepositoryQuery interface {
	FindUserByEmail(ctx context.Context, email string) (response *pb.User, err error)
	FindUserById(ctx context.Context, id string) (response *pb.User, err error)
}

type UserMongoDBRepositoryStore interface {
	UserMongoDBTransactionCreator
	UserMongoDBRepositoryCommand
	UserMongoDBRepositoryQuery
}