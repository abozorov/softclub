package main

type Account struct {
    Login    string
    Password string
    Active   bool
}

var accounts = make(map[string]Account, 10)

func main() {

	data := []Account{
		{Login: "admin", Password: "root_password", Active: true},
		{Login: "j_smith", Password: "secure_pass_2024", Active: true},
		{Login: "alice_w", Password: "wonderland_99", Active: true},
		{Login: "bob_builder", Password: "can_we_fix_it", Active: false}, // Заблокирован
		{Login: "dev_user", Password: "debug_mode_on", Active: true},
		{Login: "test_account", Password: "qwerty12345", Active: false}, // Отключен
		{Login: "manager_01", Password: "mng_password!", Active: true},
		{Login: "support_tech", Password: "help_desk_pass", Active: true},
		{Login: "guest_user", Password: "no_password", Active: true},
		{Login: "security_officer", Password: "ultra_secure_hash", Active: true},
	}
	for _, acc := range data {
		accounts[acc.Login] = acc
	}

	
}