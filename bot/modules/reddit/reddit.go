package reddit

import (
	"HeinzBotGoEdition/bot"
	"github.com/jzelinskie/geddit"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"math/rand"
	"os"
	"strings"
)

type FilteredSub struct {
	URL     string
	Caption string
	IsVideo bool
}

func RegisterReddit() {

	bot.Bot().Handle("/reddit", func(m *tb.Message) {

		session, err := geddit.NewLoginSession(
			os.Getenv("REDDIT_USERNAME"),
			os.Getenv("REDDIT_PWD"),
			"gedditAgent v1",
		)

		if err != nil {
			log.Fatal(err)
		}

		limit := 50

		// Set listing options
		subOpts := geddit.ListingOptions{
			Limit: limit,
		}

		// Get specific subreddit submissions, sorted by new
		submissions, err := session.SubredditSubmissions(m.Payload, geddit.HotSubmissions, subOpts)
		if err != nil {
			bot.ReportError(err, m)
			return
		}
		if len(submissions) == 0 {
			bot.SendText("Den Subreddit gibts nimma oda er hot kane Posts duat ma lad!", m)
			return
		}

		filteredSubs := filterSubmissions(submissions)
		randomPost := filteredSubs[rand.Intn(len(filteredSubs))]

		if randomPost.IsVideo {
			bot.SendVideoByURL(randomPost.URL, randomPost.Caption, m)
		} else {
			bot.SendPictureByURL(randomPost.URL, randomPost.Caption, m)
		}

	})
}

func filterSubmissions(submissions []*geddit.Submission) []FilteredSub {

	var filteredSubs []FilteredSub

	for _, submission := range submissions {
		//Get through every single possibility of data and source the post can be
		if submission.Domain == "i.imgur.com" {
			if strings.Contains(submission.URL, ".gifv") {
				filteredSubs = append(filteredSubs, FilteredSub{
					URL:     strings.Replace(submission.URL, ".gifv", ".gif", 1),
					Caption: submission.Title,
					IsVideo: true,
				})
			} else if strings.Contains(submission.URL, ".gif") {
				filteredSubs = append(filteredSubs, FilteredSub{
					URL:     submission.URL,
					Caption: submission.Title,
					IsVideo: true,
				})
			} else {
				filteredSubs = append(filteredSubs, FilteredSub{
					URL:     submission.URL,
					Caption: submission.Title,
					IsVideo: false,
				})
			}

		}
		if submission.Domain == "i.redd.it" {
			if strings.Contains(submission.URL, ".gif") {
				filteredSubs = append(filteredSubs, FilteredSub{
					URL:     submission.URL,
					Caption: submission.Title,
					IsVideo: true,
				})
			} else {
				filteredSubs = append(filteredSubs, FilteredSub{
					URL:     submission.URL,
					Caption: submission.Title,
					IsVideo: false,
				})

			}
		}
		if submission.Domain == "gfycat.com" || submission.Domain == "redgifs.com" {
			filteredSubs = append(filteredSubs, FilteredSub{
				URL:     submission.Preview.RedditVideoPreview.FallbackURL,
				Caption: submission.Title,
				IsVideo: true,
			})
		}
		if submission.Domain != "v.redd.it" {
			if submission.IsVideo == false {
				continue
			}
			filteredSubs = append(filteredSubs, FilteredSub{
				URL:     submission.Media.RedditVideo.FallbackURL,
				Caption: submission.Title,
				IsVideo: true,
			})
		}
	}
	return filteredSubs
}
