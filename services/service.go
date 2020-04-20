package services

import (
	"context"
)

type Servicer interface {
	ExecuteRequest(context.Context)
	ServicerName() string
}
