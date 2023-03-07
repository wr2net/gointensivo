package entity

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {

}

func (o *Order) Validate() error {
	if o.ID == "" {
		return error.New("Id is required")
	}
	if o.Price <= 0 {
		return error.New("Invalid price")
	}
	if o.Price <= 0 {
		return error.New("Invalid tax")
	}
	return nil
}
