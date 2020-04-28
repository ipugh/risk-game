package main

// board GET
type Gameboard struct {
    Board World `json:"board"`
    Turn string `json:"turn"`
    Gamestate string `json:"gamestate"`
    Turnstate string `json:"turnstate"`
    Players []Player `json:"players"`
}

type Player struct {
    Name string `json:"name"`
    Color string `json:"color"`
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

// attack POST received
type Attack struct {
    From int `json:"from"`
    To int `json:"to"`
}

// attack POST returned 
type AttackReturn struct {
    Success bool `json:"success"`
    Remaining int `json:"remaining"`
    Attacking [3]int `json:"attacking"`
    Defending [2]int `json:"defending"`
}

