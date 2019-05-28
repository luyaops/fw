package types

import (
	"database/sql"
	pb "sinaclouds/user/service"
)

type Injector struct {
	DB   *sql.DB
	User *pb.User
}
