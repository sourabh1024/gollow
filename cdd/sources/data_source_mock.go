package sources

//type MockDataSource struct {
//	mock.Mock
//}
//
//func (m MockDataSource) NewEntity() data.Entity {
//	return &MockDataSource{}
//}
//
////GetNameSpace implements the data.Entity interface
//func (m MockDataSource) GetNameSpace() string {
//	args := m.Called()
//	return args.String(0)
//}
//
////GetPrimaryKey implements the data.Entity interface
//func (m MockDataSource) GetPrimaryKey() string {
//	args := m.Called()
//	return args.String(0)
//}
//
////GetDataName implements the data.Entity interface
//func (m MockDataSource) GetDataName() string {
//	args := m.Called()
//	return args.String(0)
//}
//
//func (m MockDataSource) NewDataRef() DataModel {
//	return &MockDataSource{}
//}
//
//func (m MockDataSource) CacheDuration() int64 {
//	args := m.Called()
//	return int64(args.Int(0))
//}
//
//func (m MockDataSource) LoadAll() ([]DataModel, error) {
//
//	args := m.Called()
//
//	var r0 []DataModel
//	if rf, ok := args.Get(0).(func() []DataModel); ok {
//		r0 = rf()
//	} else {
//		if args.Get(0) != nil {
//			r0 = args.Get(0).([]DataModel)
//		}
//	}
//
//	var r1 error
//
//	if rf, ok := args.Get(1).(func() error); ok {
//		r1 = rf()
//	} else {
//		r1 = args.Error(1)
//	}
//
//	return r0, r1
//}
