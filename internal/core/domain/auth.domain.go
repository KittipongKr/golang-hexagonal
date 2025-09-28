package domain

type OneAuthResb struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type AuthResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccountID    string `json:"account_id"`
}

type AuthAccountResp struct {
	ID                  string      `json:"id"`
	FirstNameTh         string      `json:"first_name_th"`
	MiddleNameTh        interface{} `json:"middle_name_th"`
	LastNameTh          string      `json:"last_name_th"`
	FirstNameEng        string      `json:"first_name_eng"`
	MiddleNameEng       interface{} `json:"middle_name_eng"`
	LastNameEng         string      `json:"last_name_eng"`
	SpecialTitleNameTh  string      `json:"special_title_name_th"`
	AccountTitleTh      string      `json:"account_title_th"`
	SpecialTitleNameEng interface{} `json:"special_title_name_eng"`
	AccountTitleEng     string      `json:"account_title_eng"`
	IDCardType          string      `json:"id_card_type"`
	IDCardNum           string      `json:"id_card_num"`
	HashIDCardNum       string      `json:"hash_id_card_num"`
	AccountCategory     string      `json:"account_category"`
	AccountSubCategory  string      `json:"account_sub_category"`
	ThaiEmail           string      `json:"thai_email"`
	ThaiEmail2          string      `json:"thai_email2"`
	ThaiEmail3          interface{} `json:"thai_email3"`
	StatusCd            string      `json:"status_cd"`
	BirthDate           string      `json:"birth_date"`
	StatusDt            string      `json:"status_dt"`
	RegisterDt          string      `json:"register_dt"`
	AddressID           interface{} `json:"address_id"`
	CreatedAt           string      `json:"created_at"`
	CreatedBy           string      `json:"created_by"`
	UpdatedAt           string      `json:"updated_at"`
	UpdatedBy           string      `json:"updated_by"`
	Reason              interface{} `json:"reason"`
	TelNo               interface{} `json:"tel_no"`
	NameOnDocumentTh    interface{} `json:"name_on_document_th"`
	NameOnDocumentEng   interface{} `json:"name_on_document_eng"`
	BlockchainFlg       interface{} `json:"blockchain_flg"`
	NicknameEng         string      `json:"nickname_eng"`
	NicknameTh          string      `json:"nickname_th"`
	FullIDCardNum       interface{} `json:"full_id_card_num"`
	TrustLevel          interface{} `json:"trust_level"`
	Mobile              []struct {
		ID        string      `json:"id"`
		MobileNo  string      `json:"mobile_no"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
		Pivot     struct {
			AccountID             string `json:"account_id"`
			MobileID              string `json:"mobile_id"`
			CreatedAt             string `json:"created_at"`
			UpdatedAt             string `json:"updated_at"`
			StatusCd              string `json:"status_cd"`
			PrimaryFlg            string `json:"primary_flg"`
			MobileConfirmFlg      string `json:"mobile_confirm_flg"`
			MobileConfirmDt       string `json:"mobile_confirm_dt"`
			MobileLoginConfirmFlg string `json:"mobile_login_confirm_flg"`
			MobileLoginConfirmDt  string `json:"mobile_login_confirm_dt"`
			Type                  string `json:"type"`
		} `json:"pivot"`
	} `json:"mobile"`
	Email []struct {
		ID        string      `json:"id"`
		Email     string      `json:"email"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
		Pivot     struct {
			AccountID            string      `json:"account_id"`
			EmailID              string      `json:"email_id"`
			CreatedAt            string      `json:"created_at"`
			UpdatedAt            string      `json:"updated_at"`
			StatusCd             string      `json:"status_cd"`
			PrimaryFlg           string      `json:"primary_flg"`
			EmailConfirmFlg      string      `json:"email_confirm_flg"`
			EmailConfirmDt       string      `json:"email_confirm_dt"`
			EmailLoginConfirmFlg string      `json:"email_login_confirm_flg"`
			EmailLoginConfirmDt  interface{} `json:"email_login_confirm_dt"`
		} `json:"pivot"`
	} `json:"email"`
	Address          []interface{} `json:"address"`
	AccountAttribute interface{}   `json:"account_attribute"`
	Status           string        `json:"status"`
	LastUpdate       string        `json:"last_update"`
}
