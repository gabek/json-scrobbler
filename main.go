package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

var config = getConfig()
var api = lastfm.New(config.LastFM.Key, config.LastFM.Secret)
var recentTrack Track

func main() {
	// Authenticate to LastFM
	err := api.Login(config.LastFM.Username, config.LastFM.Password)
	if err != nil {
		panic(err)
	}

	// Start polling
	pollingInterval, err := time.ParseDuration(config.PollingInterval)

	startTimer(pollingInterval)
}

func startTimer(pollingInterval time.Duration) {
	run()

	tick := time.Tick(pollingInterval)
	for range tick {
		run()
	}
}

func run() {
	track := getTrackFromJSON(config.Input)

	if isTrackEqual(track, recentTrack) {
		return
	}

	log.Printf("Scrobbling: %+v\n", track)

	recentTrack = track

	// Update now playing
	params := lastfm.P{"artist": track.Artist, "track": track.Name}
	_, err := api.Track.UpdateNowPlaying(params)

	if err != nil {
		fmt.Println(err)
	}

	// Scrobble
	params["timestamp"] = time.Now().Unix()
	_, err = api.Track.Scrobble(params)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func isTrackEqual(t1 Track, t2 Track) bool {
	if t1.Name == t2.Name && t1.Artist == t2.Artist {
		return true
	}
	return false
}

func getTrackFromJSON(filepath string) Track {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("Cannot open input file: " + filepath)
	}

	var track Track
	json.Unmarshal(f, &track)

	return track
}
