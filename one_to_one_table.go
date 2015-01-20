package cmagic 

import(
	"fmt"
	"errors"
)

type OneToOneT struct {
	Table
	idField string
}

func (o *OneToOneT) Update(id interface{}, m map[string]interface{}) error {
	return o.Table.Where(Eq(o.idField, id)).Update(m)
}

func (o *OneToOneT) Delete(id interface{}) error {
	return o.Table.Where(Eq(o.idField, id)).Delete()
}

func (o *OneToOneT) Read(id interface{}) (interface{}, error) {
	res, err := o.Table.Where(Eq(o.idField, id)).Query().Read()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New(fmt.Sprintf("Row with id %v not found", id))
	}
	return res[0], nil
}