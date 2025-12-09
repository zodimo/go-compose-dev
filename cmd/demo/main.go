package main

import (
	"go-compose-dev/compose"
	"go-compose-dev/compose/runtime"
	"go-compose-dev/internal/state"
	"go-compose-dev/internal/store"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
)

func main() {

	go func() {
		w := new(app.Window)
		w.Option(app.Title("Pure Compose"))
		w.Option(app.Size(unit.Dp(1024), unit.Dp(768)))

		if err := Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()

}

func Run(window *app.Window) error {

	var ops op.Ops

	store := store.NewPersistentState(map[string]state.MutableValue{})
	runtime := runtime.NewRuntime()

	for {
		switch frameEvent := window.Event().(type) {
		case app.DestroyEvent:
			return frameEvent.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, frameEvent)

			// Do the stuff here
			composer := compose.NewComposer(store)
			layout := UI(composer)
			runtime.Run(gtx, layout)

			frameEvent.Frame(gtx.Ops)

		}
	}

}
