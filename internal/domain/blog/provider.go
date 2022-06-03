package blog

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUseCase)
