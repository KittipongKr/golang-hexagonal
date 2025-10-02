package helpers

import (
	m "csat-servay/internal/adapter/calls/models"
	d "csat-servay/internal/core/domain"
)

func ConvRaUserToUser(raUsers []m.RaUserResult) []d.User {

	users := make([]d.User, 0)
	for _, user := range raUsers {
		lists := make([]d.UserList, 0)
		if user.Type != "outer" {
			if len(user.List) > 0 {
				for _, list := range user.List {
					lists = append(lists, d.UserList{
						CompanyID:    list.CompanyID,
						ID:           list.ID,
						OrgChartName: list.OrgChartName,
						OrgChartType: list.OrgChartType,
						Name:         list.Name,
					})
				}
			}

			sections := make([]d.UserSection, 0)
			if len(user.Section) > 0 {
				for _, section := range user.Section {
					sections = append(sections, d.UserSection{
						ID:           section.ID,
						OrgChartName: section.OrgChartName,
						OrgChartType: section.OrgChartType,
						Name:         section.Name,
					})
				}
			}

			users = append(users, d.User{
				UserID:              user.UserID,
				AccountID:           user.AccountID,
				TitleTh:             user.TitleTh,
				FirstNameTh:         user.FirstNameTh,
				LastNameTh:          user.LastNameTh,
				NameTh:              user.NameTh,
				TitleEn:             user.TitleEn,
				FirstNameEn:         user.FirstNameEn,
				LastNameEn:          user.LastNameEn,
				NameEn:              user.NameEn,
				Email:               user.Email,
				EmailOneID:          user.EmailOneID,
				NickName:            user.NickName,
				Tel:                 user.Tel,
				EmployeeID:          user.EmployeeID,
				PositionID:          user.PositionID,
				PositionName:        user.PositionName,
				PositionLevel:       user.PositionLevel,
				TaxID:               user.TaxID,
				CompanyID:           user.CompanyID,
				CompanyFullNameTh:   user.CompanyFullNameTh,
				CompanyFullNameEng:  user.CompanyFullNameEng,
				CompanyShortNameTh:  user.CompanyShortNameTh,
				CompanyShortNameEng: user.CompanyShortNameEng,
				Station:             user.Station,
				ContractType:        user.ContractType,
				Type:                user.Type,
				Lists:               lists,
				Sections:            sections,
			})

		}
	}

	return users
}
