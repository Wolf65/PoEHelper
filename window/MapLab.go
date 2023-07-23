package window

import (
	"encoding/json"
	"fmt"
	"os"
	"poehelper/config"
	"poehelper/misc"
	"strconv"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/sirupsen/logrus"
)

var (
	labyrinth     = Labyrinth{}
	idCurrentRoom = "13"
)

func init() {
	misc.Log().Debug("init map lab")
	registrationSettingsPages(settingPage{"Lab map", "pegeLabMap", pageLabMap})
	maps()
	misc.Log().Debug("register map lab pages")
}

func LabMapWindow(isOpen bool) {
	if isOpen || config.TestLog == "You have entered Aspirants' Plaza." {
		imgui.SetNextWindowSize(config.App.LabMap.WindowSize)
		imgui.BeginV(config.App.LabMap.Title, &config.App.LabMap.IsOpen, config.App.LabMap.WindowFlags)

		imgui.BeginGroup()
		imgui.BeginChildStrV("map", imgui.NewVec2(imgui.ContentRegionAvail().X-220, 143), true, 0) //88

		plotMap()

		imgui.EndChild()
		imgui.BeginChildStrV("mapInfo2", imgui.NewVec2(imgui.ContentRegionAvail().X-220, 0), true, 0)
		imgui.Text(fmt.Sprintf("Section 1: %s", labyrinth.Section1Layout))
		imgui.Text(fmt.Sprintf("Section 2: %s", labyrinth.Section2Layout))
		imgui.Text(fmt.Sprintf("Section 3: %s", labyrinth.Section3Layout))
		imgui.Text(config.TestLog)
		imgui.EndChild()
		imgui.EndGroup()
		imgui.SameLine()
		imgui.BeginChildStrV("mapInfo", imgui.ContentRegionAvail(), true, 0)
		imgui.Text("Izaro phases")
		imgui.Separator()
		imgui.Text(labyrinth.Weapon)
		imgui.Separator()
		imgui.Text(labyrinth.Phase1)
		imgui.Text(labyrinth.Phase2)
		imgui.Spacing()
		imgui.Text(labyrinth.Trap1)
		imgui.Text(labyrinth.Trap2)
		imgui.Separator()

		imgui.EndChild()

		// imgui.BeginGroup()
		// imgui.ButtonV("Status", imgui.NewVec2(0, 0))
		// imgui.SameLine()
		// // imgui.ProgressBarV(0.3, imgui.NewVec2(0, 0), "textProgress")
		// imgui.Dummy(imgui.Vec2{X: imgui.ContentRegionAvail().X - config.App.Vars.BaseButton.X - config.App.Vars.ItemSpacing.X, Y: config.App.Vars.BaseButton.Y})
		// imgui.SameLine()
		// imgui.ButtonV("Settings", config.App.Vars.BaseButton)
		// imgui.EndGroup()

		imgui.End()
	}
}

func pageLabMap() {
	imgui.Text("Lab map")
}

func plotMap() {
	drawList := imgui.WindowDrawList()
	pos := imgui.CursorScreenPos()
	sizeRoom := float32(14)
	colorRoom := imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1})
	colorRoomSelect := imgui.ColorU32Vec4(imgui.Vec4{X: 0, Y: 1, Z: 0, W: 1})
	sizeExit := float32(4)
	colorExit := imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1})
	colorLine := imgui.ColorU32Vec4(imgui.Vec4{X: 0.3, Y: 0.3, Z: 1, W: 1})

	// selectName := "sanitorium annex"

	for _, room := range labyrinth.Rooms {
		// 38(2) x 793 (446)
		// 40(4) y 216 (107)
		roomPos := stringPosToFloat32(room.X, room.Y)
		drawList.AddCircleFilled(imgui.Vec2{X: pos.X + sizeRoom + roomPos.X, Y: pos.Y + sizeRoom + roomPos.Y}, sizeRoom, colorRoom)

		//fixme >1 room name
		// if room.Name == selectName {
		if room.ID == idCurrentRoom {
			// fmt.Println(roomPos)
			drawList.AddCircleV(imgui.Vec2{X: pos.X + sizeRoom + roomPos.X, Y: pos.Y + sizeRoom + roomPos.Y}, sizeRoom, colorRoomSelect, 0, 3)
		}
	}

	for _, room := range labyrinth.Rooms {
		roomPos := stringPosToFloat32(room.X, room.Y)
		if room.Exits.Nw != "" {
			posExit := imgui.Vec2{X: pos.X + sizeRoom + roomPos.X - sizeRoom/1.4, Y: pos.Y + sizeRoom + roomPos.Y - sizeRoom/1.4}
			drawList.AddLineV(posExit, nextRoomPos(pos, sizeRoom, room.Exits.Nw), colorLine, 2)
			drawList.AddCircleFilled(posExit, sizeExit, colorExit)
		}
		if room.Exits.N != "" {
			posExit := imgui.Vec2{X: pos.X + sizeRoom + roomPos.X, Y: pos.Y + sizeRoom + roomPos.Y - sizeRoom}
			drawList.AddLineV(posExit, nextRoomPos(pos, sizeRoom, room.Exits.N), colorLine, 2)
			drawList.AddCircleFilled(posExit, sizeExit, colorExit)
		}
		if room.Exits.Ne != "" {
			posExit := imgui.Vec2{X: pos.X + sizeRoom + roomPos.X + sizeRoom/1.4, Y: pos.Y + sizeRoom + roomPos.Y - sizeRoom/1.4}
			drawList.AddLineV(posExit, nextRoomPos(pos, sizeRoom, room.Exits.Ne), colorLine, 2)
			drawList.AddCircleFilled(posExit, sizeExit, colorExit)
		}
		if room.Exits.C != "" {
			posExit := imgui.Vec2{X: pos.X + sizeRoom + roomPos.X, Y: pos.Y + sizeRoom + roomPos.Y}
			drawList.AddLineV(posExit, nextRoomPos(pos, sizeRoom, room.Exits.C), colorLine, 2)
			drawList.AddCircleFilled(posExit, sizeExit, colorExit)
		}
	}
}

func stringPosToFloat32(x, y string) imgui.Vec2 {
	mapScale := float32(1.8)
	mapPosOffset := imgui.Vec2{X: 38, Y: 40}
	resX, _ := strconv.ParseFloat(x, 32)
	resY, _ := strconv.ParseFloat(y, 32)

	return imgui.Vec2{X: (float32(resX) - mapPosOffset.X) / mapScale, Y: (float32(resY) - mapPosOffset.Y) / mapScale}
}

func nextRoomPos(cursorPos imgui.Vec2, sizeRoom float32, idRoom string) imgui.Vec2 {
	i, _ := strconv.Atoi(idRoom)
	nextRoom := labyrinth.Rooms[i-1]
	nextRoomPos := stringPosToFloat32(nextRoom.X, nextRoom.Y)
	return imgui.Vec2{X: cursorPos.X + sizeRoom + nextRoomPos.X, Y: cursorPos.Y + sizeRoom + nextRoomPos.Y}
}

type Labyrinth struct {
	Difficulty     string `json:"difficulty"`
	Date           string `json:"date"`
	Weapon         string `json:"weapon"`
	Phase1         string `json:"phase1"`
	Phase2         string `json:"phase2"`
	Trap1          string `json:"trap1"`
	Trap2          string `json:"trap2"`
	Section1Layout string `json:"section1layout"`
	Section2Layout string `json:"section2layout"`
	Section3Layout string `json:"section3layout"`
	Rooms          []struct {
		Name              string   `json:"name"`
		Areacode          string   `json:"areacode"`
		ID                string   `json:"id"`
		X                 string   `json:"x"`
		Y                 string   `json:"y"`
		Dangerous         string   `json:"dangerous"`
		Contens           []string `json:"contents"`
		ContentDirections []string `json:"content_directions"`
		Exits             struct {
			Nw string `json:"NW"`
			N  string `json:"N"`
			Ne string `json:"NE"`
			C  string `json:"C"`
		} `json:"exits,omitempty"`
	} `json:"rooms"`
}

func maps() {
	files, err := os.ReadFile("uber-2023-06-24.json")
	if err != nil {
		misc.Log().WithFields(logrus.Fields{
			"err": err,
		}).Error("Read json lab file")
	}
	err = json.Unmarshal(files, &labyrinth)
	if err != nil {
		misc.Log().WithFields(logrus.Fields{
			"err": err,
		}).Error("Unmarshal json lab file")
	}
}

func LabCompassWindow(isOpen bool) {
	if isOpen {
		imgui.SetNextWindowSize(config.App.LabCompass.WindowSize)
		imgui.BeginV(config.App.LabCompass.Title, &config.App.LabCompass.IsOpen, config.App.LabCompass.WindowFlags)

		drawList := imgui.WindowDrawList()
		pos := imgui.CursorScreenPos()
		i, _ := strconv.ParseInt(idCurrentRoom, 10, 0)
		room := labyrinth.Rooms[i-1]
		{
			drawList.AddCircle(imgui.Vec2{X: pos.X + 75, Y: pos.Y + 60}, 50, imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1}))
			// imgui.Text(fmt.Sprintf("%v", imgui.ContentRegionAvail()))

			if room.Exits.Nw != "" {
				posDotNW := imgui.Vec2{X: pos.X + 75 + (50 * -0.7), Y: pos.Y + 60 + (50 * -0.7)}
				drawList.AddCircleFilled(posDotNW, 8, imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1}))
			}
			if room.Exits.N != "" {
				posDotN := imgui.Vec2{X: pos.X + 75 + (50 * 0), Y: pos.Y + 60 + (50 * -1)}
				drawList.AddCircleFilled(posDotN, 8, imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1}))
			}
			if room.Exits.Ne != "" {
				posDotNE := imgui.Vec2{X: pos.X + 75 + (50 * 0.7), Y: pos.Y + 60 + (50 * -0.7)}
				drawList.AddCircleFilled(posDotNE, 8, imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1}))
			}
			if room.Exits.C != "" {
				drawList.AddCircleFilled(imgui.Vec2{X: pos.X + 75, Y: pos.Y + 60}, 8, imgui.ColorU32Vec4(imgui.Vec4{X: 1, Y: 0, Z: 0, W: 1}))
			}
		}
		imgui.Text("Loot:")
		for _, c := range room.Contens {
			imgui.Text(c)
		}

		imgui.End()
	}
}
