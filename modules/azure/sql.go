package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/preview/sql/mgmt/sql"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetSQLServerClient is a helper function that will setup a sql server client
func GetSQLServerClient(subscriptionID string) (*sql.ServersClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a sql server client
	sqlClient := sql.NewServersClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	sqlClient.Authorizer = *authorizer

	return &sqlClient, nil
}

// GetSQLServerID is a helper function that gets the sql server ID
func GetSQLServerID(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) string {
	id, err := GetSQLServerIDE(t, subscriptionID, resGroupName, serverName)
	require.NoError(t, err)

	return id
}

// GetSQLServerIDE is a helper function that gets the sql server ID
func GetSQLServerIDE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string) (string, error) {
	// Create a SQl Server client
	sqlClient, err := GetSQLServerClient(subscriptionID)
	if err != nil {
		return "", err
	}

	//Get the corresponding server client
	sqlServer, err := sqlClient.Get(context.Background(), resGroupName, serverName)
	if err != nil {
		return "", err
	}

	//Return server ID
	return *sqlServer.ID, nil
}

// GetSQLServerFullDomainName is a helper function that gets the sql server full domain name
func GetSQLServerFullDomainName(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) string {
	fullDomainName, err := GetSQLServerFullDomainNameE(t, subscriptionID, resGroupName, serverName)
	require.NoError(t, err)

	return fullDomainName
}

// GetSQLServerFullDomainNameE is a helper function that gets the sql server full domain name
func GetSQLServerFullDomainNameE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string) (string, error) {
	// Create a SQl Server client
	sqlClient, err := GetSQLServerClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the corresponding server client
	sqlServer, err := sqlClient.Get(context.Background(), resGroupName, serverName)
	if err != nil {
		return "", err
	}

	//Return server full domain name
	return *sqlServer.FullyQualifiedDomainName, nil
}

// GetSQLServerState is a helper function that gets the sql server state
func GetSQLServerState(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) sql.ServerState {
	serverState, err := GetSQLServerStateE(t, subscriptionID, resGroupName, serverName)
	require.NoError(t, err)

	return serverState
}

// GetSQLServerStateE is a helper function that gets the sql server state
func GetSQLServerStateE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string) (sql.ServerState, error) {
	// Create a SQl Server client
	sqlClient, err := GetSQLServerClient(subscriptionID)
	if err != nil {
		return "", err
	}

	// Get the corresponding server client
	sqlServer, err := sqlClient.Get(context.Background(), resGroupName, serverName)
	if err != nil {
		return "", err
	}

	//Return server state
	return sqlServer.State, nil
}

// GetDatabaseClient  is a helper function that will setup a sql DB client
func GetDatabaseClient(subscriptionID string) (*sql.DatabasesClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a sql DB client
	sqlDBClient := sql.NewDatabasesClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	sqlDBClient.Authorizer = *authorizer

	return &sqlDBClient, nil
}

//ListSQLServerDatabases is a helper function that gets a list of databases on a sql server
func ListSQLServerDatabases(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) *[]sql.Database {
	dbList, err := ListSQLServerDatabasesE(t, resGroupName, serverName, subscriptionID)
	require.NoError(t, err)

	return dbList
}

//ListSQLServerDatabasesE is a helper function that gets a list of databases on a sql server
func ListSQLServerDatabasesE(t testing.TestingT, resGroupName string, serverName string, subscriptionID string) (*[]sql.Database, error) {
	// Create a SQl db client
	sqlClient, err := GetDatabaseClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	//Get the corresponding DB client
	sqlDbs, err := sqlClient.ListByServer(context.Background(), resGroupName, serverName, "", "")
	if err != nil {
		return nil, err
	}

	// Return DB ID
	return sqlDbs.Value, nil
}

// GetDatabaseID is a helper function that gets the sql db id
func GetDatabaseID(t testing.TestingT, resGroupName string, serverName string, dbName string, subscriptionID string) string {
	dbID, err := GetDatabaseIDE(t, subscriptionID, resGroupName, serverName, dbName)
	require.NoError(t, err)

	return dbID
}

// GetDatabaseIDE is a helper function that gets the sql db id
func GetDatabaseIDE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string, dbName string) (string, error) {
	// Create a SQl db client
	sqlClient, err := GetDatabaseClient(subscriptionID)
	if err != nil {
		return "", err
	}

	//Get the corresponding DB client
	sqlDb, err := sqlClient.Get(context.Background(), resGroupName, serverName, dbName, "")
	if err != nil {
		return "", err
	}

	// Return DB ID
	return *sqlDb.ID, nil
}

// GetDatabaseStatus is a helper function that gets the sql db state
func GetDatabaseStatus(t testing.TestingT, resGroupName string, serverName string, dbName string, subscriptionID string) string {
	dbStatus, err := GetDatabaseStatusE(t, subscriptionID, resGroupName, serverName, dbName)
	require.NoError(t, err)

	return dbStatus
}

// GetDatabaseStatusE is a helper function that gets the sql db state
func GetDatabaseStatusE(t testing.TestingT, subscriptionID string, resGroupName string, serverName string, dbName string) (string, error) {
	// Create a SQl db client
	sqlClient, err := GetDatabaseClient(subscriptionID)
	if err != nil {
		return "", err
	}

	//Get corresponding DB client
	sqlDb, err := sqlClient.Get(context.Background(), resGroupName, serverName, dbName, "")
	if err != nil {
		return "", err
	}

	//Return DB status
	return *sqlDb.Status, nil
}
