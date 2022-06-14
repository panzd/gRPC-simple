package sample

import "pb"

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}
