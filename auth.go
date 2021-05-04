package core

import (
	"context"
	"errors"
	"github.com/whatshowing/core/keys"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegAuthCtx struct {
	ID          string
	Email       string
	AccountType string
	ObjectId    primitive.ObjectID
}

type AuthCtx struct {
	ID       string
	Email    string
	ObjectId primitive.ObjectID
}

var UserStatuses = newUserStatusRegistry()

type UserStatus struct {
	Name string
}

type userStatusRegistry struct {
	Registration *UserStatus
	Blocked      *UserStatus
	Disabled     *UserStatus
	Enabled      *UserStatus

	statuses []*UserStatus
}

func newUserStatusRegistry() *userStatusRegistry {

	registration := &UserStatus{Name: "registration"}
	blocked := &UserStatus{Name: "blocked"}
	disabled := &UserStatus{Name: "disabled"}
	enabled := &UserStatus{Name: "enabled"}

	return &userStatusRegistry{
		Registration: registration,
		Blocked:      blocked,
		Disabled:     disabled,
		Enabled:      enabled,

		statuses: []*UserStatus{registration, blocked, disabled, enabled},
	}
}

func (s *userStatusRegistry) List() []*UserStatus {
	return s.statuses
}

func (s *userStatusRegistry) Parse(status string) (*UserStatus, error) {

	for _, st := range s.List() {
		if st.Name == status {
			return st, nil
		}
	}

	return nil, errors.New("cloud not parse user status")
}

var RpcHeaders = newRpcHeadersRegistry()

type RpcHeader struct {
	Name string
}

type rpcHeaderRegistry struct {
	Auth             *RpcHeader
	SetAuth          *RpcHeader
	AuthRefreshTk    *RpcHeader
	SetAuthRefreshTk *RpcHeader
	RegAuth          *RpcHeader
	SetRegAuth       *RpcHeader
	DeviceTk         *RpcHeader
	SetDeviceTk      *RpcHeader

	statuses []*RpcHeader
}

func newRpcHeadersRegistry() *rpcHeaderRegistry {

	auth := &RpcHeader{Name: "_a"}
	setAuth := &RpcHeader{Name: "set-_a"}
	authRefreshTk := &RpcHeader{Name: "_rf_tk"}
	setAuthRefreshTk := &RpcHeader{Name: "set-_rf_tk"}
	regAuth := &RpcHeader{Name: "r_a"}
	setRegAuth := &RpcHeader{Name: "set-r_a"}
	deviceTk := &RpcHeader{Name: "d_tk"}
	setDeviceTk := &RpcHeader{Name: "set-d_tk"}

	return &rpcHeaderRegistry{
		Auth:             auth,
		SetAuth:          setAuth,
		AuthRefreshTk:    authRefreshTk,
		SetAuthRefreshTk: setAuthRefreshTk,
		RegAuth:          regAuth,
		SetRegAuth:       setRegAuth,
		DeviceTk:         deviceTk,
		SetDeviceTk:      setDeviceTk,
		statuses:         []*RpcHeader{auth, setRegAuth, regAuth, setRegAuth, deviceTk, setDeviceTk},
	}
}

func (s *rpcHeaderRegistry) List() []*RpcHeader {
	return s.statuses
}

func (s *rpcHeaderRegistry) Parse(status string) (*RpcHeader, error) {
	for _, st := range s.List() {
		if st.Name == status {
			return st, nil
		}
	}
	return nil, errors.New("cloud not parse user status")
}

func GetAuthUser(ctx context.Context, authCtx *AuthCtx) error {
	uv := ctx.Value(keys.UserCtxKey)
	if uv == nil {
		return status.Error(codes.PermissionDenied, "user not allowed")
	}
	u := uv.(AuthCtx)

	objectId, err := ParseObjectId(u.ID)

	if err != nil {
		return err
	}

	authCtx.ID = u.ID
	authCtx.Email = u.Email
	authCtx.ObjectId = objectId
	return nil
}
