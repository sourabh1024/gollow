package producer

//import "gollow/logging"
//
//// DataModelsManager manages the DataModels
//type DataProducerManager interface {
//	Add(producer DataProducer)
//	GetAllDataModels() map[string][]DataProducer
//	GetDMInNameSpace(namespace string) ([]DataProducer, bool)
//}
//
//// DataModels stores all the data Models inside a given namespace
//type DataProducers struct {
//	Models map[string][]DataProducer
//}
//
//func (d *DataProducers) Add(model DataProducer, namespace string) {
//	d.Models[namespace] = append(d.Models[namespace], model)
//}
//
//func (d *DataProducers) GetAllDataModels() map[string][]DataProducer {
//	return d.Models
//}
//
//func (d *DataProducers) GetDMInNameSpace(namespace string) ([]DataProducer, bool) {
//
//	dataModels, ok := d.Models[namespace]
//
//	if !ok {
//		logging.GetLogger().Error("No data producer  under namespace : ", namespace)
//		return nil, ok
//	}
//
//	return dataModels, ok
//}
