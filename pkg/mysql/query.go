package mysql

type QueryOption func(options *QueryOptions)

type QueryOptions struct {
	AuthorName string
	Tag        string
	Favorited  string
	Offset     int64
	Limit      int64
}

func WithOffset(offset int64) QueryOption {
	return func(o *QueryOptions) {
		o.Offset = offset
	}
}

func WithLimit(limit int64) QueryOption {
	return func(o *QueryOptions) {
		o.Limit = limit
	}
}

func WithFavorited(Favorited string) QueryOption {
	return func(o *QueryOptions) {
		o.Favorited = Favorited
	}
}

func WithTag(tag string) QueryOption {
	return func(o *QueryOptions) {
		o.Tag = tag
	}
}

func WithAuthorName(name string) QueryOption {
	return func(o *QueryOptions) {
		o.AuthorName = name
	}
}
