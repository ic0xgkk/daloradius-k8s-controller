package secret

import v1 "k8s.io/api/core/v1"

func PatchUsername(sr *v1.Secret, username string) {
	sr.Data["username"] = []byte(username)
}

func PatchPassword(sr *v1.Secret, password string) {
	sr.Data["password"] = []byte(password)
}
