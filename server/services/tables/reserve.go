package tables

import "auth-server/interfaces/repository"

type ReserveService interface {
	Create(mailAddress string) (int64, error)
}

type reserveServiceImpl struct {
	r repository.ReserveRepository
}

func NewReserveService(r repository.ReserveRepository) *reserveServiceImpl {
	return &reserveServiceImpl{r: r}
}

func (r *reserveServiceImpl) Create(mailAddress string) (int64, error) {
	id, err := r.r.Create(mailAddress)
	if err != nil {
		return 0, err
	}
	return id, nil
}
