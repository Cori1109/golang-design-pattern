package builder

//Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
	GetResult() string
}

type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
