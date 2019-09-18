// 用户微博数据
package factory

var (
	userIDs = []uint{1,2,3,4,5}
)

func blogFactory(i int) *factory.Factory {
	r := randomdata.Number(0,len(userIDs)-1)
	s := &blogModel.Blog{
		UserIDs: userIDs[r],
	} 

	return factory.NewFactory(
		s,
	).Attr("Content", func(args factory.Args) (interface{}, error) {
		return randomdata.Paragraph(), nil
	})
}