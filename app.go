package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/kindlyfire/go-keylogger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	SET   = true
	RESET = false
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
				runtime.WindowReload(ctx)
				runtime.WindowShow(ctx)
				updateMonitor(SET)
			case "-":
				runtime.WindowHide(ctx)
				updateMonitor(RESET)

			case "!":
				runtime.Quit(ctx)

			}
		}

		// emptyCount++

		// fmt.Printf("Empty count: %d\r", emptyCount)

		time.Sleep(5 * time.Millisecond)
	}

}

/*
	Выведем банер на все мониторы

	DisplaySwitch.exe/internal

/internal используется для переключения вашего ПК на использование только основного дисплея.
Совет: вы можете попробовать эти параметры прямо в диалоговом окне «Выполнить». Откройте его с помощью ярлыка Win + R и введите указанную выше команду в поле «Выполнить».

DisplaySwitch.exe/external

Используйте эту команду для переключения только на внешний дисплей.

	DisplaySwitch.exe/clone

Дублирует основной дисплей

	DisplaySwitch.exe/extend

Расширяет рабочий стол до второго дисплея
*/
func updateMonitor(state bool) {

	var command string
	if state {
		command = "DisplaySwitch.exe/clone"
	} else {
		command = "DisplaySwitch.exe/extend"

	}

	exec.Command("cmd", "/C", command).Run()

}
