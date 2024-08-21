const db = db.getSiblingDB("example");

db.runCommand({
  collMod: "example",
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["exampleId", "title"],
      properties: {
        exampleId: {
          bsonType: "int",
        },
        title: {
          bsonType: "string",
        },
      },
    },
  },
  validationLevel: "strict",
  validationAction: "error",
});

db.example.createIndex({ exampleId: 1 }, { unique: true, background: true });
