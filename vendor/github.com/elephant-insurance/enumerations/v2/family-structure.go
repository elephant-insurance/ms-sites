package enumerations

// FamilyStructure captures the four boolean properties describing family structure.
// It is meant to be included as an anonymous member of a parent struct, enabling
// us to treat these as direct properties of the parent struct
type FamilyStructure struct {
	HasLittleOnes  *bool
	HasPreTeens    *bool
	HasTeens       *bool
	HasYoungAdults *bool
}

const (
	EnumFamilyHasLittleOnes  = 1
	EnumFamilyHasPreTeens    = 2
	EnumFamilyHasTeens       = 4
	EnumFamilyHasYoungAdults = 8
)

func (fs *FamilyStructure) ToInt() int {
	rtn := 0
	if fs == nil {
		return rtn
	}
	if fs.HasLittleOnes != nil && *fs.HasLittleOnes {
		rtn += EnumFamilyHasLittleOnes
	}
	if fs.HasPreTeens != nil && *fs.HasPreTeens {
		rtn += EnumFamilyHasPreTeens
	}
	if fs.HasTeens != nil && *fs.HasTeens {
		rtn += EnumFamilyHasTeens
	}
	if fs.HasYoungAdults != nil && *fs.HasYoungAdults {
		rtn += EnumFamilyHasYoungAdults
	}

	return rtn
}

func (fs *FamilyStructure) FromInt(i int) {
	if fs == nil {
		return
	}

	fs.HasLittleOnes = ptrTo((i & EnumFamilyHasLittleOnes) > 0)
	fs.HasPreTeens = ptrTo((i & EnumFamilyHasPreTeens) > 0)
	fs.HasTeens = ptrTo((i&EnumFamilyHasTeens > 0))
	fs.HasYoungAdults = ptrTo((i&EnumFamilyHasYoungAdults > 0))
}

func ptrTo(b bool) *bool {
	rtn := b
	return &rtn
}
