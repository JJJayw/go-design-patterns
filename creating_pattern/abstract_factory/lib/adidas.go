package lib

type Adidas struct {
}

var _ ISportsFactory = (*Adidas)(nil)

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
