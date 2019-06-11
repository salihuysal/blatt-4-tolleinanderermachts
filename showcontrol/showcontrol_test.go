package main

import (
	"context"
	"testing"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"github.com/stretchr/testify/assert"
)

func TestShowAdd(t *testing.T) {

	ShowControl := ShowControl{NextID: 0, Shows: make(map[int]Show)}

	response := proto.RequestResponse{}

	_ = ShowControl.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 0, RoomId: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestShowDelete(t *testing.T) {

	ShowControl := ShowControl{NextID: 0, Shows: make(map[int]Show)}

	response := proto.RequestResponse{}

	_ = ShowControl.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 0, RoomId: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = ShowControl.DeleteShow(context.TODO(), &proto.DeleteShowRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestShowDeleteEmpty(t *testing.T) {

	ShowControl := ShowControl{NextID: 0, Shows: make(map[int]Show)}

	response := proto.RequestResponse{}

	_ = ShowControl.DeleteShow(context.TODO(), &proto.DeleteShowRequest{Id: 0}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestShowDeleteNegative(t *testing.T) {

	ShowControl := ShowControl{NextID: 0, Shows: make(map[int]Show)}

	response := proto.RequestResponse{}

	_ = ShowControl.DeleteShow(context.TODO(), &proto.DeleteShowRequest{Id: -1}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestShowDeleteNotFound(t *testing.T) {

	ShowControl := ShowControl{NextID: 0, Shows: make(map[int]Show)}

	response := proto.RequestResponse{}

	_ = ShowControl.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 0, RoomId: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = ShowControl.DeleteShow(context.TODO(), &proto.DeleteShowRequest{Id: 2}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}
