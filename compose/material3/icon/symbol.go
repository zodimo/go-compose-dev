package icon

// Symbol renders a Material Symbol icon using the Material Symbols Outlined font.
// Use SymbolName constants (e.g. SymbolSearch, SymbolHome) to avoid typos.
//
// Deprecated: Use Icon(symbolName, ...) directly instead.
// Symbol is kept for backward compatibility.
func Symbol(name SymbolName, options ...IconOption) Composable {
	return Icon(name, options...)
}
