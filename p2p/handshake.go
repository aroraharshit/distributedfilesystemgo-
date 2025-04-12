package p2p

type HandShakeFunc func(Peer) error

func NopHandShakeFunc(Peer) error {return nil}