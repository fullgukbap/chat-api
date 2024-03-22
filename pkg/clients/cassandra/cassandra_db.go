package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

// ConnectDatabase 함수는 새 연결된 클러스터 및 *gocql.Session을 반환합니다.
func ConnectDatabase(url string, keyspace string) *gocql.Session {
	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		logrus.Fatal(err)
	}

	return session
}
