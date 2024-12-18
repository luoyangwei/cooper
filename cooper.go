package cooper

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// Command is a command that can be executed.
type Command interface {
	Execute()
}

var _ Command = (*Cooper)(nil)
var _ flag.Value = (*Cooper)(nil)

type Cooper struct {
	Files []string

	readCh chan []byte
}

// Set implements flag.Value.
func (c *Cooper) Set(value string) error {
	c.Files = append(c.Files, value)
	return nil
}

// String implements flag.Value.
func (c *Cooper) String() string {
	return fmt.Sprint(c.Files)
}

func (c *Cooper) readFile(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file line by line
	reader := bufio.NewReader(file)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done reading file: ", filePath)
		default:
			// Read line
			line, err := reader.ReadBytes('\n')
			if errors.Is(err, io.EOF) {
				continue
			}
			if err != nil {
				log.Panicln("Error reading file: ", err)
				return err
			}
			c.readCh <- line
		}
	}
}

// Execute executes the command.
func (c *Cooper) Execute() {
	c.readCh = make(chan []byte)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, filePath := range c.Files {
		go func() {
			if err := c.readFile(ctx, filePath); err != nil {
				cancel()
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			log.Panicln("Done reading files")
		case data := <-c.readCh:
			fmt.Print(string(data))
		}
	}
}
