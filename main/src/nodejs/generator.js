let co = require("co");

co(function* () {
    let db, collection, result;
    let person = { name: "yika" };
    try {
        db = yield mongoDb.open();
        collection = yield db.collection("users");
        result = yield collection.insert(person);
    } catch (e) {
        console.error(e.message);
    }
    console.log(result);
});