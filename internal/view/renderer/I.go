package renderer

type I interface {
	RenderBrief(*Brief)
	RenderDetail(*Detail)
	Reset()
}
