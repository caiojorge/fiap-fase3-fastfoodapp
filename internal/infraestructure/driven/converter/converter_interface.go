package converter

type Converter[E any, M any] interface {
	FromEntity(entity *E) *M
	ToEntity(model *M) *E
}
