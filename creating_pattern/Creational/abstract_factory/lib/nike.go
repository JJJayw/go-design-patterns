package lib

type Nike struct {
}

var _ ISportsFactory = (*Nike)(nil)

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
