package domain

import "titiktopup-core/pb"

func ToProtoCategory(c Category) *pb.Category {
	var products []*pb.Product
	for _, p := range c.Products {
		products = append(products, ToProtoProduct(p))
	}

	return &pb.Category{
		Id:       uint32(c.ID),
		Name:     c.Name,
		Slug:     c.Slug,
		Products: products,
	}
}

func ToProtoProduct(p Product) *pb.Product {
	return &pb.Product{
		Id:        uint32(p.ID),
		Sku:       p.SKU,
		Name:      p.Name,
		PriceSell: p.PriceSell,
	}
}
