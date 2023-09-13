package collectors

type Nginx_log struct {
	Active_connections string
	Server_accepts     string
	Server_handled     string
	Server_requests    string
	Reading            string
	Writing            string
	Waiting            string
}
