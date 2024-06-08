package jobs

type OnStartUpJob interface {
	Execute()
}

func ProvideOnStartUpJob(
	migrationJob MigrationJob,
) []OnStartUpJob {

	return []OnStartUpJob{&migrationJob}
}
