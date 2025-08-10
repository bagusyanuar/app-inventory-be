package dto

import "github.com/bagusyanuar/app-inventory-be/internal/domain/entity"

type BranchAddressDTO struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

func ToBranchAddress(data *entity.BranchAddress) *BranchAddressDTO {
	return &BranchAddressDTO{
		ID:      data.ID.String(),
		Address: data.Address,
	}
}

func ToBranchAddresses(data []entity.BranchAddress) []BranchAddressDTO {
	branchAddresses := make([]BranchAddressDTO, 0)
	for _, datum := range data {
		branchAddress := *ToBranchAddress(&datum)
		branchAddresses = append(branchAddresses, branchAddress)
	}
	return branchAddresses
}
