# Mongo Container
mongocontainer is a Go package that provides a simple interface for interacting with MongoDB. It provides a MongoImpl interface that defines the basic MongoDB operations, such as Find, Upsert, and Delete, as well as a Disconnect method to gracefully disconnect from the database.

The package also provides a Setup function that takes a MongoDB connection URI as an argument and returns a MongoImpl object for use in your application. This makes it easy to set up a MongoDB connection in your Go application without having to worry about the low-level details of managing the connection.

# Installation
To use mongocontainer in your Go project, you can simply import it as follows:
```go
import "github.com/6d6577/mongocontainer"
```
You can then use the Setup function to establish a MongoDB connection:
```go
mongoImpl, err := mongocontainer.Setup("mongodb://localhost:27017")
if err != nil {
    log.Fatal(err)
}
defer mongoImpl.Disconnect()
```
# Usage
Here's an example of how you can use mongocontainer to perform a Find operation on a MongoDB collection:
```go
result := mongoImpl.Find("users", bson.M{"name": "John Doe"})
var user User
err := result.Decode(&user)
if err != nil {
    log.Fatal(err)
}
```
In this example, we're finding a user with the name "John Doe" in the users collection. We then decode the result into a User struct using the Decode method provided by the MongoDB driver.

# Contributing
Contributions to mongocontainer are welcome and encouraged! If you find a bug or would like to suggest a new feature, please open an issue on the GitHub repository.

# License
[This package is licensed under the MIT License. See the LICENSE file for more information.](LICENSE.md)
