package validator

type AssertTrue interface {
	IsTrue(value string) bool
}
