package common

import "github.com/zodimo/go-compose/pkg/api"

func CommonDecorationBox() api.Composable {
	return func(c Composer) Composer {
		return c
	}
}

// internal fun CommonDecorationBox(
//     type: TextFieldType,
//     visualText: CharSequence,
//     innerTextField: @Composable () -> Unit,
//     labelPosition: TextFieldLabelPosition,
//     label: @Composable (TextFieldLabelScope.() -> Unit)?,
//     placeholder: @Composable (() -> Unit)?,
//     leadingIcon: @Composable (() -> Unit)?,
//     trailingIcon: @Composable (() -> Unit)?,
//     prefix: @Composable (() -> Unit)?,
//     suffix: @Composable (() -> Unit)?,
//     supportingText: @Composable (() -> Unit)?,
//     singleLine: Boolean,
//     enabled: Boolean,
//     isError: Boolean,
//     interactionSource: InteractionSource,
//     contentPadding: PaddingValues,
//     colors: TextFieldColors,
//     container: @Composable () -> Unit,
// ) {
//     val isFocused = interactionSource.collectIsFocusedAsState().value
//     val inputState =
//         when {
//             isFocused -> InputPhase.Focused
//             visualText.isEmpty() -> InputPhase.UnfocusedEmpty
//             else -> InputPhase.UnfocusedNotEmpty
//         }

//     val labelColor = colors.labelColor(enabled, isError, isFocused)

//     val typography = MaterialTheme.typography
//     val bodyLarge = typography.bodyLarge
//     val bodySmall = typography.bodySmall
//     val overrideLabelTextStyleColor =
//         (bodyLarge.color == Color.Unspecified && bodySmall.color != Color.Unspecified) ||
//             (bodyLarge.color != Color.Unspecified && bodySmall.color == Color.Unspecified)

//     TextFieldTransitionScope(
//         inputState = inputState,
//         focusedLabelTextStyleColor =
//             with(bodySmall.color) {
//                 if (overrideLabelTextStyleColor) this.takeOrElse { labelColor } else this
//             },
//         unfocusedLabelTextStyleColor =
//             with(bodyLarge.color) {
//                 if (overrideLabelTextStyleColor) this.takeOrElse { labelColor } else this
//             },
//         labelColor = labelColor,
//         showExpandedLabel = label != null && labelPosition.showExpandedLabel,
//     ) { labelProgress, labelTextStyleColor, labelContentColor, placeholderAlpha, prefixSuffixAlpha
//         ->
//         val labelScope = remember {
//             object : TextFieldLabelScope {
//                 override val labelMinimizedProgress: Float
//                     get() = labelProgress.value
//             }
//         }
//         val decoratedLabel: @Composable (() -> Unit)? =
//             label?.let { label ->
//                 @Composable {
//                     val labelTextStyle =
//                         lerp(bodyLarge, bodySmall, labelProgress.value).let { textStyle ->
//                             if (overrideLabelTextStyleColor) {
//                                 textStyle.copy(color = labelTextStyleColor.value)
//                             } else {
//                                 textStyle
//                             }
//                         }
//                     Decoration(labelContentColor.value, labelTextStyle) { labelScope.label() }
//                 }
//             }

//         // Transparent components interfere with Talkback (b/261061240), so if any components below
//         // have alpha == 0, we set the component to null instead.

//         val placeholderColor = colors.placeholderColor(enabled, isError, isFocused)
//         val showPlaceholder by remember {
//             derivedStateOf(structuralEqualityPolicy()) { placeholderAlpha.value > 0f }
//         }
//         val decoratedPlaceholder: @Composable ((Modifier) -> Unit)? =
//             if (placeholder != null && visualText.isEmpty() && showPlaceholder) {
//                 @Composable { modifier ->
//                     Box(modifier.graphicsLayer { alpha = placeholderAlpha.value }) {
//                         Decoration(
//                             contentColor = placeholderColor,
//                             textStyle = bodyLarge,
//                             content = placeholder,
//                         )
//                     }
//                 }
//             } else null

//         val prefixColor = colors.prefixColor(enabled, isError, isFocused)
//         val showPrefixSuffix by remember {
//             derivedStateOf(structuralEqualityPolicy()) { prefixSuffixAlpha.value > 0f }
//         }
//         val decoratedPrefix: @Composable (() -> Unit)? =
//             if (prefix != null && showPrefixSuffix) {
//                 @Composable {
//                     Box(Modifier.graphicsLayer { alpha = prefixSuffixAlpha.value }) {
//                         Decoration(
//                             contentColor = prefixColor,
//                             textStyle = bodyLarge,
//                             content = prefix,
//                         )
//                     }
//                 }
//             } else null

//         val suffixColor = colors.suffixColor(enabled, isError, isFocused)
//         val decoratedSuffix: @Composable (() -> Unit)? =
//             if (suffix != null && showPrefixSuffix) {
//                 @Composable {
//                     Box(Modifier.graphicsLayer { alpha = prefixSuffixAlpha.value }) {
//                         Decoration(
//                             contentColor = suffixColor,
//                             textStyle = bodyLarge,
//                             content = suffix,
//                         )
//                     }
//                 }
//             } else null

//         val leadingIconColor = colors.leadingIconColor(enabled, isError, isFocused)
//         val decoratedLeading: @Composable (() -> Unit)? =
//             leadingIcon?.let {
//                 @Composable { Decoration(contentColor = leadingIconColor, content = it) }
//             }

//         val trailingIconColor = colors.trailingIconColor(enabled, isError, isFocused)
//         val decoratedTrailing: @Composable (() -> Unit)? =
//             trailingIcon?.let {
//                 @Composable { Decoration(contentColor = trailingIconColor, content = it) }
//             }

//         val supportingTextColor = colors.supportingTextColor(enabled, isError, isFocused)
//         val decoratedSupporting: @Composable (() -> Unit)? =
//             supportingText?.let {
//                 @Composable {
//                     Decoration(
//                         contentColor = supportingTextColor,
//                         textStyle = bodySmall,
//                         content = it,
//                     )
//                 }
//             }

//         when (type) {
//             TextFieldType.Filled -> {
//                 val containerWithId: @Composable () -> Unit = {
//                     Box(Modifier.layoutId(ContainerId), propagateMinConstraints = true) {
//                         container()
//                     }
//                 }

//                 TextFieldLayout(
//                     modifier = Modifier,
//                     textField = innerTextField,
//                     placeholder = decoratedPlaceholder,
//                     label = decoratedLabel,
//                     leading = decoratedLeading,
//                     trailing = decoratedTrailing,
//                     prefix = decoratedPrefix,
//                     suffix = decoratedSuffix,
//                     container = containerWithId,
//                     supporting = decoratedSupporting,
//                     singleLine = singleLine,
//                     labelPosition = labelPosition,
//                     labelProgress = labelProgress::value,
//                     paddingValues = contentPadding,
//                 )
//             }
//             TextFieldType.Outlined -> {
//                 // Outlined cutout
//                 val cutoutSize = remember { mutableStateOf(Size.Zero) }
//                 val borderContainerWithId: @Composable () -> Unit = {
//                     Box(
//                         Modifier.layoutId(ContainerId)
//                             .outlineCutout(
//                                 labelSize = cutoutSize::value,
//                                 alignment = labelPosition.minimizedAlignment,
//                                 paddingValues = contentPadding,
//                             ),
//                         propagateMinConstraints = true,
//                     ) {
//                         container()
//                     }
//                 }

//                 OutlinedTextFieldLayout(
//                     modifier = Modifier,
//                     textField = innerTextField,
//                     placeholder = decoratedPlaceholder,
//                     label = decoratedLabel,
//                     leading = decoratedLeading,
//                     trailing = decoratedTrailing,
//                     prefix = decoratedPrefix,
//                     suffix = decoratedSuffix,
//                     supporting = decoratedSupporting,
//                     singleLine = singleLine,
//                     onLabelMeasured = {
//                         if (labelPosition is TextFieldLabelPosition.Above) {
//                             return@OutlinedTextFieldLayout
//                         }
//                         val progress = labelProgress.value
//                         val labelWidth = it.width * progress
//                         val labelHeight = it.height * progress
//                         if (
//                             cutoutSize.value.width != labelWidth ||
//                                 cutoutSize.value.height != labelHeight
//                         ) {
//                             cutoutSize.value = Size(labelWidth, labelHeight)
//                         }
//                     },
//                     labelPosition = labelPosition,
//                     labelProgress = labelProgress::value,
//                     container = borderContainerWithId,
//                     paddingValues = contentPadding,
//                 )
//             }
//         }
//     }
// }
