package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
)

const VERSION = "0.0.1"

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "token, t",
		Usage: "Connect to Slack api with `TOKEN`",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "slack-tail"
	app.Usage = "following messages of channels."
	app.Version = VERSION
	app.Flags = flags
	app.Action = slackTailMain

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func slackTailMain(c *cli.Context) error {
	token := c.String("token")
	if token == "" {
		token = os.Getenv("SLACK_TOKEN")
	}

	rtm, err := newRTM(token)
	if err != nil {
		return err
	}

	ws, _, err := websocket.DefaultDialer.Dial(rtm.getWSSUrl(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	sig_ch := make(chan os.Signal, 1)
	signal.Notify(sig_ch, os.Interrupt)

	exit_ch := make(chan int)

	go listenWS(ws, rtm)

	go func() {
		sig := <-sig_ch
		fmt.Println("signal!")
		fmt.Println(sig)
		exit_ch <- 1
	}()

	code := <-exit_ch
	fmt.Printf("%d\n", code)
	return nil
}

func listenWS(ws *websocket.Conn, rtm *RTM) {
	defer ws.Close()
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		event, err := trimEvent(string(msg), rtm)
		if event.isPrint() {
			fmt.Printf("%s\r\n", event.toString())
		}
	}
}
