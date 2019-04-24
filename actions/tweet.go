package actions

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/buffalo"
)

var creds = Credentials{
	APIKey:    os.Getenv("API_KEY"),
	APISecret: os.Getenv("API_SECRET"),
}

// TweetHandler is the handler for the Tweet page
func TweetHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("tweet.html"))
}

//SendHandler Function used by tweet.html to send the tweet/
//Takes in variables from twitter/server.go
func SendHandler(c buffalo.Context) error {

	//Calls http.Request package to get the form values. Assigns it to Req
	Req := c.Request()
	//ParseForm decodes the form
	Req.ParseForm()

	//Gets the value from the form, assigns it to body.
	body := Req.FormValue("tweetbody")

	fmt.Printf("%+v\n", creds)

	//Gets the client from the twitter package
	client, err := GetClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	//Sending the actual tweet. Body from the form
	tweet, resp, err := client.Statuses.Update(body, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
	return c.Redirect(302, "/tweet/confirm")

}
