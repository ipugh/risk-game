package main

// board GET
type Gameboard struct {
    Board World `json:"board"`
    Turn string `json:"turn"`
    Gamestate string `json:"gamestate"`
    Turnstate string `json:"turnstate"`
}

type World struct {
    Territories []Territory `json:"territories"`
}

type Territory struct {
    Player string `json:"player"`
    Armies int `json:"armies"`
    TerritoryId int `json:"territoryid"`
}

// reinforce POST
type Reinforce struct {
    TerritoryId int `json:"territoryid"`
    Armies int `json:"armies"`
}

// regroup POST
type Regroup struct {
    From int `json:"from"`
    To int `json:"to"`
    Armies int `json:"armies"`
}

// pick POST
type Pick struct {
    TerritoryId int `json:"territoryid"`
}

// pickreinforce POST
type PickReinforce struct {
    TerritoryId int `json:"territoryid"`
}
