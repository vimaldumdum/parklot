package models

type Vehicle struct {
	vehicle_type string
	reg_no       string
	color        string
}

func NewVehicle(vehicle_type, reg_no, color string) *Vehicle {
	return &Vehicle{
		vehicle_type: vehicle_type,
		reg_no:       reg_no,
		color:        color,
	}
}

//getters
func (v *Vehicle) GetType() string {
	return v.vehicle_type
}

func (v *Vehicle) GetReg() string {
	return v.reg_no
}

func (v *Vehicle) GetColor() string {
	return v.color
}
