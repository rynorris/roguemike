package api

type ClientAction interface {
    _implementsClientAction()
}

type MoveAction struct {
    X int
    Y int
}

func (action MoveAction) _implementsClientAction(){}
