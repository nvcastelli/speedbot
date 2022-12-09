package main
import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
)
//This the token you receive from botfather
const token string = "5961795051:AAGjuNF-bSjLb43S5Wr-R4GW81ERZ4A4xoU"
//The instance of the bot
var bot *bt.Bot
func main() {
	//Update configs
	up := cfg.DefaultUpdateConfigs()
   //Pushing bot configs next 
    //Bot configs
cf := cfg.BotConfigs{
      BotAPI: cfg.DefaultBotAPI, 
      APIKey: token, UpdateConfigs: up, 
      Webhook: false, 
      LogFileAddress: cfg.DefaultLogFile
    }
	
}
