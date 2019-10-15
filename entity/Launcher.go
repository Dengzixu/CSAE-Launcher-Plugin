package entity

type LaunchConfig struct {
	Host   string
	Path   string
	Option string
	Token  string
}

type ResponseData struct {
	Success bool
	Message string
}

type Host struct {
	Ip   string
	Port string
}

type ApiPasswordResp struct {
	Code    int
	Status  string
	Message string
	Data    struct {
		ServerPass string `json:"server_pass"`
		Remain     string
	}
}

type ApiUserInfoResp struct {
	Code    int
	Status  string
	Message string
	Data    struct {
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
