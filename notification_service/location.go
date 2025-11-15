package notification_service

type Location interface {
	Send(m Message) error
}

type ParseLocationFunc func(settings []byte) (Location, error)

var KNOWN_LOCATIONS = map[string]ParseLocationFunc{}

func RegisterLocation(name string, loc ParseLocationFunc) {
	KNOWN_LOCATIONS[name] = loc
}
