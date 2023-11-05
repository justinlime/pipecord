package main

import (
    "github.com/bwmarrin/discordgo"

    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "os/signal"
    "syscall"
)

var Message string

func main() {
    nBytes, nChunks := int64(0), int64(0)
    r := bufio.NewReader(os.Stdin)
    buf := make([]byte, 0, 4*1024)

    for {
        n, err := r.Read(buf[:cap(buf)])
        buf = buf[:n]
    
        if n == 0 {

            if err == nil {
                continue
            }

            if err == io.EOF {
                break
            }

            log.Fatal(err)
        }

        nChunks++
        nBytes += int64(len(buf))

        fmt.Println(string(buf))
        Message = string(buf)  
        sendToBot()

        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
    }

    fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
}

func sendToBot(){
    dg, err := discordgo.New("Bot " + "")
    if err != nil {
        fmt.Printf("Error!")
    }    
    dg.AddHandler(messageCreate)
    dg.Identify.Intents = discordgo.IntentsGuildMessages
 	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

    dg.Close()  
}      

 func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, Message)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
