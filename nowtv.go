package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

type Genre struct {
	GenreId    string `json:"genreId"`
	GenreColor string `json:"color"`
}

type ProgramInformation struct {
	StartDateTime      string `json:"startDateTime"`
	EndDateTime        string `json:"endDateTime"`
	ProgramId          string `json:"programId"`
	ProgramCode        string `json:"programCode"`
	OfficialUrl        string `json:"officialurl"`
	ModeCode           string `json:"modecode"`
	ProgramTitle       string `json:"programTitle"`
	ProgramSubTitle    string `json:"programSubTitle"`
	ProgramExplanation string `json:"programExplanation"`
	Gcode              string `json:"gcode"`
	Genre              Genre
}

type StationData struct {
	StationCode              string `json:"stationCode"`
	StationUrl               string `json:"stationUrl"`
	AdditionalDisplayChannel string `json:"additionalDisplayChannel"`
	StationDispName          string `json:"stationDispName"`
	Channel                  string `json:"channel"`
	ChannelId                string `json:"channelId"`
	ProgramNumber            string `json:"programNumber"`
	ProgramInformation       ProgramInformation
}

type MediaLocation struct {
	StationLocation []StationData
}

type Name struct {
	PackName string `json:"packName"`
	PackUrl  string `json:"packUrl"`
}

type ProgramScheduleInformartion struct {
	MediaId       string `json:"mediaId"`
	PackId        string `json:"PackId"`
	StationNumber int    `json:"stationnumber"`
	Name          Name
	MediaLocation MediaLocation
}

type SystemInformation struct {
	ID int `json:"id"`
}

type NowtvJsonData struct {
	SystemInformation  SystemInformation
	ProgramInformation ProgramInformation
}

func main() {
	app := cli.NewApp()
	app.Name = "nowtv"
	app.Version = Version
	app.Usage = ""
	app.Author = "dtan4"
	app.Email = "dtanshi45@gmail.com"
	app.Action = doMain
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
}
