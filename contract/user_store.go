package contract

import "app/entity"

type UserWriteStore interface {
	Save(u entity.User)
}
type UserReadStore interface {
	Load(serializationMode string) []entity.User
}
