package sources

/*
DataProducer interfaces all the methods a data should implement
for being produced
*/

type DataProducer interface {

	//CacheDuration should return how long data can be cached in nanoseconds
	CacheDuration() int64

	//Load All loads all the dataModel
	LoadAll() (interface{}, error)
}
