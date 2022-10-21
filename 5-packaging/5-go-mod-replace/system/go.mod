module github.com/dmcardoso/go-expert/5-packaging/5-go-mod-replace/system

go 1.18

replace github.com/dmcardoso/go-expert/5-packaging/5-go-mod-replace/math => ../math

require github.com/dmcardoso/go-expert/5-packaging/5-go-mod-replace/math v0.0.0-00010101000000-000000000000
