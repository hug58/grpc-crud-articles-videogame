package main

import (
	"api-grpc-articles-videogame/internal/data"
	"api-grpc-articles-videogame/pkg/article"
	"os"

	pb "api-grpc-articles-videogame/proto"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedArticleServiceServer
	Repository article.Repository
}

var articles []*pb.Article

//Listar Tareas
func (s *server) ListArticle(ctx context.Context, req *empty.Empty) (*pb.ListArticlesResponse, error) {
	articles_all, err := s.Repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return &pb.ListArticlesResponse{
		Articles: articles_all,
	}, nil
}

//Crear una tarea
func (s *server) CreateArticle(ctx context.Context, req *pb.CreateArticlerRequest) (*pb.ArticleId, error) {
	log.Printf("Creating Article Videogame %v", req)
	article, err := s.Repository.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	//articles = append(articles, req.Article)
	return &pb.ArticleId{
		ArticleId: article.Id,
	}, nil
}

//Filtrar por Id de user
func (s *server) ListArticleByUser(ctx context.Context, req *pb.ListArticlesByUserRequest) (*pb.ListArticlesResponse, error) {
	log.Printf("Get Article By User %v", req)

	article, err := s.Repository.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.ListArticlesResponse{Articles: article}, nil
}

//Delete

func (s *server) DeleteArticle(ctx context.Context, req *pb.ArticleId) (*pb.Response, error) {
	err := s.Repository.Delete(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Msg: fmt.Sprintf("DELETE %v", req.ArticleId)}, nil
}

//GET ARTICLE BY ID

func (s *server) GetOneArticle(ctx context.Context, req *pb.ArticleId) (*pb.CreateArticlerRequest, error) {
	log.Printf("Get Article By ID %v", req.ArticleId)

	article, err := s.Repository.GetOne(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	return article, nil
}

//UPDATE ARTICLE

func (s *server) UpdateArticle(ctx context.Context, req *pb.CreateArticlerRequest) (*pb.CreateArticlerRequest, error) {

	article, err := s.Repository.Update(ctx, req.Id, req)
	if err != nil {
		return nil, err
	}

	return article, nil
}

//MAIN

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	fmt.Println("API V1 CATALOGO VIDEOJUEGOS")
	fmt.Println("CON MICROSERVICIOS INCLUIDOS ")

	if err != nil {
		log.Fatalf("Error cannot create tcp connection %v", err)
	}

	//conexion a la base de datos
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	/*
		gRPC
	*/

	if err != nil {
		log.Fatalf("Error cannot create tcp connection %v", err)
	}
	log.Printf("Connection established running on port %v", port)

	ser := grpc.NewServer()
	pb.RegisterArticleServiceServer(ser, &server{
		Repository: &data.ArticlesRepository{
			Data: data.New(),
		},
	})

	if err := ser.Serve(listen); err != nil {
		log.Fatalf("Error cannot initialize the server: %v", err.Error())
	}

	defer data.Close()
}