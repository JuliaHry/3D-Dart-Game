package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dart struct {
	Position  rl.Vector3
	Velocity  rl.Vector3
	Transform rl.Matrix
	Thrown    bool
}

// Funkcja wywoływana po wystrzeleniu strzały. Ustawia początkowe prędkości rzutki dla każdej osi
func (d *Dart) Throw(initialSpeed, angle float32) {
	d.Velocity.X = initialSpeed * float32(math.Sin(float64(angle*math.Pi/180)))
	d.Velocity.Y = initialSpeed * float32(math.Cos(float64(angle*math.Pi/180)))
	d.Velocity.Z = -initialSpeed
	d.Transform = rl.MatrixIdentity()
	d.Thrown = true
}


func (d *Dart) Update(deltaTime float32, targetPosition rl.Vector3, targetRadius float32) {
	if d.Thrown {
		d.Position.X += d.Velocity.X * deltaTime
		d.Position.Y += d.Velocity.Y * deltaTime
		d.Position.Z += d.Velocity.Z * deltaTime

		d.Velocity.Y -= 9.81 * deltaTime

		// Część napisana przez chatGPT:
		pitch := -float32(math.Atan2(float64(d.Velocity.Y), float64(-d.Velocity.Z)))
		yaw := float32(math.Atan2(float64(d.Velocity.X), float64(-d.Velocity.Z)))
		rotationPitch := rl.MatrixRotateX(pitch)
		rotationYaw := rl.MatrixRotateY(yaw)
		translation := rl.MatrixTranslate(d.Position.X, d.Position.Y, d.Position.Z)
		d.Transform = rl.MatrixMultiply(rotationYaw, rotationPitch)
		d.Transform = rl.MatrixMultiply(d.Transform, translation)
	}
}

func NewDart() *Dart {
	return &Dart{
		Position:  rl.NewVector3(0.0, 0.0, 0.0),
		Velocity:  rl.NewVector3(0.0, 0.0, 0.0),
		Transform: rl.MatrixIdentity(),
		Thrown:    false,
	}
}

// Funkcja rysująca suwak do ustawienia prędkości początkowej rzutki
func DrawSliderSpeed(position rl.Vector2, width float32, min, max, value float32) float32 {

	rl.DrawRectangleV(position, rl.NewVector2(width, 10), rl.LightGray)

	scrollDelta := rl.GetMouseWheelMove()
	if scrollDelta != 0 {
		value += scrollDelta * 0.1
	}

	if value < min {
		value = min
	} else if value > max {
		value = max
	}

	handleX := position.X + (width * ((value - min) / (max - min)))
	handleY := position.Y - 5

	rl.DrawRectangle(int32(handleX), int32(handleY), 10, 20, rl.Gray)

	return value
}

// Funckja rysująca suwak do ustawienia kąta rzutki 
func DrawSliderAngle(position rl.Vector2, width float32, min, max, value float32) float32 {

	rl.DrawRectangleV(position, rl.NewVector2(width, 10), rl.LightGray)

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mouseDelta := float32(rl.GetMouseDelta().X)

		value += mouseDelta * 0.1
	}

	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	handleX := position.X + (width * ((value - min) / (max - min)))
	handleY := position.Y - 5

	rl.DrawRectangle(int32(handleX), int32(handleY), 10, 20, rl.Gray)

	return value
}

// Funcja obliczająca punkty z rzutu
func Score(position rl.Vector3, position_1 rl.Vector3) int {

	distance := math.Sqrt(math.Pow(float64(position.X-position_1.X), 2) + math.Pow(float64(position.Y-position_1.Y), 2))
	angle_1 := math.Atan2(float64(position.Y-position_1.Y), float64(position.X-position_1.X)) * (180 / math.Pi)
	angle := 90 - angle_1

	if angle < 0 {
		angle += 360
	}

	score := 0

	if distance <= 0.15 {
		score = 50
	} else if distance <= 0.264 {
		score = 25
	} else if (distance <= 1.45 && distance > 0.3) || (distance < 2.43 && distance > 1.66) {
		if angle >= 9 && angle < 27 {
			score = 1
		} else if angle >= 27 && angle < 45 {
			score = 18
		} else if angle >= 45 && angle < 63 {
			score = 4
		} else if angle >= 63 && angle < 81 {
			score = 13
		} else if angle >= 81 && angle < 99 {
			score = 6
		} else if angle >= 99 && angle < 117 {
			score = 10
		} else if angle >= 117 && angle < 135 {
			score = 15
		} else if angle >= 135 && angle < 153 {
			score = 2
		} else if angle >= 153 && angle < 171 {
			score = 17
		} else if angle >= 171 && angle < 189 {
			score = 3
		} else if angle >= 189 && angle < 207 {
			score = 19
		} else if angle >= 207 && angle < 225 {
			score = 7
		} else if angle >= 225 && angle < 243 {
			score = 16
		} else if angle >= 243 && angle < 261 {
			score = 8
		} else if angle >= 261 && angle < 279 {
			score = 11
		} else if angle >= 279 && angle < 297 {
			score = 14
		} else if angle >= 297 && angle < 315 {
			score = 9
		} else if angle >= 315 && angle < 333 {
			score = 12
		} else if angle >= 333 && angle < 351 {
			score = 5
		} else if angle >= 351 || angle < 9 {
			score = 20
		}
	} else if distance <= 1.66 {
		if angle >= 9 && angle < 27 {
			score = 3
		} else if angle >= 27 && angle < 45 {
			score = 54
		} else if angle >= 45 && angle < 63 {
			score = 12
		} else if angle >= 63 && angle < 81 {
			score = 39
		} else if angle >= 81 && angle < 99 {
			score = 18
		} else if angle >= 99 && angle < 117 {
			score = 30
		} else if angle >= 117 && angle < 135 {
			score = 45
		} else if angle >= 135 && angle < 153 {
			score = 6
		} else if angle >= 153 && angle < 171 {
			score = 51
		} else if angle >= 171 && angle < 189 {
			score = 9
		} else if angle >= 189 && angle < 207 {
			score = 57
		} else if angle >= 207 && angle < 225 {
			score = 21
		} else if angle >= 225 && angle < 243 {
			score = 48
		} else if angle >= 243 && angle < 261 {
			score = 24
		} else if angle >= 261 && angle < 279 {
			score = 33
		} else if angle >= 279 && angle < 297 {
			score = 42
		} else if angle >= 297 && angle < 315 {
			score = 27
		} else if angle >= 315 && angle < 333 {
			score = 36
		} else if angle >= 333 && angle < 351 {
			score = 15
		} else if angle >= 351 || angle < 9 {
			score = 60
		}
	} else if distance <= 2.64 && distance > 2.43 {
		if angle >= 9 && angle < 27 {
			score = 2
		} else if angle >= 27 && angle < 45 {
			score = 36
		} else if angle >= 45 && angle < 63 {
			score = 8
		} else if angle >= 63 && angle < 81 {
			score = 26
		} else if angle >= 81 && angle < 99 {
			score = 12
		} else if angle >= 99 && angle < 117 {
			score = 20
		} else if angle >= 117 && angle < 135 {
			score = 30
		} else if angle >= 135 && angle < 153 {
			score = 4
		} else if angle >= 153 && angle < 171 {
			score = 34
		} else if angle >= 171 && angle < 189 {
			score = 6
		} else if angle >= 189 && angle < 207 {
			score = 38
		} else if angle >= 207 && angle < 225 {
			score = 14
		} else if angle >= 225 && angle < 243 {
			score = 32
		} else if angle >= 243 && angle < 261 {
			score = 16
		} else if angle >= 261 && angle < 279 {
			score = 22
		} else if angle >= 279 && angle < 297 {
			score = 28
		} else if angle >= 297 && angle < 315 {
			score = 18
		} else if angle >= 315 && angle < 333 {
			score = 24
		} else if angle >= 333 && angle < 351 {
			score = 10
		} else if angle >= 351 || angle < 9 {
			score = 40
		}
	}
	return score
}

// Funkcja napisana przez chatGPT, rysuje okręgi za rzutką, aby wizualnie pokazać z jaką prędkością będzie lecieć rzutka
func DrawTrailingCircles(dart *Dart, speed float32) {
	circleCount := 5
	distanceFactor := speed * 0.05

	endPos := dart.EndPosition()
	velocityDir := rl.Vector3Normalize(dart.Velocity)

	for i := 0; i < circleCount; i++ {
		offset := rl.Vector3Scale(velocityDir, -float32(i+1)*distanceFactor)
		circlePos := rl.Vector3Add(endPos, offset)
		scaleFactor := float32(i) / float32(circleCount) * 2
		circleRadius := 0.1 * scaleFactor
		rl.DrawSphere(circlePos, circleRadius, rl.LightGray)
	}
}

// Funkcja do śledzenia pozycji początku rzutki (żeby móc wykryć kolizję oraz wyliczyć punkty ze strzału)
func (d *Dart) TipPosition() rl.Vector3 {
	tipPos := rl.NewVector3(0.0, 0.0, -1.0548)
	return rl.Vector3Transform(tipPos, d.Transform)
}

// Funkcja do śledzenia pozycji końca rzutki (żeby móc wyświetlać trajektorię)
func (d *Dart) EndPosition() rl.Vector3 {
	endPos := rl.NewVector3(0.0, 0.0, 1.4177)
	return rl.Vector3Transform(endPos, d.Transform)
}

// Funkcja do zmiany pozycji rzutki przy zmianie parametru initialAngle
func (d *Dart) SettingPosition(initialSpeed, initialAngle float32) {

	d.Position = rl.NewVector3(0.0, 0.0, 0.0)
	d.Velocity.X = initialSpeed * float32(math.Sin(float64(initialAngle*math.Pi/180)))
	d.Velocity.Y = initialSpeed * float32(math.Cos(float64(initialAngle*math.Pi/180)))
	d.Velocity.Z = -initialSpeed

	pitch := -float32(math.Atan2(float64(d.Velocity.Y), float64(-d.Velocity.Z)))
	yaw := float32(math.Atan2(float64(d.Velocity.X), float64(-d.Velocity.Z)))

	rotationPitch := rl.MatrixRotateX(pitch)
	rotationYaw := rl.MatrixRotateY(yaw)

	translation := rl.MatrixTranslate(d.Position.X, d.Position.Y, d.Position.Z)
	d.Transform = rl.MatrixMultiply(rotationYaw, rotationPitch)
	d.Transform = rl.MatrixMultiply(d.Transform, translation)
}

func MoveTarget(position rl.Vector3, velocity float32, minX, maxX float32) (rl.Vector3, float32) {
	position.X += velocity

	if position.X >= maxX || position.X <= minX {
		velocity *= -1
	}
	return position, velocity
}

// Funkcja wyświetlająca ekran do wyboru trudności poziomu
func DifficultyChoosing(mode int, targetVelocity float32) (int, float32) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("DIFFICULTY LEVEL", 710, 350, 50, rl.LightGray)
		mousePosition := rl.GetMousePosition()

		rl.DrawRectangle(600, 450, 700, 100, rl.LightGray)
		rl.DrawRectangle(600, 600, 700, 100, rl.LightGray)
		rl.DrawRectangle(600, 750, 700, 100, rl.LightGray)

		if mousePosition.X >= 600 && mousePosition.X <= 1300 {
			if mousePosition.Y >= 450 && mousePosition.Y <= 550 {
				rl.DrawRectangle(600, 450, 700, 100, rl.NewColor(126, 185, 140, 255))
			} else if mousePosition.Y >= 600 && mousePosition.Y <= 700 {
				rl.DrawRectangle(600, 600, 700, 100, rl.NewColor(126, 185, 140, 255))
			} else if mousePosition.Y >= 750 && mousePosition.Y <= 850 {
				rl.DrawRectangle(600, 750, 700, 100, rl.NewColor(126, 185, 140, 255))
			}
		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if mousePosition.X >= 600 && mousePosition.X <= 1300 {
				if mousePosition.Y >= 450 && mousePosition.Y <= 550 {
					mode = 1
					targetVelocity = 0 // Easy: tarcza nieruchoma
				} else if mousePosition.Y >= 600 && mousePosition.Y <= 700 {
					mode = 2
					targetVelocity = 0 // Medium: zaczyna od nieruchomości
				} else if mousePosition.Y >= 750 && mousePosition.Y <= 850 {
					mode = 3
					targetVelocity = 0.25 // Hard: tarcza porusza się szybko
				}
			}
		}

		rl.DrawText("EASY", 870, 480, 50, rl.Black)
		rl.DrawText("MEDIUM", 845, 630, 50, rl.Black)
		rl.DrawText("HARD", 870, 780, 50, rl.Black)

		if rl.IsKeyPressed(rl.KeyQ) || mode != 0 {
			break
		}

		rl.EndDrawing()
	}
	return mode, targetVelocity
}

// Funckja wyświetlająca ekran z wynikami i pytaniem o chęć ponownej gry
func ShowingScore(scoreList []int, total_score int, play bool, showScore bool, bestScoreEasy int, bestScoreMedium int, bestScoreHard int, mode int) bool {
	for showScore {

		mousePosition := rl.GetMousePosition()

		rl.BeginDrawing()
		rl.DrawRectangleLines(550, 450, 800, 200, rl.Black)
		rl.DrawLine(750, 450, 750, 650, rl.Black)
		rl.DrawLine(870, 450, 870, 650, rl.Black)
		rl.DrawLine(990, 450, 990, 650, rl.Black)
		rl.DrawLine(1110, 450, 1110, 650, rl.Black)
		rl.DrawLine(1230, 450, 1230, 650, rl.Black)

		rl.DrawLine(550, 550, 1350, 550, rl.Black)

		rl.DrawText("ROUND", 600, 490, 30, rl.Black)
		rl.DrawText("SCORE", 600, 590, 30, rl.Black)

		for i := 0; i < 5; i++ {
			rl.DrawText(fmt.Sprintf("%d", scoreList[i]), int32(685+120*(i+1)), 590, 30, rl.Black)
		}

		for i := 1; i <= 5; i++ {
			rl.DrawText(fmt.Sprintf("%d", i), int32(685+120*i), 490, 30, rl.Black)
		}

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangle(630, 800, 250, 60, rl.LightGray)
		rl.DrawRectangle(1030, 800, 250, 60, rl.LightGray)

		if mousePosition.Y >= 800 && mousePosition.Y <= 860 {
			if mousePosition.X >= 630 && mousePosition.X <= 880 {
				rl.DrawRectangle(630, 800, 250, 60, rl.NewColor(126, 185, 140, 255))
			} else if mousePosition.X >= 1030 && mousePosition.X <= 1280 {
				rl.DrawRectangle(1030, 800, 250, 60, rl.NewColor(126, 185, 140, 255))
			}
		}

		if mode == 1 {
			rl.DrawText("EASY MODE", 865, 280, 30, rl.LightGray)
			rl.DrawText(fmt.Sprintf("BEST SCORE: %d", bestScoreEasy), 820, 760, 30, rl.LightGray)
		} else if mode == 2 {
			rl.DrawText("MEDIUM MODE", 850, 280, 30, rl.LightGray)
			rl.DrawText(fmt.Sprintf("BEST SCORE: %d", bestScoreMedium), 820, 760, 30, rl.LightGray)
		} else if mode == 3 {
			rl.DrawText("HARD MODE", 865, 280, 30, rl.LightGray)
			rl.DrawText(fmt.Sprintf("BEST SCORE: %d", bestScoreHard), 820, 760, 30, rl.LightGray)
		}

		rl.DrawText("GAME OVER!", 800, 350, 50, rl.Black)
		//rl.DrawText("YOUR SCORE:", 860, 380, 30, rl.Black)

		rl.DrawText(fmt.Sprintf("TOTAL SCORE: %d", total_score), 720, 700, 50, rl.Black)

		rl.DrawText("PLAY AGAIN", 660, 815, 30, rl.Black)
		rl.DrawText("QUIT", 1115, 815, 30, rl.Black)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if mousePosition.Y >= 800 && mousePosition.Y <= 860 {
				if mousePosition.X >= 630 && mousePosition.X <= 880 {
					play = true
					break
				} else if mousePosition.X >= 1030 && mousePosition.X <= 1280 {
					play = false
					break
				}
			}
		}
		rl.EndDrawing()
	}
	return play
}

// Funckcja głównej rozgrywki
func PlayingRound(mode int, targetVelocity float32, total_score int, scoreList []int, play bool) (bool, int, []int) {
	targetPosition := rl.NewVector3(0.0, 5.0, -15.0)
	var trajectoryPoints []rl.Vector3
	score := 0
	round := 1
	initialSpeed := float32(10.0)
	initialAngle := float32(0.0)

	game := true

	camera := rl.Camera{}
	camera.Position = rl.NewVector3(0.0, 2.0, 15.0)
	camera.Target = rl.NewVector3(0.0, 2.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	rzutka := rl.LoadModel("dart.obj")
	defer rl.UnloadModel(rzutka)

	tarcza := rl.LoadModel("dartboard.obj")
	defer rl.UnloadModel(tarcza)

	dart := NewDart()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		targetRadius := float32(2.64)
		numberOfRounds := 5

		if !game {
			break
		}

		End := dart.EndPosition()
		Tip := dart.TipPosition()

		if !dart.Thrown {
			dart.SettingPosition(initialSpeed, initialAngle)
		}

		if rl.IsKeyPressed(rl.KeySpace) && !dart.Thrown {
			dart.Throw(initialSpeed, initialAngle)
		}

		if dart.Thrown {
			if len(trajectoryPoints) == 0 || rl.Vector3Distance(trajectoryPoints[len(trajectoryPoints)-1], End) > 0.25 {
				trajectoryPoints = append(trajectoryPoints, End)
			}
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			trajectoryPoints = []rl.Vector3{}
			dart = NewDart()
		}

		targetPosition, targetVelocity = MoveTarget(targetPosition, targetVelocity, -10.0, 10.0)

		if dart.TipPosition().Z <= targetPosition.Z {
			score = Score(Tip, targetPosition)
			total_score += score
			scoreList = append(scoreList, score)

			dart.Thrown = false
			rl.WaitTime(2)

			if mode == 2 {
				if targetVelocity >= 0 {
					targetVelocity += 0.025
				} else {
					targetVelocity -= 0.025
				}
			}

			trajectoryPoints = []rl.Vector3{}
			dart = NewDart()
			round += 1

			if round > numberOfRounds {
				game = false
			}
		}

		initialSpeed = DrawSliderSpeed(rl.NewVector2(250, float32(140)), 200, 5.0, 20.0, initialSpeed)
		initialAngle = DrawSliderAngle(rl.NewVector2(250, float32(240)), 200, -90.0, 90.0, initialAngle)

		dart.Update(1.0/60.0, targetPosition, targetRadius)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(fmt.Sprintf("SPEED: %.2f", initialSpeed), 50, int32(130), 30, rl.LightGray)
		rl.DrawText(fmt.Sprintf("ANGLE: %.2f", initialAngle), 50, int32(230), 30, rl.LightGray)

		rl.DrawRectangleLines(1100, 50, 800, 200, rl.Black)
		rl.DrawRectangle(1100, 50, 800, 200, rl.LightGray)
		rl.DrawLine(1300, 50, 1300, 250, rl.White)
		rl.DrawLine(1420, 50, 1420, 250, rl.White)
		rl.DrawLine(1540, 50, 1540, 250, rl.White)
		rl.DrawLine(1660, 50, 1660, 250, rl.White)
		rl.DrawLine(1780, 50, 1780, 250, rl.White)
		rl.DrawLine(1100, 150, 1900, 150, rl.White)
		rl.DrawText("ROUND", 1150, 90, 30, rl.White)
		rl.DrawText("SCORE", 1150, 190, 30, rl.White)

		for i := 0; i < numberOfRounds; i++ {
			if i < len(scoreList) {
				rl.DrawText(fmt.Sprintf("%d", scoreList[i]), int32(1235+120*(i+1)), 190, 30, rl.White)
			} else {
				rl.DrawText("-", int32(1235+120*(i+1)), 190, 30, rl.White)
			}
		}

		for i := 1; i <= numberOfRounds; i++ {
			rl.DrawText(fmt.Sprintf("%d", i), int32(1235+120*i), 90, 30, rl.White)
		}

		rl.DrawText(fmt.Sprintf("TOTAL SCORE: %d", total_score), 1400, 300, 30, rl.LightGray)

		mousePosition := rl.GetMousePosition()

		rl.DrawRectangle(1650, 900, 150, 50, rl.LightGray)

		if mousePosition.X >= 1650 && mousePosition.X <= 1800 {
			if mousePosition.Y >= 900 && mousePosition.Y <= 950 {
				rl.DrawRectangle(1650, 900, 150, 50, rl.NewColor(126, 185, 140, 255))
			}
		}

		rl.DrawText("QUIT", 1690, 910, 30, rl.White)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if mousePosition.X >= 1650 && mousePosition.X <= 1800 {
				if mousePosition.Y >= 900 && mousePosition.Y <= 950 {
					game = false
					play = false
				}
			}
		}

		rl.BeginMode3D(camera)

		rzutka.Transform = dart.Transform

		rl.DrawModel(rzutka, rl.NewVector3(0.0, 0.0, 0.0), 1.0, rl.White)

		//poprawiane przez chatGPT
		for _, point := range trajectoryPoints {
			rl.DrawSphere(point, 0.05, rl.Red)
		}

		if !dart.Thrown {
			DrawTrailingCircles(dart, initialSpeed)
		}

		rl.DrawModel(tarcza, targetPosition, 1.0, rl.White)
		rl.DrawGrid(10, 3.0)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	return play, total_score, scoreList
}

// Funkcja wyznaczająca najlepszy wynik
func settingBestScore(mode int, total_score int, bestScoreEasy int, bestScoreMedium int, bestScoreHard int) (int, int, int) {
	if mode == 1 && total_score > bestScoreEasy {
		bestScoreEasy = total_score
	} else if mode == 2 && total_score > bestScoreMedium {
		bestScoreMedium = total_score
	} else if mode == 3 && total_score > bestScoreHard {
		bestScoreHard = total_score
	}
	return bestScoreEasy, bestScoreMedium, bestScoreHard
}


func main() {
	screenWidth := int32(1920)
	screenHeight := int32(1080)

	play := true
	bestScoreEasy := 0
	bestScoreMedium := 0
	bestScoreHard := 0

	for play {

		targetVelocity := float32(0)
		mode := 0
		var scoreList []int
		total_score := 0
		showScore := true

		rl.InitWindow(screenWidth, screenHeight, "Gra w rzutki")

		mode, targetVelocity = DifficultyChoosing(mode, targetVelocity)
		play, total_score, scoreList = PlayingRound(mode, targetVelocity, total_score, scoreList, play)
		bestScoreEasy, bestScoreMedium, bestScoreHard = settingBestScore(mode, total_score, bestScoreEasy, bestScoreMedium, bestScoreHard)
		play = ShowingScore(scoreList, total_score, play, showScore, bestScoreEasy, bestScoreMedium, bestScoreHard, mode)
	}
	rl.CloseWindow()
}
