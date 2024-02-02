package main

import (
	"context"
	"fmt"
	"go-smtp-rcv/internal/config"
	"go-smtp-rcv/internal/parser"
	"go-smtp-rcv/internal/rfc821"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	conf := config.NewConfig()
	//conf.Read("config.yaml")
	conf.GenerateExampleConfig()
	conf.Write("config.yml")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	defer stop()

	var wg sync.WaitGroup
	//
	for _, srv := range conf.Bind {
		l, err := net.Listen(srv.S_type, srv.S_addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Add(1)
		defer l.Close()
		fmt.Printf("Start listen: %s %s\n", l.Addr().Network(), l.Addr().String())
		go func() {
			for {
				select {
				case <-ctx.Done():
					{
						fmt.Printf("Recieved interrupt: %s %s\n", l.Addr().Network(), l.Addr().String())
						l.Close()
						wg.Done()
						stop()
						return
					}
				}
			}
		}()
		go func() {
			for {
				c, err := l.Accept()
				if err != nil {
					fmt.Println(err)
					return
				}
				stype := rfc821.NewSMTP_rfc821(c, srv)
				proces := parser.Parser{Ver: stype}
				go proces.Parse()
			}
		}()
	}

	fmt.Println("Server started")
	wg.Wait()
	fmt.Println("Server gracefull stopped")
}
