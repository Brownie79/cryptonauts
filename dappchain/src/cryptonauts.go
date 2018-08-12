package main

import (
	"github.com/loomnetwork/go-loom/plugin"
	contract "github.com/loomnetwork/go-loom/plugin/contractpb"
)

func main() {
	plugin.Serve(Contract)
}

type Cryptonauts struct{}
type Building struct {
	name      string
	level     int
	resType   string // "Food" "Metal" "Coin" "Research" "Essence"
	resValue  int
	hitpoints int
}
type Turret struct {
	hitpoints int
	damage    int
}
type Planet struct {
	name      string //will be the blockhash in which it was discovered
	coords    string // (0,0) is Earth
	buildings []Building
	turrets   []Turret
}
type Resources struct {
	resFood     int
	resMetal    int
	resCoin     int
	resResearch int
}
type GameState struct {
	planets   []Planet
	resources Resources
}
type Upgrade struct{}
type Hero struct {
	level     int
	hitpoints int
	defense   int
	speed     int
	dps       int
	cost      int
}

type Player struct {
	privatekey  string
	essence     int
	currentGame GameState
	pastGames   []GameState
	inventory   []Upgrade
	heroes      []Hero
}

func (e *Cryptonauts) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "Cryptonauts",
		Version: "0.0.1",
	}, nil
}

func (e *Cryptonauts) Init(ctx contract.Context, req *plugin.Request) error {
	return nil
}

var Contract plugin.Contract = contract.MakePluginContract(&Cryptonauts{})
