// 用户微博数据
package factory
import (
	"fmt"
	blogModel "go_simpleweibo/app/models/blog"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var (
	userIDs = []uint{1,2,3,4,5}
)

func blogFactory(i int) *factory.Factory {
	r := randomdata.Number(0,len(userIDs)-1)
	s := &blogModel.Blog{
		UserID: userIDs[r],
	} 

	return factory.NewFactory(
		s,
	).Attr("Content", func(args factory.Args) (interface{}, error) {
		return randomdata.Paragraph(), nil
	})
}

// blogTableSeeder
func blogTableSeeder(needCleanTable bool) {
	if needCleanTable {
		DropAndCreateTable(&blogModel.Blog{})
	}

	for i := 0; i < 100; i++ {
		blog := blogFactory(i).MustCreate().(*blogModel.Blog)
		if err := blog.Create(); err != nil {
			fmt.Println("mock status error: %v\n",err)
		}
	}
}