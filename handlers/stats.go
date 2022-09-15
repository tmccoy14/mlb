package handlers

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/labstack/echo/v4"
	"net/url"
)

type Stat struct {
	Name    string `json:"name" validate:"required"`
	Season  string `json:"season" validate:"required"`
}

type StatResults struct {
	SportHittingTm struct {
		CopyRight    string `json:"copyRight"`
		QueryResults struct {
			TotalSize string `json:"totalSize"`
			Created   string `json:"created"`
			Row       struct {
				SportID     string `json:"sport_id"`
				LeagueShort string `json:"league_short"`
				Hr          string `json:"hr"`
				TeamID      string `json:"team_id"`
				Season      string `json:"season"`
				Ab          string `json:"ab"`
				Hldr        string `json:"hldr"`
				League      string `json:"league"`
				SportCode   string `json:"sport_code"`
				Ao          string `json:"ao"`
				Slg         string `json:"slg"`
				TeamFull    string `json:"team_full"`
				Ops         string `json:"ops"`
				TeamAbbrev  string `json:"team_abbrev"`
				Hbp         string `json:"hbp"`
				Rbi         string `json:"rbi"`
				GoAo        string `json:"go_ao"`
				Hfly        string `json:"hfly"`
				Lob         string `json:"lob"`
				Xbh         string `json:"xbh"`
				EndDate     string `json:"end_date"`
				Bb          string `json:"bb"`
				Np          string `json:"np"`
				Hgnd        string `json:"hgnd"`
				Roe         string `json:"roe"`
				Sb          string `json:"sb"`
				PlayerID    string `json:"player_id"`
				Avg         string `json:"avg"`
				Sf          string `json:"sf"`
				Sac         string `json:"sac"`
				Wo          string `json:"wo"`
				TeamShort   string `json:"team_short"`
				Hpop        string `json:"hpop"`
				So          string `json:"so"`
				GidpOpp     string `json:"gidp_opp"`
				Gidp        string `json:"gidp"`
				Ppa         string `json:"ppa"`
				D           string `json:"d"`
				Tpa         string `json:"tpa"`
				LeagueFull  string `json:"league_full"`
				G           string `json:"g"`
				H           string `json:"h"`
				Ibb         string `json:"ibb"`
				Go          string `json:"go"`
				TeamSeq     string `json:"team_seq"`
				Tb          string `json:"tb"`
				Cs          string `json:"cs"`
				R           string `json:"r"`
				T           string `json:"t"`
				Babip       string `json:"babip"`
				Obp         string `json:"obp"`
				Sport       string `json:"sport"`
				LeagueID    string `json:"league_id"`
			} `json:"row"`
		} `json:"queryResults"`
	} `json:"sport_hitting_tm"`
}

func Stats(c echo.Context) (err error) {

	// Get the player name from the parameters
	season := c.QueryParam("season")
	playerName := c.QueryParam("name")

	// Validate required parameters
	s := &Stat{
		Name: playerName,
		Season: season,
	}
	if err = c.Bind(s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(s); err != nil {
		return err
	}

	// Format the URL with the player name provided
	base, err := url.Parse("http://lookup-service-prod.mlb.com/json/named.search_player_all.bam")
	if err != nil {
		return
	}

	// Query params
	params := url.Values{}
	params.Add("sport_code", "'mlb'")
	params.Add("active_sw", "'Y'")
	params.Add("name_part", "'" + playerName + "'")
	base.RawQuery = params.Encode()

	// Get the player information from api
	resp, err := http.Get(base.String())
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// Response body is []byte
	body, err := ioutil.ReadAll(resp.Body)

	// Parse []byte to the go struct pointer
    var playerResult PlayerResults
    if err := json.Unmarshal(body, &playerResult); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	// Get the player id from player results
	playerId := playerResult.SearchPlayerAll.QueryResults.Row.PlayerID

	// Format the URL with the season year and player id
	base, err = url.Parse("http://lookup-service-prod.mlb.com/json/named.sport_hitting_tm.bam")

	// Query params
	params = url.Values{}
	params.Add("league_list_id", "'mlb'")
	params.Add("game_type", "'R'")
	params.Add("season", "'" + season + "'")
	params.Add("player_id", "'" + playerId + "'")
	base.RawQuery = params.Encode()

	// Get the team the player played for in a specific year from api
	resp, err = http.Get(base.String())
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// Response body is []byte
	body, err = ioutil.ReadAll(resp.Body)

	// Parse []byte to the go struct pointer
    var statResult StatResults
    if err := json.Unmarshal(body, &statResult); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	// Return JSON formatted player statistics information results
	return c.JSON(http.StatusOK, statResult)
}
