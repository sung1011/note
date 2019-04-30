let person = { name: "yika" };
mongoDb
    .open()
    .then(function (database) {
        return database.collection("users");
    })
    .then(function (collection) {
        return collection.insert(person);
    })
    .then(function (result) {
        console.log(result);
    })
    .catch(function (e) {
        throw new Error(e);
    })