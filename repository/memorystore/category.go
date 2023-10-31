package memorystore

import (
	"app/entity"
)

type Category struct {
	categories []entity.Category
}

func (c Category) DoesThisUserThisCategoryID(userID, CategoryID int) bool {
	isFound := false
	for _, c := range c.categories {
		if c.ID == CategoryID && c.UserID == userID {
			isFound = true
			break
		}
	}
	return isFound
}
