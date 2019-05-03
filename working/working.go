package working

type Ok interface {
	MyFunc()
}

func GetOK() Ok {
	return &notnotok{}
}

type notnotok struct {}

func (notnotok) MyFunc() {}
