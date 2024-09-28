package service

import "wall-backend/internal/dao"

type CommunityService struct {
	communityDao dao.CommunityDao
}

func NewCommunityService(communityDao dao.CommunityDao) CommunityService {
	return CommunityService{
		communityDao: communityDao,
	}
}

var da dao.CommunityDao

func (service CommunityService) AllExpression() error {
	return da.AllExpression()
}
