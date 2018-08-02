package sources

import "github.com/stretchr/testify/mock"

type MockDataModel struct {
	mock.Mock
}

//GetDataName implements the DataModel interface
func (m MockDataModel) GetDataName() string {
	args := m.Called()
	return args.String(0)
}

//CacheDuration implements the DataModel Interface
func (m MockDataModel) CacheDuration() int64 {
	args := m.Called()
	return int64(args.Int(0))
}

func (m MockDataModel) NewBag() Bag {
	args := m.Called()

	var r0 Bag
	if rf, ok := args.Get(0).(func() Bag); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(Bag)
		}
	}

	return r0
}

//LoadAll implements the DataModel interface
func (m MockDataModel) LoadAll() (Bag, error) {

	args := m.Called()

	var r0 Bag
	if rf, ok := args.Get(0).(func() Bag); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(Bag)
		}
	}

	var r1 error

	if rf, ok := args.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
