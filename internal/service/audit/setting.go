package audit

type Setting struct {
	Content      any
	OriginalData any
	NewData      any
}

type Option func(setting *Setting)

type CM map[string]any

func WithContent(contentMap CM) Option {
	return func(setting *Setting) {
		setting.Content = contentMap
	}
}

func WithOriginalData(originalData any) Option {
	return func(setting *Setting) {
		setting.OriginalData = originalData
	}
}

func WithNewData(newData any) Option {
	return func(setting *Setting) {
		setting.NewData = newData
	}
}

func WithData(originalData, newData any) Option {
	return func(setting *Setting) {
		setting.OriginalData = originalData
		setting.NewData = newData
	}
}
