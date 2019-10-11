package entity

type LaunchParam struct {
    Host  string
    Path  string
    Param string
    Token string
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
