package experiments

type TestDto struct {
	ID      int64
	Name    string
	Email   string
	Country string
}

func (t *TestDto) ToPB() Message {
	return &Test{
		Id:    t.ID,
		Name:  t.Name,
		Email: t.Email,
	}
}

func (T *TestDto) LoadAll() Bag {

	t1 := &TestDto{
		ID: 1,
	}

	t2 := &TestDto{
		ID: 2,
	}

	bags := &TestBag{}

	bags.AddEntry(t1.ToPB())
	bags.AddEntry(t2.ToPB())

	return bags
}
