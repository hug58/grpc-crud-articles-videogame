package article

import (
	pb "api-grpc-articles-videogame/proto"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*pb.Article, error)
	GetOne(ctx context.Context, id uint32) (*pb.CreateArticlerRequest, error)
	GetByUser(ctx context.Context, userID uint32) ([]*pb.Article, error)
	Create(ctx context.Context, article *pb.CreateArticlerRequest) (*pb.CreateArticlerRequest, error)
	Update(ctx context.Context, id uint32, article *pb.CreateArticlerRequest) (*pb.CreateArticlerRequest, error)
	Delete(ctx context.Context, id uint32) error
}
