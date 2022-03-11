package main

type JSONRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Id int `json:"id"`
	Auth string `json:"auth,omitempty"`
	Params interface{} `json:"params"`
}

type JSONResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id int `json:"id"`
}

type Hostgroup struct {
	GroupId string `json:"groupid"`
	Name string `json:"name"`
}


type HostgroupRespSimple struct {
	JSONResponse
	Result []Hostgroup `json:"result"`
}

type Permission struct {
	Id string `json:"id"`
	Permission string `json:"permission"`
}

type Usergroup struct {
	UsergrpId string `json:"usrgrpid"`
	Name string `json:"name"`
	Rights []Permission `json:"rights"`
}

type UsegroupRespSimple struct {
	JSONResponse
	Result []Usergroup `json:"result"`
}
