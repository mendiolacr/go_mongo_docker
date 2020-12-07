db.createUser(
    {
        user: "root",
        pwd: "rootadmin",
        roles: [
            {
                role: "readWrite",
                db: "dockerdb"
            }
        ]
    }
);