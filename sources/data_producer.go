package sources

/*
DataProducer interfaces all the methods a data should implement
for being produced
*/

type DataProducer interface {

	//CacheDuration should return how long data can be cached in seconds
	CacheDuration() int64

	//Load All loads all the dataModel
	LoadAll() interface{}
}

func ProduceDataSource() {

	/*
	 1. Get the previous announced version of the data.
	 2. De-Serialize and create back the old object
	 3. Load the current data from the data source
	 4. Get the diff of the data.
	 5. Serialize and create the new snapshot
	 6. Create the diff
	 7. Update the announced version
	*/

}
