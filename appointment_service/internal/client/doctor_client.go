package client

type DoctorClient interface {
	Exists(doctorID string) (bool, error)
}
