package prices_watcher

type IDataSource interface {
	get() pagesArray
}
