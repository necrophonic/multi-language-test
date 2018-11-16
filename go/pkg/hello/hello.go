// Package hello provides functions for returning "hello" or "hi" (or a variant thereof)
// in a variety of languages.
// Each function is named in it's own language.
package hello

// English returns "hello" in English
func English() string {
	return "hello"
}

// Français returns "hello" in French
func Français() string {
	return "Salut"
}

// Deutsch returns "hello" in German
func Deutsch() string {
	return "hallo"
}

// 中文 returns "hello" in Chinese Mandarin
func 中文() string {
	return "你好"
}
