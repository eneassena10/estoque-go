package entities

type Product struct {
	ID    int
	Name  string
	Price float64
	Count int
}

type ProductRequest struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty" binding:"required"`
	Price float64 `json:"price,omitempty" binding:"required"`
	Count int     `json:"count,omitempty" binding:"required"`
}

func NewProduct() *ProductRequest {
	return &ProductRequest{}
}

func (p *ProductRequest) WithID(id int) *ProductRequest {
	if id != 0 {
		p.ID = id
	}
	return p
}

func (p *ProductRequest) WithName(name string) *ProductRequest {
	if name != "" {
		p.Name = name
	}
	return p
}

func (p *ProductRequest) WithPrice(price float64) *ProductRequest {
	if price != 0 {
		p.Price = price
	}
	return p
}

func (p *ProductRequest) WithCount(count int) *ProductRequest {
	if count != 0 {
		p.Count = count
	}
	return p
}
