package apis

import "context"

type Executor = func(context.Context, string)
