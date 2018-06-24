package sources

// DataSource is the interface every datamodel needs to implement
// to be fetched by source data retriever
type DataSource interface {
	LoadAll()
}

type dataSourceImpl struct {
}

func NewDataSource() *dataSourceImpl {
	return &dataSourceImpl{}
}
