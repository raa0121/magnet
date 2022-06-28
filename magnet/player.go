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
	collisionLeftUp, collisionRightDown Point
	frameNum int
	leftUp Point
	isJump, isSlide bool
	objectHit [][]bool
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
	player.collisionLeftUp = Point{60, 0}
	player.collisionRightDown = Point{200, 235}

	player.objectHit = make([][]bool, len(maps.Maps))
	for i, m := range maps.Maps {
		player.objectHit[i] = make([]bool, len(m.Objects))
		for j := range m.Objects {
			player.objectHit[i][j] = false
		}
	}
}
