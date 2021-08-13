package main

import "fmt"

const globalLimit = 100
const (
	CacheKeyBook = "book_"
	CacheKeyCD = "cd_"
)
const MaxCacheSize int = 10 * globalLimit
//Declare a map that has a string for a key and a string for its values as our cache:
var cache map[string]string

func main() {
	cache = make(map[string] string)
	//add a Book to cache
	SetBook("1234-5678","Get Ready To Go")
	//add a CD to cache
	SetCd("1234-5678","Get Ready To Go Audio Book")

	//Get and Print that Book from the Cache
	fmt.Println("Book :",GetBook("1234-5678"))

	// get and print that CD from the Cache
	fmt.Println("CD :",GetCd("1234-5678"))


}
func chacheGet(Key string) string {
	return cache[Key]
}
func cacheSet(Key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[Key] = val
}

//Use the book cache prefix to create a unique key:
func GetBook(isbn string)string {
	return chacheGet(CacheKeyBook + isbn)
}
//Use the book cache prefix to create a unique key:
func SetBook(isbn string,name string) {
	cacheSet(CacheKeyBook+isbn,name)
}
//Create a function to get CD data from the cache:
func GetCd(sku string) string{
	return chacheGet(CacheKeyCD + sku)
}
//Create a function to add CDs to the shared cache:
func SetCd(sku string , title string) {
	cacheSet(CacheKeyCD+sku, title)
}