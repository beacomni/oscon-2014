package commands

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	//rss aliases the package
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//public/Visible type because of cap
type Config struct {
	//fields
	//Feeds is a slice, because it doesnt have length defined
	Feeds []string
	Port  int
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch feeds",
	Long:  `Dagobah will fetch feeds listed in config`,
	Run:   fetchRun,
}

//this allows a clean shutdown
func fetchRun(cmd *cobra.Command, args []string) {
	//note, generate an interrupt with ctrl-c
	//for example
	Fetcher()

	//1 is size of the queue / channel
	//so it can hold 2 things.
	//this is a buffered channel.
	//create a channel
	sigChan := make(chan os.Signal, 1)
	//signal is a package.
	//tell signal package to put an interrupt signal
	//on this channel (a signal-type channel) if that signal received
	signal.Notify(sigChan, os.Interrupt)
	//read on the channel (blocks until
	//something happens)
	<-sigChan
}

func Fetcher() {
	var config Config

	//marshal into a variable of the named type
	//we created before
	if err := viper.Marshal(&config); err != nil {
		fmt.Println(err)
	}

	for _, feed := range config.Feeds {
		//creates a goroutine
		//concurrent but not asynchronous!!
		//look for a rob pike presentation
		//the function is run, but execution
		//continues immediately
		//will do one for each feed we put in config
		//simultaneously
		go PollFeed(feed)
	}
}

func PollFeed(uri string) {
	// check that timeout is a valid value
	timeout := viper.GetInt("RSSTimeout")
	if timeout < 1 {
		timeout = 1
	}
	//handlers called when new item detected
	//channels here are about go channels, not rss channels
	feed := rss.New(timeout, true, chanHandler, itemHandler)

	for {
		if err := feed.Fetch(uri, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
			return
		}

		fmt.Printf("Sleeping for %d seconds on %s\n", feed.SecondsTillUpdate(), uri)
		//sleep is ok from in this goroutine because it only pauses it.
		//we only want to check feeds at an interval (not constantly)
		//sleep is through scheduler, not thread sleep
		time.Sleep(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

//here Channel is a channel type within rss
func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}

func init() {
	fetchCmd.Flags().Int("rsstimeout", 5, "Timeout (in min) for rss retrieval")
	viper.BindPFlag("rsstimeout", fetchCmd.Flags().Lookup("rsstimeout"))
}
