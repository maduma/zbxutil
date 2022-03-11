package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func createJsreqUsergroupGet(usergroupId int) JSONRequest {
	params := map[string]interface{}{
		"usrgrpids":    usergroupId,
		"selectRights": "*",
	}
	return JSONRequest{"2.0", "usergroup.get", 0, zabbi_api_token, params}
}

func createJsreqUsergroupUpdate(usergroupId int, rights []Permission) JSONRequest {
	params := map[string]interface{}{
		"usrgrpid": usergroupId,
		"rights":   rights,
	}
	return JSONRequest{"2.0", "usergroup.update", 0, zabbi_api_token, params}
}

func usergroupGet(usergroupId int) (Usergroup, error) {
	jr := createJsreqUsergroupGet(usergroupId)
	body := api_call(jr)
	resp := UsegroupRespSimple{}
	json.Unmarshal(body, &resp)
	if len(resp.Result) == 1 {
		ugrp := resp.Result[0]
		return ugrp, nil
	}
	err := fmt.Errorf("Cannot found usegroup with id %d", usergroupId)
	return Usergroup{}, err
}

func usergroupUpdate(usergroupId int, rights []Permission) {
	jr := createJsreqUsergroupUpdate(usergroupId, rights)
	api_call(jr)
}

func sortRights(rights []Permission) {
	sort.Slice(rights, func(i, j int) bool {
		i, _ = strconv.Atoi(rights[i].Id)
		j, _ = strconv.Atoi(rights[j].Id)
		return i < j
	})
}

func isSamePermission(a []Permission, b []Permission) bool {
	sortRights(a)
	sortRights(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func setUgrpRight(usergroupId int, rights []Permission) {
	ugrp, err := usergroupGet(usergroupId)
	if err != nil {
		panic(err)
	}
	same := isSamePermission(ugrp.Rights, rights)
	if !same {
		logger.Printf("Update Rights (permissions) of usergroupid %d", usergroupId)
		usergroupUpdate(usergroupId, rights)
	}
}
