package dto

import "github.com/bagusyanuar/app-inventory-be/internal/domain/entity"

type BranchContactDTO struct {
	ID    string  `json:"id"`
	Type  string  `json:"type"`
	Name  *string `json:"name"`
	Value string  `json:"value"`
}

func ToBranchContact(data *entity.BranchContact) *BranchContactDTO {
	return &BranchContactDTO{
		ID:    data.ID.String(),
		Type:  data.Type,
		Name:  data.Name,
		Value: data.Value,
	}
}

func ToBranchContacts(data []entity.BranchContact) []BranchContactDTO {
	branchContacts := make([]BranchContactDTO, 0)
	for _, datum := range data {
		branchContact := *ToBranchContact(&datum)
		branchContacts = append(branchContacts, branchContact)
	}
	return branchContacts
}
