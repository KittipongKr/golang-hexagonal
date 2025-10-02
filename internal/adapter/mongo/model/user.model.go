package mongo_model

import "time"

type User struct {
	BaseModel           `bson:",inline"`
	UserID              string        `bson:"user_id"`
	AccountID           string        `bson:"account_id"`
	TitleTh             string        `bson:"title_th"`
	FirstNameTh         string        `bson:"first_name_th"`
	LastNameTh          string        `bson:"last_name_th"`
	NameTh              string        `bson:"name_th"`
	TitleEn             string        `bson:"title_en"`
	FirstNameEn         string        `bson:"first_name_en"`
	LastNameEn          string        `bson:"last_name_en"`
	NameEn              string        `bson:"name_en"`
	Email               string        `bson:"email"`
	EmailOneID          string        `bson:"email_one_id"`
	NickName            string        `bson:"nick_name"`
	Tel                 string        `bson:"tel"`
	EmployeeID          string        `bson:"employee_id"`
	PositionID          string        `bson:"position_id"`
	PositionName        string        `bson:"position_name"`
	PositionLevel       string        `bson:"position_level"`
	TaxID               string        `bson:"tax_id"`
	CompanyID           string        `bson:"company_id"`
	CompanyFullNameTh   string        `bson:"company_full_name_th"`
	CompanyFullNameEng  string        `bson:"company_full_name_eng"`
	CompanyShortNameTh  string        `bson:"company_short_name_th"`
	CompanyShortNameEng string        `bson:"company_short_name_eng"`
	Station             string        `bson:"station"`
	Type                string        `bson:"type"`
	Lists               []UserList    `bson:"list"`
	Sections            []UserSection `bson:"section"`
}

type UserList struct {
	CompanyID    string `bson:"company_id"`
	ID           string `bson:"_id"`
	OrgChartName string `bson:"org_chart_name"`
	OrgChartType string `bson:"org_chart_type"`
	Name         string `bson:"name"`
}

type UserSection struct {
	ID           string `bson:"_id"`
	OrgChartName string `bson:"org_chart_name"`
	OrgChartType string `bson:"org_chart_type"`
	Name         string `bson:"name"`
}

type UserHistory struct {
	OldDetail       string    `bson:"old_detail"`
	NewDetail       string    `bson:"new_detail"`
	OldCompany      string    `bson:"old_company"`
	OldPosition     string    `bson:"old_position"`
	OldOrgChartName string    `bson:"old_org_chart_name"`
	CreateAt        time.Time `bson:"create_at"`
}
