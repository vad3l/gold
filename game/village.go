package game

import (
	"math/rand"
	"time"
	"log"

	"strings"
	"strconv"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Tower = iota
	House
	Tavern
	Church
)

type Village struct {
	TabBuild	[][2]int 
	TabPositionBuild	[][4]float64
	TabImg	[][]*ebiten.Image
}

func (v *Village) SaveVillage () string {
	s := ""
	for i := 0; i < len(v.TabBuild); i++ {
		s += strconv.Itoa(v.TabBuild[i][0])
		s += ";"
		s += strconv.Itoa(v.TabBuild[i][1])
		s += ";"
	}
	return s
}

func LoadVillage (vdata string) *Village {
	v := NewVillage()
	builds := strings.Split(vdata, ";")
	for i := 0; i < (len(builds) - 1); i += 2 {
		x, err := strconv.Atoi(builds[i])
		CheckError(err)
		y, err := strconv.Atoi(builds[i+1])
		CheckError(err)
		v.TabBuild = append(v.TabBuild, [2]int{ x, y })
	}
	return v
}

func NewVillage () *Village {
	var TabPositionBuild  [][4]float64

	TabPositionBuild = append(TabPositionBuild,[4]float64{1100, 350, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{1250,400, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{900,200, 0.25,0.25})
	TabPositionBuild = append(TabPositionBuild,[4]float64{50,300, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{-50,400, 0.50,0.50})
	TabPositionBuild = append(TabPositionBuild,[4]float64{300,200, 0.25,0.25})
	shuffle(TabPositionBuild) 
	
	imgTower0, _, err := ebitenutil.NewImageFromFile("data/image/building/tower0.png")
	imgTower1, _, erre := ebitenutil.NewImageFromFile("data/image/building/tower1.png")
	imgTower2, _, errer := ebitenutil.NewImageFromFile("data/image/building/tower2.png")
	imgHouse1, _, errere := ebitenutil.NewImageFromFile("data/image/building/house1.png")
	imgTavern1, _, errerer := ebitenutil.NewImageFromFile("data/image/building/tavern1.png")
	if err != nil || erre != nil || errer != nil || errere != nil || errerer != nil{
		log.Fatalf("Failed to village build load image: ")
	}

	TabImg := make([][]*ebiten.Image, 3)
	TabImg[0] = []*ebiten.Image{imgTower0, imgTower1, imgTower2}
	TabImg[1] = []*ebiten.Image{nil, imgHouse1, nil}
	TabImg[2] = []*ebiten.Image{nil, imgTavern1, nil}
	
	return &Village{
		[][2]int{},
		TabPositionBuild,
		TabImg,
	}
}

func shuffle(data [][4]float64) [][4]float64 {
	rand.Seed(time.Now().UnixNano())
	n := len(data)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func (villa *Village) ChooseBuildImg(build [2]int) *ebiten.Image {
	return villa.TabImg[build[0]][build[1]]
}
