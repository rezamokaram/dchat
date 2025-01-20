package common

import (
	"context"
)

type ServiceGetter[T any] func(context.Context) T
