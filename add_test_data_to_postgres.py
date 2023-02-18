"""
Install pyyaml using 'python3 install pyyaml
You can run this script from 'weber' folder using python3 scripts/add_test_data_to_postgres.py. 
You don't have to run weber or web-ambri for this script. This script inserts the data you mentioned
in test_data.yaml to your postgres DB.
"""

import yaml
import psycopg2
import json
import base64
from csv import DictReader

def get_db_conn():
    conn = psycopg2.connect(
        host="localhost",
        database="test1",
        user="postgres", # Replace postgres user name
        password="hibye123") # Postgres password
    return conn

def cleanup_all_tables(db_conn, table_list):
    print("Cleaning up: ", ','.join(table_list))
    cursor = db_conn.cursor()
    for table in table_list:
        query = "TRUNCATE TABLE {table} CASCADE".format(table = table)
        print(query)
        cursor.execute(query)
        db_conn.commit()

# def add_data_to_database(record, table_name):
#     db_conn = get_db_conn()
#     cursor = db_conn.cursor()
#     column_names = []
#     column_values = []
#     column_names.extend(record.keys())
#     column_values.extend(record.values())
#     # print("Column Names: ", column_names)
#     # print("Column Values: ", column_values)             
#     query = "INSERT INTO {table_name}({column_names_str}) VALUES ({column_values_str});".format(
#         table_name=table_name,
#         column_names_str=','.join(column_names),
#         column_values_str=','.join(repr(value) for value in column_values)
#     )
#     print(query)
#     cursor.execute(query)
#     db_conn.commit()


# def read_yaml():
#     # Fix the file path thingy
#     with open('./scripts/test_data.yaml', 'r') as file:
#         test_data = yaml.safe_load(file)
#         table_list = test_data['tables']
#         #cleanup_all_tables(db_conn, table_list)
        
#         tables_records = {key:value for key,value in test_data.items() if key != "tables"}
#         for table_name,records in tables_records.items():
#             for record in records:
#                 add_data_to_database(record, table_name)
                

# def read_csv():
#     table_name = 'templates'
#     with open('./scripts/'+table_name+'.csv', 'r', encoding='utf-8-sig') as file:
#         dict_reader = DictReader(file)
#         records = list(dict_reader)
#         for record in records:
#             record = {key:val.strip() for key, val in record.items()}
#             #print(record)
#             add_data_to_database(record, table_name)
 
# def add_data(table_name):
#     with open('./scripts/test_data.yaml', 'r') as file:
#         test_data = yaml.safe_load(file)
#         tables_records = {key:value for key,value in test_data.items() if key == table_name}
#         for table_name,records in tables_records.items():
#                 for record in records:
#                     add_data_to_database(record, table_name)


if __name__ == "__main__":
    #read_yaml()
    db_conn = get_db_conn()
    cleanup_all_tables(db_conn, ["user_flows", "flows", "templates", "pods", "cohorts", "publish_tasks"])
    add_data('pods')
    # add_data('pmm_users')
    add_data('cohorts')
    read_csv()
