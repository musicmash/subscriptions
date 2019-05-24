package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/musicmash/artists/internal/config"
	"github.com/musicmash/artists/internal/db"
	"github.com/musicmash/artists/internal/log"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	storeName = "spotify"
)

var (
	limit = 50

	clientID     string
	clientSecret string

	searchQuery string
)

func init() {
	flag.StringVar(&clientID, "id", "", "spotify app id")
	flag.StringVar(&clientSecret, "secret", "", "spotify app secret")
	flag.StringVar(&searchQuery, "query", "adept", "name to search")
	configPath := flag.String("config", "/etc/musicmash/artists/artists.yaml", "Path to artists.yaml config")
	flag.Parse()

	if len(clientID) == 0 || len(clientSecret) == 0 {
		log.Error("id or secret not provided")
		os.Exit(2)
	}

	if err := config.InitConfig(*configPath); err != nil {
		panic(err)
	}
	if config.Config.Log.Level == "" {
		config.Config.Log.Level = "info"
	}

	log.SetLogFormatter(&log.DefaultFormatter)
	log.ConfigureStdLogger(config.Config.Log.Level)

	db.DbMgr = db.NewMainDatabaseMgr()

	log.Info("Ensuring that 'spotify' exists...")
	if err := db.DbMgr.EnsureStoreExists(storeName); err != nil {
		log.Panic(err)
	}
}

func main() {
	credentials := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := credentials.Token(context.Background())
	if err != nil {
		log.Panic(fmt.Errorf("couldn't get token: %v", err))
	}

	client := spotify.Authenticator{}.NewClient(token)
	results, err := client.SearchOpt(searchQuery, spotify.SearchTypeArtist, &spotify.Options{
		Limit: &limit,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Debugf("limit %v offset %v total %v", results.Artists.Limit, results.Artists.Offset, results.Artists.Total)
	processArtists(client, sortArtistsByPopularity(results.Artists.Artists))

	// load next part
	for results.Artists.Total > results.Artists.Limit+results.Artists.Offset {
		log.Info("getting next artists...")
		if err = client.NextArtistResults(results); err != nil {
			log.Panic(err)
		}

		log.Debugf("limit %v offset %v total %v", results.Artists.Limit, results.Artists.Offset, results.Artists.Total)
		processArtists(client, sortArtistsByPopularity(results.Artists.Artists))
	}
}

func sortArtistsByPopularity(artists []spotify.FullArtist) []spotify.FullArtist {
	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Popularity > artists[j].Popularity
	})
	return artists
}

func processArtists(client spotify.Client, artists []spotify.FullArtist) {
	for _, artist := range artists {
		processArtist(client, artist)
	}
}

func processArtist(client spotify.Client, artist spotify.FullArtist) {
	if exists := db.DbMgr.IsArtistExistsInStore(storeName, artist.ID.String()); exists {
		log.Warn(artist.ID, artist.Name, "already exists")
		return
	}

	newArtist := &db.Artist{
		Name:       artist.Name,
		Popularity: artist.Popularity,
		Followers:  artist.Followers.Count,
	}
	if len(artist.Images) > 0 {
		newArtist.Poster = artist.Images[0].URL
	}

	log.Info("creating new artist", newArtist.Name)
	if err := db.DbMgr.EnsureArtistExists(newArtist); err != nil {
		log.Error("can't create new artist")
	}

	log.Info("save spotify id for new artist", newArtist.ID)
	if err := db.DbMgr.EnsureArtistExistsInStore(newArtist.ID, storeName, artist.ID.String()); err != nil {
		log.Error("can't save spotify id for new artist")
	}

	log.Info("loading and processing albums from", artist.Name)
	loadAndProcessAlbums(client, artist.ID, newArtist.ID)
}

func loadAndProcessAlbums(client spotify.Client, artistID spotify.ID, dbArtistID int64) {
	albumPage, err := client.GetArtistAlbums(artistID)
	if err != nil {
		log.Error(err)
	}

	processAlbums(client, albumPage.Albums, dbArtistID)

	for albumPage.Total > albumPage.Limit+albumPage.Offset {
		albumPage.Offset += albumPage.Limit
		log.Info("getting next albums...")

		opts := spotify.Options{
			Limit:  &limit,
			Offset: &albumPage.Offset,
		}
		albumPage, err = client.GetArtistAlbumsOpt(artistID, &opts, nil)
		if err != nil {
			log.Panic(err)
		}

		processAlbums(client, albumPage.Albums, dbArtistID)
	}
}

func processAlbums(client spotify.Client, albums []spotify.SimpleAlbum, dbArtistID int64) {
	for _, album := range albums {
		log.Debugf("process albums from %s", album.Artists[0].Name)
		processAlbum(client, album, dbArtistID)
	}
}

func processAlbum(client spotify.Client, album spotify.SimpleAlbum, dbArtistID int64) {
	log.Infof("saving album %s", album.Name)
	err := db.DbMgr.EnsureAlbumExists(&db.Album{
		ArtistID: dbArtistID,
		Name:     album.Name,
	})

	if err != nil {
		log.Error(err)
	}

	// handle other artists mentioned in this album
	//for _, artist := range album.Artists {
	//	processArtist(client, spotify.FullArtist{SimpleArtist:artist})
	//}
}
