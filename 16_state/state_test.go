package state

func ExampleWeek() {
	ctx := NewDayContext()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	ctx.Today()
	ctx.Next()
	// Output:
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
	// Saturday
	// Sunday
}
