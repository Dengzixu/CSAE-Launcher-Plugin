package entity

type ApiBasic struct {
	Code    int
	Status  string
	Message string
}

type ApiPasswordResp struct {
	ApiBasic
	Data struct {
		ServerPass string `json:"server_pass"`
		Remain     string
	}
}

type ApiUserInfoResp struct {
	ApiBasic
	Data struct {
		Id            int64
		Username      string
		Nickname      string
		Status        int
		Point         int
		discipline    string
		rank          int
		RegistTime    string `json:"regist_time"`
		LastLoginTime string `json:"last_login_time"`
		Global        int
		RankTimes     int `json:"rank_times"`
	}
}
