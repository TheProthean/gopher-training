package structs

//Options ...
type Options struct {
	Text string
	Arc  string
}

//ArcStory ...
type ArcStory struct {
	Name    string
	ArcData ArcStoryData
}

//ArcStoryData ...
type ArcStoryData struct {
	Title   string
	Story   []string
	Options []Options
}
