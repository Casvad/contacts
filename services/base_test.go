package services

import (
	"gorm.io/gorm"
)

func CleanDatabase(gorm *gorm.DB) {

	gorm.Exec(`
		CREATE
		OR REPLACE FUNCTION truncate_tables() RETURNS void AS $$
		DECLARE
		statements CURSOR FOR
		SELECT tablename, schemaname
		FROM pg_tables
		WHERE schemaname in ('public')
		  and tablename not in (
								'schema_migrations'
								);
		BEGIN
		FOR stmt IN statements LOOP
				EXECUTE 'Truncate TABLE ' || quote_ident(stmt.schemaname ) || '.' || quote_ident(stmt.tablename) || ' CASCADE;';
		END LOOP;
		END;
		$$
		LANGUAGE plpgsql;
		
		select truncate_tables();
	`)
}
