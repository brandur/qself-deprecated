package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)


type Conf struct {
	TwitterConsumerKey string `env:"TWITTER_CONSUMER_KEY,required"`
	TwitterConsumerSecret string `env:"TWITTER_CONSUMER_SECRET,required"`
}

func main() {
	var conf Conf
	err := envdecode.StrictDecode(&conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding configuration: %v", err)
		os.Exit(1)
	}

	twitterConf := &clientcredentials.Config{
		ClientID:     conf.TwitterConsumerKey,
		ClientSecret: conf.TwitterConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := twitterConf.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	tweet, _, err := client.Statuses.Show(585613041028431872, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error showing tweet: %v", err)
		os.Exit(1)
	}

	fmt.Printf("tweet = %+v\n", tweet)
}
