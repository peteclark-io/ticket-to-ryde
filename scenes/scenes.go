package scenes

import (
	"context"

	"github.com/faiface/pixel/pixelgl"
)

type Scene func(ctx context.Context, win *pixelgl.Window)
