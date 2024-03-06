package service

import (
	"context"

	pb "idgen/api/idgen/v1"
	"idgen/internal/biz"
)

type SnowflakeService struct {
	pb.UnimplementedSnowflakeServer

	uc *biz.SnowflakeUsecase
}

func NewSnowflakeService(uc *biz.SnowflakeUsecase) *SnowflakeService {
	return &SnowflakeService{
		uc: uc,
	}
}

func (s *SnowflakeService) Generate(ctx context.Context, req *pb.GenerateRequest) (*pb.GenerateReply, error) {
	snowflake, _ := s.uc.Generate(ctx)
	return &pb.GenerateReply{Id: snowflake.ID}, nil
}
