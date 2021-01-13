package mpu6050

import (
	"github.com/d2r2/go-i2c"
)

const (
	WHO_AM_I byte = 0x75
)

// MPU6050 instance
type MPU6050 struct {
	device  *i2c.I2C
	address byte
	bus     int
}

// NewMPU6050 create a new MPU6050 instance
func NewMPU6050() *MPU6050 {
	return &MPU6050{
		address: 0x68,
		bus:     1,
	}
}

// Open opens the i2c connection with the MPU6050
// func (mpu *MPU6050) Open() error {
// 	var err error
// 	mpu.device, err = i2c.NewI2C(mpu.address, mpu.bus)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Close closes the i2c connection with the MPU6050
// func (mpu *MPU6050) Close() error {
// 	return nil
// }

// ret, err := i2c.ReadRegU8(WHO_AM_I)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(hex.EncodeToString([]byte{ret}))
