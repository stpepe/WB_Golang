package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type StorageService struct{
}

func (s *StorageService) PackOrder(){
}

type CourierService struct {
}

func (c *CourierService) PickUpOrder(){
}

func (c *CourierService) Delivery(){
}

func (c *CourierService) GiveAnOrder(){
}

type TrackerService struct {
}

func (t *TrackerService) StartTreckingOrder(){
}

func (t *TrackerService) FinishTreckingOrder(){
}

type Facade struct{
	*StorageService
	*CourierService
	*TrackerService
}

func CreateFacade() *Facade {
	return &Facade{
		StorageService: &StorageService{},
		CourierService: &CourierService{},
		TrackerService: &TrackerService{},
	}
}

func (f *Facade) Delivery(){
	f.TrackerService.StartTreckingOrder()
	f.StorageService.PackOrder()
	f.CourierService.PickUpOrder()
	f.CourierService.Delivery()
	f.CourierService.GiveAnOrder()
	f.TrackerService.FinishTreckingOrder()
}

func main (){

}