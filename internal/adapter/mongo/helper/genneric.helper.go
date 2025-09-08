package helper

type InsertMetaSetter interface{ SetInsertMeta() }

type UpdateMetaSetter interface{ SetUpdateMeta() }

type DeleteMetaSetter interface{ SetDeleteMeta() }

func InitInsert[T InsertMetaSetter](m T) T {
	m.SetInsertMeta()
	return m
}

func TouchUpdate[T UpdateMetaSetter](m T) T {
	m.SetUpdateMeta()
	return m
}

func MarkDeleted[T DeleteMetaSetter](m T) T {
	m.SetDeleteMeta()
	return m
}
