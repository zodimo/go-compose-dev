package progress

type LoadingContainerOption func(*LoadingContainerOptions)

type LoadingContainerOptions struct {
	LoadingIndicator Composable
}

func DefaultLoadingContainerOptions() LoadingContainerOptions {
	return LoadingContainerOptions{
		LoadingIndicator: nil,
	}
}

func WithLoadingIndicator(indicator Composable) LoadingContainerOption {
	return func(opts *LoadingContainerOptions) {
		opts.LoadingIndicator = indicator
	}
}
