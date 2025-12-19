package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "client",
	Short: "Start client to interact with teamserver",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
        app.EnableMouse(true)

		// Make sure the screen itself uses the terminal default background
		app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
			screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
			return false
		})

		var agentsList *tview.List 
		agentsList = tview.NewList()
		agentsList.SetTitle("Sessions (Agents)").SetBorder(true)
		agentsList.SetBackgroundColor(tcell.ColorDefault)

		var listenerList *tview.List
		listenerList = tview.NewList()
		listenerList.SetTitle("Listener").SetBorder(true)
		listenerList.SetBackgroundColor(tcell.ColorDefault)

		var logView *tview.TextView
		logView = tview.NewTextView()
		logView.SetTitle("Logs").SetBorder(true)
		logView.SetBackgroundColor(tcell.ColorDefault)

		input := tview.NewInputField().SetLabel("Command > ")
		input.SetBackgroundColor(tcell.ColorDefault)
		input.SetLabelColor(tcell.ColorDefault)
		input.SetFieldBackgroundColor(tcell.ColorValid)

		flex := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(agentsList, 0, 1, false).
			AddItem(listenerList, 0, 1, false),
			0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(logView, 0, 3, false).
			AddItem(input, 3, 1, true),
			0, 2, false)
		flex.SetBackgroundColor(tcell.ColorDefault)

		conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
		if err != nil {
			log.Fatal("[!] WebSocket connection failed:", err)
		}

		go func() {
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					app.QueueUpdateDraw(func() {
						fmt.Fprint(logView, "[x] Lost connection ...\n")
					})
					return
				}

				parts := strings.SplitN(string(msg), ":", 2)
				if len(parts) == 2 && parts[0] == "client" {
					id := parts[1]
					app.QueueUpdateDraw(func() {
						agentsList.AddItem(id, "", 0, nil)
					})
				} else {
					app.QueueUpdateDraw(func() {
						fmt.Fprintf(logView, "[Server] " + string(msg) + "\n")
					})
				}
			}
		}()

		input.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				//currentIndex := agentsList.GetCurrentItem()
				//if currentIndex < 0 {
				//	return
				//}
				//agent, _ := agentsList.GetItemText(currentIndex)
				task := input.GetText()
				input.SetText("")
				//msg := fmt.Sprintf("send:%s:%s", agent, task)
				msg := fmt.Sprintf("%s", task)
				conn.WriteMessage(websocket.TextMessage, []byte(msg))
				//fmt.Fprintf(logView, "[>] Send to %s: %s\n", agent, task)
				fmt.Fprintf(logView, "[>] Send to %s: %s\n", "TEST", task)
			}
		})

		if err := app.SetRoot(flex, true).SetFocus(input).Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
