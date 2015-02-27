package io

import "github.com/discoviking/roguemike/common"

type UpdateBundle struct {
	Player   *PlayerData
	Entities []*EntityData
}

type PlayerData struct{}

type EntityType uint32

type EntityData struct {
	Type EntityType
	common.Coords
}

type Manager struct {
	output chan<- *UpdateBundle
}

func (mgr *Manager) SetOutput(output chan<- *UpdateBundle) {
	mgr.output = output
}

func (mgr *Manager) Update(bundle *UpdateBundle) {
	mgr.output <- bundle
}
