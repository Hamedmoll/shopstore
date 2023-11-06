package organize

import "shopstoretest/entity"

//error always nil

func OrganizeBasketItem(items []entity.BasketItem) (map[uint]entity.BasketItem, error) {
	organizeItems := make(map[uint]entity.BasketItem)

	for _, item := range items {
		_, ok := organizeItems[item.ProductID]
		if !ok {
			organizeItems[item.ProductID] = item
		} else {
			organizeItem, _ := organizeItems[item.ProductID]
			organizeItem.Count += item.Count
			organizeItems[item.ProductID] = organizeItem
		}
	}

	return organizeItems, nil
}
