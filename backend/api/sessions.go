package api

import "github.com/yanando/StudentLink/datamanager"

type SessionManager struct {
	Sessions map[string]*datamanager.User
}

func GetSession()
