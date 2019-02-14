package session

type Session struct {
	UserInfo         map[string]string
	SomeThingYouNeed map[string]interface{}
}

func NewSession() *Session {
	session := &Session{
		UserInfo:         make(map[string]string),
		SomeThingYouNeed: make(map[string]interface{}),
	}
	return session
}

func (this *Session) Set(key string, value interface{}) {
	this.SomeThingYouNeed[key] = value
}

//Get方法返回interface是需要断言的(我记得我在前面一两章说过用法)
//或者你也在Set里直接存Json然后用utils.GetJsonStruct(object)
//写法随便你，按你的需求来定
func (this *Session) Get(key string) interface{} {

	if _, ok := this.SomeThingYouNeed[key]; ok {

		return this.SomeThingYouNeed[key]

	} else {
		return nil
	}
}
