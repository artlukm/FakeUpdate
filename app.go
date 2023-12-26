package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kindlyfire/go-keylogger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go trackKey(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func trackKey(ctx context.Context) {

	kl := keylogger.NewKeylogger()
	// emptyCount := 0

	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)

			keyChar := fmt.Sprintf("%c", key.Rune)

			switch keyChar {
			case "+":
				runtime.WindowShow(ctx)
			case "-":
				runtime.WindowHide(ctx)
			case "!":
				runtime.Quit(ctx)

			}
			// if keyChar == "+" {
			// 	WindowHide(ctx)
			// }+

			// if keyChar == "_" {
			// 	WindowShow(ctx)
			// }
		}

		// emptyCount++

		// fmt.Printf("Empty count: %d\r", emptyCount)

		time.Sleep(5 * time.Millisecond)
	}

}
