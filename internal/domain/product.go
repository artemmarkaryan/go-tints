package domain

import (
	"context"
)

//type CategoryDto struct {
//	ID           int             `json:"id"`
//	Name         string          `json:"name"`
//	TranslitName string          `json:"translit"`
//	Sku          []SkuPreviewDto `json:"sku"`
//}
//
//type SkuPreviewDto struct {
//	ID          int     `json:"id"`
//	ProductID   int     `json:"productId"`
//	CategoryID  int     `json:"categoryId"`
//	Name        string  `json:"name"`
//	TanslitName string  `json:"translit"`
//	VendorCode  string  `json:"vendorCode"`
//	OldPrice    float32 `json:"oldPrice"`
//	Price       float32 `json:"price"`
//	Image       string  `json:"image"`
//	New         bool    `json:"new"`
//	Top         bool    `json:"top"`
//}
const query = `
select json_build_object(
               'id', cat.id,
               'name', cat.name,
               'translit', cat.translit_name,
               'sku', json_agg(json_build_object(
                'id', sku.id,
                'productId', prod.id,
                'name', prod.name,
                'translit', prod.translit_name,
                'vendorCode', sku.vendor_code,
                'oldPrice', sku.old_price,
                'price', sku.price,
                'image', image.image,
                'new', prod.new,
                'top', prod.top,
                'shadeId', sku.shade_id
            ))
           )

from app_sku sku
         left join
     app_product prod on prod.id = sku.product_id
         left join
     app_category cat on cat.id = prod.category_id
         left join
     app_skuimage image on sku.id = image.sku_id
group by cat.id;

select cat.id, cat.name, cat.translit_name
from app_category cat
`

func GetProducts(ctx context.Context) (resultJson string, err error) {
	db := GetDB()
	defer func() { _ = db.Close() }()
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(&resultJson)
	}
	return
}
