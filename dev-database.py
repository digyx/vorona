import sqlite3

conn = sqlite3.Connection('vorona.db')
cur = conn.cursor()

cur.execute("DROP TABLE books")

cur.execute('''
CREATE TABLE books (
    book_id      INTEGER PRIMARY KEY,
    slug         TEXT    NOT NULL UNIQUE,
    title        TEXT    NOT NULL,
    description	 TEXT    NOT NULL,
    release_time INTEGER NOT NULL
)
''')

cur.executemany(
    "INSERT INTO books (slug, title, description, release_time) VALUES (?, ?, ?, ?)",
    [
        ("AzureWitch", "Death of the Azure Witch", "This is a real book.", 1646006400),
        ("BloodOath", "Blood Oath", "Sometimes, someone needs to die.", 1644796800),
    ]
)

conn.commit()
