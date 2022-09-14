package handlers

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/labstack/echo/v4"
)

type Player struct {
	Name  string `json:"name" validate:"required"`
}

type PlayerResults struct {
	SearchPlayerAll struct {
		CopyRight    string `json:"copyRight"`
		QueryResults struct {
			Created   string `json:"created"`
			TotalSize string `json:"totalSize"`
			Row       struct {
				Position             string `json:"position"`
				BirthCountry         string `json:"birth_country"`
				Weight               string `json:"weight"`
				BirthState           string `json:"birth_state"`
				NameDisplayFirstLast string `json:"name_display_first_last"`
				College              string `json:"college"`
				HeightInches         string `json:"height_inches"`
				NameDisplayRoster    string `json:"name_display_roster"`
				SportCode            string `json:"sport_code"`
				Bats                 string `json:"bats"`
				NameFirst            string `json:"name_first"`
				TeamCode             string `json:"team_code"`
				BirthCity            string `json:"birth_city"`
				HeightFeet           string `json:"height_feet"`
				ProDebutDate         string `json:"pro_debut_date"`
				TeamFull             string `json:"team_full"`
				TeamAbbrev           string `json:"team_abbrev"`
				BirthDate            string `json:"birth_date"`
				Throws               string `json:"throws"`
				League               string `json:"league"`
				NameDisplayLastFirst string `json:"name_display_last_first"`
				PositionID           string `json:"position_id"`
				HighSchool           string `json:"high_school"`
				NameUse              string `json:"name_use"`
				PlayerID             string `json:"player_id"`
				NameLast             string `json:"name_last"`
				TeamID               string `json:"team_id"`
				ServiceYears         string `json:"service_years"`
				ActiveSw             string `json:"active_sw"`
			} `json:"row"`
		} `json:"queryResults"`
	} `json:"search_player_all"`
}

func Players(c echo.Context) (err error) {
	
	// Get the player name from the parameters
	playerName := c.QueryParam("name")

	// Validate required parameters
	p := &Player{
		Name: playerName,
	}
	if err = c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(p); err != nil {
		return err
	}

	// Format the URL with the player name provided
	url := fmt.Sprintf("http://lookup-service-prod.mlb.com/json/named.search_player_all.bam?sport_code='mlb'&active_sw='Y'&name_part='%s%s'", playerName, "%25")

	// Get the player information from api
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// Response body is []byte
	body, err := ioutil.ReadAll(resp.Body)

	// Parse []byte to the go struct pointer
    var result PlayerResults
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	// Return JSON formatted player information results
	return c.JSON(http.StatusOK, result)
}
