module github.com/jonilsonds9/goexpert-modulo-5-packaging/3-trabalhando-com-go-mod-replace/sistema

go 1.24.1

replace github.com/jonilsonds9/goexpert-modulo-5-packaging/3-trabalhando-com-go-mod-replace/math => ../math

require github.com/jonilsonds9/goexpert-modulo-5-packaging/3-trabalhando-com-go-mod-replace/math v0.0.0-00010101000000-000000000000
