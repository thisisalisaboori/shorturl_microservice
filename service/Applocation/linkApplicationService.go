package applocation

import (
	//"fmt"
	_ "fmt"
	pb "shorturl/services"

	"github.com/labstack/gommon/random"
)

type ILinkApplicationService interface {
	SetLink(newLink pb.CreateRequest) pb.Response
	GetLink(shortLink string) string
}

func GetAppLicationService(config uint8) ILinkApplicationService {
	da := LinkApplicationService{}
	return &da
}

type LinkApplicationService struct {
	cn LinkContext
}

func makeHash() string {
	const ALPHABAT = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	return random.String(5, ALPHABAT)
}

func (T *LinkApplicationService) SetLink(newLink pb.CreateRequest) pb.Response {

	shortLink := makeHash()
	for {
		check := T.cn.Get(shortLink)
		if check == "" {
			break
		}
		shortLink = makeHash()
	}

	_newLink := Link{Link: newLink.Url, ShortLink: shortLink}
	
	T.cn.Set(_newLink)
	return pb.Response{Url: _newLink.Link, Shorturl: shortLink}
}

func (T *LinkApplicationService) GetLink(shortLink string) string {
	return T.cn.Get(shortLink)
}
