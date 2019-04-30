async function insertData(person) {
    let db, collection, result;
    try {
        db = await mongoDb.open();
        collection = await db.collection("users");
        result = await collection.insert(person);
    } catch (e) {
        console.error(e.message);
    }
    console.log(result);
}

insertData({ name: "yika" });