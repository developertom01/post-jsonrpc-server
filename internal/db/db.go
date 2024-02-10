package db

type (
	Database interface {
		GetClient() any
		GetDatabase() any
		Disconnect()

		//User queries
		UserQueries
	}
)
