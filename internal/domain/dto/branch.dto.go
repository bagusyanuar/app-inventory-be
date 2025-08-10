package dto

import "github.com/bagusyanuar/app-inventory-be/internal/domain/entity"

type BranchDTO struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Address  *BranchAddressDTO  `json:"address"`
	Contacts []BranchContactDTO `json:"contacts"`
}

func ToBranch(data *entity.Branch) *BranchDTO {
	var address *BranchAddressDTO

	if data.Address != nil {
		address = ToBranchAddress(data.Address)
	}

	contacts := ToBranchContacts(data.Contacts)
	return &BranchDTO{
		ID:       data.ID.String(),
		Name:     data.Name,
		Address:  address,
		Contacts: contacts,
	}
}

func ToBranches(data []entity.Branch) []BranchDTO {
	branchs := make([]BranchDTO, 0)
	for _, datum := range data {
		branch := *ToBranch(&datum)
		branchs = append(branchs, branch)
	}
	return branchs
}
