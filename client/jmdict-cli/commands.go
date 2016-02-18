package main

import (
	"encoding/json"
	"fmt"
	"github.com/nii236/jmdict/client"
	"github.com/spf13/cobra"
)

type (
	// TranslateTranslateCommand is the command line data structure for the translate action of Translate
	TranslateTranslateCommand struct {
		Payload string
	}
)

// Run makes the HTTP request corresponding to the TranslateTranslateCommand command.
func (cmd *TranslateTranslateCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/translate"
	}
	var payload client.TranslateTranslatePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	resp, err := c.TranslateTranslate(path, &payload)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TranslateTranslateCommand) RegisterFlags(cc *cobra.Command) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request JSON body")
}
