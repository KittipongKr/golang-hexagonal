package domain

import (
	"time"
)

type User struct {
	UserID              string        `json:"userId"`
	AccountID           string        `json:"accountId"`
	TitleTh             string        `json:"titleTh"`
	FirstNameTh         string        `json:"firstNameTh"`
	LastNameTh          string        `json:"lastNameTh"`
	NameTh              string        `json:"nameTh"`
	TitleEn             string        `json:"titleEn"`
	FirstNameEn         string        `json:"firstNameEn"`
	LastNameEn          string        `json:"lastNameEn"`
	NameEn              string        `json:"nameEn"`
	Email               string        `json:"email"`
	EmailOneID          string        `json:"emailOneId"`
	NickName            string        `json:"nickName,omitempty"`
	Tel                 string        `json:"tel"`
	EmployeeID          string        `json:"employeeId"`
	PositionID          string        `json:"positionId"`
	PositionName        string        `json:"positionName"`
	PositionLevel       string        `json:"positionLevel"`
	TaxID               string        `json:"taxId"`
	CompanyID           string        `json:"companyId"`
	CompanyFullNameTh   string        `json:"companyFullNameTh"`
	CompanyFullNameEng  string        `json:"companyFullNameEng"`
	CompanyShortNameTh  string        `json:"companyShortNameTh"`
	CompanyShortNameEng string        `json:"companyShortNameEng"`
	Station             string        `json:"station"`
	ContractType        string        `json:"contractType"`
	Type                string        `json:"type"`
	Lists               []UserList    `json:"list"`
	Sections            []UserSection `json:"section"`
}

type UserList struct {
	CompanyID    string `json:"companyId"`
	ID           string `json:"_id"`
	OrgChartName string `json:"orgChartName"`
	OrgChartType string `json:"orgChartType"`
	Name         string `json:"name"`
}

type UserSection struct {
	ID           string `json:"_id"`
	OrgChartName string `json:"orgChartName"`
	OrgChartType string `json:"orgChartType"`
	Name         string `json:"name"`
}

type UserHistory struct {
	OldDetail       string    `json:"oldDetail"`
	NewDetail       string    `json:"newDetail"`
	OldCompany      string    `json:"oldCompany"`
	OldPosition     string    `json:"oldPosition"`
	OldOrgChartName string    `json:"oldOrgChartName"`
	CreateAt        time.Time `json:"createAt"`
}
