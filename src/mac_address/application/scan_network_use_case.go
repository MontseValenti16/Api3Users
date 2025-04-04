package application

import "API3/src/mac_address/domain"



type ScanNetworkUseCase struct {
	Repo domain.NetworkRepository
}

func NewScanNetworkUseCase(repo domain.NetworkRepository) *ScanNetworkUseCase {
	return &ScanNetworkUseCase{Repo: repo}
}

func (s *ScanNetworkUseCase) Execute() ([]domain.Device, error) {
	return s.Repo.ScanNetwork()
}
