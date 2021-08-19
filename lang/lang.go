package lang

import (
    "golang-master/constants"
)

type Tags map[string]Texts
type Texts struct {
	texts map[string]string
}

var message = map[string]Texts {"en":en_texts}

func Get(msgkey string) string{
    messages := (message)[constants.CURREBTLANG]
    if message, ok := messages.texts[msgkey]; ok {
        return message
    }
    return msgkey
}


