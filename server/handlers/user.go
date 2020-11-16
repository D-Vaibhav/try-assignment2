package handlers

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/vaibhav/try-assignment2/server/protos/user"
)

type userServer struct {
	log hclog.Logger
}

// CONSTRUCTOR
func NewUserServer(l hclog.Logger) *userServer {
	return &userServer{l}
}

// IMPLEMENTING INTERFACE
func (s *userServer) NameUsername(ctx context.Context, req *user.NameUsernameReq) (*user.NameUsernameRes, error) {
	// getting data from req method
	name, userName := req.GetName(), req.GetUsername()
	s.log.Info("name:", name, " username:", userName)

	// pushing it to kafka instance

	response := &user.NameUsernameRes{IsSent: true}
	return response, nil
}
