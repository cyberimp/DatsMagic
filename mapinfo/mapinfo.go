package mapinfo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"math"
	"math/rand/v2"
	"net/http"
	"sort"
	"time"
)

const server = "https://games-test.datsteam.dev"

type Mapinfo struct {
	rounds   []Round
	nowRound int
	curMap   Map
	clients  []Client
	Token    string
}

type Round struct {
	Duration int    `json:"duration"`
	EndAt    string `json:"endAt"`
	Name     string `json:"name"`
	Repeat   int    `json:"repeat"`
	StartAt  string `json:"startAt"`
	Status   string `json:"status"`
}

type Schedule struct {
	GameName string  `json:"gameName"`
	Now      string  `json:"now"`
	Rounds   []Round `json:"rounds"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type CommandTransport struct {
	Acceleration   Vector      `json:"acceleration"`
	ActivateShield bool        `json:"activateShield"`
	Attack         Coordinates `json:"attack"`
	Id             string      `json:"id"`
}

type Anomaly struct {
	EffectiveRadius float64 `json:"effectiveRadius"`
	Id              string  `json:"id"`
	Radius          float64 `json:"radius"`
	Strength        float64 `json:"strength"`
	Velocity        Vector  `json:"velocity"`
	X               int     `json:"x"`
	Y               int     `json:"y"`
}

type Bounty struct {
	Points int `json:"points"`
	Radius int `json:"radius"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type Enemy struct {
	Health       int    `json:"health"`
	KillBounty   int    `json:"killBounty"`
	ShieldLeftMs int    `json:"shieldLeftMs"`
	Status       string `json:"status"`
	Velocity     Vector `json:"velocity"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type Transport struct {
	AnomalyAcceleration Vector `json:"anomalyAcceleration"`
	AttackCooldownMs    int    `json:"attackCooldownMs"`
	DeathCount          int    `json:"deathCount"`
	Health              int    `json:"health"`
	Id                  string `json:"id"`
	SelfAcceleration    Vector `json:"selfAcceleration"`
	ShieldCooldownMs    int    `json:"shieldCooldownMs"`
	ShieldLeftMs        int    `json:"shieldLeftMs"`
	Status              string `json:"status"`
	Velocity            Vector `json:"velocity"`
	X                   int    `json:"x"`
	Y                   int    `json:"y"`
}

type command struct {
	Transports []CommandTransport `json:"transports"`
}

type Map struct {
	Anomalies             []Anomaly   `json:"anomalies"`
	AttackCooldownMs      int         `json:"attackCooldownMs"`
	AttackDamage          int         `json:"attackDamage"`
	AttackExplosionRadius float64     `json:"attackExplosionRadius"`
	AttackRange           float64     `json:"attackRange"`
	Bounties              []Bounty    `json:"bounties"`
	Enemies               []Enemy     `json:"enemies"`
	MapSize               Coordinates `json:"mapSize"`
	MaxAccel              float64     `json:"maxAccel"`
	MaxSpeed              float64     `json:"maxSpeed"`
	Name                  string      `json:"name"`
	Points                int         `json:"points"`
	ReviveTimeoutSec      int         `json:"reviveTimeoutSec"`
	ShieldCooldownMs      int         `json:"shieldCooldownMs"`
	ShieldTimeMs          int         `json:"shieldTimeMs"`
	TransportRadius       int         `json:"transportRadius"`
	Transports            []Transport `json:"transports"`
	WantedList            []Enemy     `json:"wantedList"`
}

func (r *Mapinfo) Update() {
	const endpoint = "/play/magcarp/player/move"
	client := &http.Client{}
	var emptyTransport command
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(&emptyTransport)
	req, err := http.NewRequest("POST", server+endpoint, &body)
	if err != nil {
		panic(err)
	}
	var resp *http.Response
	req.Header.Add("X-Auth-Token", r.Token)
	resp, err = client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		_ = json.NewDecoder(resp.Body).Decode(&r.curMap)
		fmt.Printf("%+v\n", r.curMap)
		fmt.Printf("now:%v\n", r.nowRound)
	}
}

func (r *Mapinfo) RandomMoves() {
	const endpoint = "/play/magcarp/player/move"
	client := &http.Client{}
	var transport []CommandTransport

	if r.nowRound == -1 {
		r.curMap.Transports = []Transport{}
	}

	for _, carpet := range r.curMap.Transports {
		if carpet.Status == "alive" {
			minRange := math.MaxInt
			closest := -1
			for i, bounty := range r.curMap.Bounties {
				diff := (bounty.X-carpet.X)*(bounty.X-carpet.X) + (bounty.Y-carpet.Y)*(bounty.Y-carpet.Y)
				if diff < minRange {
					minRange = diff
					closest = i
				}
			}

			enemyRange := math.MaxInt
			nearEnemy := -1
			safeAttack := r.curMap.AttackExplosionRadius
			validAttack := r.curMap.AttackRange

			attack := Coordinates{0, 0}

			if carpet.AttackCooldownMs == 0 {
				for i, enemy := range r.curMap.WantedList {
					diff := (enemy.X-carpet.X)*(enemy.X-carpet.X) + (enemy.Y-carpet.Y)*(enemy.Y-carpet.Y)
					if diff < enemyRange && math.Sqrt(float64(diff)) > safeAttack && math.Sqrt(float64(diff)) < validAttack && enemy.Status == "alive" && enemy.ShieldLeftMs == 0 {
						nearEnemy = i
						attack.X = enemy.X
						attack.Y = enemy.Y
					}
				}
				if nearEnemy < 0 {
					for i, enemy := range r.curMap.Enemies {
						diff := (enemy.X-carpet.X)*(enemy.X-carpet.X) + (enemy.Y-carpet.Y)*(enemy.Y-carpet.Y)
						if diff < enemyRange && math.Sqrt(float64(diff)) > safeAttack && math.Sqrt(float64(diff)) < validAttack && enemy.Status == "alive" && enemy.ShieldLeftMs == 0 {
							nearEnemy = i
							attack.X = enemy.X
							attack.Y = enemy.Y
						}
					}
				}
			}

			acc := Vector{1, 1}
			if closest >= 0 {
				bounty := r.curMap.Bounties[closest]
				acc = Vector{float64(bounty.X - int(carpet.Velocity.X*3) - carpet.X), float64(bounty.Y - int(carpet.Velocity.Y*3) - carpet.Y)}
			}

			maxAcc := r.curMap.MaxAccel
			length := math.Sqrt(acc.X*acc.X + acc.Y*acc.Y)

			acc.X = acc.X * (maxAcc / length)
			acc.Y = acc.Y * (maxAcc / length)

			//dangers := r.closestDangers(&carpet)

			fmt.Println(acc)
			newTransport := CommandTransport{
				Id:           carpet.Id,
				Acceleration: acc,
			}

			if nearEnemy >= 0 {
				newTransport.ActivateShield = carpet.ShieldCooldownMs == 0
				newTransport.Attack = attack
				fmt.Println("Attack!")

			}

			transport = append(transport, newTransport)

		}
	}
	fmt.Printf("%+v", transport)

	marshalled, err := json.Marshal(command{transport})

	fmt.Printf("\n%+v\n", string(marshalled))

	req, err := http.NewRequest("POST", server+endpoint, bytes.NewReader(marshalled))
	if err != nil {
		panic(err)
	}
	var resp *http.Response
	req.Header.Add("X-Auth-Token", r.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	type BadRequest struct {
		Error     string `json:"error"`
		ErrorCode int    `json:"errCode"`
	}

	if resp.StatusCode == 400 {
		var err BadRequest
		defer resp.Body.Close()
		_ = json.NewDecoder(resp.Body).Decode(&err)
		fmt.Printf("\n%+v\n", err)
		fmt.Printf("now:%v\n", r.nowRound)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		_ = json.NewDecoder(resp.Body).Decode(&r.curMap)
		fmt.Printf("%+v\n", r.curMap)
		fmt.Printf("now:%v\n", r.nowRound)
	}

}

func (r *Mapinfo) closestDangers(carpet *Transport) []*Danger {
	// Initialize the closest dangers and distances
	closestDangers := make([]*Danger, 0)
	closestDistances := make([]float64, 0)

	distanceToBorder := float64(r.curMap.TransportRadius - carpet.X)
	closestDangers = append(closestDangers, &Danger{Type: "screen border", X: 0, Y: carpet.Y})
	closestDistances = append(closestDistances, distanceToBorder)

	distanceToBorder = float64(carpet.X + r.curMap.TransportRadius - r.curMap.MapSize.X)
	closestDangers = append(closestDangers, &Danger{Type: "screen border", X: r.curMap.MapSize.X, Y: carpet.Y})
	closestDistances = append(closestDistances, distanceToBorder)

	distanceToBorder = float64(r.curMap.TransportRadius - carpet.Y)
	closestDangers = append(closestDangers, &Danger{Type: "screen border", X: carpet.X, Y: 0})
	closestDistances = append(closestDistances, distanceToBorder)

	distanceToBorder = float64(carpet.Y + r.curMap.TransportRadius - r.curMap.MapSize.Y)
	closestDangers = append(closestDangers, &Danger{Type: "screen border", X: carpet.X, Y: r.curMap.MapSize.Y})
	closestDistances = append(closestDistances, distanceToBorder)

	// Check anomalies
	for _, anomaly := range r.curMap.Anomalies {
		distance := math.Sqrt(float64((anomaly.X-carpet.X)*(anomaly.X-carpet.X) + (anomaly.Y-carpet.Y)*(anomaly.Y-carpet.Y)))
		closestDangers = append(closestDangers, &Danger{Type: "anomaly", X: anomaly.X, Y: anomaly.Y})
		closestDistances = append(closestDistances, distance)
	}

	// Check enemies
	for _, enemy := range r.curMap.Enemies {
		distance := math.Sqrt(float64((enemy.X-carpet.X)*(enemy.X-carpet.X) + (enemy.Y-carpet.Y)*(enemy.Y-carpet.Y)))
		closestDangers = append(closestDangers, &Danger{Type: "enemy", X: enemy.X, Y: enemy.Y})
		closestDistances = append(closestDistances, distance)
	}
	for _, enemy := range r.curMap.WantedList {
		distance := math.Sqrt(float64((enemy.X-carpet.X)*(enemy.X-carpet.X) + (enemy.Y-carpet.Y)*(enemy.Y-carpet.Y)))
		closestDangers = append(closestDangers, &Danger{Type: "wanted enemy", X: enemy.X, Y: enemy.Y})
		closestDistances = append(closestDistances, distance)
	}

	// Sort dangers by distance
	sort.Slice(closestDangers, func(i, j int) bool {
		return closestDistances[i] < closestDistances[j]
	})

	// Return the 5 closest dangers
	if len(closestDangers) > 5 {
		closestDangers = closestDangers[:5]
		closestDistances = closestDistances[:5]
	}

	return closestDangers
}

type Danger struct {
	Type string
	X    int
	Y    int
}

func (r *Mapinfo) GetCactus(w http.ResponseWriter, req *http.Request) {
	type cactus struct {
		Img string `json:"img"`
		X   int    `json:"x"`
		Y   int    `json:"y"`
	}
	var data []cactus
	images := []string{"palm.png", "rocks.png", "cactus0.png", "cactus1.png"}
	c, err := websocket.Accept(w, req, nil)
	if err != nil {
		panic(err)
	}
	defer c.CloseNow()

	ctx, cancel := context.WithTimeout(req.Context(), time.Minute*10)
	defer cancel()

	ctx = c.CloseRead(ctx)

	for i := 0; i < 25; i++ {
		im := rand.IntN(len(images))
		data = append(data, cactus{images[im], rand.IntN(1920), rand.IntN(900)})
	}

	err = wsjson.Write(ctx, c, data)
	if err != nil {
		panic(err)
	}
}

func (r *Mapinfo) UpdateRounds() {
	const endpoint = "/rounds/magcarp/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", server+endpoint, nil)
	if err != nil {
		panic(err)
	}
	var resp *http.Response
	req.Header.Add("X-Auth-Token", r.Token)
	resp, err = client.Do(req)
	if err == nil {
		var target = new(Schedule)
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(target)
		if err != nil {
			panic(err)
		}

		r.rounds = target.Rounds
		r.nowRound = -1
		for i, round := range target.Rounds {
			if round.Status == "active" {
				r.nowRound = i
				break
			}
		}
	}
}

type Client struct {
	c   *websocket.Conn
	ctx context.Context
}

func (r *Mapinfo) GetMapHandle(w http.ResponseWriter, req *http.Request) {
	c, err := websocket.Accept(w, req, nil)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	r.clients = append(r.clients, Client{c, ctx})
}

func (r *Mapinfo) Loop() {
	defer func() {
		for _, client := range r.clients {
			fmt.Printf("closing client: %+v\n", client)
			client.c.CloseNow()
		}
	}()
	var closed []int
	for {
		closed = []int{}
		r.UpdateRounds()
		//r.Update()
		r.RandomMoves()
		for i, client := range r.clients {
			err := wsjson.Write(client.ctx, client.c, r.curMap)
			if err != nil {
				closed = append(closed, i)
			}
		}

		for num, i := range closed {
			r.clients = append(r.clients[:i-num], r.clients[i+1-num:]...)
		}

		time.Sleep(time.Second / 3)
	}
}
