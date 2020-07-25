// 定义关于dictionary中的一些错误
package ch07_maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for.")
	ErrWordExists       = DictionaryErr("cannot add the word because it already exists.")
	ErrWordDoesNotExist = DictionaryErr("cannot update the word because it does not exist.")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
