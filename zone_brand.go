package libzone

import (
	"strconv"
)

type brand byte

const (
	Ipkg brand = iota
	Lipkg
	Sparse
	Pkgsrc
	Lx
	Bhyve
	Kvm
	Illumos
	Emu
	S10
)

// Function to return the human readable version of a brand
func (b brand) String() string {
	//Convert brand byte to string
	switch b {
	case Ipkg:
		return "Ipkg"
	case Lipkg:
		return "Lipkg"
	case Sparse:
		return "Sparse"
	case Pkgsrc:
		return "Pkgsrc"
	case Lx:
		return "Lx"
	case Bhyve:
		return "Bhyve"
	case Kvm:
		return "Kvm"
	case Illumos:
		return "Illumos"
	case Emu:
		return "Emu"
	case S10:
		return "S10"
	}

	return "Unknown: " + strconv.Itoa(int(b))
}

func (b brand) Ipkg() brand {
	return Ipkg
}

func (b brand) Lipkg() brand {
	return Lipkg
}

func (b brand) Sparse() brand {
	return Sparse
}

func (b brand) Pkgsrc() brand {
	return Pkgsrc
}

func (b brand) Lx() brand {
	return Lx
}

func (b brand) Bhyve() brand {
	return Bhyve
}

func (b brand) Kvm() brand {
	return Kvm
}

func (b brand) Illumos() brand {
	return Illumos
}

func (b brand) Emu() brand {
	return Emu
}

func (b brand) S10() brand {
	return S10
}

const Brand brand = 255
