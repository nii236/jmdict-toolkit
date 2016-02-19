package main

import (
	"encoding/json"
	"fmt"
	"github.com/nii236/jmdict-toolkit/serve/client"
	"github.com/spf13/cobra"
)

type (
	// TranslateWordCommand is the command line data structure for the translate action of Word
	TranslateWordCommand struct {
		Payload string
	}
)

// Run makes the HTTP request corresponding to the TranslateWordCommand command.
func (cmd *TranslateWordCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/translate"
	}
	var payload client.TranslateWordPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.TranslateWord(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TranslateWordCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}
