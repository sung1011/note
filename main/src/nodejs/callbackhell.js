mongoDb.open(function (err, db) {
    if (!err) {
        db.collection("users", function (err, collection) {
            if (!err) {
                let person = { name: "yika", age: 20 };
                collection.insert(person, function (err, result) {
                    if (!err) {
                        console.log(result);
                    }
                });
            }
        })
    }
});