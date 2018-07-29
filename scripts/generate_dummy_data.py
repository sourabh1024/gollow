"""
Script insert some dummy data in mysql
To use install dependencies:
pip install numpy
pip install shapely
pip install mysql-client
"""

import argparse
from datetime import datetime, timedelta
import MySQLdb

usage = '''
python generate_dummy_data.py [number_of_rows_to_generate]
'''



if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Insert some dummy data in  table for given input params')

    parser.add_argument('rows', help='rows count')

    args = parser.parse_args()

    db = MySQLdb.connect(host='127.0.0.1',
                         port=3307,
                         user='root',
                         passwd='password',
                         db='test')

    cur = db.cursor()

    utctime = datetime.utcnow()

    for num_rows in range(0, int(args.rows)):

        start = utctime.strftime('%Y-%m-%d %H:%M:00')

        if num_rows % 100 == 0:
            print("Inserting row number : " + str(num_rows))
        try:
            cur.execute("""INSERT INTO dummy_data_large_5 (pid,first_name, last_name,
            balance, max_credit, max_debit, score, is_active) VALUES ({},'{}',
            '{}',{},{},{},{},{})""".format(int(num_rows), "firstname", "lastname",10, 12, 2, 1.4, 1))
        except Exception as e:
            print("Exception in adding data to test db dummy_data table", e)
            db.rollback()

    db.commit()
    db.close()
