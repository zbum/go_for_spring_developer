package bean

type ScoreService struct {
	scoreRepository *ScoreRepository
}

func NewScoreService(scoreRepository *ScoreRepository) *ScoreService {
	return &ScoreService{scoreRepository}
}

func (s *ScoreService) GetScores() []Score {
	return s.scoreRepository.FindAll()
}
