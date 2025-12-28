package textfield

type TextFieldDecorator struct {
	innerTextField Composable
}

func decorator() Composable {
	return func(c Composer) Composer {
		return c
	}
}

// state: TextFieldState,
// enabled: Boolean,
// lineLimits: TextFieldLineLimits,
// outputTransformation: OutputTransformation?,
// interactionSource: InteractionSource,
// labelPosition: TextFieldLabelPosition = TextFieldLabelPosition.Attached(),
// label: @Composable (TextFieldLabelScope.() -> Unit)? = null,
// placeholder: @Composable (() -> Unit)? = null,
// leadingIcon: @Composable (() -> Unit)? = null,
// trailingIcon: @Composable (() -> Unit)? = null,
// prefix: @Composable (() -> Unit)? = null,
// suffix: @Composable (() -> Unit)? = null,
// supportingText: @Composable (() -> Unit)? = null,
// isError: Boolean = false,
// colors: TextFieldColors = colors(),
// contentPadding: PaddingValues =
// 	if (label == null || labelPosition is TextFieldLabelPosition.Above) {
// 		contentPaddingWithoutLabel()
// 	} else {
// 		contentPaddingWithLabel()
// 	},
// container: @Composable () -> Unit = {
// 	Container(
// 		enabled = enabled,
// 		isError = isError,
// 		interactionSource = interactionSource,
// 		colors = colors,
// 		shape = shape,
// 		focusedIndicatorLineThickness = FocusedIndicatorThickness,
// 		unfocusedIndicatorLineThickness = UnfocusedIndicatorThickness,
// 	)
// },
