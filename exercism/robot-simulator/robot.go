package robotsimulator

import (
	"fmt"
)


const N = Dir(0)
const E = Dir(1)
const S = Dir(2)
const W = Dir(3)

func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case S:
		return "South"
	case E:
		return "East"
	default:
		return "West"
	}
}


func Shift(p Pos, d Dir) Pos {
    shiftedDir := p

    switch d {
    case N:
		shiftedDir.Northing++
	case E:
		shiftedDir.Easting++
	case W:
		shiftedDir.Easting--
	case S:
		shiftedDir.Northing--
    }
	return shiftedDir
}

func Right() {
	Step1Robot.Dir.turnRight()
}

func (d *Dir) turnRight() {
	*d = (*d + 1) % 4
}

func Left() {
	Step1Robot.Dir.turnLeft()
}

func (d *Dir) turnLeft() {
	*d--
	if *d < 0 {
		*d = W
	}
}

func Advance() {
	to := Shift(Pos{
		Easting:  RU(Step1Robot.X),
		Northing: RU(Step1Robot.Y),
	}, Step1Robot.Dir)
	Step1Robot.Y = int(to.Northing)
	Step1Robot.X = int(to.Easting)
}


type Action Command

func StartRobot(command chan Command, action chan Action) {
	for i := range command {
		action <- Action(i)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	robot.Easting = extent.Min.Easting
    robot.Northing = extent.Min.Northing

    for i := range action {
		switch i {
		case 'A':
			newPos := Shift(robot.Pos, robot.Dir)
			if isValidForRoom(extent, newPos) {
				robot.Pos = newPos
			}
		case 'R':
			robot.Dir.turnRight()
		case 'L':
			robot.Dir.turnLeft()
		}
	}
	report <- robot
}

func isValidForRoom(r Rect, p Pos) bool {
	return p.Easting <= r.Max.Easting &&
		p.Easting >= r.Min.Easting &&
		p.Northing <= r.Max.Northing &&
		p.Northing >= r.Min.Northing
}

type Action3 struct {
	name string
	cmd  Command
}

var validCommand = map[int32]bool{
	'R': true,
	'L': true,
	'A': true,
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, i := range script {
		if _, isValid := validCommand[i]; !isValid {
			log <- "bad command"
			break
		}
		action <- Action3{
			name: name,
			cmd:  Command(i),
		}
	}
	action <- Action3{
		name: name,
		cmd:  'Q', 
	}
}

func isPositionAvailable(robots []Step3Robot, p Pos) bool {
	for _, robot := range robots {
		if robot.Pos == p {
			return false
		}
	}

	return true
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	running := 0
	theRobots := map[string]int{}
	var tmpRobots []Step3Robot
	for idx, j := range robots {
		running++
		if j.Name == "" {
			log <- "no name"
			continue
		}
		if _, exists := theRobots[j.Name]; exists {
			log <- "already exists"
			continue
		}
		if isValidForRoom(extent, j.Pos) && isPositionAvailable(tmpRobots, j.Pos) {
			tmpRobots = append(tmpRobots, j)
			theRobots[j.Name] = idx
		} else {
			log <- "bump"
		}
	}
	for a := range action {
		robotIdx := theRobots[a.name]
		switch a.cmd {
		case 'Q':
			running--
		case 'A':
			newPos := Shift(robots[robotIdx].Pos, robots[robotIdx].Dir)

			if isValidForRoom(extent, newPos) && isPositionAvailable(robots, newPos) {
				robots[robotIdx].Pos = newPos
			} else {
				log <- "invalid"
			}
		case 'R':
			robots[robotIdx].Dir.turnRight()
		case 'L':
			robots[robotIdx].Dir.turnLeft()
		}
		if running <= 0 {
			close(action)
		}
	}

	rep <- robots
}





package robot

import "fmt"

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}

