package main

import (
	"io"

	"os"

	"github.com/joho/godotenv"

	"net/http"

	"encoding/json"

	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file" + err.Error())
	}
	key := os.Getenv("steamapikey")
	id := os.Getenv("steamid")
	var count string = "20" // not required, default is 20 represent how many games played in the last 2 weeks to return
	url := "http://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v1/?key=" + key + "&steamid=" + id + "&count=" + count
	resp, err := http.Get(url)
	if err != nil {
		panic("Error fetching Steam API" + err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error reading response body" + err.Error())
	}
	println("Response body: " + string(body))
	type SteamResponse struct {
		Response struct {
			TotalCount int `json:"total_count"`
			Games      []struct {
				AppID                  int    `json:"appid"`
				Name                   string `json:"name"`
				Playtime2Weeks         int    `json:"playtime_2weeks"`
				PlaytimeForever        int    `json:"playtime_forever"`
				ImgIconURL             string `json:"img_icon_url"`
				PlaytimeWindowsForever int    `json:"playtime_windows_forever"`
				PlaytimeMacForever     int    `json:"playtime_mac_forever"`
				PlaytimeLinuxForever   int    `json:"playtime_linux_forever"`
				PlaytimeDeckForever    int    `json:"playtime_deck_forever"`
			} `json:"games"`
		} `json:"response"`
	}
	var data SteamResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic("Error unmarshaling JSON" + err.Error())
	}
	totalPlaytime := 0
	for _, game := range data.Response.Games {
		totalPlaytime += game.Playtime2Weeks
		println("Game: " + game.Name + ", Playtime in last 2 weeks: " + strconv.Itoa(game.Playtime2Weeks) + " minutes")
	}
	println("Total playtime in last 2 weeks: " + strconv.Itoa(totalPlaytime) + " minutes")
}
