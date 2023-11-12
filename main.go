package main

import (
	"github.com/bwmarrin/discordgo"

	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var(
    Message string
    Token string
    Channel string
    Log string
)      

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&Channel, "c", "", "Channel ID")
    flag.StringVar(&Log, "l", "Log", "Log Message Title")    
	flag.Parse()
    Log = "## " + Log + "\n### " + time.Now().Format(time.RFC1123)
    if Token == "" {
        fmt.Printf("No token passed, pass one with the -t flag\n\n") 
        os.Exit(1)   
    }    
    if Channel == "" {
        fmt.Printf("No ChannelID passed, pass one with the -c flag\n\n")
        os.Exit(1)   
    } 
}

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

    // fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
}

func sendToBot(){
    dg, err := discordgo.New("Bot " + Token)
    if err != nil {
        fmt.Printf("Error!")
    }    
    dg.Identify.Intents = discordgo.IntentsGuildMessages
 	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
    dg.ChannelMessageSend(Channel, Log + "\n```" + Message + "```")

    dg.Close()  
    os.Exit(0)   
}      
