package identities

import "fmt"

// NOTE: upstream PR #47 ("Add fido2/ctap2 + user auth") referenced a passkey
// "export" feature (ExportPasskeysArchive / ExportKeePassPasskey /
// PasskeyExportMetadata) but never actually included its implementation, so the
// PR did not compile. We don't use passkey export in this build (the goal is a
// CTAP2/PIN passkey gated by the macropad), so these are minimal stubs that
// satisfy the compiler and return a clear error if the `export` command is run.

type PasskeyExportMetadata struct {
	Exporter        string
	ExporterVersion string
}

func ExportPasskeysArchive(sources []CredentialSource, meta PasskeyExportMetadata) ([]byte, error) {
	return nil, fmt.Errorf("passkey export is not supported in this build")
}

func ExportKeePassPasskey(source CredentialSource) ([]byte, error) {
	return nil, fmt.Errorf("passkey export is not supported in this build")
}
