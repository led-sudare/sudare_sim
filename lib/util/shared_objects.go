package util

type SharedByteData interface {
	AddByteData(id string, byteData ByteData)
	GetByteData(id string) ByteData
	EditByteData(id string, editableBlock func(editable ByteData) error) error
	RemoveSharedByteData(id string)
}

var sharedByteData SharedByteData

func AddSharedByteData(id string, byteData ByteData) {
	getSharedByteDataInstance().AddByteData(id, byteData)
}

func GetSharedByteData(id string) ImmutableByteData {
	return getSharedByteDataInstance().GetByteData(id)
}

func RemoveSharedByteData(id string) {
	getSharedByteDataInstance().RemoveSharedByteData(id)
}

func EditSharedByteData(id string, editableBlock func(editable ByteData) error) error {
	return getSharedByteDataInstance().EditByteData(id, editableBlock)
}

/**
private
*/

func getSharedByteDataInstance() SharedByteData {
	if sharedByteData == nil {
		sharedByteData = newSharedByteData()
	}
	return sharedByteData
}

func newSharedByteData() SharedByteData {
	return &sharedByteDataImpl{make(map[string]ByteData)}
}

type sharedByteDataImpl struct {
	byteDatas map[string]ByteData
}

func (o *sharedByteDataImpl) AddByteData(id string, byteData ByteData) {
	o.byteDatas[id] = byteData
}

func (o *sharedByteDataImpl) GetByteData(id string) ByteData {
	if i, ok := o.byteDatas[id]; !ok {
		return nil
	} else {
		return i
	}
}

func (o *sharedByteDataImpl) RemoveSharedByteData(id string) {
	delete(o.byteDatas, id)
}

func (o *sharedByteDataImpl) EditByteData(id string, editableBlock func(editable ByteData) error) error {
	return o.GetByteData(id).EditSafe(func(editable ByteData) error {
		if err := editableBlock(editable); err != nil {
			return err
		}
		return nil
	})
}
