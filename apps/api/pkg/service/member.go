package service

import (
	"context"
	"time"

	"github.com/vincentkdeli/vinance-backend/entity"
	"github.com/vincentkdeli/vinance-backend/model"
	"github.com/vincentkdeli/vinance-backend/repository"
)

type MemberService struct {
	memberRepo repository.Member
}

func NewMemberService(memberRepo repository.Member) *MemberService {
	return &MemberService{
		memberRepo: memberRepo,
	}
}

func (s *MemberService) CreateMember(ctx context.Context, request *model.CreateMemberRequest) (bool, error) {
	res, err := s.memberRepo.InsertOne(ctx, &entity.Member{
		Email:       request.Email,
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return false, nil
	}

	return res != nil, nil
}
