package success

import "gohook/utils"

type ThenSuccess struct {
	IsError error
}

func (then *ThenSuccess) ShowSuccess(msg string) {
	if then.IsError == nil {
		utils.InfoShow(msg)
	}
}
