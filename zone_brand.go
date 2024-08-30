package libzone

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
