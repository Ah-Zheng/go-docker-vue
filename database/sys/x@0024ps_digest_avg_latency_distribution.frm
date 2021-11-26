TYPE=VIEW
query=select count(0) AS `cnt`,round(`performance_schema`.`events_statements_summary_by_digest`.`AVG_TIMER_WAIT` / 1000000,0) AS `avg_us` from `performance_schema`.`events_statements_summary_by_digest` group by round(`performance_schema`.`events_statements_summary_by_digest`.`AVG_TIMER_WAIT` / 1000000,0)
md5=90f26794b9a8e64fa4b20b5972595230
updatable=0
algorithm=2
definer_user=mariadb.sys
definer_host=localhost
suid=0
with_check_option=0
timestamp=2021-11-23 09:44:14
create-version=2
source=SELECT COUNT(*) cnt,\n       ROUND(avg_timer_wait/1000000) AS avg_us\n  FROM performance_schema.events_statements_summary_by_digest\n GROUP BY avg_us;
client_cs_name=utf8mb3
connection_cl_name=utf8mb3_general_ci
view_body_utf8=select count(0) AS `cnt`,round(`performance_schema`.`events_statements_summary_by_digest`.`AVG_TIMER_WAIT` / 1000000,0) AS `avg_us` from `performance_schema`.`events_statements_summary_by_digest` group by round(`performance_schema`.`events_statements_summary_by_digest`.`AVG_TIMER_WAIT` / 1000000,0)
mariadb-version=100605
