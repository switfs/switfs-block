package service

const (
	AttoFIL = 18
	NanoFIL = 9
)

func NanoOrAttoToFIL(fil string, filtype int) string {
	//大于18or9位
	if len(fil) > filtype {
		str := fil[0:len(fil)-filtype] + "." + fil[len(fil)-filtype:]

		return str
	}
	//小于18or9位
	str := "0."
	for i := 0; i < filtype-len(fil); i++ {
		str += "0"
	}

	return str + fil
}
