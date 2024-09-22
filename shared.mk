define get_remote_secret
	kubectl get secret -n remote-dev $(1) -o json | jq -r '.data | to_entries[] | "\(.key)=\(.value | @base64d)"' > .env.generated
endef
