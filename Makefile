build-server:
	mysql -h "${SQL_SERVER}" -u "root" "hospital_costs" < "./database/costs.sql"