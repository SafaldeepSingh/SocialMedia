package Model

type Post struct {
	Id int
	UserId int
	Caption string
	Photos []string
}

func (p *Post) getId() int{
	return p.Id
}
func (p *Post) getUserId() int{
	return p.UserId
}
func (p *Post) getCaption() string{
	return p.Caption
}
func (p *Post) getPhotos() []string{
	return p.Photos
}
