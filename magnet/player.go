package magnet

const (
	playerFrame0X = 0
	playerFrame0Y = 0
	playerFrameWidth = 256
	playerFrameHeight = 256
	playerFrameNum = 8
	playerFootY = 790
)

type Player struct {
	frame0 Point
	frameSize Point
	frameNum int
	leftUp Point
	isJump, isSlide bool
}

func init() {
	player = &Player{
		frame0: Point{0, 0},
		frameSize: Point{playerFrameWidth, playerFrameHeight},
		frameNum: playerFrameNum,
		isJump: false,
		isSlide: false,
	}
	player.leftUp = Point{
		(ScreenWidth - player.frameSize.X) / 2,
		(playerFootY - player.frameSize.Y),
	}
}
