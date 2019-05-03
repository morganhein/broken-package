package broken

type Ok interface {
	MyFunc()
}

func GetOK() Ok {
	return &notok{}
}

type notok struct {}