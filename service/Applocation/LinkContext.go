package applocation

import (
	//"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type IContext interface {
	Init() error
	Set(link Link) bool
	Get(link string) string
}
type Link struct {
	Link      string
	ShortLink string
}

type LinkContext struct {
	LDB *leveldb.DB
}

func (T *LinkContext) Init() error {
	var err error
	T.LDB, err = leveldb.OpenFile("/tmp/links", nil)
	return err
}

func (T *LinkContext) Set(link Link) bool {
	//fmt.Println(T.LDB)
	T.Init()
	err := T.LDB.Put([]byte(link.ShortLink), []byte(link.Link), nil)
	//fmt.Println(err)
	defer T.LDB.Close()
	return err == nil
}

func (T *LinkContext) Get(link string) string {
	T.Init()
	data, _ := T.LDB.Get([]byte(link), nil)
	defer T.LDB.Close()
	//fmt.Printf("leveldata=%s\n", string(data))
	return string(data)
}
