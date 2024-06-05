CREATE TABLE IF NOT EXISTS insurance_companies (
    insurance_company_id INT NOT NULL AUTO_INCREMENT,
    insurance_company_name VARCHAR(255) NOT NULL,
    PRIMARY KEY(insurance_company_id)
);

CREATE TABLE IF NOT EXISTS insurance_plans (
    insurance_plan_id INT NOT NULL AUTO_INCREMENT,
    insurance_plan_name VARCHAR(255) NOT NULL,
    company_id INT NOT NULL,
    PRIMARY KEY(insurance_plan_id),
    FOREIGN KEY(company_id) REFERENCES insurance_companies(insurance_company_id)
);

CREATE TABLE IF NOT EXISTS procedures (
    procedure_id VARCHAR(255) NOT NULL,
    procedure_code VARCHAR(255) NOT NULL,
    procedure_desc VARCHAR(255) NOT NULL,
    PRIMARY KEY(procedure_id)
);

CREATE TABLE IF NOT EXISTS hospital_groups (
    group_id INT NOT NULL AUTO_INCREMENT,
    procedure_code VARCHAR(255) NOT NULL,
    procedure_desc VARCHAR(255) NOT NULL,
    PRIMARY KEY(group_id)
);

CREATE TABLE IF NOT EXISTS hospitals (
    hospital_id INT NOT NULL AUTO_INCREMENT,
    hospital_name VARCHAR(255) NOT NULL,
    hospital_address VARCHAR(255) NOT NULL,
    hospital_zip_code VARCHAR(32) NOT NULL,
    hospital_group_id INT NOT NULL,
    PRIMARY KEY(hospital_id),
    FOREIGN KEY(hospital_group_id) REFERENCES hospital_groups(group_id)
);

CREATE TABLE IF NOT EXISTS base_costs (
    procedure_id VARCHAR(255) NOT NULL,
    hospital_id INT NOT NULL,
    gross_charge INT,
    discounted_cash_price INT,
    PRIMARY KEY(procedure_id, hospital_id),
    FOREIGN KEY(procedure_id) REFERENCES procedures(procedure_id),
    FOREIGN KEY(hospital_id) REFERENCES hospitals(hospital_id)
);

CREATE TABLE IF NOT EXISTS costs (
    procedure_id VARCHAR(255) NOT NULL,
    hospital_id INT NOT NULL,
    insurance_plan_id INT NOT NULL,
    rate INT,
    PRIMARY KEY(procedure_id, hospital_id, insurance_plan_id),
    FOREIGN KEY(procedure_id) REFERENCES procedures(procedure_id),
    FOREIGN KEY(hospital_id) REFERENCES hospitals(hospital_id),
    FOREIGN KEY(insurance_plan_id) REFERENCES insurance_plans(insurance_plan_id)
);




