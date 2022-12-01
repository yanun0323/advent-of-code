package service

var _input = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func (su *ServiceSuite) Test_Day1A() {
	ans, ok := su.Day1A(_input).(int)
	su.True(ok)
	su.Equal(24000, ans)
}

func (su *ServiceSuite) Test_Day1B() {
	ans, ok := su.Day1B(_input).(int)
	su.True(ok)
	su.Equal(45000, ans)
}
