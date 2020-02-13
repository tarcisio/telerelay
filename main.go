package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("informar o user ID do dono do bot e o token")
	}

	dono, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	tkn := os.Args[2]

	b, err := tb.NewBot(tb.Settings{
		Token:  tkn,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}
	b.Handle(tb.OnText, func(m *tb.Message) {
		fmt.Println("USER ID:")
		fmt.Println(m.Sender.ID)
		
		if m.Sender.ID == dono {
			index := strings.Index(m.Text, " ")
			usr, err := strconv.Atoi(m.Text[:index])
			if err != nil {
				b.Send(&tb.User{ID: dono}, "erro ao pegar o id do usuario")
				return
			}
			b.Send(&tb.User{ID: usr}, m.Text[index+1:])
		} else {
			b.Send(&tb.User{ID: dono}, "["+strconv.Itoa(m.Sender.ID)+"] (- "+m.Sender.Username+" - "+m.Sender.FirstName+" "+m.Sender.LastName+") "+m.Text)
		}
	})

	b.Start()
}
