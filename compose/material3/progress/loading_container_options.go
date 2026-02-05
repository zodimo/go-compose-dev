package progress

type LoadingContainerOption func(*LoadingContainerOptions)

type LoadingContainerOptions struct {
	LoadingIndicator Composable
	IndicatorOptions []IndicatorOption
}

func DefaultLoadingContainerOptions() LoadingContainerOptions {
	return LoadingContainerOptions{
		LoadingIndicator: nil,
		IndicatorOptions: []IndicatorOption{},
	}
}

func WithLoadingIndicator(indicator Composable) LoadingContainerOption {
	return func(opts *LoadingContainerOptions) {
		opts.LoadingIndicator = indicator
	}
}

func WithIndicatorOptions(options ...IndicatorOption) LoadingContainerOption {
	return func(opts *LoadingContainerOptions) {
		opts.IndicatorOptions = options
	}
}
