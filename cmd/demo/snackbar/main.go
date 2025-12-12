package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/unit"

	"go-compose-dev/compose"
	"go-compose-dev/compose/runtime"
	"go-compose-dev/internal/state"
	"go-compose-dev/internal/store"
	"go-compose-dev/internal/theme"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Snackbar Demo"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	enLocale := system.Locale{Language: "en", Direction: system.LTR}
	var ops op.Ops

	store := store.NewPersistentState(map[string]state.MutableValue{})
	runtime := runtime.NewRuntime()
	themeManager := theme.GetThemeManager()

	for {
		switch frameEvent := w.Event().(type) {
		case app.DestroyEvent:
			return frameEvent.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, frameEvent)
			gtx.Locale = enLocale
			gtx = themeManager.Material3ThemeInit(gtx)

			composer := compose.NewComposer(store)
			layoutNode := UI(composer)

			callOp := runtime.Run(gtx, layoutNode)
			callOp.Add(gtx.Ops)
			frameEvent.Frame(gtx.Ops)
		}
	}
}
