package main

import (
	"github.com/haapjari/caliber/internal/pkg/log"
	"github.com/haapjari/caliber/internal/pkg/ui"
)

func main() {
	log.NewLogger()

	cli := ui.New()
	cli.Listen()
}
