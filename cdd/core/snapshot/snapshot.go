package snapshot

type snapshot interface {
	GetVersion(keyName string) (string, error)

	UpdateVersion(keyName, newVersion string) error
}
