package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type Team struct {
	Name   string `json:"name" validate:"required"`
	Season string `json:"season" validate:"required"`
}

type TeamResults struct {
	PlayerTeams struct {
		CopyRight    string `json:"copyRight"`
		QueryResults struct {
			TotalSize string `json:"totalSize"`
			Created   string `json:"created"`
			Row       struct {
				EndDate         string `json:"end_date"`
				SportID         string `json:"sport_id"`
				StatusDate      string `json:"status_date"`
				StatusCode      string `json:"status_code"`
				LeagueShort     string `json:"league_short"`
				ClassID         string `json:"class_id"`
				TeamID          string `json:"team_id"`
				SportFull       string `json:"sport_full"`
				TeamBrief       string `json:"team_brief"`
				PlayerID        string `json:"player_id"`
				OrgShort        string `json:"org_short"`
				TeamShort       string `json:"team_short"`
				PrimaryPosition string `json:"primary_position"`
				JerseyNumber    string `json:"jersey_number"`
				PitchingSeason  string `json:"pitching_season"`
				OrgFull         string `json:"org_full"`
				Class           string `json:"class"`
				StartDate       string `json:"start_date"`
				HittingSeason   string `json:"hitting_season"`
				Org             string `json:"org"`
				SeasonState     string `json:"season_state"`
				LeagueFull      string `json:"league_full"`
				League          string `json:"league"`
				FortyManSw      string `json:"forty_man_sw"`
				LeagueSeason    string `json:"league_season"`
				SportCode       string `json:"sport_code"`
				Team            string `json:"team"`
				OrgAbbrev       string `json:"org_abbrev"`
				SportShort      string `json:"sport_short"`
				TeamAbbrev      string `json:"team_abbrev"`
				OrgID           string `json:"org_id"`
				FortyManSwOld   string `json:"forty_man_sw_old"`
				PrimaryStatType string `json:"primary_stat_type"`
				Sport           string `json:"sport"`
				CurrentSw       string `json:"current_sw"`
				FieldingSeason  string `json:"fielding_season"`
				LeagueID        string `json:"league_id"`
				Status          string `json:"status"`
			} `json:"row"`
		} `json:"queryResults"`
	} `json:"player_teams"`
}

func Teams(c echo.Context) (err error) {

	// Get the player name from the parameters
	season := c.QueryParam("season")
	playerName := c.QueryParam("name")

	// Validate required parameters
	t := &Team{
		Name:   playerName,
		Season: season,
	}
	if err = c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(t); err != nil {
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
	params.Add("name_part", "'"+playerName+"'")
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
	base, err = url.Parse("http://lookup-service-prod.mlb.com/json/named.player_teams.bam")

	// Query params
	params = url.Values{}
	params.Add("season", "'"+season+"'")
	params.Add("player_id", "'"+playerId+"'")
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
	var teamResult TeamResults
	if err := json.Unmarshal(body, &teamResult); err != nil {
		fmt.Println("Can not unmarshal JSON111")
	}

	// Return JSON formatted team information results
	return c.JSON(http.StatusOK, teamResult)
}
