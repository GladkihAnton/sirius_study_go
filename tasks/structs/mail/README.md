Дан набор классов, описывающих работу гипотетической почтовой системы.

Для начала рассмотрим код, описывающий все используемые сущности.

```go
// Интерфейс: сущность, которую можно отправить по почте.
// У такой сущности можно получить от кого и кому направляется письмо.
type SendableI interface {
	GetFrom() string
	GetTo() string
}


//Абстрактный класс,который позволяет абстрагировать логику хранения
//источника и получателя письма в соответствующих полях класса.
type abstractSendable struct {
	from string
	to   string
}

func (sendable *abstractSendable) GetFrom() string {
	return sendable.from
}

func (sendable *abstractSendable) GetTo() string {
	return sendable.to
}
```

У abstractSendable есть два наследника:

Первый класс описывает обычное письмо, в котором находится только текстовое сообщение.

```go
// Письмо, у которого есть текст, который можно получить с помощью метода `getMessage`
type Message struct {
    message string

    abstractSendable
}

func (message *Message) GetMessage() string {
    return message.message
}
```

Второй класс описывает почтовую посылку:

```go
// Посылка, содержимое которой можно получить с помощью метода `getContent`
type Package struct {
	content Content

	abstractSendable
}

func (pkg *Package) GetContent() Content {
	return pkg.content
}

```

При этом сама посылка описывается следующим классом:

```go
// Класс, который задает посылку. У посылки есть текстовое описание содержимого и целочисленная ценность.
type Content struct {
    content string
    price   int
}

func (content *Content) GetContent() string {
    return content.content
}

func (content *Content) GetPrice() int {
    return content.price
}
```

Теперь рассмотрим классы, которые моделируют работу почтового сервиса:

```go
// Интерфейс, который задает класс, который может каким-либо образом обработать почтовый объект.
type ServiceI interface {
	ProcessMail(sendable SendableI) SendableI
}

// Класс, в котором скрыта логика настоящей почты
type RealService struct{}

func (service RealService) processMail(mail SendableI) SendableI {
    return mail
}
```

Вам необходимо описать набор классов, каждый из которых является ServiceI:

1) UntrustworthyMailWorker – класс, моделирующий ненадежного работника почты, который вместо того, 
    чтобы передать почтовый объект непосредственно в сервис почты, последовательно передает этот объект набору третьих лиц,
    а затем, в конце концов, передает получившийся объект непосредственно экземпляру RealMailService.
    У UntrustworthyMailWorker должен быть конструктор от массива ServiceI
    ( результат вызова processMail первого элемента массива передается на вход processMail второго элемента, и т. д.) 
    и метод GetRealMailService, который возвращает ссылку на внутренний экземпляр RealMailService.

2) Spy – шпион, который логгирует о всей почтовой переписке, которая проходит через его руки. 
    Он следит только за объектами класса Message и пишет в логгер следующие сообщения  
    (в выражениях нужно заменить части в фигурных скобках на значения полей почты):
   2.1) Если в качестве отправителя или получателя указан "Austin Powers", 
        то нужно написать в лог сообщение с уровнем WARN: Detected target mail correspondence: from {from} to {to} "{message}"
   2.2) Иначе, необходимо написать в лог сообщение с уровнем INFO: Usual correspondence: from {from} to {to}

3) Thief – вор, который ворует самые ценные посылки и игнорирует все остальное. 
   Вор принимает в конструкторе переменную int – минимальную стоимость посылки, которую он будет воровать. 
   Также, в данном классе должен присутствовать метод GetStolenValue, который возвращает суммарную стоимость всех посылок, которые он своровал.
   Воровство происходит следующим образом: вместо посылки, которая пришла вору, он отдает новую, такую же,
   только с нулевой ценностью и содержимым посылки "stones instead of {content}".

4) Inspector – Инспектор, который следит за запрещенными и украденными посылками и бьет тревогу в виде исключения,
   если была обнаружена подобная посылка. Если он заметил запрещенную посылку с одним из запрещенных содержимым ("weapons" и "banned substance"), 
   то логгирует сообщение "Warn: The package is dangerous"
   Если он находит посылку, состоящую из камней (содержит слово "stones"), то логгирует сообщение "Warn: The package was stolen"



T: - (но какой-то пример юзаджа хочется)
B: -