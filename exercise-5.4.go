/*
You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth The coins will be distributed based on the vowels contained in each name where:

a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

and a user can't get more than 10 coins. Print a map with each user's name and the amount of coins distributed. After distributing all the coins, you should have 2 coins left.

The output should look something like that:

map[Matthew:2 Peter:2 Giana:4 Adriano:7 Elizabeth:5 Sarah:2 Augustus:10 Heidi:5 Emilie:6 Aaron:5]
Coins left: 2
*/
package main

import "fmt"

var (
    coins = 50
    users = []string{
        "Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
        "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
    }
    distribution = make(map[string]int, len(users))
)

func main() {
    for _, user := range users {
        for _, str := range user {
            if distribution[user] <= 10 {
                switch s := string(str); s {
                case "A", "E", "a", "e":
                    distribution[user]++
                    coins--
                case "i", "I":
                    distribution[user] = distribution[user] + 2
                    coins = coins - 2
                case "o", "O":
                    distribution[user] = distribution[user] + 3
                    coins = coins - 3
                case "u", "U":
                    distribution[user] = distribution[user] + 4
                    coins = coins - 4
                }
            }
        }
        if y := distribution[user] - 10; y > 0 {
            distribution[user] = 10
            coins = coins + y
        }
    }
    fmt.Println(distribution)
    fmt.Println("Coins left:", coins)
}
