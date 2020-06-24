/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store_order_summary
 * @Version: 1.0.0
 * @Date: 2020-06-24 09:52
 */
package serializer

type StoreOrderSummary struct {
	//MaterialId   int    `json:"material_id"`
	StoreNum      int    `json:"store_num"`
	SumNum        int    `json:"sum_num"`
	MaterialName  string `json:"material_name"`
	Img           string `json:"img"`
	MaterialSpecs string `json:"material_specs"`
	MaterialBrand string `json:"material_brand"`
}
