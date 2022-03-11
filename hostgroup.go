package main

import (
	"encoding/json"
)

func create_jsreq_hostgroup_get() JSONRequest {
	var empty_map = map[string]interface{}{}
	return JSONRequest{"2.0", "hostgroup.get", 0, zabbi_api_token, empty_map}
}

func hostgroup_get() []Hostgroup {
	jr := create_jsreq_hostgroup_get()
	body := api_call(jr)
	resp := HostgroupRespSimple{}
	json.Unmarshal(body, &resp)
	return resp.Result
}

func allGrpRoRights() []Permission {
	hgrps := hostgroup_get()
	rights := make([]Permission, len(hgrps))
	for i, hgrp := range hgrps {
		rights[i] = Permission{hgrp.GroupId, "2"}
	}
	return rights
}
