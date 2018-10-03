package main

type IDataSource interface {
	get() pagesArray
}
