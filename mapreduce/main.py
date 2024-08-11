from pyspark.sql import SparkSession
from pyspark.sql.functions import col, row_number, collect_list, sort_array, struct
from pyspark.sql.window import Window


# Initialize Spark session
spark = SparkSession.builder.appName("AutocompleteDemo").getOrCreate()

# Sample search queries and timestamps (simulating the last 24 hours)
data = [
    ("bat", "2024-08-10 12:00:00"),
    ("bad", "2024-08-10 12:10:00"),
    ("bat", "2024-08-10 12:15:00"),
    ("better", "2024-08-10 12:20:00"),
    ("better", "2024-08-10 12:25:00"),
    ("bat", "2024-08-10 12:30:00"),
]

# Create DataFrame
df = spark.createDataFrame(data, ["query", "timestamp"])

df.show()
# +------+-------------------+
# | query|          timestamp|
# +------+-------------------+
# |   bat|2024-08-10 12:00:00|
# |   bad|2024-08-10 12:10:00|
# |   bat|2024-08-10 12:15:00|
# |better|2024-08-10 12:20:00|
# |better|2024-08-10 12:25:00|
# |   bat|2024-08-10 12:30:00|
# +------+-------------------+

# Generate prefixes (return a list of tuples with prefix and query)
# "bat" -> [("b", "bat"), ("ba", "bat"), ("bat", "bat")]
def generate_prefixes(query):
    return [(query[:i+1], query) for i in range(len(query))]

# Apply the function to generate prefixes
df_with_prefixes = df.rdd.flatMap(lambda row: generate_prefixes(row['query'])).toDF(["prefix", "query"])

df_with_prefixes.show()
# +------+------+
# |prefix| query|
# +------+------+
# |     b|   bat|
# |    ba|   bat|
# |   bat|   bat|
# |     b|   bad|
# |    ba|   bad|
# |   bad|   bad|
# |     b|   bat|
# |    ba|   bat|
# |   bat|   bat|
# |     b|better|
# |    be|better|
# |   bet|better|
# |  bett|better|
# | bette|better|
# |better|better|
# |     b|better|
# |    be|better|
# |   bet|better|
# |  bett|better|
# | bette|better|
# +------+------+

# Aggregate and count the occurrences
agg_df = df_with_prefixes.groupBy("prefix", "query").count()

# +------+------+-----+
# |prefix| query|count|
# +------+------+-----+
# |     b|   bat|    3|
# |    ba|   bat|    3|
# |   bat|   bat|    3|
# |   bad|   bad|    1|
# |    ba|   bad|    1|
# |     b|   bad|    1|
# |     b|better|    2|
# |  bett|better|    2|
# |   bet|better|    2|
# |better|better|    2|
# |    be|better|    2|
# | bette|better|    2|
# +------+------+-----+

# Define a window specification
windowSpec = Window.partitionBy("prefix").orderBy(col("count").desc())

# Rank the queries by count within each prefix
ranked_df = agg_df.withColumn("rank", row_number().over(windowSpec))

# Filter to keep only the top k (e.g., 2) queries for each prefix
top_k_df = ranked_df.filter(col("rank") <= 2)

top_k_df.show()

# Group by prefix and collect the queries as a list, ordered by the count
result_df = top_k_df.groupBy("prefix").agg(
    sort_array(collect_list(struct(col("rank"), col("query")))).alias("ranked_queries")
).select("prefix", col("ranked_queries.query").alias("queries"))

# Show the result
result_df.show(truncate=False)

# +------+-------------+
# |prefix|queries      |
# +------+-------------+
# |b     |[bat, better]|
# |ba    |[bat, bad]   |
# |bad   |[bad]        |
# |bat   |[bat]        |
# |be    |[better]     |
# |bet   |[better]     |
# |bett  |[better]     |
# |bette |[better]     |
# |better|[better]     |
# +------+-------------+

# Simulate storing the results
result_df.write.mode("overwrite").parquet("/path/to/autocomplete/results")

print("Autocomplete suggestions stored successfully.")