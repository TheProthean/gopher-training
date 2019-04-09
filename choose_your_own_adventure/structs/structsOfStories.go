package structs

//Options ...
type Options struct {
	Text string
	Arc  string
}

//ArcStory ...
type ArcStory struct {
	Name       string
	Title      string
	Story      []string
	arcOptions Options
}
