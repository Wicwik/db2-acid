import psycopg2
import secrets
import random
import timeit

conn = psycopg2.connect('postgresql://postgres@localhost:5432/oz')
cur = conn.cursor()

def insert():
    doc_name = f"Generated doc {secrets.token_hex(10)}"
    department = f"Department {secrets.token_hex(10)}"
    contracted_amount = int(random.random() * 1_000_000)
    cur.execute(f"INSERT INTO documents(name, type, created_at, department, contracted_amount) VALUES ('{doc_name}', 'MyType', NOW(), '{department}', {contracted_amount})")
    conn.commit()

time = timeit.timeit(insert, number=1000)
print(f"Execution time for 1000 inserts: {time} seconds.")