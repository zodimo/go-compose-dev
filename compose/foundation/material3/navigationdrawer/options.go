package navigationdrawer

import "github.com/zodimo/go-compose/internal/modifier"

type ModalNavigationDrawerOptions struct {
	Modifier Modifier
	IsOpen   bool
	OnClose  func()
}

type ModalNavigationDrawerOption func(*ModalNavigationDrawerOptions)

func WithModifier(modifier Modifier) ModalNavigationDrawerOption {
	return func(o *ModalNavigationDrawerOptions) {
		o.Modifier = o.Modifier.Then(modifier)
	}
}

func WithIsOpen(isOpen bool) ModalNavigationDrawerOption {
	return func(o *ModalNavigationDrawerOptions) {
		o.IsOpen = isOpen
	}
}

func WithOnClose(onClose func()) ModalNavigationDrawerOption {
	return func(o *ModalNavigationDrawerOptions) {
		o.OnClose = onClose
	}
}

func DefaultModalNavigationDrawerOptions() ModalNavigationDrawerOptions {
	return ModalNavigationDrawerOptions{
		Modifier: modifier.EmptyModifier,
		IsOpen:   false,
		OnClose:  func() {},
	}
}
