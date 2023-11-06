package main

import (
    "github.com/bwmarrin/discordgo"

    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
)

var(
    Message string
    Token string
    Channel string
)      

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&Channel, "c", "", "Channel")
	flag.Parse()
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
    if Token == "" {
        log.Fatal("No token passed, pass one with the -t flag") 
        return
    }    
    if Channel == "" {
        log.Fatal("No ChannelID passed, pass one with the -c flag")
        return
    } 
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
    dg.ChannelMessageSend(Channel, Message)

    dg.Close()  
}      
