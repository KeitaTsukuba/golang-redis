// sort/sort.go
package sort
 
import (
	"golang-redis/models"
	"sort"
)
 
func RankingSort(userList []models.User) []models.User {
	sort.Slice(
		userList,
		func(i, j int) bool {
			return userList[i].AccountID > userList[j].AccountID
		},
	)
 
	return userList
}