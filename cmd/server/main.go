package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
    conn_str := "amqp://guest:guest@localhost:5672/"
    conn, err := amqp.Dial(conn_str)
    if err != nil {
	return
    }
    defer conn.Close()
    fmt.Println("AMQP connection successful.")
    fmt.Println("Starting Peril server...")
    
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT)

    done := make(chan bool, 1)

    go func() {
        <-sigs
        fmt.Println("\nReceived SIGINT - Shutting down...")
        done <- true
    }()

    <-done
    fmt.Println("Program exited gracefully.")
}
