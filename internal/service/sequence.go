package service

import (
	"context"

	pb "idgen/api/idgen/v1"
	"idgen/internal/biz"
)

type SequenceService struct {
	pb.UnimplementedSequenceServer

	uc *biz.SequenceUsecase
}

func NewSequenceService(uc *biz.SequenceUsecase) *SequenceService {
	return &SequenceService{
		uc: uc,
	}
}

func (s *SequenceService) Next(ctx context.Context, req *pb.NextRequest) (*pb.NextReply, error) {
	sequence, _ := s.uc.Next(ctx)
	return &pb.NextReply{Id: sequence.ID}, nil
}
