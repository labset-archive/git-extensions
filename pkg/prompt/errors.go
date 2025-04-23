package prompt

import "errors"

var ErrPromptCancelled = errors.New("prompt cancelled")
var ErrPromptEmpty = errors.New("prompt empty")
