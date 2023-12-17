package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"

)

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6380590892000-6370528131153-U79XBWAdZSlGzAvw4cTT8j9q")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06AWFE0JJV-6360343392292-68c1dfe020fd3d230f98c0842cba6b7e89f872ec7535e3dbed423191ea76f64b")
	os.Setenv("CHANNEL_ID","C06A37ZDFMM")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr  := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Konstantinos-Kolios-Devops-2023-cv.pdf"}

	for i := 0; i< len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file , err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n",err)
		}
		fmt.Printf("Name: %s, Url: %s \n",file.Name, file.URL)
		
	}
}