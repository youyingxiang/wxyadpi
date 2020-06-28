/**
 * @Author: youxingxiang
 * @Description:
 * @File:  MaterialStore
 * @Version: 1.0.0
 * @Date: 2020-06-28 09:40
 */
package serializer

type MaterialStore struct {
	OrderItemNo  string  `json:"order_item_no"`
	ShouldNumber float64 `json:"should_number"`
	ActualNumber float64 `json:"actual_number"`
	StoreName    string  `json:"store_name"`
	StoreId      int     `json:"store_id"`
	MaterialName string  `json:"material_name"`
}
