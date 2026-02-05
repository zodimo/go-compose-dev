package dialog

import (
	"github.com/zodimo/go-compose/pkg/api"
)

// AlertDialog creates a Material3 alert dialog.
//
// The dialog displays optional icon, title, and text content, with a required
// confirm button and optional dismiss button. All content slots accept composables,
// allowing for rich custom content.
//
// Note: This dialog should be wrapped in an overlay.Overlay to display with
// a scrim and handle click-outside-to-dismiss behavior:
//
//	overlay.Overlay(
//	    dialog.AlertDialog(
//	        onDismissRequest,
//	        button.Text(onConfirm, "Confirm"),
//	        dialog.WithTitleText("Delete Item?"),
//	        dialog.WithTextContent("This action cannot be undone."),
//	        dialog.WithDismissButton(button.Text(onDismiss, "Cancel")),
//	    ),
//	    overlay.WithOnDismiss(onDismissRequest),
//	)
//
// Parameters:
//   - onDismissRequest: Called when the user tries to dismiss the dialog by clicking outside
//     or pressing the back button. This is NOT called when the dismiss button is clicked.
//   - confirmButton: A composable for the confirm action button (typically button.Text).
//   - options: Optional configuration including icon, title, text, and dismiss button.
func AlertDialog(
	onDismissRequest func(),
	confirmButton api.Composable,
	content api.Composable,
	options ...DialogOption,
) Composable {
	opts := DefaultDialogOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opts)
	}

	// Build the button row: [dismiss button (optional), confirm button]
	var buttonItems []api.Composable
	if opts.DismissButton != nil {
		buttonItems = append(buttonItems, opts.DismissButton)
	}
	if confirmButton != nil {
		buttonItems = append(buttonItems, confirmButton)
	}

	buttons := DialogButtonRow(buttonItems...)

	// Return the dialog content layout directly - no need for a custom block
	return DialogSurface(
		DialogContent(
			opts.Icon,
			opts.Title,
			content,
			buttons,
		),
		onDismissRequest,
	)
}

// BasicAlertDialog creates a bare dialog surface that accepts arbitrary content.
// Use this when you need full control over the dialog layout.
//
// Note: This dialog should be wrapped in an overlay.Overlay to display with
// a scrim and handle click-outside-to-dismiss behavior.
func BasicAlertDialog(
	onDismissRequest func(),
	content api.Composable,
) Composable {
	// BasicAlertDialog just wraps content in the dialog surface styling
	return DialogSurface(content, onDismissRequest)
}
