module example.com/main
go 1.16

replace example.com/sub => ./pkg

replace example.com/follow => ./pkg/follow

replace example.com/integration => ./test/integration

require (
	example.com/integration v0.0.0-00010101000000-000000000000
	example.com/sub v0.0.0-00010101000000-000000000000
)
