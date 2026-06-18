package main

import "github.com/bulwarkid/virtual-fido/fido_client"

// NOTE: upstream PR #47 referenced platform fingerprint user-verification
// (platformSupportsUserVerification / platformVerifyUser) and
// fingerprintAvailabilityHints but never implemented them, so the PR did not
// compile. This build does user verification via CTAP2 ClientPIN (handled in the
// ctap layer) and user presence via the macropad APPROVE key, so built-in
// (fingerprint) UV is simply unavailable. These stubs report that.

func (support *ClientSupport) platformSupportsUserVerification() bool {
	return false
}

func (support *ClientSupport) platformVerifyUser(action fido_client.ClientAction, params fido_client.ClientActionRequestParams) bool {
	return false
}

func fingerprintAvailabilityHints() []string {
	return []string{"Fingerprint verification is not available in this build (use PIN)."}
}
